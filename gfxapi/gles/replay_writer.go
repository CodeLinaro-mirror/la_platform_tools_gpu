////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func loadRemap(ϟb *builder.Builder, key interface{}, ty protocol.Type, val value.Value) {
	if ptr, found := ϟb.Remappings[key]; found {
		ϟb.Load(ty, ptr)
	} else {
		ptr = ϟb.AllocateMemory(uint64(ty.Size(ϟb.Architecture().PointerSize)))
		ϟb.Push(val) // We have an input to an unknown id, use the unmapped value.
		ϟb.Clone(0)
		ϟb.Store(ptr)
		ϟb.Remappings[key] = ptr
	}
}

var funcInfoGlBlendBarrierKHR = builder.FunctionInfo{ID: 0, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlBlendEquationSeparateiEXT = builder.FunctionInfo{ID: 1, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBlendEquationiEXT = builder.FunctionInfo{ID: 2, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendFuncSeparateiEXT = builder.FunctionInfo{ID: 3, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlBlendFunciEXT = builder.FunctionInfo{ID: 4, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlColorMaskiEXT = builder.FunctionInfo{ID: 5, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlCopyImageSubDataEXT = builder.FunctionInfo{ID: 6, ReturnType: protocol.TypeVoid, Parameters: 15}
var funcInfoGlDebugMessageCallbackKHR = builder.FunctionInfo{ID: 7, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDebugMessageControlKHR = builder.FunctionInfo{ID: 8, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlDebugMessageInsertKHR = builder.FunctionInfo{ID: 9, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlDisableiEXT = builder.FunctionInfo{ID: 10, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEnableiEXT = builder.FunctionInfo{ID: 11, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlFramebufferTextureEXT = builder.FunctionInfo{ID: 12, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetDebugMessageLogKHR = builder.FunctionInfo{ID: 13, ReturnType: protocol.TypeUint32, Parameters: 8}
var funcInfoGlGetObjectLabelKHR = builder.FunctionInfo{ID: 14, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetObjectPtrLabelKHR = builder.FunctionInfo{ID: 15, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetPointervKHR = builder.FunctionInfo{ID: 16, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetSamplerParameterIivEXT = builder.FunctionInfo{ID: 17, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetSamplerParameterIuivEXT = builder.FunctionInfo{ID: 18, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameterIivEXT = builder.FunctionInfo{ID: 19, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameterIuivEXT = builder.FunctionInfo{ID: 20, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlIsEnablediEXT = builder.FunctionInfo{ID: 21, ReturnType: protocol.TypeBool, Parameters: 2}
var funcInfoGlMinSampleShadingOES = builder.FunctionInfo{ID: 22, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlObjectLabelKHR = builder.FunctionInfo{ID: 23, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlObjectPtrLabelKHR = builder.FunctionInfo{ID: 24, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlPatchParameteriEXT = builder.FunctionInfo{ID: 25, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPopDebugGroupKHR = builder.FunctionInfo{ID: 26, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlPrimitiveBoundingBoxEXT = builder.FunctionInfo{ID: 27, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlPushDebugGroupKHR = builder.FunctionInfo{ID: 28, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlSamplerParameterIivEXT = builder.FunctionInfo{ID: 29, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlSamplerParameterIuivEXT = builder.FunctionInfo{ID: 30, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexBufferEXT = builder.FunctionInfo{ID: 31, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexBufferRangeEXT = builder.FunctionInfo{ID: 32, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTexParameterIivEXT = builder.FunctionInfo{ID: 33, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameterIuivEXT = builder.FunctionInfo{ID: 34, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexStorage3DMultisampleOES = builder.FunctionInfo{ID: 35, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlBeginQuery = builder.FunctionInfo{ID: 36, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteQueries = builder.FunctionInfo{ID: 37, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndQuery = builder.FunctionInfo{ID: 38, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGenQueries = builder.FunctionInfo{ID: 39, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetQueryObjectuiv = builder.FunctionInfo{ID: 40, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryiv = builder.FunctionInfo{ID: 41, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlIsQuery = builder.FunctionInfo{ID: 42, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlBindBuffer = builder.FunctionInfo{ID: 43, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindBufferBase = builder.FunctionInfo{ID: 44, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBindBufferRange = builder.FunctionInfo{ID: 45, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlBufferData = builder.FunctionInfo{ID: 46, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBufferSubData = builder.FunctionInfo{ID: 47, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlCopyBufferSubData = builder.FunctionInfo{ID: 48, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDeleteBuffers = builder.FunctionInfo{ID: 49, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenBuffers = builder.FunctionInfo{ID: 50, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetBufferParameteri64v = builder.FunctionInfo{ID: 51, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetBufferParameteriv = builder.FunctionInfo{ID: 52, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetBufferPointerv = builder.FunctionInfo{ID: 53, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlIsBuffer = builder.FunctionInfo{ID: 54, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlMapBufferRange = builder.FunctionInfo{ID: 55, ReturnType: protocol.TypeVolatilePointer, Parameters: 4}
var funcInfoGlUnmapBuffer = builder.FunctionInfo{ID: 56, ReturnType: protocol.TypeUint8, Parameters: 1}
var funcInfoGlDrawArrays = builder.FunctionInfo{ID: 57, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlDrawArraysIndirect = builder.FunctionInfo{ID: 58, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDrawArraysInstanced = builder.FunctionInfo{ID: 59, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawBuffers = builder.FunctionInfo{ID: 60, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDrawElements = builder.FunctionInfo{ID: 61, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawElementsIndirect = builder.FunctionInfo{ID: 62, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlDrawElementsInstanced = builder.FunctionInfo{ID: 63, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDrawRangeElements = builder.FunctionInfo{ID: 64, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlActiveShaderProgramEXT = builder.FunctionInfo{ID: 65, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlAlphaFuncQCOM = builder.FunctionInfo{ID: 66, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginConditionalRenderNV = builder.FunctionInfo{ID: 67, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBeginPerfMonitorAMD = builder.FunctionInfo{ID: 68, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBeginPerfQueryINTEL = builder.FunctionInfo{ID: 69, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBeginQueryEXT = builder.FunctionInfo{ID: 70, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindProgramPipelineEXT = builder.FunctionInfo{ID: 71, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBindVertexArrayOES = builder.FunctionInfo{ID: 72, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBlendBarrierNV = builder.FunctionInfo{ID: 73, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlBlendEquationSeparateiOES = builder.FunctionInfo{ID: 74, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBlendEquationiOES = builder.FunctionInfo{ID: 75, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendFuncSeparateiOES = builder.FunctionInfo{ID: 76, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlBlendFunciOES = builder.FunctionInfo{ID: 77, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBlendParameteriNV = builder.FunctionInfo{ID: 78, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlitFramebufferANGLE = builder.FunctionInfo{ID: 79, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlBlitFramebufferNV = builder.FunctionInfo{ID: 80, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlBufferStorageEXT = builder.FunctionInfo{ID: 81, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlClientWaitSyncAPPLE = builder.FunctionInfo{ID: 82, ReturnType: protocol.TypeUint32, Parameters: 3}
var funcInfoGlColorMaskiOES = builder.FunctionInfo{ID: 83, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlCompressedTexImage3DOES = builder.FunctionInfo{ID: 84, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlCompressedTexSubImage3DOES = builder.FunctionInfo{ID: 85, ReturnType: protocol.TypeVoid, Parameters: 11}
var funcInfoGlCopyBufferSubDataNV = builder.FunctionInfo{ID: 86, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlCopyImageSubDataOES = builder.FunctionInfo{ID: 87, ReturnType: protocol.TypeVoid, Parameters: 15}
var funcInfoGlCopyPathNV = builder.FunctionInfo{ID: 88, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCopyTexSubImage3DOES = builder.FunctionInfo{ID: 89, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlCopyTextureLevelsAPPLE = builder.FunctionInfo{ID: 90, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlCoverFillPathInstancedNV = builder.FunctionInfo{ID: 91, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlCoverFillPathNV = builder.FunctionInfo{ID: 92, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCoverStrokePathInstancedNV = builder.FunctionInfo{ID: 93, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlCoverStrokePathNV = builder.FunctionInfo{ID: 94, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCoverageMaskNV = builder.FunctionInfo{ID: 95, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCoverageOperationNV = builder.FunctionInfo{ID: 96, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCreatePerfQueryINTEL = builder.FunctionInfo{ID: 97, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCreateShaderProgramvEXT = builder.FunctionInfo{ID: 98, ReturnType: protocol.TypeUint32, Parameters: 3}
var funcInfoGlDeleteFencesNV = builder.FunctionInfo{ID: 99, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeletePathsNV = builder.FunctionInfo{ID: 100, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeletePerfMonitorsAMD = builder.FunctionInfo{ID: 101, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeletePerfQueryINTEL = builder.FunctionInfo{ID: 102, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteProgramPipelinesEXT = builder.FunctionInfo{ID: 103, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteQueriesEXT = builder.FunctionInfo{ID: 104, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteSyncAPPLE = builder.FunctionInfo{ID: 105, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteVertexArraysOES = builder.FunctionInfo{ID: 106, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDepthRangeArrayfvNV = builder.FunctionInfo{ID: 107, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlDepthRangeIndexedfNV = builder.FunctionInfo{ID: 108, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlDisableDriverControlQCOM = builder.FunctionInfo{ID: 109, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDisableiNV = builder.FunctionInfo{ID: 110, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDisableiOES = builder.FunctionInfo{ID: 111, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDiscardFramebufferEXT = builder.FunctionInfo{ID: 112, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlDrawArraysInstancedANGLE = builder.FunctionInfo{ID: 113, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawArraysInstancedBaseInstanceEXT = builder.FunctionInfo{ID: 114, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDrawArraysInstancedEXT = builder.FunctionInfo{ID: 115, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawArraysInstancedNV = builder.FunctionInfo{ID: 116, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDrawBuffersEXT = builder.FunctionInfo{ID: 117, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDrawBuffersIndexedEXT = builder.FunctionInfo{ID: 118, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlDrawBuffersNV = builder.FunctionInfo{ID: 119, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDrawElementsBaseVertexEXT = builder.FunctionInfo{ID: 120, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDrawElementsBaseVertexOES = builder.FunctionInfo{ID: 121, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDrawElementsInstancedANGLE = builder.FunctionInfo{ID: 122, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDrawElementsInstancedBaseInstanceEXT = builder.FunctionInfo{ID: 123, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlDrawElementsInstancedBaseVertexEXT = builder.FunctionInfo{ID: 124, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlDrawElementsInstancedBaseVertexOES = builder.FunctionInfo{ID: 125, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlDrawElementsInstancedBaseVertexBaseInstanceEXT = builder.FunctionInfo{ID: 126, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlDrawElementsInstancedEXT = builder.FunctionInfo{ID: 127, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDrawElementsInstancedNV = builder.FunctionInfo{ID: 128, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlDrawRangeElementsBaseVertexEXT = builder.FunctionInfo{ID: 129, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlDrawRangeElementsBaseVertexOES = builder.FunctionInfo{ID: 130, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlEGLImageTargetRenderbufferStorageOES = builder.FunctionInfo{ID: 131, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEGLImageTargetTexture2DOES = builder.FunctionInfo{ID: 132, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEnableDriverControlQCOM = builder.FunctionInfo{ID: 133, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlEnableiNV = builder.FunctionInfo{ID: 134, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEnableiOES = builder.FunctionInfo{ID: 135, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndConditionalRenderNV = builder.FunctionInfo{ID: 136, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlEndPerfMonitorAMD = builder.FunctionInfo{ID: 137, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlEndPerfQueryINTEL = builder.FunctionInfo{ID: 138, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlEndQueryEXT = builder.FunctionInfo{ID: 139, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlEndTilingQCOM = builder.FunctionInfo{ID: 140, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlExtGetBufferPointervQCOM = builder.FunctionInfo{ID: 141, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlExtGetBuffersQCOM = builder.FunctionInfo{ID: 142, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlExtGetFramebuffersQCOM = builder.FunctionInfo{ID: 143, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlExtGetProgramBinarySourceQCOM = builder.FunctionInfo{ID: 144, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlExtGetProgramsQCOM = builder.FunctionInfo{ID: 145, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlExtGetRenderbuffersQCOM = builder.FunctionInfo{ID: 146, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlExtGetShadersQCOM = builder.FunctionInfo{ID: 147, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlExtGetTexLevelParameterivQCOM = builder.FunctionInfo{ID: 148, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlExtGetTexSubImageQCOM = builder.FunctionInfo{ID: 149, ReturnType: protocol.TypeVoid, Parameters: 11}
var funcInfoGlExtGetTexturesQCOM = builder.FunctionInfo{ID: 150, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlExtIsProgramBinaryQCOM = builder.FunctionInfo{ID: 151, ReturnType: protocol.TypeUint8, Parameters: 1}
var funcInfoGlExtTexObjectStateOverrideiQCOM = builder.FunctionInfo{ID: 152, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlFenceSyncAPPLE = builder.FunctionInfo{ID: 153, ReturnType: protocol.TypeUint64, Parameters: 2}
var funcInfoGlFinishFenceNV = builder.FunctionInfo{ID: 154, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlFlushMappedBufferRangeEXT = builder.FunctionInfo{ID: 155, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlFramebufferTexture2DMultisampleEXT = builder.FunctionInfo{ID: 156, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlFramebufferTexture2DMultisampleIMG = builder.FunctionInfo{ID: 157, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlFramebufferTexture3DOES = builder.FunctionInfo{ID: 158, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlFramebufferTextureOES = builder.FunctionInfo{ID: 159, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFramebufferTextureMultiviewOVR = builder.FunctionInfo{ID: 160, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlGenFencesNV = builder.FunctionInfo{ID: 161, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenPathsNV = builder.FunctionInfo{ID: 162, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlGenPerfMonitorsAMD = builder.FunctionInfo{ID: 163, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenProgramPipelinesEXT = builder.FunctionInfo{ID: 164, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenQueriesEXT = builder.FunctionInfo{ID: 165, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenVertexArraysOES = builder.FunctionInfo{ID: 166, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetBufferPointervOES = builder.FunctionInfo{ID: 167, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetDriverControlStringQCOM = builder.FunctionInfo{ID: 168, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetDriverControlsQCOM = builder.FunctionInfo{ID: 169, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetFenceivNV = builder.FunctionInfo{ID: 170, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetFirstPerfQueryIdINTEL = builder.FunctionInfo{ID: 171, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetFloati_vNV = builder.FunctionInfo{ID: 172, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetGraphicsResetStatusEXT = builder.FunctionInfo{ID: 173, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlGetGraphicsResetStatusKHR = builder.FunctionInfo{ID: 174, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlGetImageHandleNV = builder.FunctionInfo{ID: 175, ReturnType: protocol.TypeUint64, Parameters: 5}
var funcInfoGlGetInteger64vAPPLE = builder.FunctionInfo{ID: 176, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetIntegeri_vEXT = builder.FunctionInfo{ID: 177, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetInternalformatSampleivNV = builder.FunctionInfo{ID: 178, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlGetNextPerfQueryIdINTEL = builder.FunctionInfo{ID: 179, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetObjectLabelEXT = builder.FunctionInfo{ID: 180, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetPathCommandsNV = builder.FunctionInfo{ID: 181, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetPathCoordsNV = builder.FunctionInfo{ID: 182, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetPathDashArrayNV = builder.FunctionInfo{ID: 183, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetPathLengthNV = builder.FunctionInfo{ID: 184, ReturnType: protocol.TypeFloat, Parameters: 3}
var funcInfoGlGetPathMetricRangeNV = builder.FunctionInfo{ID: 185, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetPathMetricsNV = builder.FunctionInfo{ID: 186, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetPathParameterfvNV = builder.FunctionInfo{ID: 187, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetPathParameterivNV = builder.FunctionInfo{ID: 188, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetPathSpacingNV = builder.FunctionInfo{ID: 189, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlGetPerfCounterInfoINTEL = builder.FunctionInfo{ID: 190, ReturnType: protocol.TypeVoid, Parameters: 11}
var funcInfoGlGetPerfMonitorCounterDataAMD = builder.FunctionInfo{ID: 191, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetPerfMonitorCounterInfoAMD = builder.FunctionInfo{ID: 192, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetPerfMonitorCounterStringAMD = builder.FunctionInfo{ID: 193, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetPerfMonitorCountersAMD = builder.FunctionInfo{ID: 194, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetPerfMonitorGroupStringAMD = builder.FunctionInfo{ID: 195, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetPerfMonitorGroupsAMD = builder.FunctionInfo{ID: 196, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetPerfQueryDataINTEL = builder.FunctionInfo{ID: 197, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetPerfQueryIdByNameINTEL = builder.FunctionInfo{ID: 198, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetPerfQueryInfoINTEL = builder.FunctionInfo{ID: 199, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetProgramBinaryOES = builder.FunctionInfo{ID: 200, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetProgramPipelineInfoLogEXT = builder.FunctionInfo{ID: 201, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetProgramPipelineivEXT = builder.FunctionInfo{ID: 202, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetProgramResourcefvNV = builder.FunctionInfo{ID: 203, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlGetQueryObjecti64vEXT = builder.FunctionInfo{ID: 204, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectivEXT = builder.FunctionInfo{ID: 205, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectui64vEXT = builder.FunctionInfo{ID: 206, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectuivEXT = builder.FunctionInfo{ID: 207, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryivEXT = builder.FunctionInfo{ID: 208, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetSamplerParameterIivOES = builder.FunctionInfo{ID: 209, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetSamplerParameterIuivOES = builder.FunctionInfo{ID: 210, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetSyncivAPPLE = builder.FunctionInfo{ID: 211, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetTexParameterIivOES = builder.FunctionInfo{ID: 212, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameterIuivOES = builder.FunctionInfo{ID: 213, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTextureHandleNV = builder.FunctionInfo{ID: 214, ReturnType: protocol.TypeUint64, Parameters: 1}
var funcInfoGlGetTextureSamplerHandleNV = builder.FunctionInfo{ID: 215, ReturnType: protocol.TypeUint64, Parameters: 2}
var funcInfoGlGetTranslatedShaderSourceANGLE = builder.FunctionInfo{ID: 216, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetnUniformfvEXT = builder.FunctionInfo{ID: 217, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetnUniformfvKHR = builder.FunctionInfo{ID: 218, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetnUniformivEXT = builder.FunctionInfo{ID: 219, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetnUniformivKHR = builder.FunctionInfo{ID: 220, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetnUniformuivKHR = builder.FunctionInfo{ID: 221, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlInsertEventMarkerEXT = builder.FunctionInfo{ID: 222, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlInterpolatePathsNV = builder.FunctionInfo{ID: 223, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlIsEnablediOES = builder.FunctionInfo{ID: 224, ReturnType: protocol.TypeBool, Parameters: 2}
var funcInfoGlIsEnablediNV = builder.FunctionInfo{ID: 225, ReturnType: protocol.TypeBool, Parameters: 2}
var funcInfoGlIsFenceNV = builder.FunctionInfo{ID: 226, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsImageHandleResidentNV = builder.FunctionInfo{ID: 227, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsPathNV = builder.FunctionInfo{ID: 228, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsPointInFillPathNV = builder.FunctionInfo{ID: 229, ReturnType: protocol.TypeBool, Parameters: 4}
var funcInfoGlIsPointInStrokePathNV = builder.FunctionInfo{ID: 230, ReturnType: protocol.TypeBool, Parameters: 3}
var funcInfoGlIsProgramPipelineEXT = builder.FunctionInfo{ID: 231, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsQueryEXT = builder.FunctionInfo{ID: 232, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsSyncAPPLE = builder.FunctionInfo{ID: 233, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsTextureHandleResidentNV = builder.FunctionInfo{ID: 234, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsVertexArrayOES = builder.FunctionInfo{ID: 235, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlLabelObjectEXT = builder.FunctionInfo{ID: 236, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlMakeImageHandleNonResidentNV = builder.FunctionInfo{ID: 237, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlMakeImageHandleResidentNV = builder.FunctionInfo{ID: 238, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlMakeTextureHandleNonResidentNV = builder.FunctionInfo{ID: 239, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlMakeTextureHandleResidentNV = builder.FunctionInfo{ID: 240, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlMapBufferOES = builder.FunctionInfo{ID: 241, ReturnType: protocol.TypeVolatilePointer, Parameters: 2}
var funcInfoGlMapBufferRangeEXT = builder.FunctionInfo{ID: 242, ReturnType: protocol.TypeVolatilePointer, Parameters: 4}
var funcInfoGlMatrixLoad3x2fNV = builder.FunctionInfo{ID: 243, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlMatrixLoad3x3fNV = builder.FunctionInfo{ID: 244, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlMatrixLoadTranspose3x3fNV = builder.FunctionInfo{ID: 245, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlMatrixMult3x2fNV = builder.FunctionInfo{ID: 246, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlMatrixMult3x3fNV = builder.FunctionInfo{ID: 247, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlMatrixMultTranspose3x3fNV = builder.FunctionInfo{ID: 248, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlMultiDrawArraysEXT = builder.FunctionInfo{ID: 249, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlMultiDrawArraysIndirectEXT = builder.FunctionInfo{ID: 250, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlMultiDrawElementsBaseVertexEXT = builder.FunctionInfo{ID: 251, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlMultiDrawElementsBaseVertexOES = builder.FunctionInfo{ID: 252, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlMultiDrawElementsEXT = builder.FunctionInfo{ID: 253, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlMultiDrawElementsIndirectEXT = builder.FunctionInfo{ID: 254, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlPatchParameteriOES = builder.FunctionInfo{ID: 255, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPathCommandsNV = builder.FunctionInfo{ID: 256, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlPathCoordsNV = builder.FunctionInfo{ID: 257, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlPathCoverDepthFuncNV = builder.FunctionInfo{ID: 258, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlPathDashArrayNV = builder.FunctionInfo{ID: 259, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlPathGlyphIndexArrayNV = builder.FunctionInfo{ID: 260, ReturnType: protocol.TypeUint32, Parameters: 8}
var funcInfoGlPathGlyphIndexRangeNV = builder.FunctionInfo{ID: 261, ReturnType: protocol.TypeUint32, Parameters: 6}
var funcInfoGlPathGlyphRangeNV = builder.FunctionInfo{ID: 262, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlPathGlyphsNV = builder.FunctionInfo{ID: 263, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlPathMemoryGlyphIndexArrayNV = builder.FunctionInfo{ID: 264, ReturnType: protocol.TypeUint32, Parameters: 9}
var funcInfoGlPathParameterfNV = builder.FunctionInfo{ID: 265, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlPathParameterfvNV = builder.FunctionInfo{ID: 266, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlPathParameteriNV = builder.FunctionInfo{ID: 267, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlPathParameterivNV = builder.FunctionInfo{ID: 268, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlPathStencilDepthOffsetNV = builder.FunctionInfo{ID: 269, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPathStencilFuncNV = builder.FunctionInfo{ID: 270, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlPathStringNV = builder.FunctionInfo{ID: 271, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlPathSubCommandsNV = builder.FunctionInfo{ID: 272, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlPathSubCoordsNV = builder.FunctionInfo{ID: 273, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlPointAlongPathNV = builder.FunctionInfo{ID: 274, ReturnType: protocol.TypeUint8, Parameters: 8}
var funcInfoGlPolygonModeNV = builder.FunctionInfo{ID: 275, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlPopGroupMarkerEXT = builder.FunctionInfo{ID: 276, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlPrimitiveBoundingBoxOES = builder.FunctionInfo{ID: 277, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlProgramBinaryOES = builder.FunctionInfo{ID: 278, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramParameteriEXT = builder.FunctionInfo{ID: 279, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramPathFragmentInputGenNV = builder.FunctionInfo{ID: 280, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniform1fEXT = builder.FunctionInfo{ID: 281, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniform1fvEXT = builder.FunctionInfo{ID: 282, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform1iEXT = builder.FunctionInfo{ID: 283, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniform1ivEXT = builder.FunctionInfo{ID: 284, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform1uiEXT = builder.FunctionInfo{ID: 285, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniform1uivEXT = builder.FunctionInfo{ID: 286, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2fEXT = builder.FunctionInfo{ID: 287, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2fvEXT = builder.FunctionInfo{ID: 288, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2iEXT = builder.FunctionInfo{ID: 289, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2ivEXT = builder.FunctionInfo{ID: 290, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2uiEXT = builder.FunctionInfo{ID: 291, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2uivEXT = builder.FunctionInfo{ID: 292, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform3fEXT = builder.FunctionInfo{ID: 293, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniform3fvEXT = builder.FunctionInfo{ID: 294, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform3iEXT = builder.FunctionInfo{ID: 295, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniform3ivEXT = builder.FunctionInfo{ID: 296, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform3uiEXT = builder.FunctionInfo{ID: 297, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniform3uivEXT = builder.FunctionInfo{ID: 298, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform4fEXT = builder.FunctionInfo{ID: 299, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlProgramUniform4fvEXT = builder.FunctionInfo{ID: 300, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform4iEXT = builder.FunctionInfo{ID: 301, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlProgramUniform4ivEXT = builder.FunctionInfo{ID: 302, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform4uiEXT = builder.FunctionInfo{ID: 303, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlProgramUniform4uivEXT = builder.FunctionInfo{ID: 304, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniformHandleui64NV = builder.FunctionInfo{ID: 305, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniformHandleui64vNV = builder.FunctionInfo{ID: 306, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniformMatrix2fvEXT = builder.FunctionInfo{ID: 307, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix2x3fvEXT = builder.FunctionInfo{ID: 308, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix2x4fvEXT = builder.FunctionInfo{ID: 309, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix3fvEXT = builder.FunctionInfo{ID: 310, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix3x2fvEXT = builder.FunctionInfo{ID: 311, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix3x4fvEXT = builder.FunctionInfo{ID: 312, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix4fvEXT = builder.FunctionInfo{ID: 313, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix4x2fvEXT = builder.FunctionInfo{ID: 314, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix4x3fvEXT = builder.FunctionInfo{ID: 315, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlPushGroupMarkerEXT = builder.FunctionInfo{ID: 316, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlQueryCounterEXT = builder.FunctionInfo{ID: 317, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlReadBufferIndexedEXT = builder.FunctionInfo{ID: 318, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlReadBufferNV = builder.FunctionInfo{ID: 319, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlReadnPixelsEXT = builder.FunctionInfo{ID: 320, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlReadnPixelsKHR = builder.FunctionInfo{ID: 321, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlRenderbufferStorageMultisampleANGLE = builder.FunctionInfo{ID: 322, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlRenderbufferStorageMultisampleAPPLE = builder.FunctionInfo{ID: 323, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlRenderbufferStorageMultisampleEXT = builder.FunctionInfo{ID: 324, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlRenderbufferStorageMultisampleIMG = builder.FunctionInfo{ID: 325, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlRenderbufferStorageMultisampleNV = builder.FunctionInfo{ID: 326, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlResolveMultisampleFramebufferAPPLE = builder.FunctionInfo{ID: 327, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlSamplerParameterIivOES = builder.FunctionInfo{ID: 328, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlSamplerParameterIuivOES = builder.FunctionInfo{ID: 329, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlScissorArrayvNV = builder.FunctionInfo{ID: 330, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlScissorIndexedNV = builder.FunctionInfo{ID: 331, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlScissorIndexedvNV = builder.FunctionInfo{ID: 332, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlSelectPerfMonitorCountersAMD = builder.FunctionInfo{ID: 333, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlSetFenceNV = builder.FunctionInfo{ID: 334, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlStartTilingQCOM = builder.FunctionInfo{ID: 335, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlStencilFillPathInstancedNV = builder.FunctionInfo{ID: 336, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlStencilFillPathNV = builder.FunctionInfo{ID: 337, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlStencilStrokePathInstancedNV = builder.FunctionInfo{ID: 338, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlStencilStrokePathNV = builder.FunctionInfo{ID: 339, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlStencilThenCoverFillPathInstancedNV = builder.FunctionInfo{ID: 340, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlStencilThenCoverFillPathNV = builder.FunctionInfo{ID: 341, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilThenCoverStrokePathInstancedNV = builder.FunctionInfo{ID: 342, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlStencilThenCoverStrokePathNV = builder.FunctionInfo{ID: 343, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlTestFenceNV = builder.FunctionInfo{ID: 344, ReturnType: protocol.TypeUint8, Parameters: 1}
var funcInfoGlTexBufferOES = builder.FunctionInfo{ID: 345, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexBufferRangeOES = builder.FunctionInfo{ID: 346, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTexImage3DOES = builder.FunctionInfo{ID: 347, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlTexPageCommitmentARB = builder.FunctionInfo{ID: 348, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlTexParameterIivOES = builder.FunctionInfo{ID: 349, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameterIuivOES = builder.FunctionInfo{ID: 350, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexStorage1DEXT = builder.FunctionInfo{ID: 351, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlTexStorage2DEXT = builder.FunctionInfo{ID: 352, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTexStorage3DEXT = builder.FunctionInfo{ID: 353, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTexSubImage3DOES = builder.FunctionInfo{ID: 354, ReturnType: protocol.TypeVoid, Parameters: 11}
var funcInfoGlTextureStorage1DEXT = builder.FunctionInfo{ID: 355, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTextureStorage2DEXT = builder.FunctionInfo{ID: 356, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTextureStorage3DEXT = builder.FunctionInfo{ID: 357, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlTextureViewEXT = builder.FunctionInfo{ID: 358, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlTextureViewOES = builder.FunctionInfo{ID: 359, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlTransformPathNV = builder.FunctionInfo{ID: 360, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformHandleui64NV = builder.FunctionInfo{ID: 361, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniformHandleui64vNV = builder.FunctionInfo{ID: 362, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniformMatrix2x3fvNV = builder.FunctionInfo{ID: 363, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix2x4fvNV = builder.FunctionInfo{ID: 364, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3x2fvNV = builder.FunctionInfo{ID: 365, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3x4fvNV = builder.FunctionInfo{ID: 366, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4x2fvNV = builder.FunctionInfo{ID: 367, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4x3fvNV = builder.FunctionInfo{ID: 368, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUnmapBufferOES = builder.FunctionInfo{ID: 369, ReturnType: protocol.TypeUint8, Parameters: 1}
var funcInfoGlUseProgramStagesEXT = builder.FunctionInfo{ID: 370, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlValidateProgramPipelineEXT = builder.FunctionInfo{ID: 371, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlVertexAttribDivisorANGLE = builder.FunctionInfo{ID: 372, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttribDivisorEXT = builder.FunctionInfo{ID: 373, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttribDivisorNV = builder.FunctionInfo{ID: 374, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlViewportArrayvNV = builder.FunctionInfo{ID: 375, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlViewportIndexedfNV = builder.FunctionInfo{ID: 376, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlViewportIndexedfvNV = builder.FunctionInfo{ID: 377, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlWaitSyncAPPLE = builder.FunctionInfo{ID: 378, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlWeightPathsNV = builder.FunctionInfo{ID: 379, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlCoverageModulationNV = builder.FunctionInfo{ID: 380, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCoverageModulationTableNV = builder.FunctionInfo{ID: 381, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlFragmentCoverageColorNV = builder.FunctionInfo{ID: 382, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlFramebufferSampleLocationsfvNV = builder.FunctionInfo{ID: 383, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetCoverageModulationTableNV = builder.FunctionInfo{ID: 384, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlNamedFramebufferSampleLocationsfvNV = builder.FunctionInfo{ID: 385, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlRasterSamplesEXT = builder.FunctionInfo{ID: 386, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlResolveDepthValuesNV = builder.FunctionInfo{ID: 387, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlSubpixelPrecisionBiasNV = builder.FunctionInfo{ID: 388, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendColor = builder.FunctionInfo{ID: 389, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBlendEquation = builder.FunctionInfo{ID: 390, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBlendEquationSeparate = builder.FunctionInfo{ID: 391, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendFunc = builder.FunctionInfo{ID: 392, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlendFuncSeparate = builder.FunctionInfo{ID: 393, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDepthFunc = builder.FunctionInfo{ID: 394, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlSampleCoverage = builder.FunctionInfo{ID: 395, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlSampleMaski = builder.FunctionInfo{ID: 396, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlScissor = builder.FunctionInfo{ID: 397, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilFunc = builder.FunctionInfo{ID: 398, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlStencilFuncSeparate = builder.FunctionInfo{ID: 399, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlStencilOp = builder.FunctionInfo{ID: 400, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlStencilOpSeparate = builder.FunctionInfo{ID: 401, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBindFramebuffer = builder.FunctionInfo{ID: 402, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindRenderbuffer = builder.FunctionInfo{ID: 403, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBlitFramebuffer = builder.FunctionInfo{ID: 404, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlCheckFramebufferStatus = builder.FunctionInfo{ID: 405, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlClear = builder.FunctionInfo{ID: 406, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearBufferfi = builder.FunctionInfo{ID: 407, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlClearBufferfv = builder.FunctionInfo{ID: 408, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlClearBufferiv = builder.FunctionInfo{ID: 409, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlClearBufferuiv = builder.FunctionInfo{ID: 410, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlClearColor = builder.FunctionInfo{ID: 411, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlClearDepthf = builder.FunctionInfo{ID: 412, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlClearStencil = builder.FunctionInfo{ID: 413, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlColorMask = builder.FunctionInfo{ID: 414, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteFramebuffers = builder.FunctionInfo{ID: 415, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteRenderbuffers = builder.FunctionInfo{ID: 416, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDepthMask = builder.FunctionInfo{ID: 417, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlFramebufferParameteri = builder.FunctionInfo{ID: 418, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlFramebufferRenderbuffer = builder.FunctionInfo{ID: 419, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlFramebufferTexture2D = builder.FunctionInfo{ID: 420, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlFramebufferTextureLayer = builder.FunctionInfo{ID: 421, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGenFramebuffers = builder.FunctionInfo{ID: 422, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenRenderbuffers = builder.FunctionInfo{ID: 423, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetFramebufferAttachmentParameteriv = builder.FunctionInfo{ID: 424, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetFramebufferParameteriv = builder.FunctionInfo{ID: 425, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetRenderbufferParameteriv = builder.FunctionInfo{ID: 426, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlInvalidateFramebuffer = builder.FunctionInfo{ID: 427, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlInvalidateSubFramebuffer = builder.FunctionInfo{ID: 428, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlIsFramebuffer = builder.FunctionInfo{ID: 429, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsRenderbuffer = builder.FunctionInfo{ID: 430, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlReadBuffer = builder.FunctionInfo{ID: 431, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlReadPixels = builder.FunctionInfo{ID: 432, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlRenderbufferStorage = builder.FunctionInfo{ID: 433, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlRenderbufferStorageMultisample = builder.FunctionInfo{ID: 434, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlStencilMask = builder.FunctionInfo{ID: 435, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlStencilMaskSeparate = builder.FunctionInfo{ID: 436, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDisable = builder.FunctionInfo{ID: 437, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlEnable = builder.FunctionInfo{ID: 438, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlFinish = builder.FunctionInfo{ID: 439, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlFlush = builder.FunctionInfo{ID: 440, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlFlushMappedBufferRange = builder.FunctionInfo{ID: 441, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetError = builder.FunctionInfo{ID: 442, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlHint = builder.FunctionInfo{ID: 443, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlActiveShaderProgram = builder.FunctionInfo{ID: 444, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlAttachShader = builder.FunctionInfo{ID: 445, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindAttribLocation = builder.FunctionInfo{ID: 446, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlBindProgramPipeline = builder.FunctionInfo{ID: 447, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCompileShader = builder.FunctionInfo{ID: 448, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCreateProgram = builder.FunctionInfo{ID: 449, ReturnType: protocol.TypeUint32, Parameters: 0}
var funcInfoGlCreateShader = builder.FunctionInfo{ID: 450, ReturnType: protocol.TypeUint32, Parameters: 1}
var funcInfoGlCreateShaderProgramv = builder.FunctionInfo{ID: 451, ReturnType: protocol.TypeUint32, Parameters: 3}
var funcInfoGlDeleteProgram = builder.FunctionInfo{ID: 452, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDeleteProgramPipelines = builder.FunctionInfo{ID: 453, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteShader = builder.FunctionInfo{ID: 454, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDetachShader = builder.FunctionInfo{ID: 455, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDispatchCompute = builder.FunctionInfo{ID: 456, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlDispatchComputeIndirect = builder.FunctionInfo{ID: 457, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGenProgramPipelines = builder.FunctionInfo{ID: 458, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetActiveAttrib = builder.FunctionInfo{ID: 459, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetActiveUniform = builder.FunctionInfo{ID: 460, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlGetActiveUniformBlockName = builder.FunctionInfo{ID: 461, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetActiveUniformBlockiv = builder.FunctionInfo{ID: 462, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetActiveUniformsiv = builder.FunctionInfo{ID: 463, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetAttachedShaders = builder.FunctionInfo{ID: 464, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetAttribLocation = builder.FunctionInfo{ID: 465, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlGetFragDataLocation = builder.FunctionInfo{ID: 466, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlGetProgramBinary = builder.FunctionInfo{ID: 467, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetProgramInfoLog = builder.FunctionInfo{ID: 468, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetProgramInterfaceiv = builder.FunctionInfo{ID: 469, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetProgramPipelineInfoLog = builder.FunctionInfo{ID: 470, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetProgramPipelineiv = builder.FunctionInfo{ID: 471, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetProgramResourceIndex = builder.FunctionInfo{ID: 472, ReturnType: protocol.TypeUint32, Parameters: 3}
var funcInfoGlGetProgramResourceLocation = builder.FunctionInfo{ID: 473, ReturnType: protocol.TypeInt32, Parameters: 3}
var funcInfoGlGetProgramResourceName = builder.FunctionInfo{ID: 474, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlGetProgramResourceiv = builder.FunctionInfo{ID: 475, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlGetProgramiv = builder.FunctionInfo{ID: 476, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetShaderInfoLog = builder.FunctionInfo{ID: 477, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetShaderPrecisionFormat = builder.FunctionInfo{ID: 478, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetShaderSource = builder.FunctionInfo{ID: 479, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetShaderiv = builder.FunctionInfo{ID: 480, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformBlockIndex = builder.FunctionInfo{ID: 481, ReturnType: protocol.TypeUint32, Parameters: 2}
var funcInfoGlGetUniformIndices = builder.FunctionInfo{ID: 482, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetUniformLocation = builder.FunctionInfo{ID: 483, ReturnType: protocol.TypeInt32, Parameters: 2}
var funcInfoGlGetUniformfv = builder.FunctionInfo{ID: 484, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformiv = builder.FunctionInfo{ID: 485, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetUniformuiv = builder.FunctionInfo{ID: 486, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlIsProgram = builder.FunctionInfo{ID: 487, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsProgramPipeline = builder.FunctionInfo{ID: 488, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsShader = builder.FunctionInfo{ID: 489, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlLinkProgram = builder.FunctionInfo{ID: 490, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlMemoryBarrier = builder.FunctionInfo{ID: 491, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlMemoryBarrierByRegion = builder.FunctionInfo{ID: 492, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlProgramBinary = builder.FunctionInfo{ID: 493, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramParameteri = builder.FunctionInfo{ID: 494, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniform1f = builder.FunctionInfo{ID: 495, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniform1fv = builder.FunctionInfo{ID: 496, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform1i = builder.FunctionInfo{ID: 497, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniform1iv = builder.FunctionInfo{ID: 498, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform1ui = builder.FunctionInfo{ID: 499, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlProgramUniform1uiv = builder.FunctionInfo{ID: 500, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2f = builder.FunctionInfo{ID: 501, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2fv = builder.FunctionInfo{ID: 502, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2i = builder.FunctionInfo{ID: 503, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2iv = builder.FunctionInfo{ID: 504, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2ui = builder.FunctionInfo{ID: 505, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform2uiv = builder.FunctionInfo{ID: 506, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform3f = builder.FunctionInfo{ID: 507, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniform3fv = builder.FunctionInfo{ID: 508, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform3i = builder.FunctionInfo{ID: 509, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniform3iv = builder.FunctionInfo{ID: 510, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform3ui = builder.FunctionInfo{ID: 511, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniform3uiv = builder.FunctionInfo{ID: 512, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform4f = builder.FunctionInfo{ID: 513, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlProgramUniform4fv = builder.FunctionInfo{ID: 514, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform4i = builder.FunctionInfo{ID: 515, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlProgramUniform4iv = builder.FunctionInfo{ID: 516, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniform4ui = builder.FunctionInfo{ID: 517, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlProgramUniform4uiv = builder.FunctionInfo{ID: 518, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlProgramUniformMatrix2fv = builder.FunctionInfo{ID: 519, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix2x3fv = builder.FunctionInfo{ID: 520, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix2x4fv = builder.FunctionInfo{ID: 521, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix3fv = builder.FunctionInfo{ID: 522, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix3x2fv = builder.FunctionInfo{ID: 523, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix3x4fv = builder.FunctionInfo{ID: 524, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix4fv = builder.FunctionInfo{ID: 525, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix4x2fv = builder.FunctionInfo{ID: 526, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlProgramUniformMatrix4x3fv = builder.FunctionInfo{ID: 527, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlReleaseShaderCompiler = builder.FunctionInfo{ID: 528, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlShaderBinary = builder.FunctionInfo{ID: 529, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlShaderSource = builder.FunctionInfo{ID: 530, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform1f = builder.FunctionInfo{ID: 531, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform1fv = builder.FunctionInfo{ID: 532, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1i = builder.FunctionInfo{ID: 533, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform1iv = builder.FunctionInfo{ID: 534, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform1ui = builder.FunctionInfo{ID: 535, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlUniform1uiv = builder.FunctionInfo{ID: 536, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2f = builder.FunctionInfo{ID: 537, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2fv = builder.FunctionInfo{ID: 538, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2i = builder.FunctionInfo{ID: 539, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2iv = builder.FunctionInfo{ID: 540, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2ui = builder.FunctionInfo{ID: 541, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform2uiv = builder.FunctionInfo{ID: 542, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3f = builder.FunctionInfo{ID: 543, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform3fv = builder.FunctionInfo{ID: 544, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3i = builder.FunctionInfo{ID: 545, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform3iv = builder.FunctionInfo{ID: 546, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform3ui = builder.FunctionInfo{ID: 547, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniform3uiv = builder.FunctionInfo{ID: 548, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4f = builder.FunctionInfo{ID: 549, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform4fv = builder.FunctionInfo{ID: 550, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4i = builder.FunctionInfo{ID: 551, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform4iv = builder.FunctionInfo{ID: 552, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniform4ui = builder.FunctionInfo{ID: 553, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlUniform4uiv = builder.FunctionInfo{ID: 554, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniformBlockBinding = builder.FunctionInfo{ID: 555, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlUniformMatrix2fv = builder.FunctionInfo{ID: 556, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix2x3fv = builder.FunctionInfo{ID: 557, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix2x4fv = builder.FunctionInfo{ID: 558, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3fv = builder.FunctionInfo{ID: 559, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3x2fv = builder.FunctionInfo{ID: 560, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix3x4fv = builder.FunctionInfo{ID: 561, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4fv = builder.FunctionInfo{ID: 562, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4x2fv = builder.FunctionInfo{ID: 563, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUniformMatrix4x3fv = builder.FunctionInfo{ID: 564, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlUseProgram = builder.FunctionInfo{ID: 565, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlUseProgramStages = builder.FunctionInfo{ID: 566, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlValidateProgram = builder.FunctionInfo{ID: 567, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlValidateProgramPipeline = builder.FunctionInfo{ID: 568, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlCullFace = builder.FunctionInfo{ID: 569, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlDepthRangef = builder.FunctionInfo{ID: 570, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlFrontFace = builder.FunctionInfo{ID: 571, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetMultisamplefv = builder.FunctionInfo{ID: 572, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlLineWidth = builder.FunctionInfo{ID: 573, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlPolygonOffset = builder.FunctionInfo{ID: 574, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlViewport = builder.FunctionInfo{ID: 575, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetBooleani_v = builder.FunctionInfo{ID: 576, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetBooleanv = builder.FunctionInfo{ID: 577, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetFloatv = builder.FunctionInfo{ID: 578, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetInteger64i_v = builder.FunctionInfo{ID: 579, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetInteger64v = builder.FunctionInfo{ID: 580, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetIntegeri_v = builder.FunctionInfo{ID: 581, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetIntegerv = builder.FunctionInfo{ID: 582, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetInternalformativ = builder.FunctionInfo{ID: 583, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlGetString = builder.FunctionInfo{ID: 584, ReturnType: protocol.TypeVolatilePointer, Parameters: 1}
var funcInfoGlGetStringi = builder.FunctionInfo{ID: 585, ReturnType: protocol.TypeVolatilePointer, Parameters: 2}
var funcInfoGlIsEnabled = builder.FunctionInfo{ID: 586, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlClientWaitSync = builder.FunctionInfo{ID: 587, ReturnType: protocol.TypeUint32, Parameters: 3}
var funcInfoGlDeleteSync = builder.FunctionInfo{ID: 588, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlFenceSync = builder.FunctionInfo{ID: 589, ReturnType: protocol.TypeUint64, Parameters: 2}
var funcInfoGlGetSynciv = builder.FunctionInfo{ID: 590, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlIsSync = builder.FunctionInfo{ID: 591, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlWaitSync = builder.FunctionInfo{ID: 592, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlActiveTexture = builder.FunctionInfo{ID: 593, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBindImageTexture = builder.FunctionInfo{ID: 594, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlBindSampler = builder.FunctionInfo{ID: 595, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlBindTexture = builder.FunctionInfo{ID: 596, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlCompressedTexImage2D = builder.FunctionInfo{ID: 597, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCompressedTexImage3D = builder.FunctionInfo{ID: 598, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlCompressedTexSubImage2D = builder.FunctionInfo{ID: 599, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlCompressedTexSubImage3D = builder.FunctionInfo{ID: 600, ReturnType: protocol.TypeVoid, Parameters: 11}
var funcInfoGlCopyTexImage2D = builder.FunctionInfo{ID: 601, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCopyTexSubImage2D = builder.FunctionInfo{ID: 602, ReturnType: protocol.TypeVoid, Parameters: 8}
var funcInfoGlCopyTexSubImage3D = builder.FunctionInfo{ID: 603, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlDeleteSamplers = builder.FunctionInfo{ID: 604, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteTextures = builder.FunctionInfo{ID: 605, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenSamplers = builder.FunctionInfo{ID: 606, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenTextures = builder.FunctionInfo{ID: 607, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGenerateMipmap = builder.FunctionInfo{ID: 608, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGetSamplerParameterfv = builder.FunctionInfo{ID: 609, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetSamplerParameteriv = builder.FunctionInfo{ID: 610, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexLevelParameterfv = builder.FunctionInfo{ID: 611, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetTexLevelParameteriv = builder.FunctionInfo{ID: 612, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlGetTexParameterfv = builder.FunctionInfo{ID: 613, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetTexParameteriv = builder.FunctionInfo{ID: 614, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlIsSampler = builder.FunctionInfo{ID: 615, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlIsTexture = builder.FunctionInfo{ID: 616, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlPixelStorei = builder.FunctionInfo{ID: 617, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlSamplerParameterf = builder.FunctionInfo{ID: 618, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlSamplerParameterfv = builder.FunctionInfo{ID: 619, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlSamplerParameteri = builder.FunctionInfo{ID: 620, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlSamplerParameteriv = builder.FunctionInfo{ID: 621, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexImage2D = builder.FunctionInfo{ID: 622, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlTexImage3D = builder.FunctionInfo{ID: 623, ReturnType: protocol.TypeVoid, Parameters: 10}
var funcInfoGlTexParameterf = builder.FunctionInfo{ID: 624, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameterfv = builder.FunctionInfo{ID: 625, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameteri = builder.FunctionInfo{ID: 626, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexParameteriv = builder.FunctionInfo{ID: 627, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlTexStorage2D = builder.FunctionInfo{ID: 628, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlTexStorage2DMultisample = builder.FunctionInfo{ID: 629, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTexStorage3D = builder.FunctionInfo{ID: 630, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlTexSubImage2D = builder.FunctionInfo{ID: 631, ReturnType: protocol.TypeVoid, Parameters: 9}
var funcInfoGlTexSubImage3D = builder.FunctionInfo{ID: 632, ReturnType: protocol.TypeVoid, Parameters: 11}
var funcInfoGlBeginTransformFeedback = builder.FunctionInfo{ID: 633, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBindTransformFeedback = builder.FunctionInfo{ID: 634, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDeleteTransformFeedbacks = builder.FunctionInfo{ID: 635, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlEndTransformFeedback = builder.FunctionInfo{ID: 636, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlGenTransformFeedbacks = builder.FunctionInfo{ID: 637, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetTransformFeedbackVarying = builder.FunctionInfo{ID: 638, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoGlIsTransformFeedback = builder.FunctionInfo{ID: 639, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlPauseTransformFeedback = builder.FunctionInfo{ID: 640, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlResumeTransformFeedback = builder.FunctionInfo{ID: 641, ReturnType: protocol.TypeVoid, Parameters: 0}
var funcInfoGlTransformFeedbackVaryings = builder.FunctionInfo{ID: 642, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlBindVertexArray = builder.FunctionInfo{ID: 643, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlBindVertexBuffer = builder.FunctionInfo{ID: 644, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlDeleteVertexArrays = builder.FunctionInfo{ID: 645, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlDisableVertexAttribArray = builder.FunctionInfo{ID: 646, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlEnableVertexAttribArray = builder.FunctionInfo{ID: 647, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoGlGenVertexArrays = builder.FunctionInfo{ID: 648, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlGetVertexAttribIiv = builder.FunctionInfo{ID: 649, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetVertexAttribIuiv = builder.FunctionInfo{ID: 650, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetVertexAttribPointerv = builder.FunctionInfo{ID: 651, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetVertexAttribfv = builder.FunctionInfo{ID: 652, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetVertexAttribiv = builder.FunctionInfo{ID: 653, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlIsVertexArray = builder.FunctionInfo{ID: 654, ReturnType: protocol.TypeBool, Parameters: 1}
var funcInfoGlVertexAttrib1f = builder.FunctionInfo{ID: 655, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib1fv = builder.FunctionInfo{ID: 656, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib2f = builder.FunctionInfo{ID: 657, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlVertexAttrib2fv = builder.FunctionInfo{ID: 658, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib3f = builder.FunctionInfo{ID: 659, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlVertexAttrib3fv = builder.FunctionInfo{ID: 660, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttrib4f = builder.FunctionInfo{ID: 661, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttrib4fv = builder.FunctionInfo{ID: 662, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttribBinding = builder.FunctionInfo{ID: 663, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttribDivisor = builder.FunctionInfo{ID: 664, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttribFormat = builder.FunctionInfo{ID: 665, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttribI4i = builder.FunctionInfo{ID: 666, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttribI4iv = builder.FunctionInfo{ID: 667, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttribI4ui = builder.FunctionInfo{ID: 668, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttribI4uiv = builder.FunctionInfo{ID: 669, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlVertexAttribIFormat = builder.FunctionInfo{ID: 670, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoGlVertexAttribIPointer = builder.FunctionInfo{ID: 671, ReturnType: protocol.TypeVoid, Parameters: 5}
var funcInfoGlVertexAttribPointer = builder.FunctionInfo{ID: 672, ReturnType: protocol.TypeVoid, Parameters: 6}
var funcInfoGlVertexBindingDivisor = builder.FunctionInfo{ID: 673, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoEglInitialize = builder.FunctionInfo{ID: 674, ReturnType: protocol.TypeInt64, Parameters: 3}
var funcInfoEglCreateContext = builder.FunctionInfo{ID: 675, ReturnType: protocol.TypeVolatilePointer, Parameters: 4}
var funcInfoEglMakeCurrent = builder.FunctionInfo{ID: 676, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoEglSwapBuffers = builder.FunctionInfo{ID: 677, ReturnType: protocol.TypeInt64, Parameters: 2}
var funcInfoEglQuerySurface = builder.FunctionInfo{ID: 678, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoGlXCreateContext = builder.FunctionInfo{ID: 679, ReturnType: protocol.TypeVolatilePointer, Parameters: 4}
var funcInfoGlXCreateNewContext = builder.FunctionInfo{ID: 680, ReturnType: protocol.TypeVolatilePointer, Parameters: 5}
var funcInfoGlXMakeContextCurrent = builder.FunctionInfo{ID: 681, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoGlXMakeCurrent = builder.FunctionInfo{ID: 682, ReturnType: protocol.TypeInt64, Parameters: 3}
var funcInfoGlXSwapBuffers = builder.FunctionInfo{ID: 683, ReturnType: protocol.TypeVoid, Parameters: 2}
var funcInfoGlXQueryDrawable = builder.FunctionInfo{ID: 684, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoWglCreateContext = builder.FunctionInfo{ID: 685, ReturnType: protocol.TypeVolatilePointer, Parameters: 1}
var funcInfoWglCreateContextAttribsARB = builder.FunctionInfo{ID: 686, ReturnType: protocol.TypeVolatilePointer, Parameters: 3}
var funcInfoWglMakeCurrent = builder.FunctionInfo{ID: 687, ReturnType: protocol.TypeInt64, Parameters: 2}
var funcInfoWglSwapBuffers = builder.FunctionInfo{ID: 688, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoCGLCreateContext = builder.FunctionInfo{ID: 689, ReturnType: protocol.TypeInt64, Parameters: 3}
var funcInfoCGLSetCurrentContext = builder.FunctionInfo{ID: 690, ReturnType: protocol.TypeInt64, Parameters: 1}
var funcInfoCGLGetSurface = builder.FunctionInfo{ID: 691, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoCGSGetSurfaceBounds = builder.FunctionInfo{ID: 692, ReturnType: protocol.TypeInt64, Parameters: 4}
var funcInfoCGLFlushDrawable = builder.FunctionInfo{ID: 693, ReturnType: protocol.TypeInt64, Parameters: 1}
var funcInfoGlGetQueryObjecti64v = builder.FunctionInfo{ID: 694, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoGlGetQueryObjectui64v = builder.FunctionInfo{ID: 695, ReturnType: protocol.TypeVoid, Parameters: 3}
var funcInfoArchitecture = builder.FunctionInfo{ID: 696, ReturnType: protocol.TypeVoid, Parameters: 4}
var funcInfoReplayCreateRenderer = builder.FunctionInfo{ID: 697, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoReplayBindRenderer = builder.FunctionInfo{ID: 698, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoSwitchThread = builder.FunctionInfo{ID: 699, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoBackbufferInfo = builder.FunctionInfo{ID: 700, ReturnType: protocol.TypeVoid, Parameters: 7}
var funcInfoStartTimer = builder.FunctionInfo{ID: 701, ReturnType: protocol.TypeVoid, Parameters: 1}
var funcInfoStopTimer = builder.FunctionInfo{ID: 702, ReturnType: protocol.TypeUint64, Parameters: 1}
var funcInfoFlushPostBuffer = builder.FunctionInfo{ID: 703, ReturnType: protocol.TypeVoid, Parameters: 0}

func (c RenderbufferId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c TextureId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c FramebufferId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c BufferId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ShaderId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ProgramId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c VertexArrayId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c QueryId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c UniformLocation) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c SamplerId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c PipelineId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c UniformBlockId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c TransformFeedbackId) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c AttributeLocation) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ContextID) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c ThreadID) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U64(uint64(c))
}
func (c EGLConfig) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c EGLContext) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c EGLDisplay) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c EGLSurface) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c GLXContext) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c GLXDrawable) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c HGLRC) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c HDC) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c CGLPixelFormatObj) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c CGLContextObj) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c CGSConnectionID) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c CGSWindowID) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c CGSSurfaceID) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c GLboolean) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U8(uint8(c))
}
func (c GLbyte) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S8(int8(c))
}
func (c GLubyte) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U8(uint8(c))
}
func (c GLshort) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S16(int16(c))
}
func (c GLushort) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U16(uint16(c))
}
func (c GLint) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c GLuint) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U32(uint32(c))
}
func (c GLint64) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S64(int64(c))
}
func (c GLuint64) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U64(uint64(c))
}
func (c GLfixed) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c GLsizei) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c GLintptr) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c GLsizeiptr) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.S32(int32(c))
}
func (c GLDEBUGPROCKHR) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return Voidᵖ(c).value()
}
func (c GLhalf) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U16(uint16(c))
}
func (c GLfloat) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.F32(float32(c))
}
func (c GLclampf) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.F32(float32(c))
}
func (c GLsync) value(ϟb *builder.Builder, ϟa atom.Atom, ϟs *gfxapi.State) value.Value {
	return value.U64(uint64(c))
}

var _ = replay.Replayer(&GlBlendBarrierKHR{}) // interface compliance check
func (ϟa *GlBlendBarrierKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_0_ext := ExtensionId_GL_KHR_blend_equation_advanced // ExtensionId
	ϟb.Call(funcInfoGlBlendBarrierKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_0_ext
	return nil
}

var _ = replay.Replayer(&GlBlendEquationSeparateiEXT{}) // interface compliance check
func (ϟa *GlBlendEquationSeparateiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_1_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.ModeRGB))
	ϟb.Push(value.U32(ϟa.ModeAlpha))
	ϟb.Call(funcInfoGlBlendEquationSeparateiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_1_ext
	return nil
}

var _ = replay.Replayer(&GlBlendEquationiEXT{}) // interface compliance check
func (ϟa *GlBlendEquationiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_2_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlBlendEquationiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_2_ext
	return nil
}

var _ = replay.Replayer(&GlBlendFuncSeparateiEXT{}) // interface compliance check
func (ϟa *GlBlendFuncSeparateiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_3_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.SrcRGB))
	ϟb.Push(value.U32(ϟa.DstRGB))
	ϟb.Push(value.U32(ϟa.SrcAlpha))
	ϟb.Push(value.U32(ϟa.DstAlpha))
	ϟb.Call(funcInfoGlBlendFuncSeparateiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_3_ext
	return nil
}

var _ = replay.Replayer(&GlBlendFunciEXT{}) // interface compliance check
func (ϟa *GlBlendFunciEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_4_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Src))
	ϟb.Push(value.U32(ϟa.Dst))
	ϟb.Call(funcInfoGlBlendFunciEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_4_ext
	return nil
}

var _ = replay.Replayer(&GlColorMaskiEXT{}) // interface compliance check
func (ϟa *GlColorMaskiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_5_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.R.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.G.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.B.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.A.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlColorMaskiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_5_ext
	return nil
}

var _ = replay.Replayer(&GlCopyImageSubDataEXT{}) // interface compliance check
func (ϟa *GlCopyImageSubDataEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_6_ext := ExtensionId_GL_EXT_copy_image // ExtensionId
	ϟb.Push(ϟa.SrcName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.SrcTarget))
	ϟb.Push(ϟa.SrcLevel.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.DstTarget))
	ϟb.Push(ϟa.DstLevel.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcWidth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcHeight.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcDepth.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyImageSubDataEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_6_ext
	return nil
}

var _ = replay.Replayer(&GlDebugMessageCallbackKHR{}) // interface compliance check
func (ϟa *GlDebugMessageCallbackKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_7_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(ϟa.Callback.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.UserParam.value())
	ϟb.Call(funcInfoGlDebugMessageCallbackKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_7_ext
	return nil
}

var _ = replay.Replayer(&GlDebugMessageControlKHR{}) // interface compliance check
func (ϟa *GlDebugMessageControlKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_8_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(value.U32(ϟa.Source))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(value.U32(ϟa.Severity))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Ids.value())
	ϟb.Push(ϟa.Enabled.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDebugMessageControlKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_8_ext
	return nil
}

var _ = replay.Replayer(&GlDebugMessageInsertKHR{}) // interface compliance check
func (ϟa *GlDebugMessageInsertKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_9_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(value.U32(ϟa.Source))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Id.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Severity))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Buf.value())
	ϟb.Call(funcInfoGlDebugMessageInsertKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_9_ext
	return nil
}

var _ = replay.Replayer(&GlDisableiEXT{}) // interface compliance check
func (ϟa *GlDisableiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_10_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDisableiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_10_ext
	return nil
}

var _ = replay.Replayer(&GlEnableiEXT{}) // interface compliance check
func (ϟa *GlEnableiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_11_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEnableiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_11_ext
	return nil
}

var _ = replay.Replayer(&GlFramebufferTextureEXT{}) // interface compliance check
func (ϟa *GlFramebufferTextureEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_12_ext := ExtensionId_GL_EXT_geometry_shader // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Attachment))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTextureEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_12_ext
	return nil
}

var _ = replay.Replayer(&GlGetDebugMessageLogKHR{}) // interface compliance check
func (ϟa *GlGetDebugMessageLogKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_13_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Sources.value())
	ϟb.Push(ϟa.Types.value())
	ϟb.Push(ϟa.Ids.value())
	ϟb.Push(ϟa.Severities.value())
	ϟb.Push(ϟa.Lengths.value())
	ϟb.Push(ϟa.MessageLog.value())
	ϟb.Call(funcInfoGlGetDebugMessageLogKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_13_ext
	return nil
}

var _ = replay.Replayer(&GlGetObjectLabelKHR{}) // interface compliance check
func (ϟa *GlGetObjectLabelKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_14_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(value.U32(ϟa.Identifier))
	ϟb.Push(ϟa.Name.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Label.value())
	ϟb.Call(funcInfoGlGetObjectLabelKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_14_ext
	return nil
}

var _ = replay.Replayer(&GlGetObjectPtrLabelKHR{}) // interface compliance check
func (ϟa *GlGetObjectPtrLabelKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_15_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(ϟa.Ptr.value())
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Label.value())
	ϟb.Call(funcInfoGlGetObjectPtrLabelKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_15_ext
	return nil
}

var _ = replay.Replayer(&GlGetPointervKHR{}) // interface compliance check
func (ϟa *GlGetPointervKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_16_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetPointervKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_16_ext
	return nil
}

var _ = replay.Replayer(&GlGetSamplerParameterIivEXT{}) // interface compliance check
func (ϟa *GlGetSamplerParameterIivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_17_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetSamplerParameterIivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_17_ext
	return nil
}

var _ = replay.Replayer(&GlGetSamplerParameterIuivEXT{}) // interface compliance check
func (ϟa *GlGetSamplerParameterIuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_18_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetSamplerParameterIuivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_18_ext
	return nil
}

var _ = replay.Replayer(&GlGetTexParameterIivEXT{}) // interface compliance check
func (ϟa *GlGetTexParameterIivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_19_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetTexParameterIivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_19_ext
	return nil
}

var _ = replay.Replayer(&GlGetTexParameterIuivEXT{}) // interface compliance check
func (ϟa *GlGetTexParameterIuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_20_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetTexParameterIuivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_20_ext
	return nil
}

var _ = replay.Replayer(&GlIsEnablediEXT{}) // interface compliance check
func (ϟa *GlIsEnablediEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_21_ext := ExtensionId_GL_EXT_draw_buffers_indexed // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsEnablediEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_21_ext
	return nil
}

var _ = replay.Replayer(&GlMinSampleShadingOES{}) // interface compliance check
func (ϟa *GlMinSampleShadingOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_22_ext := ExtensionId_GL_OES_sample_shading // ExtensionId
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMinSampleShadingOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_22_ext
	return nil
}

var _ = replay.Replayer(&GlObjectLabelKHR{}) // interface compliance check
func (ϟa *GlObjectLabelKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_23_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(value.U32(ϟa.Identifier))
	ϟb.Push(ϟa.Name.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Label.value())
	ϟb.Call(funcInfoGlObjectLabelKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_23_ext
	return nil
}

var _ = replay.Replayer(&GlObjectPtrLabelKHR{}) // interface compliance check
func (ϟa *GlObjectPtrLabelKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_24_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(ϟa.Ptr.value())
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Label.value())
	ϟb.Call(funcInfoGlObjectPtrLabelKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_24_ext
	return nil
}

var _ = replay.Replayer(&GlPatchParameteriEXT{}) // interface compliance check
func (ϟa *GlPatchParameteriEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_25_ext := ExtensionId_GL_EXT_tessellation_shader // ExtensionId
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPatchParameteriEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_25_ext
	return nil
}

var _ = replay.Replayer(&GlPopDebugGroupKHR{}) // interface compliance check
func (ϟa *GlPopDebugGroupKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_26_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Call(funcInfoGlPopDebugGroupKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_26_ext
	return nil
}

var _ = replay.Replayer(&GlPrimitiveBoundingBoxEXT{}) // interface compliance check
func (ϟa *GlPrimitiveBoundingBoxEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_27_ext := ExtensionId_GL_EXT_primitive_bounding_box // ExtensionId
	ϟb.Push(ϟa.MinX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MinY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MinZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MinW.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxW.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPrimitiveBoundingBoxEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_27_ext
	return nil
}

var _ = replay.Replayer(&GlPushDebugGroupKHR{}) // interface compliance check
func (ϟa *GlPushDebugGroupKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_28_ext := ExtensionId_GL_KHR_debug // ExtensionId
	ϟb.Push(value.U32(ϟa.Source))
	ϟb.Push(ϟa.Id.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Message.value())
	ϟb.Call(funcInfoGlPushDebugGroupKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_28_ext
	return nil
}

var _ = replay.Replayer(&GlSamplerParameterIivEXT{}) // interface compliance check
func (ϟa *GlSamplerParameterIivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_29_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value())
	ϟb.Call(funcInfoGlSamplerParameterIivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_29_ext
	return nil
}

var _ = replay.Replayer(&GlSamplerParameterIuivEXT{}) // interface compliance check
func (ϟa *GlSamplerParameterIuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_30_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value())
	ϟb.Call(funcInfoGlSamplerParameterIuivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_30_ext
	return nil
}

var _ = replay.Replayer(&GlTexBufferEXT{}) // interface compliance check
func (ϟa *GlTexBufferEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_31_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Internalformat))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlTexBufferEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_31_ext
	return nil
}

var _ = replay.Replayer(&GlTexBufferRangeEXT{}) // interface compliance check
func (ϟa *GlTexBufferRangeEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_32_ext := ExtensionId_GL_EXT_texture_buffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Internalformat))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexBufferRangeEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_32_ext
	return nil
}

var _ = replay.Replayer(&GlTexParameterIivEXT{}) // interface compliance check
func (ϟa *GlTexParameterIivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_33_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlTexParameterIivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_33_ext
	return nil
}

var _ = replay.Replayer(&GlTexParameterIuivEXT{}) // interface compliance check
func (ϟa *GlTexParameterIuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_34_ext := ExtensionId_GL_EXT_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlTexParameterIuivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_34_ext
	return nil
}

var _ = replay.Replayer(&GlTexStorage3DMultisampleOES{}) // interface compliance check
func (ϟa *GlTexStorage3DMultisampleOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_35_ext := ExtensionId_GL_OES_texture_storage_multisample_2d_array // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Fixedsamplelocations.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexStorage3DMultisampleOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_35_ext
	return nil
}

var _ = replay.Replayer(&GlBeginQuery{}) // interface compliance check
func (ϟa *GlBeginQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_36_major := uint32(3) // u32
	minRequiredVersion_36_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBeginQuery)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_36_major, minRequiredVersion_36_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteQueries{}) // interface compliance check
func (ϟa *GlDeleteQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_38_major := uint32(3)                               // u32
	minRequiredVersion_38_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_39_result := context                                        // Contextʳ
	ctx := GetContext_39_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlDeleteQueries)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_38_major, minRequiredVersion_38_minor, q, context, GetContext_39_result, ctx
	return nil
}

var _ = replay.Replayer(&GlEndQuery{}) // interface compliance check
func (ϟa *GlEndQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_40_major := uint32(3) // u32
	minRequiredVersion_40_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlEndQuery)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_40_major, minRequiredVersion_40_minor
	return nil
}

var _ = replay.Replayer(&GlGenQueries{}) // interface compliance check
func (ϟa *GlGenQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_42_major := uint32(3)                               // u32
	minRequiredVersion_42_minor := uint32(0)                               // u32
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_43_result := context                                        // Contextʳ
	ctx := GetContext_43_result                                            // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlGenQueries)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_42_major, minRequiredVersion_42_minor, q, context, GetContext_43_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectuiv{}) // interface compliance check
func (ϟa *GlGetQueryObjectuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_44_major := uint32(3) // u32
	minRequiredVersion_44_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_QUERY_RESULT, GLenum_GL_QUERY_RESULT_AVAILABLE:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectuiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_44_major, minRequiredVersion_44_minor
	return nil
}

var _ = replay.Replayer(&GlGetQueryiv{}) // interface compliance check
func (ϟa *GlGetQueryiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_46_major := uint32(3) // u32
	minRequiredVersion_46_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ANY_SAMPLES_PASSED, GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE, GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_CURRENT_QUERY:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_46_major, minRequiredVersion_46_minor
	return nil
}

var _ = replay.Replayer(&GlIsQuery{}) // interface compliance check
func (ϟa *GlIsQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_49_major := uint32(3)     // u32
	minRequiredVersion_49_minor := uint32(0)     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_50_result := context              // Contextʳ
	ctx := GetContext_50_result                  // Contextʳ
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsQuery)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_49_major, minRequiredVersion_49_minor, context, GetContext_50_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindBuffer{}) // interface compliance check
func (ϟa *GlBindBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_51_major := uint32(2) // u32
	minRequiredVersion_51_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_52_major := uint32(3) // u32
		minRequiredVersion_52_minor := uint32(0) // u32
		_, _ = minRequiredVersion_52_major, minRequiredVersion_52_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_53_major := uint32(3) // u32
		minRequiredVersion_53_minor := uint32(1) // u32
		_, _ = minRequiredVersion_53_major, minRequiredVersion_53_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_55_result := context              // Contextʳ
	ctx := GetContext_55_result                  // Contextʳ
	if !(ctx.Instances.Buffers.Contains(ϟa.Buffer)) {
		ctx.Instances.Buffers[ϟa.Buffer] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundBuffers[ϟa.Target] = ϟa.Buffer
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_51_major, minRequiredVersion_51_minor, context, GetContext_55_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindBufferBase{}) // interface compliance check
func (ϟa *GlBindBufferBase) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_56_major := uint32(3) // u32
	minRequiredVersion_56_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_57_major := uint32(3) // u32
		minRequiredVersion_57_minor := uint32(1) // u32
		_, _ = minRequiredVersion_57_major, minRequiredVersion_57_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindBufferBase)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_56_major, minRequiredVersion_56_minor
	return nil
}

var _ = replay.Replayer(&GlBindBufferRange{}) // interface compliance check
func (ϟa *GlBindBufferRange) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_59_major := uint32(3) // u32
	minRequiredVersion_59_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_60_major := uint32(3) // u32
		minRequiredVersion_60_minor := uint32(1) // u32
		_, _ = minRequiredVersion_60_major, minRequiredVersion_60_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBindBufferRange)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_59_major, minRequiredVersion_59_minor
	return nil
}

var _ = replay.Replayer(&GlBufferData{}) // interface compliance check
func (ϟa *GlBufferData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_62_major := uint32(2) // u32
	minRequiredVersion_62_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_63_major := uint32(3) // u32
		minRequiredVersion_63_minor := uint32(0) // u32
		_, _ = minRequiredVersion_63_major, minRequiredVersion_63_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_64_major := uint32(3) // u32
		minRequiredVersion_64_minor := uint32(1) // u32
		_, _ = minRequiredVersion_64_major, minRequiredVersion_64_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Usage {
	case GLenum_GL_DYNAMIC_DRAW, GLenum_GL_STATIC_DRAW, GLenum_GL_STREAM_DRAW:
	case GLenum_GL_DYNAMIC_COPY, GLenum_GL_DYNAMIC_READ, GLenum_GL_STATIC_COPY, GLenum_GL_STATIC_READ, GLenum_GL_STREAM_COPY, GLenum_GL_STREAM_READ:
		minRequiredVersion_66_major := uint32(3) // u32
		minRequiredVersion_66_minor := uint32(0) // u32
		_, _ = minRequiredVersion_66_major, minRequiredVersion_66_minor
	default:
		v := ϟa.Usage
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_68_result := context              // Contextʳ
	ctx := GetContext_68_result                  // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // Bufferʳ
	b.Data = func() (result U8ˢ) {
		switch (ϟa.Data) != (BufferDataPointer(Voidᶜᵖ{})) {
		case true:
			return U8ᵖ(ϟa.Data).Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
		case false:
			return MakeU8ˢ(uint64(ϟa.Size), ϟs)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (ϟa.Data) != (BufferDataPointer(Voidᶜᵖ{})), ϟa))
			return result
		}
	}()
	b.Size = ϟa.Size
	b.Usage = ϟa.Usage
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Usage))
	ϟb.Call(funcInfoGlBufferData)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_62_major, minRequiredVersion_62_minor, context, GetContext_68_result, ctx, id, b
	return nil
}

var _ = replay.Replayer(&GlBufferSubData{}) // interface compliance check
func (ϟa *GlBufferSubData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_69_major := uint32(2) // u32
	minRequiredVersion_69_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_70_major := uint32(3) // u32
		minRequiredVersion_70_minor := uint32(0) // u32
		_, _ = minRequiredVersion_70_major, minRequiredVersion_70_minor
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_71_major := uint32(3) // u32
		minRequiredVersion_71_minor := uint32(1) // u32
		_, _ = minRequiredVersion_71_major, minRequiredVersion_71_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.Data.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Size), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBufferSubData)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_69_major, minRequiredVersion_69_minor
	return nil
}

var _ = replay.Replayer(&GlCopyBufferSubData{}) // interface compliance check
func (ϟa *GlCopyBufferSubData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_73_major := uint32(3) // u32
	minRequiredVersion_73_minor := uint32(0) // u32
	switch ϟa.ReadTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.ReadTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.WriteTarget {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.WriteTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.ReadTarget))
	ϟb.Push(value.U32(ϟa.WriteTarget))
	ϟb.Push(ϟa.ReadOffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.WriteOffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyBufferSubData)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_73_major, minRequiredVersion_73_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteBuffers{}) // interface compliance check
func (ϟa *GlDeleteBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_76_major := uint32(2)                               // u32
	minRequiredVersion_76_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_77_result := context                                        // Contextʳ
	ctx := GetContext_77_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Buffers, b.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Buffers.value())
	ϟb.Call(funcInfoGlDeleteBuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_76_major, minRequiredVersion_76_minor, b, context, GetContext_77_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenBuffers{}) // interface compliance check
func (ϟa *GlGenBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_78_major := uint32(2)                               // u32
	minRequiredVersion_78_minor := uint32(0)                               // u32
	b := ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // BufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                           // Contextʳ
	GetContext_79_result := context                                        // Contextʳ
	ctx := GetContext_79_result                                            // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Buffers.value())
	ϟb.Call(funcInfoGlGenBuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := BufferId(ϟa.Buffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // BufferId
		ctx.Instances.Buffers[id] = func() *Buffer {
			s := &Buffer{}
			s.Init()
			return s
		}()
		b.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_78_major, minRequiredVersion_78_minor, b, context, GetContext_79_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetBufferParameteri64v{}) // interface compliance check
func (ϟa *GlGetBufferParameteri64v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_80_major := uint32(3) // u32
	minRequiredVersion_80_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET, GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetBufferParameteri64v)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_80_major, minRequiredVersion_80_minor
	return nil
}

var _ = replay.Replayer(&GlGetBufferParameteriv{}) // interface compliance check
func (ϟa *GlGetBufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_83_major := uint32(2) // u32
	minRequiredVersion_83_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER:
	case GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
		minRequiredVersion_84_major := uint32(3) // u32
		minRequiredVersion_84_minor := uint32(0) // u32
		_, _ = minRequiredVersion_84_major, minRequiredVersion_84_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_BUFFER_SIZE, GLenum_GL_BUFFER_USAGE:
	case GLenum_GL_BUFFER_ACCESS_FLAGS, GLenum_GL_BUFFER_MAPPED, GLenum_GL_BUFFER_MAP_LENGTH, GLenum_GL_BUFFER_MAP_OFFSET:
		minRequiredVersion_86_major := uint32(3) // u32
		minRequiredVersion_86_minor := uint32(0) // u32
		_, _ = minRequiredVersion_86_major, minRequiredVersion_86_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_88_result := context              // Contextʳ
	ctx := GetContext_88_result                  // Contextʳ
	id := ctx.BoundBuffers.Get(ϟa.Target)        // BufferId
	b := ctx.Instances.Buffers.Get(id)           // Bufferʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetBufferParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_BUFFER_SIZE:
			return GLint(b.Size)
		case GLenum_GL_BUFFER_USAGE:
			return GLint(b.Usage)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_83_major, minRequiredVersion_83_minor, context, GetContext_88_result, ctx, id, b
	return nil
}

var _ = replay.Replayer(&GlGetBufferPointerv{}) // interface compliance check
func (ϟa *GlGetBufferPointerv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_89_major := uint32(3) // u32
	minRequiredVersion_89_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_BUFFER_MAP_POINTER:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetBufferPointerv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_89_major, minRequiredVersion_89_minor
	return nil
}

var _ = replay.Replayer(&GlIsBuffer{}) // interface compliance check
func (ϟa *GlIsBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_92_major := uint32(2)     // u32
	minRequiredVersion_92_minor := uint32(0)     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_93_result := context              // Contextʳ
	ctx := GetContext_93_result                  // Contextʳ
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_92_major, minRequiredVersion_92_minor, context, GetContext_93_result, ctx
	return nil
}

var _ = replay.Replayer(&GlMapBufferRange{}) // interface compliance check
func (ϟa *GlMapBufferRange) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_94_major := uint32(3) // u32
	minRequiredVersion_94_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_95_major := uint32(3) // u32
		minRequiredVersion_95_minor := uint32(1) // u32
		_, _ = minRequiredVersion_95_major, minRequiredVersion_95_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	supportsBits_97_seenBits := ϟa.Access                                                                                                                                                                                                                                      // GLbitfield
	supportsBits_97_validBits := (GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT) | ((GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT) | ((GLbitfield_GL_MAP_READ_BIT) | ((GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT) | (GLbitfield_GL_MAP_WRITE_BIT))))) // GLbitfield
	if (GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_READ_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT)&(ϟa.Access) != 0 {
	}
	if (GLbitfield_GL_MAP_WRITE_BIT)&(ϟa.Access) != 0 {
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                    // Contextʳ
	GetContext_98_result := context                                 // Contextʳ
	ctx := GetContext_98_result                                     // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Access))
	ϟb.Call(funcInfoGlMapBufferRange)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ptr := U8ᵖ(ϟa.Result) // U8ᵖ
	b.MappingAccess = ϟa.Access
	b.MappingData = ptr.Slice(uint64(GLsizeiptr(int32(0))), uint64(ϟa.Length), ϟs)
	if (GLbitfield_GL_MAP_READ_BIT)&(ϟa.Access) != 0 {
		src := b.Data.Slice(uint64(ϟa.Offset), uint64((ϟa.Offset)+(GLintptr(ϟa.Length))), ϟs) // U8ˢ
		dst := b.MappingData                                                                  // U8ˢ
		dst.OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = src, dst
	}
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_94_major, minRequiredVersion_94_minor, supportsBits_97_seenBits, supportsBits_97_validBits, context, GetContext_98_result, ctx, b, ptr
	return nil
}

var _ = replay.Replayer(&GlUnmapBuffer{}) // interface compliance check
func (ϟa *GlUnmapBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_99_major := uint32(3) // u32
	minRequiredVersion_99_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_DISPATCH_INDIRECT_BUFFER, GLenum_GL_DRAW_INDIRECT_BUFFER, GLenum_GL_SHADER_STORAGE_BUFFER:
		minRequiredVersion_100_major := uint32(3) // u32
		minRequiredVersion_100_minor := uint32(1) // u32
		_, _ = minRequiredVersion_100_major, minRequiredVersion_100_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                    // Contextʳ
	GetContext_102_result := context                                // Contextʳ
	ctx := GetContext_102_result                                    // Contextʳ
	b := ctx.Instances.Buffers.Get(ctx.BoundBuffers.Get(ϟa.Target)) // Bufferʳ
	ϟdst, ϟsrc := b.Data.Slice(uint64(b.MappingOffset), uint64((b.MappingOffset)+(int32(b.MappingData.Count))), ϟs).Copy(b.MappingData, ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlUnmapBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _ = minRequiredVersion_99_major, minRequiredVersion_99_minor, context, GetContext_102_result, ctx, b
	return nil
}

var _ = replay.Replayer(&GlDrawArrays{}) // interface compliance check
func (ϟa *GlDrawArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_103_major := uint32(2) // u32
	minRequiredVersion_103_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.DrawMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                    // Contextʳ
	GetContext_105_result := context                                                // Contextʳ
	ctx := GetContext_105_result                                                    // Contextʳ
	last_index := (uint32(ϟa.FirstIndex)) + ((uint32(ϟa.IndexCount)) - (uint32(1))) // u32
	ReadVertexArrays_106_ctx := ctx                                                 // Contextʳ
	ReadVertexArrays_106_first_index := uint32(ϟa.FirstIndex)                       // u32
	ReadVertexArrays_106_last_index := last_index                                   // u32
	for i := int32(int32(0)); i < int32(len(ReadVertexArrays_106_ctx.VertexAttributeArrays)); i++ {
		arr := ReadVertexArrays_106_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
		if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
			vertexAttribTypeSize_107_t := arr.Type // GLenum
			vertexAttribTypeSize_107_result := func() (result uint32) {
				switch vertexAttribTypeSize_107_t {
				case GLenum_GL_BYTE:
					return uint32(1)
				case GLenum_GL_UNSIGNED_BYTE:
					return uint32(1)
				case GLenum_GL_SHORT:
					return uint32(2)
				case GLenum_GL_UNSIGNED_SHORT:
					return uint32(2)
				case GLenum_GL_FIXED:
					return uint32(4)
				case GLenum_GL_FLOAT:
					return uint32(4)
				case GLenum_GL_HALF_FLOAT_ARB:
					return uint32(2)
				case GLenum_GL_HALF_FLOAT_OES:
					return uint32(2)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_107_t, ϟa))
					return result
				}
			}() // u32
			elsize := (vertexAttribTypeSize_107_result) * (arr.Size) // u32
			elstride := func() (result uint32) {
				switch (arr.Stride) == (GLsizei(int32(0))) {
				case true:
					return elsize
				case false:
					return uint32(arr.Stride)
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (GLsizei(int32(0))), ϟa))
					return result
				}
			}() // u32
			for v := uint32(ReadVertexArrays_106_first_index); v < (ReadVertexArrays_106_last_index)+(uint32(1)); v++ {
				offset := (elstride) * (v) // u32
				arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
				_ = offset
			}
			_, _, _, _ = vertexAttribTypeSize_107_t, vertexAttribTypeSize_107_result, elsize, elstride
		}
		_ = arr
	}
	ϟb.Push(value.U32(ϟa.DrawMode))
	ϟb.Push(ϟa.FirstIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.IndexCount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawArrays)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_103_major, minRequiredVersion_103_minor, context, GetContext_105_result, ctx, last_index, ReadVertexArrays_106_ctx, ReadVertexArrays_106_first_index, ReadVertexArrays_106_last_index
	return nil
}

var _ = replay.Replayer(&GlDrawArraysIndirect{}) // interface compliance check
func (ϟa *GlDrawArraysIndirect) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_108_major := uint32(3) // u32
	minRequiredVersion_108_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Indirect.value())
	ϟb.Call(funcInfoGlDrawArraysIndirect)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_108_major, minRequiredVersion_108_minor
	return nil
}

var _ = replay.Replayer(&GlDrawArraysInstanced{}) // interface compliance check
func (ϟa *GlDrawArraysInstanced) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_110_major := uint32(3) // u32
	minRequiredVersion_110_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.First.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Instancecount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawArraysInstanced)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_110_major, minRequiredVersion_110_minor
	return nil
}

var _ = replay.Replayer(&GlDrawBuffers{}) // interface compliance check
func (ϟa *GlDrawBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_112_major := uint32(3) // u32
	minRequiredVersion_112_minor := uint32(0) // u32
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Bufs.value())
	ϟb.Call(funcInfoGlDrawBuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_112_major, minRequiredVersion_112_minor
	return nil
}

var _ = replay.Replayer(&GlDrawElements{}) // interface compliance check
func (ϟa *GlDrawElements) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_113_major := uint32(2) // u32
	minRequiredVersion_113_minor := uint32(0) // u32
	switch ϟa.DrawMode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.DrawMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.IndicesType {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_UNSIGNED_INT:
		minRequiredVersion_115_major := uint32(3) // u32
		minRequiredVersion_115_minor := uint32(0) // u32
		_, _ = minRequiredVersion_115_major, minRequiredVersion_115_minor
	default:
		v := ϟa.IndicesType
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)               // Contextʳ
	GetContext_117_result := context                           // Contextʳ
	ctx := GetContext_117_result                               // Contextʳ
	count := uint32(ϟa.ElementCount)                           // u32
	id := ctx.BoundBuffers.Get(GLenum_GL_ELEMENT_ARRAY_BUFFER) // BufferId
	if (id) != (BufferId(uint32(0))) {
		index_data := ctx.Instances.Buffers.Get(id).Data                                                   // U8ˢ
		offset := uint32(uint64(ϟa.Indices.Address))                                                       // u32
		first := externs{ϟs, ϟd, ϟl}.minIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count) // u32
		last := externs{ϟs, ϟd, ϟl}.maxIndex(U8ᵖ(index_data.Index(0, ϟs)), ϟa.IndicesType, offset, count)  // u32
		ReadVertexArrays_118_ctx := ctx                                                                    // Contextʳ
		ReadVertexArrays_118_first_index := first                                                          // u32
		ReadVertexArrays_118_last_index := last                                                            // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_118_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_118_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_119_t := arr.Type // GLenum
				vertexAttribTypeSize_119_result := func() (result uint32) {
					switch vertexAttribTypeSize_119_t {
					case GLenum_GL_BYTE:
						return uint32(1)
					case GLenum_GL_UNSIGNED_BYTE:
						return uint32(1)
					case GLenum_GL_SHORT:
						return uint32(2)
					case GLenum_GL_UNSIGNED_SHORT:
						return uint32(2)
					case GLenum_GL_FIXED:
						return uint32(4)
					case GLenum_GL_FLOAT:
						return uint32(4)
					case GLenum_GL_HALF_FLOAT_ARB:
						return uint32(2)
					case GLenum_GL_HALF_FLOAT_OES:
						return uint32(2)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_119_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_119_result) * (arr.Size) // u32
				elstride := func() (result uint32) {
					switch (arr.Stride) == (GLsizei(int32(0))) {
					case true:
						return elsize
					case false:
						return uint32(arr.Stride)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (GLsizei(int32(0))), ϟa))
						return result
					}
				}() // u32
				for v := uint32(ReadVertexArrays_118_first_index); v < (ReadVertexArrays_118_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_119_t, vertexAttribTypeSize_119_result, elsize, elstride
			}
			_ = arr
		}
		_, _, _, _, _, _, _ = index_data, offset, first, last, ReadVertexArrays_118_ctx, ReadVertexArrays_118_first_index, ReadVertexArrays_118_last_index
	} else {
		index_data := U8ᵖ(ϟa.Indices)                                                       // U8ᵖ
		first := externs{ϟs, ϟd, ϟl}.minIndex(index_data, ϟa.IndicesType, uint32(0), count) // u32
		last := externs{ϟs, ϟd, ϟl}.maxIndex(index_data, ϟa.IndicesType, uint32(0), count)  // u32
		ReadVertexArrays_120_ctx := ctx                                                     // Contextʳ
		ReadVertexArrays_120_first_index := first                                           // u32
		ReadVertexArrays_120_last_index := last                                             // u32
		for i := int32(int32(0)); i < int32(len(ReadVertexArrays_120_ctx.VertexAttributeArrays)); i++ {
			arr := ReadVertexArrays_120_ctx.VertexAttributeArrays.Get(AttributeLocation(i)) // VertexAttributeArrayʳ
			if (arr.Enabled) && ((arr.Buffer) == (BufferId(uint32(0)))) {
				vertexAttribTypeSize_121_t := arr.Type // GLenum
				vertexAttribTypeSize_121_result := func() (result uint32) {
					switch vertexAttribTypeSize_121_t {
					case GLenum_GL_BYTE:
						return uint32(1)
					case GLenum_GL_UNSIGNED_BYTE:
						return uint32(1)
					case GLenum_GL_SHORT:
						return uint32(2)
					case GLenum_GL_UNSIGNED_SHORT:
						return uint32(2)
					case GLenum_GL_FIXED:
						return uint32(4)
					case GLenum_GL_FLOAT:
						return uint32(4)
					case GLenum_GL_HALF_FLOAT_ARB:
						return uint32(2)
					case GLenum_GL_HALF_FLOAT_OES:
						return uint32(2)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", vertexAttribTypeSize_121_t, ϟa))
						return result
					}
				}() // u32
				elsize := (vertexAttribTypeSize_121_result) * (arr.Size) // u32
				elstride := func() (result uint32) {
					switch (arr.Stride) == (GLsizei(int32(0))) {
					case true:
						return elsize
					case false:
						return uint32(arr.Stride)
					default:
						// TODO: better unmatched handling
						panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (arr.Stride) == (GLsizei(int32(0))), ϟa))
						return result
					}
				}() // u32
				for v := uint32(ReadVertexArrays_120_first_index); v < (ReadVertexArrays_120_last_index)+(uint32(1)); v++ {
					offset := (elstride) * (v) // u32
					arr.Pointer.Slice(uint64(offset), uint64((offset)+(elsize)), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
					_ = offset
				}
				_, _, _, _ = vertexAttribTypeSize_121_t, vertexAttribTypeSize_121_result, elsize, elstride
			}
			_ = arr
		}
		IndexSize_122_indices_type := ϟa.IndicesType // GLenum
		IndexSize_122_result := func() (result uint32) {
			switch IndexSize_122_indices_type {
			case GLenum_GL_UNSIGNED_BYTE:
				return uint32(1)
			case GLenum_GL_UNSIGNED_SHORT:
				return uint32(2)
			case GLenum_GL_UNSIGNED_INT:
				return uint32(4)
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", IndexSize_122_indices_type, ϟa))
				return result
			}
		}() // u32
		index_data.Slice(uint64(uint32(0)), uint64((uint32(ϟa.ElementCount))*(IndexSize_122_result)), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _, _, _, _, _, _, _ = index_data, first, last, ReadVertexArrays_120_ctx, ReadVertexArrays_120_first_index, ReadVertexArrays_120_last_index, IndexSize_122_indices_type, IndexSize_122_result
	}
	ϟb.Push(value.U32(ϟa.DrawMode))
	ϟb.Push(ϟa.ElementCount.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.IndicesType))
	ϟb.Push(ϟa.Indices.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElements)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_113_major, minRequiredVersion_113_minor, context, GetContext_117_result, ctx, count, id
	return nil
}

var _ = replay.Replayer(&GlDrawElementsIndirect{}) // interface compliance check
func (ϟa *GlDrawElementsIndirect) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_123_major := uint32(3) // u32
	minRequiredVersion_123_minor := uint32(1) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indirect.value())
	ϟb.Call(funcInfoGlDrawElementsIndirect)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_123_major, minRequiredVersion_123_minor
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstanced{}) // interface compliance check
func (ϟa *GlDrawElementsInstanced) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_126_major := uint32(3) // u32
	minRequiredVersion_126_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Instancecount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstanced)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_126_major, minRequiredVersion_126_minor
	return nil
}

var _ = replay.Replayer(&GlDrawRangeElements{}) // interface compliance check
func (ϟa *GlDrawRangeElements) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_129_major := uint32(3) // u32
	minRequiredVersion_129_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_LINES, GLenum_GL_LINE_LOOP, GLenum_GL_LINE_STRIP, GLenum_GL_POINTS, GLenum_GL_TRIANGLES, GLenum_GL_TRIANGLE_FAN, GLenum_GL_TRIANGLE_STRIP:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Start.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.End.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Call(funcInfoGlDrawRangeElements)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_129_major, minRequiredVersion_129_minor
	return nil
}

var _ = replay.Replayer(&GlActiveShaderProgramEXT{}) // interface compliance check
func (ϟa *GlActiveShaderProgramEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_132_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlActiveShaderProgramEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_132_ext
	return nil
}

var _ = replay.Replayer(&GlAlphaFuncQCOM{}) // interface compliance check
func (ϟa *GlAlphaFuncQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_133_ext := ExtensionId_GL_QCOM_alpha_test // ExtensionId
	ϟb.Push(value.U32(ϟa.Func))
	ϟb.Push(ϟa.Ref.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlAlphaFuncQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_133_ext
	return nil
}

var _ = replay.Replayer(&GlBeginConditionalRenderNV{}) // interface compliance check
func (ϟa *GlBeginConditionalRenderNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_134_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟb.Push(ϟa.Id.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlBeginConditionalRenderNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_134_ext
	return nil
}

var _ = replay.Replayer(&GlBeginPerfMonitorAMD{}) // interface compliance check
func (ϟa *GlBeginPerfMonitorAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_135_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Monitor.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBeginPerfMonitorAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_135_ext
	return nil
}

var _ = replay.Replayer(&GlBeginPerfQueryINTEL{}) // interface compliance check
func (ϟa *GlBeginPerfQueryINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_136_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryHandle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBeginPerfQueryINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_136_ext
	return nil
}

var _ = replay.Replayer(&GlBeginQueryEXT{}) // interface compliance check
func (ϟa *GlBeginQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_137_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_138_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBeginQueryEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_137_ext, requiresExtension_138_ext
	return nil
}

var _ = replay.Replayer(&GlBindProgramPipelineEXT{}) // interface compliance check
func (ϟa *GlBindProgramPipelineEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_139_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBindProgramPipelineEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_139_ext
	return nil
}

var _ = replay.Replayer(&GlBindVertexArrayOES{}) // interface compliance check
func (ϟa *GlBindVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_140_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_141_result := context                                    // Contextʳ
	ctx := GetContext_141_result                                        // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ctx.BoundVertexArray = ϟa.Array
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindVertexArrayOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = requiresExtension_140_ext, context, GetContext_141_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendBarrierNV{}) // interface compliance check
func (ϟa *GlBlendBarrierNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_142_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟb.Call(funcInfoGlBlendBarrierNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_142_ext
	return nil
}

var _ = replay.Replayer(&GlBlendEquationSeparateiOES{}) // interface compliance check
func (ϟa *GlBlendEquationSeparateiOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_143_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.ModeRGB))
	ϟb.Push(value.U32(ϟa.ModeAlpha))
	ϟb.Call(funcInfoGlBlendEquationSeparateiOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_143_ext
	return nil
}

var _ = replay.Replayer(&GlBlendEquationiOES{}) // interface compliance check
func (ϟa *GlBlendEquationiOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_144_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlBlendEquationiOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_144_ext
	return nil
}

var _ = replay.Replayer(&GlBlendFuncSeparateiOES{}) // interface compliance check
func (ϟa *GlBlendFuncSeparateiOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_145_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.SrcRGB))
	ϟb.Push(value.U32(ϟa.DstRGB))
	ϟb.Push(value.U32(ϟa.SrcAlpha))
	ϟb.Push(value.U32(ϟa.DstAlpha))
	ϟb.Call(funcInfoGlBlendFuncSeparateiOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_145_ext
	return nil
}

var _ = replay.Replayer(&GlBlendFunciOES{}) // interface compliance check
func (ϟa *GlBlendFunciOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_146_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Buf.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Src))
	ϟb.Push(value.U32(ϟa.Dst))
	ϟb.Call(funcInfoGlBlendFunciOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_146_ext
	return nil
}

var _ = replay.Replayer(&GlBlendParameteriNV{}) // interface compliance check
func (ϟa *GlBlendParameteriNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_147_ext := ExtensionId_GL_NV_blend_equation_advanced // ExtensionId
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBlendParameteriNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_147_ext
	return nil
}

var _ = replay.Replayer(&GlBlitFramebufferANGLE{}) // interface compliance check
func (ϟa *GlBlitFramebufferANGLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_148_ext := ExtensionId_GL_ANGLE_framebuffer_blit // ExtensionId
	ϟb.Push(ϟa.SrcX0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcX1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY1.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Push(value.U32(ϟa.Filter))
	ϟb.Call(funcInfoGlBlitFramebufferANGLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_148_ext
	return nil
}

var _ = replay.Replayer(&GlBlitFramebufferNV{}) // interface compliance check
func (ϟa *GlBlitFramebufferNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_149_ext := ExtensionId_GL_NV_framebuffer_blit // ExtensionId
	ϟb.Push(ϟa.SrcX0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcX1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY1.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Push(value.U32(ϟa.Filter))
	ϟb.Call(funcInfoGlBlitFramebufferNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_149_ext
	return nil
}

var _ = replay.Replayer(&GlBufferStorageEXT{}) // interface compliance check
func (ϟa *GlBufferStorageEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_150_ext := ExtensionId_GL_EXT_buffer_storage // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Push(value.U32(ϟa.Flag))
	ϟb.Call(funcInfoGlBufferStorageEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_150_ext
	return nil
}

var _ = replay.Replayer(&GlClientWaitSyncAPPLE{}) // interface compliance check
func (ϟa *GlClientWaitSyncAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_151_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Flag))
	ϟb.Push(ϟa.Timeout.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlClientWaitSyncAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_151_ext
	return nil
}

var _ = replay.Replayer(&GlColorMaskiOES{}) // interface compliance check
func (ϟa *GlColorMaskiOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_152_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.R.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.G.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.B.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.A.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlColorMaskiOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_152_ext
	return nil
}

var _ = replay.Replayer(&GlCompressedTexImage3DOES{}) // interface compliance check
func (ϟa *GlCompressedTexImage3DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_153_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Border.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.ImageSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlCompressedTexImage3DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_153_ext
	return nil
}

var _ = replay.Replayer(&GlCompressedTexSubImage3DOES{}) // interface compliance check
func (ϟa *GlCompressedTexSubImage3DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_154_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.ImageSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlCompressedTexSubImage3DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_154_ext
	return nil
}

var _ = replay.Replayer(&GlCopyBufferSubDataNV{}) // interface compliance check
func (ϟa *GlCopyBufferSubDataNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_155_ext := ExtensionId_GL_NV_copy_buffer // ExtensionId
	ϟb.Push(value.U32(ϟa.ReadTarget))
	ϟb.Push(value.U32(ϟa.WriteTarget))
	ϟb.Push(ϟa.ReadOffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.WriteOffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyBufferSubDataNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_155_ext
	return nil
}

var _ = replay.Replayer(&GlCopyImageSubDataOES{}) // interface compliance check
func (ϟa *GlCopyImageSubDataOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_156_ext := ExtensionId_GL_OES_copy_image // ExtensionId
	ϟb.Push(ϟa.SrcName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.SrcTarget))
	ϟb.Push(ϟa.SrcLevel.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.DstTarget))
	ϟb.Push(ϟa.DstLevel.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcWidth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcHeight.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcDepth.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyImageSubDataOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_156_ext
	return nil
}

var _ = replay.Replayer(&GlCopyPathNV{}) // interface compliance check
func (ϟa *GlCopyPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_157_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.ResultPath.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcPath.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_157_ext
	return nil
}

var _ = replay.Replayer(&GlCopyTexSubImage3DOES{}) // interface compliance check
func (ϟa *GlCopyTexSubImage3DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_158_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyTexSubImage3DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_158_ext
	return nil
}

var _ = replay.Replayer(&GlCopyTextureLevelsAPPLE{}) // interface compliance check
func (ϟa *GlCopyTextureLevelsAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_159_ext := ExtensionId_GL_APPLE_copy_texture_levels // ExtensionId
	ϟb.Push(ϟa.DestinationTexture.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SourceTexture.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SourceBaseLevel.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SourceLevelCount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyTextureLevelsAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_159_ext
	return nil
}

var _ = replay.Replayer(&GlCoverFillPathInstancedNV{}) // interface compliance check
func (ϟa *GlCoverFillPathInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_160_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.TransformValues.value())
	ϟb.Call(funcInfoGlCoverFillPathInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_160_ext
	return nil
}

var _ = replay.Replayer(&GlCoverFillPathNV{}) // interface compliance check
func (ϟa *GlCoverFillPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_161_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Call(funcInfoGlCoverFillPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_161_ext
	return nil
}

var _ = replay.Replayer(&GlCoverStrokePathInstancedNV{}) // interface compliance check
func (ϟa *GlCoverStrokePathInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_162_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.TransformValues.value())
	ϟb.Call(funcInfoGlCoverStrokePathInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_162_ext
	return nil
}

var _ = replay.Replayer(&GlCoverStrokePathNV{}) // interface compliance check
func (ϟa *GlCoverStrokePathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_163_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Call(funcInfoGlCoverStrokePathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_163_ext
	return nil
}

var _ = replay.Replayer(&GlCoverageMaskNV{}) // interface compliance check
func (ϟa *GlCoverageMaskNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_164_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCoverageMaskNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_164_ext
	return nil
}

var _ = replay.Replayer(&GlCoverageOperationNV{}) // interface compliance check
func (ϟa *GlCoverageOperationNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_165_ext := ExtensionId_GL_NV_coverage_sample // ExtensionId
	ϟb.Push(value.U32(ϟa.Operation))
	ϟb.Call(funcInfoGlCoverageOperationNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_165_ext
	return nil
}

var _ = replay.Replayer(&GlCreatePerfQueryINTEL{}) // interface compliance check
func (ϟa *GlCreatePerfQueryINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_166_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryId.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.QueryHandle.value())
	ϟb.Call(funcInfoGlCreatePerfQueryINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_166_ext
	return nil
}

var _ = replay.Replayer(&GlCreateShaderProgramvEXT{}) // interface compliance check
func (ϟa *GlCreateShaderProgramvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_167_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Strings.value())
	ϟb.Call(funcInfoGlCreateShaderProgramvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_167_ext
	return nil
}

var _ = replay.Replayer(&GlDeleteFencesNV{}) // interface compliance check
func (ϟa *GlDeleteFencesNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_168_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Fences.value())
	ϟb.Call(funcInfoGlDeleteFencesNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_168_ext
	return nil
}

var _ = replay.Replayer(&GlDeletePathsNV{}) // interface compliance check
func (ϟa *GlDeletePathsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_169_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Range.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDeletePathsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_169_ext
	return nil
}

var _ = replay.Replayer(&GlDeletePerfMonitorsAMD{}) // interface compliance check
func (ϟa *GlDeletePerfMonitorsAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_170_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Monitors.value())
	ϟb.Call(funcInfoGlDeletePerfMonitorsAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_170_ext
	return nil
}

var _ = replay.Replayer(&GlDeletePerfQueryINTEL{}) // interface compliance check
func (ϟa *GlDeletePerfQueryINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_171_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryHandle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDeletePerfQueryINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_171_ext
	return nil
}

var _ = replay.Replayer(&GlDeleteProgramPipelinesEXT{}) // interface compliance check
func (ϟa *GlDeleteProgramPipelinesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_172_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Pipelines.value())
	ϟb.Call(funcInfoGlDeleteProgramPipelinesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_172_ext
	return nil
}

var _ = replay.Replayer(&GlDeleteQueriesEXT{}) // interface compliance check
func (ϟa *GlDeleteQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_173_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_174_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_175_result := context                                        // Contextʳ
	ctx := GetContext_175_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Queries, q.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlDeleteQueriesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = requiresExtension_173_ext, requiresExtension_174_ext, q, context, GetContext_175_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDeleteSyncAPPLE{}) // interface compliance check
func (ϟa *GlDeleteSyncAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_176_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDeleteSyncAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_176_ext
	return nil
}

var _ = replay.Replayer(&GlDeleteVertexArraysOES{}) // interface compliance check
func (ϟa *GlDeleteVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_177_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_178_result := context                                      // Contextʳ
	ctx := GetContext_178_result                                          // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Arrays.value())
	ϟb.Call(funcInfoGlDeleteVertexArraysOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = requiresExtension_177_ext, context, GetContext_178_result, ctx, a
	return nil
}

var _ = replay.Replayer(&GlDepthRangeArrayfvNV{}) // interface compliance check
func (ϟa *GlDepthRangeArrayfvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_179_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.First.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlDepthRangeArrayfvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_179_ext
	return nil
}

var _ = replay.Replayer(&GlDepthRangeIndexedfNV{}) // interface compliance check
func (ϟa *GlDepthRangeIndexedfNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_180_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.F.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDepthRangeIndexedfNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_180_ext
	return nil
}

var _ = replay.Replayer(&GlDisableDriverControlQCOM{}) // interface compliance check
func (ϟa *GlDisableDriverControlQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_181_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟb.Push(ϟa.DriverControl.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDisableDriverControlQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_181_ext
	return nil
}

var _ = replay.Replayer(&GlDisableiNV{}) // interface compliance check
func (ϟa *GlDisableiNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_182_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDisableiNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_182_ext
	return nil
}

var _ = replay.Replayer(&GlDisableiOES{}) // interface compliance check
func (ϟa *GlDisableiOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_183_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDisableiOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_183_ext
	return nil
}

var _ = replay.Replayer(&GlDiscardFramebufferEXT{}) // interface compliance check
func (ϟa *GlDiscardFramebufferEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_184_ext := ExtensionId_GL_EXT_discard_framebuffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.NumAttachments.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Attachments.value())
	ϟb.Call(funcInfoGlDiscardFramebufferEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_184_ext
	return nil
}

var _ = replay.Replayer(&GlDrawArraysInstancedANGLE{}) // interface compliance check
func (ϟa *GlDrawArraysInstancedANGLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_185_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.First.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawArraysInstancedANGLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_185_ext
	return nil
}

var _ = replay.Replayer(&GlDrawArraysInstancedBaseInstanceEXT{}) // interface compliance check
func (ϟa *GlDrawArraysInstancedBaseInstanceEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_186_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.First.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Instancecount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Baseinstance.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawArraysInstancedBaseInstanceEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_186_ext
	return nil
}

var _ = replay.Replayer(&GlDrawArraysInstancedEXT{}) // interface compliance check
func (ϟa *GlDrawArraysInstancedEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_187_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_188_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Start.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawArraysInstancedEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_187_ext, requiresExtension_188_ext
	return nil
}

var _ = replay.Replayer(&GlDrawArraysInstancedNV{}) // interface compliance check
func (ϟa *GlDrawArraysInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_189_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.First.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawArraysInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_189_ext
	return nil
}

var _ = replay.Replayer(&GlDrawBuffersEXT{}) // interface compliance check
func (ϟa *GlDrawBuffersEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_190_ext := ExtensionId_GL_EXT_draw_buffers // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Bufs.value())
	ϟb.Call(funcInfoGlDrawBuffersEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_190_ext
	return nil
}

var _ = replay.Replayer(&GlDrawBuffersIndexedEXT{}) // interface compliance check
func (ϟa *GlDrawBuffersIndexedEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_191_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Location.value())
	ϟb.Push(ϟa.Indices.value())
	ϟb.Call(funcInfoGlDrawBuffersIndexedEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_191_ext
	return nil
}

var _ = replay.Replayer(&GlDrawBuffersNV{}) // interface compliance check
func (ϟa *GlDrawBuffersNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_192_ext := ExtensionId_GL_NV_draw_buffers // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Bufs.value())
	ϟb.Call(funcInfoGlDrawBuffersNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_192_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsBaseVertexEXT{}) // interface compliance check
func (ϟa *GlDrawElementsBaseVertexEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_193_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Basevertex.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsBaseVertexEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_193_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsBaseVertexOES{}) // interface compliance check
func (ϟa *GlDrawElementsBaseVertexOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_194_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Basevertex.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsBaseVertexOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_194_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstancedANGLE{}) // interface compliance check
func (ϟa *GlDrawElementsInstancedANGLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_195_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstancedANGLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_195_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstancedBaseInstanceEXT{}) // interface compliance check
func (ϟa *GlDrawElementsInstancedBaseInstanceEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_196_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Instancecount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Baseinstance.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstancedBaseInstanceEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_196_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstancedBaseVertexEXT{}) // interface compliance check
func (ϟa *GlDrawElementsInstancedBaseVertexEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_197_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Instancecount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Basevertex.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstancedBaseVertexEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_197_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstancedBaseVertexOES{}) // interface compliance check
func (ϟa *GlDrawElementsInstancedBaseVertexOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_198_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Instancecount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Basevertex.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstancedBaseVertexOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_198_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstancedBaseVertexBaseInstanceEXT{}) // interface compliance check
func (ϟa *GlDrawElementsInstancedBaseVertexBaseInstanceEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_199_ext := ExtensionId_GL_EXT_base_instance // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Instancecount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Basevertex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Baseinstance.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstancedBaseVertexBaseInstanceEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_199_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstancedEXT{}) // interface compliance check
func (ϟa *GlDrawElementsInstancedEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_200_ext := ExtensionId_GL_EXT_draw_instanced   // ExtensionId
	requiresExtension_201_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstancedEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_200_ext, requiresExtension_201_ext
	return nil
}

var _ = replay.Replayer(&GlDrawElementsInstancedNV{}) // interface compliance check
func (ϟa *GlDrawElementsInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_202_ext := ExtensionId_GL_NV_draw_instanced // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawElementsInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_202_ext
	return nil
}

var _ = replay.Replayer(&GlDrawRangeElementsBaseVertexEXT{}) // interface compliance check
func (ϟa *GlDrawRangeElementsBaseVertexEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_203_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Start.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.End.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Basevertex.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawRangeElementsBaseVertexEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_203_ext
	return nil
}

var _ = replay.Replayer(&GlDrawRangeElementsBaseVertexOES{}) // interface compliance check
func (ϟa *GlDrawRangeElementsBaseVertexOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_204_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Start.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.End.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Basevertex.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDrawRangeElementsBaseVertexOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_204_ext
	return nil
}

var _ = replay.Replayer(&GlEGLImageTargetRenderbufferStorageOES{}) // interface compliance check
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_205_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Image.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEGLImageTargetRenderbufferStorageOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_205_ext
	return nil
}

var _ = replay.Replayer(&GlEGLImageTargetTexture2DOES{}) // interface compliance check
func (ϟa *GlEGLImageTargetTexture2DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_206_ext := ExtensionId_GL_OES_EGL_image // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Image.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEGLImageTargetTexture2DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_206_ext
	return nil
}

var _ = replay.Replayer(&GlEnableDriverControlQCOM{}) // interface compliance check
func (ϟa *GlEnableDriverControlQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_207_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟb.Push(ϟa.DriverControl.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEnableDriverControlQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_207_ext
	return nil
}

var _ = replay.Replayer(&GlEnableiNV{}) // interface compliance check
func (ϟa *GlEnableiNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_208_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEnableiNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_208_ext
	return nil
}

var _ = replay.Replayer(&GlEnableiOES{}) // interface compliance check
func (ϟa *GlEnableiOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_209_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEnableiOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_209_ext
	return nil
}

var _ = replay.Replayer(&GlEndConditionalRenderNV{}) // interface compliance check
func (ϟa *GlEndConditionalRenderNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_210_ext := ExtensionId_GL_NV_conditional_render // ExtensionId
	ϟb.Call(funcInfoGlEndConditionalRenderNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_210_ext
	return nil
}

var _ = replay.Replayer(&GlEndPerfMonitorAMD{}) // interface compliance check
func (ϟa *GlEndPerfMonitorAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_211_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Monitor.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEndPerfMonitorAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_211_ext
	return nil
}

var _ = replay.Replayer(&GlEndPerfQueryINTEL{}) // interface compliance check
func (ϟa *GlEndPerfQueryINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_212_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryHandle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEndPerfQueryINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_212_ext
	return nil
}

var _ = replay.Replayer(&GlEndQueryEXT{}) // interface compliance check
func (ϟa *GlEndQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_213_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_214_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlEndQueryEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = requiresExtension_213_ext, requiresExtension_214_ext
	return nil
}

var _ = replay.Replayer(&GlEndTilingQCOM{}) // interface compliance check
func (ϟa *GlEndTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_215_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.PreserveMask))
	ϟb.Call(funcInfoGlEndTilingQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_215_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetBufferPointervQCOM{}) // interface compliance check
func (ϟa *GlExtGetBufferPointervQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_216_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlExtGetBufferPointervQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_216_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetBuffersQCOM{}) // interface compliance check
func (ϟa *GlExtGetBuffersQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_217_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟb.Push(ϟa.Buffers.value())
	ϟb.Push(ϟa.MaxBuffers.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumBuffers.value())
	ϟb.Call(funcInfoGlExtGetBuffersQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_217_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetFramebuffersQCOM{}) // interface compliance check
func (ϟa *GlExtGetFramebuffersQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_218_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟb.Push(ϟa.Framebuffers.value())
	ϟb.Push(ϟa.MaxFramebuffers.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumFramebuffers.value())
	ϟb.Call(funcInfoGlExtGetFramebuffersQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_218_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetProgramBinarySourceQCOM{}) // interface compliance check
func (ϟa *GlExtGetProgramBinarySourceQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_219_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Shadertype))
	ϟb.Push(ϟa.Source.value())
	ϟb.Push(ϟa.Length.value())
	ϟb.Call(funcInfoGlExtGetProgramBinarySourceQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_219_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetProgramsQCOM{}) // interface compliance check
func (ϟa *GlExtGetProgramsQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_220_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟb.Push(ϟa.Programs.value())
	ϟb.Push(ϟa.MaxPrograms.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumPrograms.value())
	ϟb.Call(funcInfoGlExtGetProgramsQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_220_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetRenderbuffersQCOM{}) // interface compliance check
func (ϟa *GlExtGetRenderbuffersQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_221_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟb.Push(ϟa.Renderbuffers.value())
	ϟb.Push(ϟa.MaxRenderbuffers.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumRenderbuffers.value())
	ϟb.Call(funcInfoGlExtGetRenderbuffersQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_221_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetShadersQCOM{}) // interface compliance check
func (ϟa *GlExtGetShadersQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_222_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	ϟb.Push(ϟa.Shaders.value())
	ϟb.Push(ϟa.MaxShaders.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumShaders.value())
	ϟb.Call(funcInfoGlExtGetShadersQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_222_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetTexLevelParameterivQCOM{}) // interface compliance check
func (ϟa *GlExtGetTexLevelParameterivQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_223_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlExtGetTexLevelParameterivQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_223_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetTexSubImageQCOM{}) // interface compliance check
func (ϟa *GlExtGetTexSubImageQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_224_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Texels.value())
	ϟb.Call(funcInfoGlExtGetTexSubImageQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_224_ext
	return nil
}

var _ = replay.Replayer(&GlExtGetTexturesQCOM{}) // interface compliance check
func (ϟa *GlExtGetTexturesQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_225_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟb.Push(ϟa.Textures.value())
	ϟb.Push(ϟa.MaxTextures.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumTextures.value())
	ϟb.Call(funcInfoGlExtGetTexturesQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_225_ext
	return nil
}

var _ = replay.Replayer(&GlExtIsProgramBinaryQCOM{}) // interface compliance check
func (ϟa *GlExtIsProgramBinaryQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_226_ext := ExtensionId_GL_QCOM_extended_get2 // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlExtIsProgramBinaryQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_226_ext
	return nil
}

var _ = replay.Replayer(&GlExtTexObjectStateOverrideiQCOM{}) // interface compliance check
func (ϟa *GlExtTexObjectStateOverrideiQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_227_ext := ExtensionId_GL_QCOM_extended_get // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlExtTexObjectStateOverrideiQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_227_ext
	return nil
}

var _ = replay.Replayer(&GlFenceSyncAPPLE{}) // interface compliance check
func (ϟa *GlFenceSyncAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_228_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟb.Push(value.U32(ϟa.Condition))
	ϟb.Push(value.U32(ϟa.Flag))
	ϟb.Call(funcInfoGlFenceSyncAPPLE)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		ptr, found := ϟb.Remappings[key]
		if !found {
			ptr = ϟb.AllocateMemory(uint64(8))
			ϟb.Remappings[key] = ptr
		}
		ϟb.Clone(0)
		ϟb.Store(ptr)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_228_ext
	return nil
}

var _ = replay.Replayer(&GlFinishFenceNV{}) // interface compliance check
func (ϟa *GlFinishFenceNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_229_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟb.Push(ϟa.Fence.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFinishFenceNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_229_ext
	return nil
}

var _ = replay.Replayer(&GlFlushMappedBufferRangeEXT{}) // interface compliance check
func (ϟa *GlFlushMappedBufferRangeEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_230_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFlushMappedBufferRangeEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_230_ext
	return nil
}

var _ = replay.Replayer(&GlFramebufferTexture2DMultisampleEXT{}) // interface compliance check
func (ϟa *GlFramebufferTexture2DMultisampleEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_231_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Attachment))
	ϟb.Push(value.U32(ϟa.Textarget))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTexture2DMultisampleEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_231_ext
	return nil
}

var _ = replay.Replayer(&GlFramebufferTexture2DMultisampleIMG{}) // interface compliance check
func (ϟa *GlFramebufferTexture2DMultisampleIMG) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_232_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Attachment))
	ϟb.Push(value.U32(ϟa.Textarget))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTexture2DMultisampleIMG)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_232_ext
	return nil
}

var _ = replay.Replayer(&GlFramebufferTexture3DOES{}) // interface compliance check
func (ϟa *GlFramebufferTexture3DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_233_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Attachment))
	ϟb.Push(value.U32(ϟa.Textarget))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTexture3DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_233_ext
	return nil
}

var _ = replay.Replayer(&GlFramebufferTextureOES{}) // interface compliance check
func (ϟa *GlFramebufferTextureOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_234_ext := ExtensionId_GL_OES_geometry_shader // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Attachment))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTextureOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_234_ext
	return nil
}

var _ = replay.Replayer(&GlFramebufferTextureMultiviewOVR{}) // interface compliance check
func (ϟa *GlFramebufferTextureMultiviewOVR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_235_ext := ExtensionId_GL_OVR_multiview // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Attachment))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BaseViewIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumViews.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTextureMultiviewOVR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_235_ext
	return nil
}

var _ = replay.Replayer(&GlGenFencesNV{}) // interface compliance check
func (ϟa *GlGenFencesNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_236_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Fences.value())
	ϟb.Call(funcInfoGlGenFencesNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_236_ext
	return nil
}

var _ = replay.Replayer(&GlGenPathsNV{}) // interface compliance check
func (ϟa *GlGenPathsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_237_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Range.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlGenPathsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_237_ext
	return nil
}

var _ = replay.Replayer(&GlGenPerfMonitorsAMD{}) // interface compliance check
func (ϟa *GlGenPerfMonitorsAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_238_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Monitors.value())
	ϟb.Call(funcInfoGlGenPerfMonitorsAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_238_ext
	return nil
}

var _ = replay.Replayer(&GlGenProgramPipelinesEXT{}) // interface compliance check
func (ϟa *GlGenProgramPipelinesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_239_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Pipelines.value())
	ϟb.Call(funcInfoGlGenProgramPipelinesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_239_ext
	return nil
}

var _ = replay.Replayer(&GlGenQueriesEXT{}) // interface compliance check
func (ϟa *GlGenQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_240_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_241_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	q := ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs)  // QueryIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_242_result := context                                        // Contextʳ
	ctx := GetContext_242_result                                            // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Queries.value())
	ϟb.Call(funcInfoGlGenQueriesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := QueryId(ϟa.Queries.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // QueryId
		ctx.Instances.Queries[id] = func() *Query {
			s := &Query{}
			s.Init()
			return s
		}()
		q.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = requiresExtension_240_ext, requiresExtension_241_ext, q, context, GetContext_242_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenVertexArraysOES{}) // interface compliance check
func (ϟa *GlGenVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_243_ext := ExtensionId_GL_OES_vertex_array_object   // ExtensionId
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_244_result := context                                      // Contextʳ
	ctx := GetContext_244_result                                          // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Arrays.value())
	ϟb.Call(funcInfoGlGenVertexArraysOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		a.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _ = requiresExtension_243_ext, a, context, GetContext_244_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetBufferPointervOES{}) // interface compliance check
func (ϟa *GlGetBufferPointervOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_245_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetBufferPointervOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_245_ext
	return nil
}

var _ = replay.Replayer(&GlGetDriverControlStringQCOM{}) // interface compliance check
func (ϟa *GlGetDriverControlStringQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_246_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟb.Push(ϟa.DriverControl.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.DriverControlString.value())
	ϟb.Call(funcInfoGlGetDriverControlStringQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_246_ext
	return nil
}

var _ = replay.Replayer(&GlGetDriverControlsQCOM{}) // interface compliance check
func (ϟa *GlGetDriverControlsQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_247_ext := ExtensionId_GL_QCOM_driver_control // ExtensionId
	ϟb.Push(ϟa.Num.value())
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DriverControls.value())
	ϟb.Call(funcInfoGlGetDriverControlsQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_247_ext
	return nil
}

var _ = replay.Replayer(&GlGetFenceivNV{}) // interface compliance check
func (ϟa *GlGetFenceivNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_248_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟb.Push(ϟa.Fence.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetFenceivNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_248_ext
	return nil
}

var _ = replay.Replayer(&GlGetFirstPerfQueryIdINTEL{}) // interface compliance check
func (ϟa *GlGetFirstPerfQueryIdINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_249_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryId.value())
	ϟb.Call(funcInfoGlGetFirstPerfQueryIdINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_249_ext
	return nil
}

var _ = replay.Replayer(&GlGetFloati_vNV{}) // interface compliance check
func (ϟa *GlGetFloati_vNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_250_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlGetFloati_vNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_250_ext
	return nil
}

var _ = replay.Replayer(&GlGetGraphicsResetStatusEXT{}) // interface compliance check
func (ϟa *GlGetGraphicsResetStatusEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_251_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟb.Call(funcInfoGlGetGraphicsResetStatusEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_251_ext
	return nil
}

var _ = replay.Replayer(&GlGetGraphicsResetStatusKHR{}) // interface compliance check
func (ϟa *GlGetGraphicsResetStatusKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_252_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟb.Call(funcInfoGlGetGraphicsResetStatusKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_252_ext
	return nil
}

var _ = replay.Replayer(&GlGetImageHandleNV{}) // interface compliance check
func (ϟa *GlGetImageHandleNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_253_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Layered.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Layer.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Call(funcInfoGlGetImageHandleNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_253_ext
	return nil
}

var _ = replay.Replayer(&GlGetInteger64vAPPLE{}) // interface compliance check
func (ϟa *GlGetInteger64vAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_254_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetInteger64vAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_254_ext
	return nil
}

var _ = replay.Replayer(&GlGetIntegeri_vEXT{}) // interface compliance check
func (ϟa *GlGetIntegeri_vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_255_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlGetIntegeri_vEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_255_ext
	return nil
}

var _ = replay.Replayer(&GlGetInternalformatSampleivNV{}) // interface compliance check
func (ϟa *GlGetInternalformatSampleivNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_256_ext := ExtensionId_GL_NV_internalformat_sample_query // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetInternalformatSampleivNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_256_ext
	return nil
}

var _ = replay.Replayer(&GlGetNextPerfQueryIdINTEL{}) // interface compliance check
func (ϟa *GlGetNextPerfQueryIdINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_257_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryId.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NextQueryId.value())
	ϟb.Call(funcInfoGlGetNextPerfQueryIdINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_257_ext
	return nil
}

var _ = replay.Replayer(&GlGetObjectLabelEXT{}) // interface compliance check
func (ϟa *GlGetObjectLabelEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_258_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Object.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Label.value())
	ϟb.Call(funcInfoGlGetObjectLabelEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_258_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathCommandsNV{}) // interface compliance check
func (ϟa *GlGetPathCommandsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_259_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Commands.value())
	ϟb.Call(funcInfoGlGetPathCommandsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_259_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathCoordsNV{}) // interface compliance check
func (ϟa *GlGetPathCoordsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_260_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Coords.value())
	ϟb.Call(funcInfoGlGetPathCoordsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_260_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathDashArrayNV{}) // interface compliance check
func (ϟa *GlGetPathDashArrayNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_261_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DashArray.value())
	ϟb.Call(funcInfoGlGetPathDashArrayNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_261_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathLengthNV{}) // interface compliance check
func (ϟa *GlGetPathLengthNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_262_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.StartSegment.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumSegments.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlGetPathLengthNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_262_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathMetricRangeNV{}) // interface compliance check
func (ϟa *GlGetPathMetricRangeNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_263_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MetricQueryMask))
	ϟb.Push(ϟa.FirstPathName.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Stride.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Metrics.value())
	ϟb.Call(funcInfoGlGetPathMetricRangeNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_263_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathMetricsNV{}) // interface compliance check
func (ϟa *GlGetPathMetricsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_264_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MetricQueryMask))
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Stride.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Metrics.value())
	ϟb.Call(funcInfoGlGetPathMetricsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_264_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathParameterfvNV{}) // interface compliance check
func (ϟa *GlGetPathParameterfvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_265_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetPathParameterfvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_265_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathParameterivNV{}) // interface compliance check
func (ϟa *GlGetPathParameterivNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_266_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetPathParameterivNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_266_ext
	return nil
}

var _ = replay.Replayer(&GlGetPathSpacingNV{}) // interface compliance check
func (ϟa *GlGetPathSpacingNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_267_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.PathListMode))
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.AdvanceScale.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.KerningScale.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.ReturnedSpacing.value())
	ϟb.Call(funcInfoGlGetPathSpacingNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_267_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfCounterInfoINTEL{}) // interface compliance check
func (ϟa *GlGetPerfCounterInfoINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_268_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryId.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CounterId.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CounterNameLength.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CounterName.value())
	ϟb.Push(ϟa.CounterDescLength.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CounterDesc.value())
	ϟb.Push(ϟa.CounterOffset.value())
	ϟb.Push(ϟa.CounterDataSize.value())
	ϟb.Push(ϟa.CounterTypeEnum.value())
	ϟb.Push(ϟa.CounterDataTypeEnum.value())
	ϟb.Push(ϟa.RawCounterMaxValue.value())
	ϟb.Call(funcInfoGlGetPerfCounterInfoINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_268_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfMonitorCounterDataAMD{}) // interface compliance check
func (ϟa *GlGetPerfMonitorCounterDataAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_269_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Monitor.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.DataSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Push(ϟa.BytesWritten.value())
	ϟb.Call(funcInfoGlGetPerfMonitorCounterDataAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_269_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfMonitorCounterInfoAMD{}) // interface compliance check
func (ϟa *GlGetPerfMonitorCounterInfoAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_270_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Group.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Counter.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlGetPerfMonitorCounterInfoAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_270_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfMonitorCounterStringAMD{}) // interface compliance check
func (ϟa *GlGetPerfMonitorCounterStringAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_271_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Group.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Counter.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.CounterString.value())
	ϟb.Call(funcInfoGlGetPerfMonitorCounterStringAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_271_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfMonitorCountersAMD{}) // interface compliance check
func (ϟa *GlGetPerfMonitorCountersAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_272_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Group.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumCounters.value())
	ϟb.Push(ϟa.MaxActiveCounters.value())
	ϟb.Push(ϟa.CounterSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Counters.value())
	ϟb.Call(funcInfoGlGetPerfMonitorCountersAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_272_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfMonitorGroupStringAMD{}) // interface compliance check
func (ϟa *GlGetPerfMonitorGroupStringAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_273_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Group.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.GroupString.value())
	ϟb.Call(funcInfoGlGetPerfMonitorGroupStringAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_273_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfMonitorGroupsAMD{}) // interface compliance check
func (ϟa *GlGetPerfMonitorGroupsAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_274_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.NumGroups.value())
	ϟb.Push(ϟa.GroupsSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Groups.value())
	ϟb.Call(funcInfoGlGetPerfMonitorGroupsAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_274_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfQueryDataINTEL{}) // interface compliance check
func (ϟa *GlGetPerfQueryDataINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_275_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryHandle.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Flag.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DataSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Push(ϟa.BytesWritten.value())
	ϟb.Call(funcInfoGlGetPerfQueryDataINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_275_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfQueryIdByNameINTEL{}) // interface compliance check
func (ϟa *GlGetPerfQueryIdByNameINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_276_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryName.value())
	ϟb.Push(ϟa.QueryId.value())
	ϟb.Call(funcInfoGlGetPerfQueryIdByNameINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_276_ext
	return nil
}

var _ = replay.Replayer(&GlGetPerfQueryInfoINTEL{}) // interface compliance check
func (ϟa *GlGetPerfQueryInfoINTEL) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_277_ext := ExtensionId_GL_INTEL_performance_query // ExtensionId
	ϟb.Push(ϟa.QueryId.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.QueryNameLength.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.QueryName.value())
	ϟb.Push(ϟa.DataSize.value())
	ϟb.Push(ϟa.NoCounters.value())
	ϟb.Push(ϟa.NoInstances.value())
	ϟb.Push(ϟa.CapsMask.value())
	ϟb.Call(funcInfoGlGetPerfQueryInfoINTEL)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_277_ext
	return nil
}

var _ = replay.Replayer(&GlGetProgramBinaryOES{}) // interface compliance check
func (ϟa *GlGetProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_278_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufferSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BytesWritten.value())
	ϟb.Push(ϟa.BinaryFormat.value())
	ϟb.Push(ϟa.Binary.value())
	ϟb.Call(funcInfoGlGetProgramBinaryOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // GLsizei
	ϟa.BytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.BinaryFormat.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Binary.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_278_ext, l
	return nil
}

var _ = replay.Replayer(&GlGetProgramPipelineInfoLogEXT{}) // interface compliance check
func (ϟa *GlGetProgramPipelineInfoLogEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_279_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.InfoLog.value())
	ϟb.Call(funcInfoGlGetProgramPipelineInfoLogEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_279_ext
	return nil
}

var _ = replay.Replayer(&GlGetProgramPipelineivEXT{}) // interface compliance check
func (ϟa *GlGetProgramPipelineivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_280_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetProgramPipelineivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_280_ext
	return nil
}

var _ = replay.Replayer(&GlGetProgramResourcefvNV{}) // interface compliance check
func (ϟa *GlGetProgramResourcefvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_281_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.ProgramInterface))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.PropCount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Props.value())
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetProgramResourcefvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_281_ext
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjecti64vEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjecti64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_282_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjecti64vEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_282_ext
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectivEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_283_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_283_ext
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectui64vEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectui64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_284_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectui64vEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_ = requiresExtension_284_ext
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectuivEXT{}) // interface compliance check
func (ϟa *GlGetQueryObjectuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_285_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_286_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectuivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_285_ext, requiresExtension_286_ext
	return nil
}

var _ = replay.Replayer(&GlGetQueryivEXT{}) // interface compliance check
func (ϟa *GlGetQueryivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_287_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_288_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = requiresExtension_287_ext, requiresExtension_288_ext
	return nil
}

var _ = replay.Replayer(&GlGetSamplerParameterIivOES{}) // interface compliance check
func (ϟa *GlGetSamplerParameterIivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_289_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetSamplerParameterIivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_289_ext
	return nil
}

var _ = replay.Replayer(&GlGetSamplerParameterIuivOES{}) // interface compliance check
func (ϟa *GlGetSamplerParameterIuivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_290_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetSamplerParameterIuivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_290_ext
	return nil
}

var _ = replay.Replayer(&GlGetSyncivAPPLE{}) // interface compliance check
func (ϟa *GlGetSyncivAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_291_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetSyncivAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_291_ext
	return nil
}

var _ = replay.Replayer(&GlGetTexParameterIivOES{}) // interface compliance check
func (ϟa *GlGetTexParameterIivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_292_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetTexParameterIivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_292_ext
	return nil
}

var _ = replay.Replayer(&GlGetTexParameterIuivOES{}) // interface compliance check
func (ϟa *GlGetTexParameterIuivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_293_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetTexParameterIuivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_293_ext
	return nil
}

var _ = replay.Replayer(&GlGetTextureHandleNV{}) // interface compliance check
func (ϟa *GlGetTextureHandleNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_294_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlGetTextureHandleNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_294_ext
	return nil
}

var _ = replay.Replayer(&GlGetTextureSamplerHandleNV{}) // interface compliance check
func (ϟa *GlGetTextureSamplerHandleNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_295_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlGetTextureSamplerHandleNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_295_ext
	return nil
}

var _ = replay.Replayer(&GlGetTranslatedShaderSourceANGLE{}) // interface compliance check
func (ϟa *GlGetTranslatedShaderSourceANGLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_296_ext := ExtensionId_GL_ANGLE_translated_shader_source // ExtensionId
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Bufsize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Source.value())
	ϟb.Call(funcInfoGlGetTranslatedShaderSourceANGLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_296_ext
	return nil
}

var _ = replay.Replayer(&GlGetnUniformfvEXT{}) // interface compliance check
func (ϟa *GlGetnUniformfvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_297_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetnUniformfvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_297_ext
	return nil
}

var _ = replay.Replayer(&GlGetnUniformfvKHR{}) // interface compliance check
func (ϟa *GlGetnUniformfvKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_298_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetnUniformfvKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_298_ext
	return nil
}

var _ = replay.Replayer(&GlGetnUniformivEXT{}) // interface compliance check
func (ϟa *GlGetnUniformivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_299_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetnUniformivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_299_ext
	return nil
}

var _ = replay.Replayer(&GlGetnUniformivKHR{}) // interface compliance check
func (ϟa *GlGetnUniformivKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_300_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetnUniformivKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_300_ext
	return nil
}

var _ = replay.Replayer(&GlGetnUniformuivKHR{}) // interface compliance check
func (ϟa *GlGetnUniformuivKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_301_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetnUniformuivKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_301_ext
	return nil
}

var _ = replay.Replayer(&GlInsertEventMarkerEXT{}) // interface compliance check
func (ϟa *GlInsertEventMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_302_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl, true).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Marker.value())
	ϟb.Call(funcInfoGlInsertEventMarkerEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_302_ext
	return nil
}

var _ = replay.Replayer(&GlInterpolatePathsNV{}) // interface compliance check
func (ϟa *GlInterpolatePathsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_303_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.ResultPath.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.PathA.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.PathB.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Weight.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlInterpolatePathsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_303_ext
	return nil
}

var _ = replay.Replayer(&GlIsEnablediOES{}) // interface compliance check
func (ϟa *GlIsEnablediOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_304_ext := ExtensionId_GL_OES_draw_buffers_indexed // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsEnablediOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_304_ext
	return nil
}

var _ = replay.Replayer(&GlIsEnablediNV{}) // interface compliance check
func (ϟa *GlIsEnablediNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_305_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsEnablediNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_305_ext
	return nil
}

var _ = replay.Replayer(&GlIsFenceNV{}) // interface compliance check
func (ϟa *GlIsFenceNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_306_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟb.Push(ϟa.Fence.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsFenceNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_306_ext
	return nil
}

var _ = replay.Replayer(&GlIsImageHandleResidentNV{}) // interface compliance check
func (ϟa *GlIsImageHandleResidentNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_307_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟb.Push(ϟa.Handle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsImageHandleResidentNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_307_ext
	return nil
}

var _ = replay.Replayer(&GlIsPathNV{}) // interface compliance check
func (ϟa *GlIsPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_308_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_308_ext
	return nil
}

var _ = replay.Replayer(&GlIsPointInFillPathNV{}) // interface compliance check
func (ϟa *GlIsPointInFillPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_309_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsPointInFillPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_309_ext
	return nil
}

var _ = replay.Replayer(&GlIsPointInStrokePathNV{}) // interface compliance check
func (ϟa *GlIsPointInStrokePathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_310_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsPointInStrokePathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_310_ext
	return nil
}

var _ = replay.Replayer(&GlIsProgramPipelineEXT{}) // interface compliance check
func (ϟa *GlIsProgramPipelineEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_311_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsProgramPipelineEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_311_ext
	return nil
}

var _ = replay.Replayer(&GlIsQueryEXT{}) // interface compliance check
func (ϟa *GlIsQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_312_ext := ExtensionId_GL_EXT_disjoint_timer_query    // ExtensionId
	requiresExtension_313_ext := ExtensionId_GL_EXT_occlusion_query_boolean // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_314_result := context                                        // Contextʳ
	ctx := GetContext_314_result                                            // Contextʳ
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsQueryEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = requiresExtension_312_ext, requiresExtension_313_ext, context, GetContext_314_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsSyncAPPLE{}) // interface compliance check
func (ϟa *GlIsSyncAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_315_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsSyncAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_315_ext
	return nil
}

var _ = replay.Replayer(&GlIsTextureHandleResidentNV{}) // interface compliance check
func (ϟa *GlIsTextureHandleResidentNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_316_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟb.Push(ϟa.Handle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsTextureHandleResidentNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_316_ext
	return nil
}

var _ = replay.Replayer(&GlIsVertexArrayOES{}) // interface compliance check
func (ϟa *GlIsVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_317_ext := ExtensionId_GL_OES_vertex_array_object // ExtensionId
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                        // Contextʳ
	GetContext_318_result := context                                    // Contextʳ
	ctx := GetContext_318_result                                        // Contextʳ
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsVertexArrayOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = requiresExtension_317_ext, context, GetContext_318_result, ctx
	return nil
}

var _ = replay.Replayer(&GlLabelObjectEXT{}) // interface compliance check
func (ϟa *GlLabelObjectEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_319_ext := ExtensionId_GL_EXT_debug_label // ExtensionId
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Object.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Label.value())
	ϟb.Call(funcInfoGlLabelObjectEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_319_ext
	return nil
}

var _ = replay.Replayer(&GlMakeImageHandleNonResidentNV{}) // interface compliance check
func (ϟa *GlMakeImageHandleNonResidentNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_320_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟb.Push(ϟa.Handle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMakeImageHandleNonResidentNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_320_ext
	return nil
}

var _ = replay.Replayer(&GlMakeImageHandleResidentNV{}) // interface compliance check
func (ϟa *GlMakeImageHandleResidentNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_321_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟb.Push(ϟa.Handle.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Access))
	ϟb.Call(funcInfoGlMakeImageHandleResidentNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_321_ext
	return nil
}

var _ = replay.Replayer(&GlMakeTextureHandleNonResidentNV{}) // interface compliance check
func (ϟa *GlMakeTextureHandleNonResidentNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_322_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟb.Push(ϟa.Handle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMakeTextureHandleNonResidentNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_322_ext
	return nil
}

var _ = replay.Replayer(&GlMakeTextureHandleResidentNV{}) // interface compliance check
func (ϟa *GlMakeTextureHandleResidentNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_323_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	ϟb.Push(ϟa.Handle.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMakeTextureHandleResidentNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_323_ext
	return nil
}

var _ = replay.Replayer(&GlMapBufferOES{}) // interface compliance check
func (ϟa *GlMapBufferOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_324_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Access))
	ϟb.Call(funcInfoGlMapBufferOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_324_ext
	return nil
}

var _ = replay.Replayer(&GlMapBufferRangeEXT{}) // interface compliance check
func (ϟa *GlMapBufferRangeEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_325_ext := ExtensionId_GL_EXT_map_buffer_range // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Access))
	ϟb.Call(funcInfoGlMapBufferRangeEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_325_ext
	return nil
}

var _ = replay.Replayer(&GlMatrixLoad3x2fNV{}) // interface compliance check
func (ϟa *GlMatrixLoad3x2fNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_326_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MatrixMode))
	ϟb.Push(ϟa.M.value())
	ϟb.Call(funcInfoGlMatrixLoad3x2fNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_326_ext
	return nil
}

var _ = replay.Replayer(&GlMatrixLoad3x3fNV{}) // interface compliance check
func (ϟa *GlMatrixLoad3x3fNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_327_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MatrixMode))
	ϟb.Push(ϟa.M.value())
	ϟb.Call(funcInfoGlMatrixLoad3x3fNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_327_ext
	return nil
}

var _ = replay.Replayer(&GlMatrixLoadTranspose3x3fNV{}) // interface compliance check
func (ϟa *GlMatrixLoadTranspose3x3fNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_328_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MatrixMode))
	ϟb.Push(ϟa.M.value())
	ϟb.Call(funcInfoGlMatrixLoadTranspose3x3fNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_328_ext
	return nil
}

var _ = replay.Replayer(&GlMatrixMult3x2fNV{}) // interface compliance check
func (ϟa *GlMatrixMult3x2fNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_329_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MatrixMode))
	ϟb.Push(ϟa.M.value())
	ϟb.Call(funcInfoGlMatrixMult3x2fNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_329_ext
	return nil
}

var _ = replay.Replayer(&GlMatrixMult3x3fNV{}) // interface compliance check
func (ϟa *GlMatrixMult3x3fNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_330_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MatrixMode))
	ϟb.Push(ϟa.M.value())
	ϟb.Call(funcInfoGlMatrixMult3x3fNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_330_ext
	return nil
}

var _ = replay.Replayer(&GlMatrixMultTranspose3x3fNV{}) // interface compliance check
func (ϟa *GlMatrixMultTranspose3x3fNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_331_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.MatrixMode))
	ϟb.Push(ϟa.M.value())
	ϟb.Call(funcInfoGlMatrixMultTranspose3x3fNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_331_ext
	return nil
}

var _ = replay.Replayer(&GlMultiDrawArraysEXT{}) // interface compliance check
func (ϟa *GlMultiDrawArraysEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_332_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.First.value())
	ϟb.Push(ϟa.Count.value())
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMultiDrawArraysEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_332_ext
	return nil
}

var _ = replay.Replayer(&GlMultiDrawArraysIndirectEXT{}) // interface compliance check
func (ϟa *GlMultiDrawArraysIndirectEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_333_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Indirect.value())
	ϟb.Push(ϟa.Drawcount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Stride.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMultiDrawArraysIndirectEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_333_ext
	return nil
}

var _ = replay.Replayer(&GlMultiDrawElementsBaseVertexEXT{}) // interface compliance check
func (ϟa *GlMultiDrawElementsBaseVertexEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_334_ext := ExtensionId_GL_EXT_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value())
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Basevertex.value())
	ϟb.Call(funcInfoGlMultiDrawElementsBaseVertexEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_334_ext
	return nil
}

var _ = replay.Replayer(&GlMultiDrawElementsBaseVertexOES{}) // interface compliance check
func (ϟa *GlMultiDrawElementsBaseVertexOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_335_ext := ExtensionId_GL_OES_draw_elements_base_vertex // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value())
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Basevertex.value())
	ϟb.Call(funcInfoGlMultiDrawElementsBaseVertexOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_335_ext
	return nil
}

var _ = replay.Replayer(&GlMultiDrawElementsEXT{}) // interface compliance check
func (ϟa *GlMultiDrawElementsEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_336_ext := ExtensionId_GL_EXT_multi_draw_arrays // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(ϟa.Count.value())
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indices.value())
	ϟb.Push(ϟa.Primcount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMultiDrawElementsEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_336_ext
	return nil
}

var _ = replay.Replayer(&GlMultiDrawElementsIndirectEXT{}) // interface compliance check
func (ϟa *GlMultiDrawElementsIndirectEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_337_ext := ExtensionId_GL_EXT_multi_draw_indirect // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Indirect.value())
	ϟb.Push(ϟa.Drawcount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Stride.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlMultiDrawElementsIndirectEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_337_ext
	return nil
}

var _ = replay.Replayer(&GlPatchParameteriOES{}) // interface compliance check
func (ϟa *GlPatchParameteriOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_338_ext := ExtensionId_GL_OES_tessellation_shader // ExtensionId
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPatchParameteriOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_338_ext
	return nil
}

var _ = replay.Replayer(&GlPathCommandsNV{}) // interface compliance check
func (ϟa *GlPathCommandsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_339_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumCommands.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Commands.value())
	ϟb.Push(ϟa.NumCoords.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoordType))
	ϟb.Push(ϟa.Coords.value())
	ϟb.Call(funcInfoGlPathCommandsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_339_ext
	return nil
}

var _ = replay.Replayer(&GlPathCoordsNV{}) // interface compliance check
func (ϟa *GlPathCoordsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_340_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumCoords.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoordType))
	ϟb.Push(ϟa.Coords.value())
	ϟb.Call(funcInfoGlPathCoordsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_340_ext
	return nil
}

var _ = replay.Replayer(&GlPathCoverDepthFuncNV{}) // interface compliance check
func (ϟa *GlPathCoverDepthFuncNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_341_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.Func))
	ϟb.Call(funcInfoGlPathCoverDepthFuncNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_341_ext
	return nil
}

var _ = replay.Replayer(&GlPathDashArrayNV{}) // interface compliance check
func (ϟa *GlPathDashArrayNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_342_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DashCount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DashArray.value())
	ϟb.Call(funcInfoGlPathDashArrayNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_342_ext
	return nil
}

var _ = replay.Replayer(&GlPathGlyphIndexArrayNV{}) // interface compliance check
func (ϟa *GlPathGlyphIndexArrayNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_343_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.FirstPathName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FontTarget))
	ϟb.Push(ϟa.FontName.value())
	ϟb.Push(value.U32(ϟa.FontStyle))
	ϟb.Push(ϟa.FirstGlyphIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumGlyphs.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.PathParameterTemplate.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.EmScale.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathGlyphIndexArrayNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_343_ext
	return nil
}

var _ = replay.Replayer(&GlPathGlyphIndexRangeNV{}) // interface compliance check
func (ϟa *GlPathGlyphIndexRangeNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_344_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.FontTarget))
	ϟb.Push(ϟa.FontName.value())
	ϟb.Push(value.U32(ϟa.FontStyle))
	ϟb.Push(ϟa.PathParameterTemplate.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.EmScale.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BaseAndCount.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathGlyphIndexRangeNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_344_ext
	return nil
}

var _ = replay.Replayer(&GlPathGlyphRangeNV{}) // interface compliance check
func (ϟa *GlPathGlyphRangeNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_345_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.FirstPathName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FontTarget))
	ϟb.Push(ϟa.FontName.value())
	ϟb.Push(value.U32(ϟa.FontStyle))
	ϟb.Push(ϟa.FirstGlyph.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumGlyphs.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.HandleMissingGlyphs))
	ϟb.Push(ϟa.PathParameterTemplate.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.EmScale.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathGlyphRangeNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_345_ext
	return nil
}

var _ = replay.Replayer(&GlPathGlyphsNV{}) // interface compliance check
func (ϟa *GlPathGlyphsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_346_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.FirstPathName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FontTarget))
	ϟb.Push(ϟa.FontName.value())
	ϟb.Push(value.U32(ϟa.FontStyle))
	ϟb.Push(ϟa.NumGlyphs.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Charcodes.value())
	ϟb.Push(value.U32(ϟa.HandleMissingGlyphs))
	ϟb.Push(ϟa.PathParameterTemplate.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.EmScale.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathGlyphsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_346_ext
	return nil
}

var _ = replay.Replayer(&GlPathMemoryGlyphIndexArrayNV{}) // interface compliance check
func (ϟa *GlPathMemoryGlyphIndexArrayNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_347_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.FirstPathName.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FontTarget))
	ϟb.Push(ϟa.FontSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.FontData.value())
	ϟb.Push(ϟa.FaceIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.FirstGlyphIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumGlyphs.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.PathParameterTemplate.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.EmScale.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathMemoryGlyphIndexArrayNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_347_ext
	return nil
}

var _ = replay.Replayer(&GlPathParameterfNV{}) // interface compliance check
func (ϟa *GlPathParameterfNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_348_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathParameterfNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_348_ext
	return nil
}

var _ = replay.Replayer(&GlPathParameterfvNV{}) // interface compliance check
func (ϟa *GlPathParameterfvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_349_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlPathParameterfvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_349_ext
	return nil
}

var _ = replay.Replayer(&GlPathParameteriNV{}) // interface compliance check
func (ϟa *GlPathParameteriNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_350_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathParameteriNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_350_ext
	return nil
}

var _ = replay.Replayer(&GlPathParameterivNV{}) // interface compliance check
func (ϟa *GlPathParameterivNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_351_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlPathParameterivNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_351_ext
	return nil
}

var _ = replay.Replayer(&GlPathStencilDepthOffsetNV{}) // interface compliance check
func (ϟa *GlPathStencilDepthOffsetNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_352_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Factor.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Units.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathStencilDepthOffsetNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_352_ext
	return nil
}

var _ = replay.Replayer(&GlPathStencilFuncNV{}) // interface compliance check
func (ϟa *GlPathStencilFuncNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_353_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(value.U32(ϟa.Func))
	ϟb.Push(ϟa.Ref.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPathStencilFuncNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_353_ext
	return nil
}

var _ = replay.Replayer(&GlPathStringNV{}) // interface compliance check
func (ϟa *GlPathStringNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_354_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.PathString.value())
	ϟb.Call(funcInfoGlPathStringNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_354_ext
	return nil
}

var _ = replay.Replayer(&GlPathSubCommandsNV{}) // interface compliance check
func (ϟa *GlPathSubCommandsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_355_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CommandStart.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CommandsToDelete.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumCommands.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Commands.value())
	ϟb.Push(ϟa.NumCoords.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoordType))
	ϟb.Push(ϟa.Coords.value())
	ϟb.Call(funcInfoGlPathSubCommandsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_355_ext
	return nil
}

var _ = replay.Replayer(&GlPathSubCoordsNV{}) // interface compliance check
func (ϟa *GlPathSubCoordsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_356_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CoordStart.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumCoords.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoordType))
	ϟb.Push(ϟa.Coords.value())
	ϟb.Call(funcInfoGlPathSubCoordsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_356_ext
	return nil
}

var _ = replay.Replayer(&GlPointAlongPathNV{}) // interface compliance check
func (ϟa *GlPointAlongPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_357_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.StartSegment.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumSegments.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Distance.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value())
	ϟb.Push(ϟa.Y.value())
	ϟb.Push(ϟa.TangentX.value())
	ϟb.Push(ϟa.TangentY.value())
	ϟb.Call(funcInfoGlPointAlongPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_357_ext
	return nil
}

var _ = replay.Replayer(&GlPolygonModeNV{}) // interface compliance check
func (ϟa *GlPolygonModeNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_358_ext := ExtensionId_GL_NV_polygon_mode // ExtensionId
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlPolygonModeNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_358_ext
	return nil
}

var _ = replay.Replayer(&GlPopGroupMarkerEXT{}) // interface compliance check
func (ϟa *GlPopGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_359_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	ϟb.Call(funcInfoGlPopGroupMarkerEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_359_ext
	return nil
}

var _ = replay.Replayer(&GlPrimitiveBoundingBoxOES{}) // interface compliance check
func (ϟa *GlPrimitiveBoundingBoxOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_360_ext := ExtensionId_GL_OES_primitive_bounding_box // ExtensionId
	ϟb.Push(ϟa.MinX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MinY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MinZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MinW.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxZ.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.MaxW.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPrimitiveBoundingBoxOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_360_ext
	return nil
}

var _ = replay.Replayer(&GlProgramBinaryOES{}) // interface compliance check
func (ϟa *GlProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_361_ext := ExtensionId_GL_OES_get_program_binary // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.BinaryFormat))
	ϟb.Push(ϟa.Binary.value())
	ϟb.Push(ϟa.BinarySize.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramBinaryOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_361_ext
	return nil
}

var _ = replay.Replayer(&GlProgramParameteriEXT{}) // interface compliance check
func (ϟa *GlProgramParameteriEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_362_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramParameteriEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_362_ext
	return nil
}

var _ = replay.Replayer(&GlProgramPathFragmentInputGenNV{}) // interface compliance check
func (ϟa *GlProgramPathFragmentInputGenNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_363_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.GenMode))
	ϟb.Push(ϟa.Components.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Coeffs.value())
	ϟb.Call(funcInfoGlProgramPathFragmentInputGenNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_363_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1fEXT{}) // interface compliance check
func (ϟa *GlProgramUniform1fEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_364_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform1fEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_364_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniform1fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_365_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform1fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_365_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1iEXT{}) // interface compliance check
func (ϟa *GlProgramUniform1iEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_366_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform1iEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_366_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1ivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform1ivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_367_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform1ivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_367_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1uiEXT{}) // interface compliance check
func (ϟa *GlProgramUniform1uiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_368_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform1uiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_368_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1uivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform1uivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_369_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform1uivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_369_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2fEXT{}) // interface compliance check
func (ϟa *GlProgramUniform2fEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_370_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform2fEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_370_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniform2fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_371_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform2fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_371_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2iEXT{}) // interface compliance check
func (ϟa *GlProgramUniform2iEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_372_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform2iEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_372_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2ivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform2ivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_373_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform2ivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_373_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2uiEXT{}) // interface compliance check
func (ϟa *GlProgramUniform2uiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_374_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform2uiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_374_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2uivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform2uivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_375_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform2uivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_375_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3fEXT{}) // interface compliance check
func (ϟa *GlProgramUniform3fEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_376_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform3fEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_376_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniform3fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_377_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform3fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_377_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3iEXT{}) // interface compliance check
func (ϟa *GlProgramUniform3iEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_378_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform3iEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_378_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3ivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform3ivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_379_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform3ivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_379_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3uiEXT{}) // interface compliance check
func (ϟa *GlProgramUniform3uiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_380_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform3uiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_380_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3uivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform3uivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_381_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform3uivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_381_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4fEXT{}) // interface compliance check
func (ϟa *GlProgramUniform4fEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_382_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform4fEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_382_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniform4fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_383_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform4fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_383_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4iEXT{}) // interface compliance check
func (ϟa *GlProgramUniform4iEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_384_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform4iEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_384_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4ivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform4ivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_385_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform4ivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_385_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4uiEXT{}) // interface compliance check
func (ϟa *GlProgramUniform4uiEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_386_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform4uiEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_386_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4uivEXT{}) // interface compliance check
func (ϟa *GlProgramUniform4uivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_387_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform4uivEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_387_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformHandleui64NV{}) // interface compliance check
func (ϟa *GlProgramUniformHandleui64NV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_388_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniformHandleui64NV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_388_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformHandleui64vNV{}) // interface compliance check
func (ϟa *GlProgramUniformHandleui64vNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_389_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlProgramUniformHandleui64vNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_389_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix2fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix2fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_390_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix2fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_390_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix2x3fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix2x3fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_391_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix2x3fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_391_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix2x4fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix2x4fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_392_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix2x4fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_392_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix3fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix3fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_393_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix3fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_393_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix3x2fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix3x2fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_394_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix3x2fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_394_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix3x4fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix3x4fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_395_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix3x4fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_395_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix4fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix4fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_396_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix4fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_396_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix4x2fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix4x2fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_397_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix4x2fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_397_ext
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix4x3fvEXT{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix4x3fvEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_398_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix4x3fvEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_398_ext
	return nil
}

var _ = replay.Replayer(&GlPushGroupMarkerEXT{}) // interface compliance check
func (ϟa *GlPushGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_399_ext := ExtensionId_GL_EXT_debug_marker // ExtensionId
	if (ϟa.Length) > (GLsizei(int32(0))) {
		ϟa.Marker.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Length), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	} else {
		_ = strings.TrimRight(string(Charᵖ(ϟa.Marker).StringSlice(ϟs, ϟd, ϟl, true).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
	}
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Marker.value())
	ϟb.Call(funcInfoGlPushGroupMarkerEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_399_ext
	return nil
}

var _ = replay.Replayer(&GlQueryCounterEXT{}) // interface compliance check
func (ϟa *GlQueryCounterEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_400_ext := ExtensionId_GL_EXT_disjoint_timer_query // ExtensionId
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlQueryCounterEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_400_ext
	return nil
}

var _ = replay.Replayer(&GlReadBufferIndexedEXT{}) // interface compliance check
func (ϟa *GlReadBufferIndexedEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_401_ext := ExtensionId_GL_EXT_multiview_draw_buffers // ExtensionId
	ϟb.Push(value.U32(ϟa.Src))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlReadBufferIndexedEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_401_ext
	return nil
}

var _ = replay.Replayer(&GlReadBufferNV{}) // interface compliance check
func (ϟa *GlReadBufferNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_402_ext := ExtensionId_GL_NV_read_buffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlReadBufferNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_402_ext
	return nil
}

var _ = replay.Replayer(&GlReadnPixelsEXT{}) // interface compliance check
func (ϟa *GlReadnPixelsEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_403_ext := ExtensionId_GL_EXT_robustness // ExtensionId
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlReadnPixelsEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_403_ext
	return nil
}

var _ = replay.Replayer(&GlReadnPixelsKHR{}) // interface compliance check
func (ϟa *GlReadnPixelsKHR) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_404_ext := ExtensionId_GL_KHR_robustness // ExtensionId
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlReadnPixelsKHR)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_404_ext
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisampleANGLE{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisampleANGLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_405_ext := ExtensionId_GL_ANGLE_framebuffer_multisample // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRenderbufferStorageMultisampleANGLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_405_ext
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisampleAPPLE{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisampleAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_406_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRenderbufferStorageMultisampleAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_406_ext
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisampleEXT{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisampleEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_407_ext := ExtensionId_GL_EXT_multisampled_render_to_texture // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRenderbufferStorageMultisampleEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_407_ext
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisampleIMG{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisampleIMG) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_408_ext := ExtensionId_GL_IMG_multisampled_render_to_texture // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRenderbufferStorageMultisampleIMG)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_408_ext
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisampleNV{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisampleNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_409_ext := ExtensionId_GL_NV_framebuffer_multisample // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRenderbufferStorageMultisampleNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_409_ext
	return nil
}

var _ = replay.Replayer(&GlResolveMultisampleFramebufferAPPLE{}) // interface compliance check
func (ϟa *GlResolveMultisampleFramebufferAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_410_ext := ExtensionId_GL_APPLE_framebuffer_multisample // ExtensionId
	ϟb.Call(funcInfoGlResolveMultisampleFramebufferAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_410_ext
	return nil
}

var _ = replay.Replayer(&GlSamplerParameterIivOES{}) // interface compliance check
func (ϟa *GlSamplerParameterIivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_411_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value())
	ϟb.Call(funcInfoGlSamplerParameterIivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_411_ext
	return nil
}

var _ = replay.Replayer(&GlSamplerParameterIuivOES{}) // interface compliance check
func (ϟa *GlSamplerParameterIuivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_412_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value())
	ϟb.Call(funcInfoGlSamplerParameterIuivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_412_ext
	return nil
}

var _ = replay.Replayer(&GlScissorArrayvNV{}) // interface compliance check
func (ϟa *GlScissorArrayvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_413_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.First.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlScissorArrayvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_413_ext
	return nil
}

var _ = replay.Replayer(&GlScissorIndexedNV{}) // interface compliance check
func (ϟa *GlScissorIndexedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_414_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Left.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Bottom.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlScissorIndexedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_414_ext
	return nil
}

var _ = replay.Replayer(&GlScissorIndexedvNV{}) // interface compliance check
func (ϟa *GlScissorIndexedvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_415_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlScissorIndexedvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_415_ext
	return nil
}

var _ = replay.Replayer(&GlSelectPerfMonitorCountersAMD{}) // interface compliance check
func (ϟa *GlSelectPerfMonitorCountersAMD) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_416_ext := ExtensionId_GL_AMD_performance_monitor // ExtensionId
	ϟb.Push(ϟa.Monitor.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Enable.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Group.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumCounters.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.CounterList.value())
	ϟb.Call(funcInfoGlSelectPerfMonitorCountersAMD)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_416_ext
	return nil
}

var _ = replay.Replayer(&GlSetFenceNV{}) // interface compliance check
func (ϟa *GlSetFenceNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_417_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟb.Push(ϟa.Fence.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Condition))
	ϟb.Call(funcInfoGlSetFenceNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_417_ext
	return nil
}

var _ = replay.Replayer(&GlStartTilingQCOM{}) // interface compliance check
func (ϟa *GlStartTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_418_ext := ExtensionId_GL_QCOM_tiled_rendering // ExtensionId
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PreserveMask))
	ϟb.Call(funcInfoGlStartTilingQCOM)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_418_ext
	return nil
}

var _ = replay.Replayer(&GlStencilFillPathInstancedNV{}) // interface compliance check
func (ϟa *GlStencilFillPathInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_419_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FillMode))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.TransformValues.value())
	ϟb.Call(funcInfoGlStencilFillPathInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_419_ext
	return nil
}

var _ = replay.Replayer(&GlStencilFillPathNV{}) // interface compliance check
func (ϟa *GlStencilFillPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_420_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FillMode))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlStencilFillPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_420_ext
	return nil
}

var _ = replay.Replayer(&GlStencilStrokePathInstancedNV{}) // interface compliance check
func (ϟa *GlStencilStrokePathInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_421_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Reference.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.TransformValues.value())
	ϟb.Call(funcInfoGlStencilStrokePathInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_421_ext
	return nil
}

var _ = replay.Replayer(&GlStencilStrokePathNV{}) // interface compliance check
func (ϟa *GlStencilStrokePathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_422_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Reference.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlStencilStrokePathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_422_ext
	return nil
}

var _ = replay.Replayer(&GlStencilThenCoverFillPathInstancedNV{}) // interface compliance check
func (ϟa *GlStencilThenCoverFillPathInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_423_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FillMode))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.TransformValues.value())
	ϟb.Call(funcInfoGlStencilThenCoverFillPathInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_423_ext
	return nil
}

var _ = replay.Replayer(&GlStencilThenCoverFillPathNV{}) // interface compliance check
func (ϟa *GlStencilThenCoverFillPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_424_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.FillMode))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Call(funcInfoGlStencilThenCoverFillPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_424_ext
	return nil
}

var _ = replay.Replayer(&GlStencilThenCoverStrokePathInstancedNV{}) // interface compliance check
func (ϟa *GlStencilThenCoverStrokePathInstancedNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_425_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.PathNameType))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.PathBase.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Reference.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.TransformValues.value())
	ϟb.Call(funcInfoGlStencilThenCoverStrokePathInstancedNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_425_ext
	return nil
}

var _ = replay.Replayer(&GlStencilThenCoverStrokePathNV{}) // interface compliance check
func (ϟa *GlStencilThenCoverStrokePathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_426_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.Path.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Reference.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.CoverMode))
	ϟb.Call(funcInfoGlStencilThenCoverStrokePathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_426_ext
	return nil
}

var _ = replay.Replayer(&GlTestFenceNV{}) // interface compliance check
func (ϟa *GlTestFenceNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_427_ext := ExtensionId_GL_NV_fence // ExtensionId
	ϟb.Push(ϟa.Fence.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTestFenceNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_427_ext
	return nil
}

var _ = replay.Replayer(&GlTexBufferOES{}) // interface compliance check
func (ϟa *GlTexBufferOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_428_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Internalformat))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlTexBufferOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_428_ext
	return nil
}

var _ = replay.Replayer(&GlTexBufferRangeOES{}) // interface compliance check
func (ϟa *GlTexBufferRangeOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_429_ext := ExtensionId_GL_OES_texture_buffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Internalformat))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexBufferRangeOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_429_ext
	return nil
}

var _ = replay.Replayer(&GlTexImage3DOES{}) // interface compliance check
func (ϟa *GlTexImage3DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_430_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Border.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Pixels.value())
	ϟb.Call(funcInfoGlTexImage3DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_430_ext
	return nil
}

var _ = replay.Replayer(&GlTexPageCommitmentARB{}) // interface compliance check
func (ϟa *GlTexPageCommitmentARB) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_431_ext := ExtensionId_GL_EXT_sparse_texture // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Commit.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexPageCommitmentARB)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_431_ext
	return nil
}

var _ = replay.Replayer(&GlTexParameterIivOES{}) // interface compliance check
func (ϟa *GlTexParameterIivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_432_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlTexParameterIivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_432_ext
	return nil
}

var _ = replay.Replayer(&GlTexParameterIuivOES{}) // interface compliance check
func (ϟa *GlTexParameterIuivOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_433_ext := ExtensionId_GL_OES_texture_border_clamp // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlTexParameterIuivOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_433_ext
	return nil
}

var _ = replay.Replayer(&GlTexStorage1DEXT{}) // interface compliance check
func (ϟa *GlTexStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_434_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexStorage1DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_434_ext
	return nil
}

var _ = replay.Replayer(&GlTexStorage2DEXT{}) // interface compliance check
func (ϟa *GlTexStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_435_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexStorage2DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_435_ext
	return nil
}

var _ = replay.Replayer(&GlTexStorage3DEXT{}) // interface compliance check
func (ϟa *GlTexStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_436_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexStorage3DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_436_ext
	return nil
}

var _ = replay.Replayer(&GlTexSubImage3DOES{}) // interface compliance check
func (ϟa *GlTexSubImage3DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_437_ext := ExtensionId_GL_OES_texture_3D // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Pixels.value())
	ϟb.Call(funcInfoGlTexSubImage3DOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_437_ext
	return nil
}

var _ = replay.Replayer(&GlTextureStorage1DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_438_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTextureStorage1DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_438_ext
	return nil
}

var _ = replay.Replayer(&GlTextureStorage2DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_439_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTextureStorage2DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_439_ext
	return nil
}

var _ = replay.Replayer(&GlTextureStorage3DEXT{}) // interface compliance check
func (ϟa *GlTextureStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_440_ext := ExtensionId_GL_EXT_texture_storage // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTextureStorage3DEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_440_ext
	return nil
}

var _ = replay.Replayer(&GlTextureViewEXT{}) // interface compliance check
func (ϟa *GlTextureViewEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_441_ext := ExtensionId_GL_EXT_texture_view // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Origtexture.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Minlevel.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Numlevels.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Minlayer.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Numlayers.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTextureViewEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_441_ext
	return nil
}

var _ = replay.Replayer(&GlTextureViewOES{}) // interface compliance check
func (ϟa *GlTextureViewOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_442_ext := ExtensionId_GL_OES_texture_view // ExtensionId
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Origtexture.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Minlevel.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Numlevels.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Minlayer.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Numlayers.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTextureViewOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_442_ext
	return nil
}

var _ = replay.Replayer(&GlTransformPathNV{}) // interface compliance check
func (ϟa *GlTransformPathNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_443_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.ResultPath.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcPath.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.TransformType))
	ϟb.Push(ϟa.TransformValues.value())
	ϟb.Call(funcInfoGlTransformPathNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_443_ext
	return nil
}

var _ = replay.Replayer(&GlUniformHandleui64NV{}) // interface compliance check
func (ϟa *GlUniformHandleui64NV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_444_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniformHandleui64NV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_444_ext
	return nil
}

var _ = replay.Replayer(&GlUniformHandleui64vNV{}) // interface compliance check
func (ϟa *GlUniformHandleui64vNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_445_ext := ExtensionId_GL_NV_bindless_texture // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformHandleui64vNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_445_ext
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix2x3fvNV{}) // interface compliance check
func (ϟa *GlUniformMatrix2x3fvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_446_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix2x3fvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_446_ext
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix2x4fvNV{}) // interface compliance check
func (ϟa *GlUniformMatrix2x4fvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_447_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix2x4fvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_447_ext
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix3x2fvNV{}) // interface compliance check
func (ϟa *GlUniformMatrix3x2fvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_448_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix3x2fvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_448_ext
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix3x4fvNV{}) // interface compliance check
func (ϟa *GlUniformMatrix3x4fvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_449_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix3x4fvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_449_ext
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix4x2fvNV{}) // interface compliance check
func (ϟa *GlUniformMatrix4x2fvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_450_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix4x2fvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_450_ext
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix4x3fvNV{}) // interface compliance check
func (ϟa *GlUniformMatrix4x3fvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_451_ext := ExtensionId_GL_NV_non_square_matrices // ExtensionId
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix4x3fvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_451_ext
	return nil
}

var _ = replay.Replayer(&GlUnmapBufferOES{}) // interface compliance check
func (ϟa *GlUnmapBufferOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_452_ext := ExtensionId_GL_OES_mapbuffer // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlUnmapBufferOES)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_452_ext
	return nil
}

var _ = replay.Replayer(&GlUseProgramStagesEXT{}) // interface compliance check
func (ϟa *GlUseProgramStagesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_453_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Stages))
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlUseProgramStagesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_453_ext
	return nil
}

var _ = replay.Replayer(&GlValidateProgramPipelineEXT{}) // interface compliance check
func (ϟa *GlValidateProgramPipelineEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_454_ext := ExtensionId_GL_EXT_separate_shader_objects // ExtensionId
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlValidateProgramPipelineEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_454_ext
	return nil
}

var _ = replay.Replayer(&GlVertexAttribDivisorANGLE{}) // interface compliance check
func (ϟa *GlVertexAttribDivisorANGLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_455_ext := ExtensionId_GL_ANGLE_instanced_arrays // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Divisor.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribDivisorANGLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_455_ext
	return nil
}

var _ = replay.Replayer(&GlVertexAttribDivisorEXT{}) // interface compliance check
func (ϟa *GlVertexAttribDivisorEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_456_ext := ExtensionId_GL_EXT_instanced_arrays // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Divisor.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribDivisorEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_456_ext
	return nil
}

var _ = replay.Replayer(&GlVertexAttribDivisorNV{}) // interface compliance check
func (ϟa *GlVertexAttribDivisorNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_457_ext := ExtensionId_GL_NV_instanced_arrays // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Divisor.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribDivisorNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_457_ext
	return nil
}

var _ = replay.Replayer(&GlViewportArrayvNV{}) // interface compliance check
func (ϟa *GlViewportArrayvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_458_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.First.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlViewportArrayvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_458_ext
	return nil
}

var _ = replay.Replayer(&GlViewportIndexedfNV{}) // interface compliance check
func (ϟa *GlViewportIndexedfNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_459_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.W.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.H.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlViewportIndexedfNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_459_ext
	return nil
}

var _ = replay.Replayer(&GlViewportIndexedfvNV{}) // interface compliance check
func (ϟa *GlViewportIndexedfvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_460_ext := ExtensionId_GL_NV_viewport_array // ExtensionId
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlViewportIndexedfvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_460_ext
	return nil
}

var _ = replay.Replayer(&GlWaitSyncAPPLE{}) // interface compliance check
func (ϟa *GlWaitSyncAPPLE) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_461_ext := ExtensionId_GL_APPLE_sync // ExtensionId
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Flag))
	ϟb.Push(ϟa.Timeout.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlWaitSyncAPPLE)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_461_ext
	return nil
}

var _ = replay.Replayer(&GlWeightPathsNV{}) // interface compliance check
func (ϟa *GlWeightPathsNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_462_ext := ExtensionId_GL_NV_path_rendering // ExtensionId
	ϟb.Push(ϟa.ResultPath.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumPaths.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Paths.value())
	ϟb.Push(ϟa.Weights.value())
	ϟb.Call(funcInfoGlWeightPathsNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_462_ext
	return nil
}

var _ = replay.Replayer(&GlCoverageModulationNV{}) // interface compliance check
func (ϟa *GlCoverageModulationNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_463_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟb.Push(value.U32(ϟa.Components))
	ϟb.Call(funcInfoGlCoverageModulationNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_463_ext
	return nil
}

var _ = replay.Replayer(&GlCoverageModulationTableNV{}) // interface compliance check
func (ϟa *GlCoverageModulationTableNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_464_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlCoverageModulationTableNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_464_ext
	return nil
}

var _ = replay.Replayer(&GlFragmentCoverageColorNV{}) // interface compliance check
func (ϟa *GlFragmentCoverageColorNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_465_ext := ExtensionId_GL_NV_fragment_coverage_to_color // ExtensionId
	ϟb.Push(ϟa.Color.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFragmentCoverageColorNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_465_ext
	return nil
}

var _ = replay.Replayer(&GlFramebufferSampleLocationsfvNV{}) // interface compliance check
func (ϟa *GlFramebufferSampleLocationsfvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_466_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Start.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlFramebufferSampleLocationsfvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_466_ext
	return nil
}

var _ = replay.Replayer(&GlGetCoverageModulationTableNV{}) // interface compliance check
func (ϟa *GlGetCoverageModulationTableNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_467_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟb.Push(ϟa.Bufsize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlGetCoverageModulationTableNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_467_ext
	return nil
}

var _ = replay.Replayer(&GlNamedFramebufferSampleLocationsfvNV{}) // interface compliance check
func (ϟa *GlNamedFramebufferSampleLocationsfvNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_468_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	if key, remap := ϟa.Framebuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Start.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlNamedFramebufferSampleLocationsfvNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_468_ext
	return nil
}

var _ = replay.Replayer(&GlRasterSamplesEXT{}) // interface compliance check
func (ϟa *GlRasterSamplesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_469_ext := ExtensionId_GL_EXT_raster_multisample       // ExtensionId
	requiresExtension_470_ext := ExtensionId_GL_EXT_texture_filter_minmax    // ExtensionId
	requiresExtension_471_ext := ExtensionId_GL_NV_framebuffer_mixed_samples // ExtensionId
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Fixedsamplelocations.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRasterSamplesEXT)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _ = requiresExtension_469_ext, requiresExtension_470_ext, requiresExtension_471_ext
	return nil
}

var _ = replay.Replayer(&GlResolveDepthValuesNV{}) // interface compliance check
func (ϟa *GlResolveDepthValuesNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_472_ext := ExtensionId_GL_NV_sample_locations // ExtensionId
	ϟb.Call(funcInfoGlResolveDepthValuesNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_472_ext
	return nil
}

var _ = replay.Replayer(&GlSubpixelPrecisionBiasNV{}) // interface compliance check
func (ϟa *GlSubpixelPrecisionBiasNV) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	requiresExtension_473_ext := ExtensionId_GL_NV_conservative_raster // ExtensionId
	ϟb.Push(ϟa.Xbits.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Ybits.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlSubpixelPrecisionBiasNV)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_ = requiresExtension_473_ext
	return nil
}

var _ = replay.Replayer(&GlBlendColor{}) // interface compliance check
func (ϟa *GlBlendColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_474_major := uint32(2)    // u32
	minRequiredVersion_474_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_475_result := context             // Contextʳ
	ctx := GetContext_475_result                 // Contextʳ
	ctx.Blending.BlendColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.Red
		s.Green = ϟa.Green
		s.Blue = ϟa.Blue
		s.Alpha = ϟa.Alpha
		return s
	}()
	ϟb.Push(ϟa.Red.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Green.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Blue.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Alpha.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBlendColor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_474_major, minRequiredVersion_474_minor, context, GetContext_475_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendEquation{}) // interface compliance check
func (ϟa *GlBlendEquation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_476_major := uint32(2) // u32
	minRequiredVersion_476_minor := uint32(0) // u32
	switch ϟa.Equation {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_477_major := uint32(3) // u32
		minRequiredVersion_477_minor := uint32(0) // u32
		_, _ = minRequiredVersion_477_major, minRequiredVersion_477_minor
	default:
		v := ϟa.Equation
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_479_result := context             // Contextʳ
	ctx := GetContext_479_result                 // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Equation
	ctx.Blending.BlendEquationAlpha = ϟa.Equation
	ϟb.Push(value.U32(ϟa.Equation))
	ϟb.Call(funcInfoGlBlendEquation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_476_major, minRequiredVersion_476_minor, context, GetContext_479_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendEquationSeparate{}) // interface compliance check
func (ϟa *GlBlendEquationSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_480_major := uint32(2) // u32
	minRequiredVersion_480_minor := uint32(0) // u32
	switch ϟa.Rgb {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_481_major := uint32(3) // u32
		minRequiredVersion_481_minor := uint32(0) // u32
		_, _ = minRequiredVersion_481_major, minRequiredVersion_481_minor
	default:
		v := ϟa.Rgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Alpha {
	case GLenum_GL_FUNC_ADD, GLenum_GL_FUNC_REVERSE_SUBTRACT, GLenum_GL_FUNC_SUBTRACT:
	case GLenum_GL_MAX, GLenum_GL_MIN:
		minRequiredVersion_483_major := uint32(3) // u32
		minRequiredVersion_483_minor := uint32(0) // u32
		_, _ = minRequiredVersion_483_major, minRequiredVersion_483_minor
	default:
		v := ϟa.Alpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_485_result := context             // Contextʳ
	ctx := GetContext_485_result                 // Contextʳ
	ctx.Blending.BlendEquationRgb = ϟa.Rgb
	ctx.Blending.BlendEquationAlpha = ϟa.Alpha
	ϟb.Push(value.U32(ϟa.Rgb))
	ϟb.Push(value.U32(ϟa.Alpha))
	ϟb.Call(funcInfoGlBlendEquationSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_480_major, minRequiredVersion_480_minor, context, GetContext_485_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendFunc{}) // interface compliance check
func (ϟa *GlBlendFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_486_major := uint32(2) // u32
	minRequiredVersion_486_minor := uint32(0) // u32
	switch ϟa.SrcFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactor
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactor {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactor
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_489_result := context             // Contextʳ
	ctx := GetContext_489_result                 // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactor
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactor
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactor
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactor
	ϟb.Push(value.U32(ϟa.SrcFactor))
	ϟb.Push(value.U32(ϟa.DstFactor))
	ϟb.Call(funcInfoGlBlendFunc)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_486_major, minRequiredVersion_486_minor, context, GetContext_489_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlendFuncSeparate{}) // interface compliance check
func (ϟa *GlBlendFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_490_major := uint32(2) // u32
	minRequiredVersion_490_minor := uint32(0) // u32
	switch ϟa.SrcFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactorRgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactorRgb {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactorRgb
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.SrcFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_ALPHA_SATURATE, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.SrcFactorAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.DstFactorAlpha {
	case GLenum_GL_CONSTANT_ALPHA, GLenum_GL_CONSTANT_COLOR, GLenum_GL_DST_ALPHA, GLenum_GL_DST_COLOR, GLenum_GL_ONE, GLenum_GL_ONE_MINUS_CONSTANT_ALPHA, GLenum_GL_ONE_MINUS_CONSTANT_COLOR, GLenum_GL_ONE_MINUS_DST_ALPHA, GLenum_GL_ONE_MINUS_DST_COLOR, GLenum_GL_ONE_MINUS_SRC_ALPHA, GLenum_GL_ONE_MINUS_SRC_COLOR, GLenum_GL_SRC_ALPHA, GLenum_GL_SRC_COLOR, GLenum_GL_ZERO:
	default:
		v := ϟa.DstFactorAlpha
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_495_result := context             // Contextʳ
	ctx := GetContext_495_result                 // Contextʳ
	ctx.Blending.SrcRgbBlendFactor = ϟa.SrcFactorRgb
	ctx.Blending.DstRgbBlendFactor = ϟa.DstFactorRgb
	ctx.Blending.SrcAlphaBlendFactor = ϟa.SrcFactorAlpha
	ctx.Blending.DstAlphaBlendFactor = ϟa.DstFactorAlpha
	ϟb.Push(value.U32(ϟa.SrcFactorRgb))
	ϟb.Push(value.U32(ϟa.DstFactorRgb))
	ϟb.Push(value.U32(ϟa.SrcFactorAlpha))
	ϟb.Push(value.U32(ϟa.DstFactorAlpha))
	ϟb.Call(funcInfoGlBlendFuncSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_490_major, minRequiredVersion_490_minor, context, GetContext_495_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDepthFunc{}) // interface compliance check
func (ϟa *GlDepthFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_496_major := uint32(2) // u32
	minRequiredVersion_496_minor := uint32(0) // u32
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Function
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_498_result := context             // Contextʳ
	ctx := GetContext_498_result                 // Contextʳ
	ctx.Rasterizing.DepthTestFunction = ϟa.Function
	ϟb.Push(value.U32(ϟa.Function))
	ϟb.Call(funcInfoGlDepthFunc)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_496_major, minRequiredVersion_496_minor, context, GetContext_498_result, ctx
	return nil
}

var _ = replay.Replayer(&GlSampleCoverage{}) // interface compliance check
func (ϟa *GlSampleCoverage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_499_major := uint32(2)    // u32
	minRequiredVersion_499_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_500_result := context             // Contextʳ
	ctx := GetContext_500_result                 // Contextʳ
	ctx.Rasterizing.SampleCoverageValue = ϟa.Value
	ctx.Rasterizing.SampleCoverageInvert = ϟa.Invert
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.Bool(ϟa.Invert))
	ϟb.Call(funcInfoGlSampleCoverage)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_499_major, minRequiredVersion_499_minor, context, GetContext_500_result, ctx
	return nil
}

var _ = replay.Replayer(&GlSampleMaski{}) // interface compliance check
func (ϟa *GlSampleMaski) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_501_major := uint32(3) // u32
	minRequiredVersion_501_minor := uint32(1) // u32
	ϟb.Push(ϟa.MaskNumber.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Call(funcInfoGlSampleMaski)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_501_major, minRequiredVersion_501_minor
	return nil
}

var _ = replay.Replayer(&GlScissor{}) // interface compliance check
func (ϟa *GlScissor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_502_major := uint32(2)    // u32
	minRequiredVersion_502_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_503_result := context             // Contextʳ
	ctx := GetContext_503_result                 // Contextʳ
	ctx.Rasterizing.Scissor = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlScissor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_502_major, minRequiredVersion_502_minor, context, GetContext_503_result, ctx
	return nil
}

var _ = replay.Replayer(&GlStencilFunc{}) // interface compliance check
func (ϟa *GlStencilFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_504_major := uint32(2) // u32
	minRequiredVersion_504_minor := uint32(0) // u32
	switch ϟa.Func {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Func
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Func))
	ϟb.Push(ϟa.Ref.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlStencilFunc)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_504_major, minRequiredVersion_504_minor
	return nil
}

var _ = replay.Replayer(&GlStencilFuncSeparate{}) // interface compliance check
func (ϟa *GlStencilFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_506_major := uint32(2) // u32
	minRequiredVersion_506_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Function {
	case GLenum_GL_ALWAYS, GLenum_GL_EQUAL, GLenum_GL_GEQUAL, GLenum_GL_GREATER, GLenum_GL_LEQUAL, GLenum_GL_LESS, GLenum_GL_NEVER, GLenum_GL_NOTEQUAL:
	default:
		v := ϟa.Function
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.Function))
	ϟb.Push(ϟa.ReferenceValue.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlStencilFuncSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_506_major, minRequiredVersion_506_minor
	return nil
}

var _ = replay.Replayer(&GlStencilOp{}) // interface compliance check
func (ϟa *GlStencilOp) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_509_major := uint32(2) // u32
	minRequiredVersion_509_minor := uint32(0) // u32
	switch ϟa.Fail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.Fail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Zfail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.Zfail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Zpass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.Zpass
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Fail))
	ϟb.Push(value.U32(ϟa.Zfail))
	ϟb.Push(value.U32(ϟa.Zpass))
	ϟb.Call(funcInfoGlStencilOp)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_509_major, minRequiredVersion_509_minor
	return nil
}

var _ = replay.Replayer(&GlStencilOpSeparate{}) // interface compliance check
func (ϟa *GlStencilOpSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_513_major := uint32(2) // u32
	minRequiredVersion_513_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.StencilFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.StencilFail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.StencilPassDepthFail {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.StencilPassDepthFail
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.StencilPassDepthPass {
	case GLenum_GL_DECR, GLenum_GL_DECR_WRAP, GLenum_GL_INCR, GLenum_GL_INCR_WRAP, GLenum_GL_INVERT, GLenum_GL_KEEP, GLenum_GL_REPLACE, GLenum_GL_ZERO:
	default:
		v := ϟa.StencilPassDepthPass
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(value.U32(ϟa.StencilFail))
	ϟb.Push(value.U32(ϟa.StencilPassDepthFail))
	ϟb.Push(value.U32(ϟa.StencilPassDepthPass))
	ϟb.Call(funcInfoGlStencilOpSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_513_major, minRequiredVersion_513_minor
	return nil
}

var _ = replay.Replayer(&GlBindFramebuffer{}) // interface compliance check
func (ϟa *GlBindFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_518_major := uint32(2) // u32
	minRequiredVersion_518_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_519_major := uint32(3) // u32
		minRequiredVersion_519_minor := uint32(0) // u32
		_, _ = minRequiredVersion_519_major, minRequiredVersion_519_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_521_result := context             // Contextʳ
	ctx := GetContext_521_result                 // Contextʳ
	if !(ctx.Instances.Framebuffers.Contains(ϟa.Framebuffer)) {
		ctx.Instances.Framebuffers[ϟa.Framebuffer] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
	}
	if (ϟa.Target) == (GLenum_GL_FRAMEBUFFER) {
		ctx.BoundFramebuffers[GLenum_GL_READ_FRAMEBUFFER] = ϟa.Framebuffer
		ctx.BoundFramebuffers[GLenum_GL_DRAW_FRAMEBUFFER] = ϟa.Framebuffer
	} else {
		ctx.BoundFramebuffers[ϟa.Target] = ϟa.Framebuffer
	}
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Framebuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_518_major, minRequiredVersion_518_minor, context, GetContext_521_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindRenderbuffer{}) // interface compliance check
func (ϟa *GlBindRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_522_major := uint32(2) // u32
	minRequiredVersion_522_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_524_result := context             // Contextʳ
	ctx := GetContext_524_result                 // Contextʳ
	if !(ctx.Instances.Renderbuffers.Contains(ϟa.Renderbuffer)) {
		ctx.Instances.Renderbuffers[ϟa.Renderbuffer] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
	}
	ctx.BoundRenderbuffers[ϟa.Target] = ϟa.Renderbuffer
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindRenderbuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_522_major, minRequiredVersion_522_minor, context, GetContext_524_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBlitFramebuffer{}) // interface compliance check
func (ϟa *GlBlitFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_525_major := uint32(3)                                                                                                // u32
	minRequiredVersion_525_minor := uint32(0)                                                                                                // u32
	supportsBits_526_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_526_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	switch ϟa.Filter {
	case GLenum_GL_LINEAR, GLenum_GL_NEAREST:
	default:
		v := ϟa.Filter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.SrcX0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcX1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.SrcY1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstX1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.DstY1.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Push(value.U32(ϟa.Filter))
	ϟb.Call(funcInfoGlBlitFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_525_major, minRequiredVersion_525_minor, supportsBits_526_seenBits, supportsBits_526_validBits
	return nil
}

var _ = replay.Replayer(&GlCheckFramebufferStatus{}) // interface compliance check
func (ϟa *GlCheckFramebufferStatus) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_528_major := uint32(2) // u32
	minRequiredVersion_528_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_529_major := uint32(3) // u32
		minRequiredVersion_529_minor := uint32(0) // u32
		_, _ = minRequiredVersion_529_major, minRequiredVersion_529_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlCheckFramebufferStatus)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_528_major, minRequiredVersion_528_minor
	return nil
}

var _ = replay.Replayer(&GlClear{}) // interface compliance check
func (ϟa *GlClear) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_531_major := uint32(2)                                                                                                // u32
	minRequiredVersion_531_minor := uint32(0)                                                                                                // u32
	supportsBits_532_seenBits := ϟa.Mask                                                                                                     // GLbitfield
	supportsBits_532_validBits := (GLbitfield_GL_COLOR_BUFFER_BIT) | ((GLbitfield_GL_DEPTH_BUFFER_BIT) | (GLbitfield_GL_STENCIL_BUFFER_BIT)) // GLbitfield
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_DEPTH_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_STENCIL_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	if (GLbitfield_GL_COLOR_BUFFER_BIT)&(ϟa.Mask) != 0 {
	}
	ϟb.Push(value.U32(ϟa.Mask))
	ϟb.Call(funcInfoGlClear)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_531_major, minRequiredVersion_531_minor, supportsBits_532_seenBits, supportsBits_532_validBits
	return nil
}

var _ = replay.Replayer(&GlClearBufferfi{}) // interface compliance check
func (ϟa *GlClearBufferfi) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_533_major := uint32(3) // u32
	minRequiredVersion_533_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_DEPTH_STENCIL:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Buffer))
	ϟb.Push(ϟa.Drawbuffer.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Stencil.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlClearBufferfi)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_533_major, minRequiredVersion_533_minor
	return nil
}

var _ = replay.Replayer(&GlClearBufferfv{}) // interface compliance check
func (ϟa *GlClearBufferfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_535_major := uint32(3) // u32
	minRequiredVersion_535_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_DEPTH:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Buffer))
	ϟb.Push(ϟa.Drawbuffer.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlClearBufferfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_535_major, minRequiredVersion_535_minor
	return nil
}

var _ = replay.Replayer(&GlClearBufferiv{}) // interface compliance check
func (ϟa *GlClearBufferiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_537_major := uint32(3) // u32
	minRequiredVersion_537_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR, GLenum_GL_STENCIL:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Buffer))
	ϟb.Push(ϟa.Drawbuffer.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlClearBufferiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_537_major, minRequiredVersion_537_minor
	return nil
}

var _ = replay.Replayer(&GlClearBufferuiv{}) // interface compliance check
func (ϟa *GlClearBufferuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_539_major := uint32(3) // u32
	minRequiredVersion_539_minor := uint32(0) // u32
	switch ϟa.Buffer {
	case GLenum_GL_COLOR:
	default:
		v := ϟa.Buffer
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Buffer))
	ϟb.Push(ϟa.Drawbuffer.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlClearBufferuiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_539_major, minRequiredVersion_539_minor
	return nil
}

var _ = replay.Replayer(&GlClearColor{}) // interface compliance check
func (ϟa *GlClearColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_541_major := uint32(2)    // u32
	minRequiredVersion_541_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_542_result := context             // Contextʳ
	ctx := GetContext_542_result                 // Contextʳ
	ctx.Clearing.ClearColor = func() Color {
		s := Color{}
		s.Init()
		s.Red = ϟa.R
		s.Green = ϟa.G
		s.Blue = ϟa.B
		s.Alpha = ϟa.A
		return s
	}()
	ϟb.Push(ϟa.R.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.G.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.B.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.A.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlClearColor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_541_major, minRequiredVersion_541_minor, context, GetContext_542_result, ctx
	return nil
}

var _ = replay.Replayer(&GlClearDepthf{}) // interface compliance check
func (ϟa *GlClearDepthf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_543_major := uint32(2)    // u32
	minRequiredVersion_543_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_544_result := context             // Contextʳ
	ctx := GetContext_544_result                 // Contextʳ
	ctx.Clearing.ClearDepth = ϟa.Depth
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlClearDepthf)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_543_major, minRequiredVersion_543_minor, context, GetContext_544_result, ctx
	return nil
}

var _ = replay.Replayer(&GlClearStencil{}) // interface compliance check
func (ϟa *GlClearStencil) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_545_major := uint32(2)    // u32
	minRequiredVersion_545_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_546_result := context             // Contextʳ
	ctx := GetContext_546_result                 // Contextʳ
	ctx.Clearing.ClearStencil = ϟa.Stencil
	ϟb.Push(ϟa.Stencil.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlClearStencil)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_545_major, minRequiredVersion_545_minor, context, GetContext_546_result, ctx
	return nil
}

var _ = replay.Replayer(&GlColorMask{}) // interface compliance check
func (ϟa *GlColorMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_547_major := uint32(2)    // u32
	minRequiredVersion_547_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_548_result := context             // Contextʳ
	ctx := GetContext_548_result                 // Contextʳ
	ctx.Rasterizing.ColorMaskRed = ϟa.Red
	ctx.Rasterizing.ColorMaskGreen = ϟa.Green
	ctx.Rasterizing.ColorMaskBlue = ϟa.Blue
	ctx.Rasterizing.ColorMaskAlpha = ϟa.Alpha
	ϟb.Push(value.Bool(ϟa.Red))
	ϟb.Push(value.Bool(ϟa.Green))
	ϟb.Push(value.Bool(ϟa.Blue))
	ϟb.Push(value.Bool(ϟa.Alpha))
	ϟb.Call(funcInfoGlColorMask)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_547_major, minRequiredVersion_547_minor, context, GetContext_548_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDeleteFramebuffers{}) // interface compliance check
func (ϟa *GlDeleteFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_549_major := uint32(2)                                   // u32
	minRequiredVersion_549_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_550_result := context                                            // Contextʳ
	ctx := GetContext_550_result                                                // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Framebuffers, f.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Framebuffers.value())
	ϟb.Call(funcInfoGlDeleteFramebuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_549_major, minRequiredVersion_549_minor, f, context, GetContext_550_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDeleteRenderbuffers{}) // interface compliance check
func (ϟa *GlDeleteRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_551_major := uint32(2)                                    // u32
	minRequiredVersion_551_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	GetContext_552_result := context                                             // Contextʳ
	ctx := GetContext_552_result                                                 // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Renderbuffers, r.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Renderbuffers.value())
	ϟb.Call(funcInfoGlDeleteRenderbuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_551_major, minRequiredVersion_551_minor, r, context, GetContext_552_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDepthMask{}) // interface compliance check
func (ϟa *GlDepthMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_553_major := uint32(2)    // u32
	minRequiredVersion_553_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_554_result := context             // Contextʳ
	ctx := GetContext_554_result                 // Contextʳ
	ctx.Rasterizing.DepthMask = ϟa.Enabled
	ϟb.Push(value.Bool(ϟa.Enabled))
	ϟb.Call(funcInfoGlDepthMask)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_553_major, minRequiredVersion_553_minor, context, GetContext_554_result, ctx
	return nil
}

var _ = replay.Replayer(&GlFramebufferParameteri{}) // interface compliance check
func (ϟa *GlFramebufferParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_555_major := uint32(3) // u32
	minRequiredVersion_555_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferParameteri)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_555_major, minRequiredVersion_555_minor
	return nil
}

var _ = replay.Replayer(&GlFramebufferRenderbuffer{}) // interface compliance check
func (ϟa *GlFramebufferRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_558_major := uint32(2) // u32
	minRequiredVersion_558_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_559_major := uint32(3) // u32
		minRequiredVersion_559_minor := uint32(0) // u32
		_, _ = minRequiredVersion_559_major, minRequiredVersion_559_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_561_major := uint32(3) // u32
		minRequiredVersion_561_minor := uint32(0) // u32
		_, _ = minRequiredVersion_561_major, minRequiredVersion_561_minor
	default:
		v := ϟa.FramebufferAttachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.RenderbufferTarget {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.RenderbufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_564_result := context             // Contextʳ
	ctx := GetContext_564_result                 // Contextʳ
	target := func() (result GLenum) {
		switch ϟa.FramebufferTarget {
		case GLenum_GL_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_DRAW_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_READ_FRAMEBUFFER:
			return GLenum_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // GLenum
	framebufferId := ctx.BoundFramebuffers.Get(target)                  // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // Framebufferʳ
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Renderbuffer) == (RenderbufferId(uint32(0))) {
		attachment.Type = GLenum_GL_NONE
	} else {
		attachment.Type = GLenum_GL_RENDERBUFFER
	}
	attachment.Object = uint32(ϟa.Renderbuffer)
	attachment.TextureLevel = GLint(int32(0))
	attachment.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	framebuffer.Attachments[ϟa.FramebufferAttachment] = attachment
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.FramebufferAttachment))
	ϟb.Push(value.U32(ϟa.RenderbufferTarget))
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlFramebufferRenderbuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_558_major, minRequiredVersion_558_minor, context, GetContext_564_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}

var _ = replay.Replayer(&GlFramebufferTexture2D{}) // interface compliance check
func (ϟa *GlFramebufferTexture2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_565_major := uint32(2) // u32
	minRequiredVersion_565_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_566_major := uint32(3) // u32
		minRequiredVersion_566_minor := uint32(0) // u32
		_, _ = minRequiredVersion_566_major, minRequiredVersion_566_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.FramebufferAttachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_STENCIL_ATTACHMENT:
		minRequiredVersion_568_major := uint32(3) // u32
		minRequiredVersion_568_minor := uint32(0) // u32
		_, _ = minRequiredVersion_568_major, minRequiredVersion_568_minor
	default:
		v := ϟa.FramebufferAttachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.TextureTarget {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_570_major := uint32(3) // u32
		minRequiredVersion_570_minor := uint32(1) // u32
		_, _ = minRequiredVersion_570_major, minRequiredVersion_570_minor
	default:
		v := ϟa.TextureTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_572_result := context             // Contextʳ
	ctx := GetContext_572_result                 // Contextʳ
	target := func() (result GLenum) {
		switch ϟa.FramebufferTarget {
		case GLenum_GL_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_DRAW_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_READ_FRAMEBUFFER:
			return GLenum_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // GLenum
	framebufferId := ctx.BoundFramebuffers.Get(target)                  // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId)        // Framebufferʳ
	attachment := framebuffer.Attachments.Get(ϟa.FramebufferAttachment) // FramebufferAttachmentInfo
	if (ϟa.Texture) == (TextureId(uint32(0))) {
		attachment.Type = GLenum_GL_NONE
		attachment.Object = uint32(0)
		attachment.TextureLevel = GLint(int32(0))
		attachment.CubeMapFace = GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
	} else {
		attachment.Type = GLenum_GL_TEXTURE
		attachment.Object = uint32(ϟa.Texture)
		attachment.TextureLevel = ϟa.Level
		attachment.CubeMapFace = func() (result GLenum) {
			switch ϟa.TextureTarget {
			case GLenum_GL_TEXTURE_2D:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
			case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X
			case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y
			case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
				return GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z
			case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X:
				return GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X
			case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y:
				return GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y
			case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
				return GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.TextureTarget, ϟa))
				return result
			}
		}()
	}
	framebuffer.Attachments[ϟa.FramebufferAttachment] = attachment
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.FramebufferAttachment))
	ϟb.Push(value.U32(ϟa.TextureTarget))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTexture2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_565_major, minRequiredVersion_565_minor, context, GetContext_572_result, ctx, target, framebufferId, framebuffer, attachment
	return nil
}

var _ = replay.Replayer(&GlFramebufferTextureLayer{}) // interface compliance check
func (ϟa *GlFramebufferTextureLayer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_573_major := uint32(3) // u32
	minRequiredVersion_573_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	default:
		v := ϟa.Attachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Attachment))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Layer.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFramebufferTextureLayer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_573_major, minRequiredVersion_573_minor
	return nil
}

var _ = replay.Replayer(&GlGenFramebuffers{}) // interface compliance check
func (ϟa *GlGenFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_576_major := uint32(2)                                   // u32
	minRequiredVersion_576_minor := uint32(0)                                   // u32
	f := ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // FramebufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_577_result := context                                            // Contextʳ
	ctx := GetContext_577_result                                                // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Framebuffers.value())
	ϟb.Call(funcInfoGlGenFramebuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := FramebufferId(ϟa.Framebuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // FramebufferId
		ctx.Instances.Framebuffers[id] = func() *Framebuffer {
			s := &Framebuffer{}
			s.Init()
			return s
		}()
		f.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_576_major, minRequiredVersion_576_minor, f, context, GetContext_577_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenRenderbuffers{}) // interface compliance check
func (ϟa *GlGenRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_578_major := uint32(2)                                    // u32
	minRequiredVersion_578_minor := uint32(0)                                    // u32
	r := ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // RenderbufferIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                 // Contextʳ
	GetContext_579_result := context                                             // Contextʳ
	ctx := GetContext_579_result                                                 // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Renderbuffers.value())
	ϟb.Call(funcInfoGlGenRenderbuffers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := RenderbufferId(ϟa.Renderbuffers.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // RenderbufferId
		ctx.Instances.Renderbuffers[id] = func() *Renderbuffer {
			s := &Renderbuffer{}
			s.Init()
			return s
		}()
		r.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_578_major, minRequiredVersion_578_minor, r, context, GetContext_579_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetFramebufferAttachmentParameteriv{}) // interface compliance check
func (ϟa *GlGetFramebufferAttachmentParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_580_major := uint32(2) // u32
	minRequiredVersion_580_minor := uint32(0) // u32
	switch ϟa.FramebufferTarget {
	case GLenum_GL_FRAMEBUFFER:
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
		minRequiredVersion_581_major := uint32(3) // u32
		minRequiredVersion_581_minor := uint32(0) // u32
		_, _ = minRequiredVersion_581_major, minRequiredVersion_581_minor
	default:
		v := ϟa.FramebufferTarget
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Attachment {
	case GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_DEPTH_ATTACHMENT, GLenum_GL_STENCIL_ATTACHMENT:
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_DEPTH, GLenum_GL_DEPTH_STENCIL_ATTACHMENT, GLenum_GL_STENCIL:
		minRequiredVersion_583_major := uint32(3) // u32
		minRequiredVersion_583_minor := uint32(0) // u32
		_, _ = minRequiredVersion_583_major, minRequiredVersion_583_minor
	default:
		v := ϟa.Attachment
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME, GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
	case GLenum_GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING, GLenum_GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE, GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER:
		minRequiredVersion_585_major := uint32(3) // u32
		minRequiredVersion_585_minor := uint32(0) // u32
		_, _ = minRequiredVersion_585_major, minRequiredVersion_585_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_587_result := context             // Contextʳ
	ctx := GetContext_587_result                 // Contextʳ
	target := func() (result GLenum) {
		switch ϟa.FramebufferTarget {
		case GLenum_GL_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_DRAW_FRAMEBUFFER:
			return GLenum_GL_DRAW_FRAMEBUFFER
		case GLenum_GL_READ_FRAMEBUFFER:
			return GLenum_GL_READ_FRAMEBUFFER
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.FramebufferTarget, ϟa))
			return result
		}
	}() // GLenum
	framebufferId := ctx.BoundFramebuffers.Get(target)           // FramebufferId
	framebuffer := ctx.Instances.Framebuffers.Get(framebufferId) // Framebufferʳ
	a := framebuffer.Attachments.Get(ϟa.Attachment)              // FramebufferAttachmentInfo
	ϟb.Push(value.U32(ϟa.FramebufferTarget))
	ϟb.Push(value.U32(ϟa.Attachment))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetFramebufferAttachmentParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE:
			return GLint(a.Type)
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME:
			return GLint(a.Object)
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL:
			return a.TextureLevel
		case GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE:
			return GLint(a.CubeMapFace)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _ = minRequiredVersion_580_major, minRequiredVersion_580_minor, context, GetContext_587_result, ctx, target, framebufferId, framebuffer, a
	return nil
}

var _ = replay.Replayer(&GlGetFramebufferParameteriv{}) // interface compliance check
func (ϟa *GlGetFramebufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_588_major := uint32(3) // u32
	minRequiredVersion_588_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_DRAW_FRAMEBUFFER, GLenum_GL_FRAMEBUFFER, GLenum_GL_READ_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS, GLenum_GL_FRAMEBUFFER_DEFAULT_HEIGHT, GLenum_GL_FRAMEBUFFER_DEFAULT_SAMPLES, GLenum_GL_FRAMEBUFFER_DEFAULT_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetFramebufferParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_588_major, minRequiredVersion_588_minor
	return nil
}

var _ = replay.Replayer(&GlGetRenderbufferParameteriv{}) // interface compliance check
func (ϟa *GlGetRenderbufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_591_major := uint32(2) // u32
	minRequiredVersion_591_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_RENDERBUFFER_ALPHA_SIZE, GLenum_GL_RENDERBUFFER_BLUE_SIZE, GLenum_GL_RENDERBUFFER_DEPTH_SIZE, GLenum_GL_RENDERBUFFER_GREEN_SIZE, GLenum_GL_RENDERBUFFER_HEIGHT, GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT, GLenum_GL_RENDERBUFFER_RED_SIZE, GLenum_GL_RENDERBUFFER_STENCIL_SIZE, GLenum_GL_RENDERBUFFER_WIDTH:
	case GLenum_GL_RENDERBUFFER_SAMPLES:
		minRequiredVersion_593_major := uint32(3) // u32
		minRequiredVersion_593_minor := uint32(0) // u32
		_, _ = minRequiredVersion_593_major, minRequiredVersion_593_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_595_result := context             // Contextʳ
	ctx := GetContext_595_result                 // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetRenderbufferParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_RENDERBUFFER_WIDTH:
			return GLint(rb.Width)
		case GLenum_GL_RENDERBUFFER_HEIGHT:
			return GLint(rb.Height)
		case GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT:
			return GLint(rb.Format)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_591_major, minRequiredVersion_591_minor, context, GetContext_595_result, ctx, id, rb
	return nil
}

var _ = replay.Replayer(&GlInvalidateFramebuffer{}) // interface compliance check
func (ϟa *GlInvalidateFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_596_major := uint32(3) // u32
	minRequiredVersion_596_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Attachments.value())
	ϟb.Call(funcInfoGlInvalidateFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_596_major, minRequiredVersion_596_minor
	return nil
}

var _ = replay.Replayer(&GlInvalidateSubFramebuffer{}) // interface compliance check
func (ϟa *GlInvalidateSubFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_598_major := uint32(3) // u32
	minRequiredVersion_598_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_FRAMEBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.NumAttachments.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Attachments.value())
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlInvalidateSubFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_598_major, minRequiredVersion_598_minor
	return nil
}

var _ = replay.Replayer(&GlIsFramebuffer{}) // interface compliance check
func (ϟa *GlIsFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_600_major := uint32(2)    // u32
	minRequiredVersion_600_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_601_result := context             // Contextʳ
	ctx := GetContext_601_result                 // Contextʳ
	if key, remap := ϟa.Framebuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Framebuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsFramebuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_600_major, minRequiredVersion_600_minor, context, GetContext_601_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsRenderbuffer{}) // interface compliance check
func (ϟa *GlIsRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_602_major := uint32(2)    // u32
	minRequiredVersion_602_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_603_result := context             // Contextʳ
	ctx := GetContext_603_result                 // Contextʳ
	if key, remap := ϟa.Renderbuffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Renderbuffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsRenderbuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_602_major, minRequiredVersion_602_minor, context, GetContext_603_result, ctx
	return nil
}

var _ = replay.Replayer(&GlReadBuffer{}) // interface compliance check
func (ϟa *GlReadBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_604_major := uint32(3) // u32
	minRequiredVersion_604_minor := uint32(0) // u32
	switch ϟa.Src {
	case GLenum_GL_BACK, GLenum_GL_COLOR_ATTACHMENT0, GLenum_GL_COLOR_ATTACHMENT1, GLenum_GL_COLOR_ATTACHMENT10, GLenum_GL_COLOR_ATTACHMENT11, GLenum_GL_COLOR_ATTACHMENT12, GLenum_GL_COLOR_ATTACHMENT13, GLenum_GL_COLOR_ATTACHMENT14, GLenum_GL_COLOR_ATTACHMENT15, GLenum_GL_COLOR_ATTACHMENT2, GLenum_GL_COLOR_ATTACHMENT3, GLenum_GL_COLOR_ATTACHMENT4, GLenum_GL_COLOR_ATTACHMENT5, GLenum_GL_COLOR_ATTACHMENT6, GLenum_GL_COLOR_ATTACHMENT7, GLenum_GL_COLOR_ATTACHMENT8, GLenum_GL_COLOR_ATTACHMENT9, GLenum_GL_NONE:
	default:
		v := ϟa.Src
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Src))
	ϟb.Call(funcInfoGlReadBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_604_major, minRequiredVersion_604_minor
	return nil
}

var _ = replay.Replayer(&GlReadPixels{}) // interface compliance check
func (ϟa *GlReadPixels) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_606_major := uint32(2) // u32
	minRequiredVersion_606_minor := uint32(0) // u32
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_607_major := uint32(3) // u32
		minRequiredVersion_607_minor := uint32(0) // u32
		_, _ = minRequiredVersion_607_major, minRequiredVersion_607_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV:
		minRequiredVersion_609_major := uint32(3) // u32
		minRequiredVersion_609_minor := uint32(0) // u32
		_, _ = minRequiredVersion_609_major, minRequiredVersion_609_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlReadPixels)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Data.Slice(uint64(uint32(0)), uint64(externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_606_major, minRequiredVersion_606_minor
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorage{}) // interface compliance check
func (ϟa *GlRenderbufferStorage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_611_major := uint32(2) // u32
	minRequiredVersion_611_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGBA4, GLenum_GL_STENCIL_INDEX8:
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_613_major := uint32(3) // u32
		minRequiredVersion_613_minor := uint32(0) // u32
		_, _ = minRequiredVersion_613_major, minRequiredVersion_613_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_615_result := context             // Contextʳ
	ctx := GetContext_615_result                 // Contextʳ
	id := ctx.BoundRenderbuffers.Get(ϟa.Target)  // RenderbufferId
	rb := ctx.Instances.Renderbuffers.Get(id)    // Renderbufferʳ
	rb.Format = ϟa.Format
	rb.Width = ϟa.Width
	rb.Height = ϟa.Height
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRenderbufferStorage)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_611_major, minRequiredVersion_611_minor, context, GetContext_615_result, ctx, id, rb
	return nil
}

var _ = replay.Replayer(&GlRenderbufferStorageMultisample{}) // interface compliance check
func (ϟa *GlRenderbufferStorageMultisample) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_616_major := uint32(3) // u32
	minRequiredVersion_616_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8_ALPHA8, GLenum_GL_STENCIL_INDEX8:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlRenderbufferStorageMultisample)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_616_major, minRequiredVersion_616_minor
	return nil
}

var _ = replay.Replayer(&GlStencilMask{}) // interface compliance check
func (ϟa *GlStencilMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_619_major := uint32(2)    // u32
	minRequiredVersion_619_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_620_result := context             // Contextʳ
	ctx := GetContext_620_result                 // Contextʳ
	ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
	ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlStencilMask)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_619_major, minRequiredVersion_619_minor, context, GetContext_620_result, ctx
	return nil
}

var _ = replay.Replayer(&GlStencilMaskSeparate{}) // interface compliance check
func (ϟa *GlStencilMaskSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_621_major := uint32(2) // u32
	minRequiredVersion_621_minor := uint32(0) // u32
	switch ϟa.Face {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_623_result := context             // Contextʳ
	ctx := GetContext_623_result                 // Contextʳ
	switch ϟa.Face {
	case GLenum_GL_FRONT:
		ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
	case GLenum_GL_BACK:
		ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	case GLenum_GL_FRONT_AND_BACK:
		ctx.Rasterizing.StencilMask[GLenum_GL_FRONT] = ϟa.Mask
		ctx.Rasterizing.StencilMask[GLenum_GL_BACK] = ϟa.Mask
	default:
		v := ϟa.Face
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Face))
	ϟb.Push(ϟa.Mask.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlStencilMaskSeparate)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_621_major, minRequiredVersion_621_minor, context, GetContext_623_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDisable{}) // interface compliance check
func (ϟa *GlDisable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_624_major := uint32(2) // u32
	minRequiredVersion_624_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_625_major := uint32(3) // u32
		minRequiredVersion_625_minor := uint32(0) // u32
		_, _ = minRequiredVersion_625_major, minRequiredVersion_625_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_626_major := uint32(3) // u32
		minRequiredVersion_626_minor := uint32(1) // u32
		_, _ = minRequiredVersion_626_major, minRequiredVersion_626_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_628_result := context             // Contextʳ
	ctx := GetContext_628_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = false
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.Call(funcInfoGlDisable)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_624_major, minRequiredVersion_624_minor, context, GetContext_628_result, ctx
	return nil
}

var _ = replay.Replayer(&GlEnable{}) // interface compliance check
func (ϟa *GlEnable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_629_major := uint32(2) // u32
	minRequiredVersion_629_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_630_major := uint32(3) // u32
		minRequiredVersion_630_minor := uint32(0) // u32
		_, _ = minRequiredVersion_630_major, minRequiredVersion_630_minor
	case GLenum_GL_SAMPLE_MASK:
		minRequiredVersion_631_major := uint32(3) // u32
		minRequiredVersion_631_minor := uint32(1) // u32
		_, _ = minRequiredVersion_631_major, minRequiredVersion_631_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_633_result := context             // Contextʳ
	ctx := GetContext_633_result                 // Contextʳ
	ctx.Capabilities[ϟa.Capability] = true
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.Call(funcInfoGlEnable)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_629_major, minRequiredVersion_629_minor, context, GetContext_633_result, ctx
	return nil
}

var _ = replay.Replayer(&GlFinish{}) // interface compliance check
func (ϟa *GlFinish) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_634_major := uint32(2) // u32
	minRequiredVersion_634_minor := uint32(0) // u32
	ϟb.Call(funcInfoGlFinish)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_634_major, minRequiredVersion_634_minor
	return nil
}

var _ = replay.Replayer(&GlFlush{}) // interface compliance check
func (ϟa *GlFlush) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_635_major := uint32(2) // u32
	minRequiredVersion_635_minor := uint32(0) // u32
	ϟb.Call(funcInfoGlFlush)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_635_major, minRequiredVersion_635_minor
	return nil
}

var _ = replay.Replayer(&GlFlushMappedBufferRange{}) // interface compliance check
func (ϟa *GlFlushMappedBufferRange) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_636_major := uint32(3) // u32
	minRequiredVersion_636_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ARRAY_BUFFER, GLenum_GL_COPY_READ_BUFFER, GLenum_GL_COPY_WRITE_BUFFER, GLenum_GL_ELEMENT_ARRAY_BUFFER, GLenum_GL_PIXEL_PACK_BUFFER, GLenum_GL_PIXEL_UNPACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM_BUFFER:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlFlushMappedBufferRange)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_636_major, minRequiredVersion_636_minor
	return nil
}

var _ = replay.Replayer(&GlGetError{}) // interface compliance check
func (ϟa *GlGetError) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_638_major := uint32(2) // u32
	minRequiredVersion_638_minor := uint32(0) // u32
	ϟb.Call(funcInfoGlGetError)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_638_major, minRequiredVersion_638_minor
	return nil
}

var _ = replay.Replayer(&GlHint{}) // interface compliance check
func (ϟa *GlHint) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_639_major := uint32(2) // u32
	minRequiredVersion_639_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_GENERATE_MIPMAP_HINT:
	case GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT:
		minRequiredVersion_640_major := uint32(3) // u32
		minRequiredVersion_640_minor := uint32(0) // u32
		_, _ = minRequiredVersion_640_major, minRequiredVersion_640_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Mode {
	case GLenum_GL_DONT_CARE, GLenum_GL_FASTEST, GLenum_GL_NICEST:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_643_result := context             // Contextʳ
	ctx := GetContext_643_result                 // Contextʳ
	ctx.GenerateMipmapHint = ϟa.Mode
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlHint)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_639_major, minRequiredVersion_639_minor, context, GetContext_643_result, ctx
	return nil
}

var _ = replay.Replayer(&GlActiveShaderProgram{}) // interface compliance check
func (ϟa *GlActiveShaderProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_644_major := uint32(3) // u32
	minRequiredVersion_644_minor := uint32(1) // u32
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlActiveShaderProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_644_major, minRequiredVersion_644_minor
	return nil
}

var _ = replay.Replayer(&GlAttachShader{}) // interface compliance check
func (ϟa *GlAttachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_645_major := uint32(2)    // u32
	minRequiredVersion_645_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_646_result := context             // Contextʳ
	ctx := GetContext_646_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	p.Shaders[s.Type] = ϟa.Shader
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlAttachShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_645_major, minRequiredVersion_645_minor, context, GetContext_646_result, ctx, p, s
	return nil
}

var _ = replay.Replayer(&GlBindAttribLocation{}) // interface compliance check
func (ϟa *GlBindAttribLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_647_major := uint32(2)    // u32
	minRequiredVersion_647_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_648_result := context             // Contextʳ
	ctx := GetContext_648_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	p.AttributeBindings[ϟa.Name] = ϟa.Location
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.Call(funcInfoGlBindAttribLocation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_647_major, minRequiredVersion_647_minor, context, GetContext_648_result, ctx, p
	return nil
}

var _ = replay.Replayer(&GlBindProgramPipeline{}) // interface compliance check
func (ϟa *GlBindProgramPipeline) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_649_major := uint32(3) // u32
	minRequiredVersion_649_minor := uint32(1) // u32
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBindProgramPipeline)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_649_major, minRequiredVersion_649_minor
	return nil
}

var _ = replay.Replayer(&GlCompileShader{}) // interface compliance check
func (ϟa *GlCompileShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_650_major := uint32(2) // u32
	minRequiredVersion_650_minor := uint32(0) // u32
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlCompileShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_650_major, minRequiredVersion_650_minor
	return nil
}

var _ = replay.Replayer(&GlCreateProgram{}) // interface compliance check
func (ϟa *GlCreateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_651_major := uint32(2)    // u32
	minRequiredVersion_651_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_652_result := context             // Contextʳ
	ctx := GetContext_652_result                 // Contextʳ
	ϟb.Call(funcInfoGlCreateProgram)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		ptr, found := ϟb.Remappings[key]
		if !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Remappings[key] = ptr
		}
		ϟb.Clone(0)
		ϟb.Store(ptr)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ProgramId(ϟa.Result) // ProgramId
	ctx.Instances.Programs[id] = func() *Program {
		s := &Program{}
		s.Init()
		return s
	}()
	_, _, _, _, _, _ = minRequiredVersion_651_major, minRequiredVersion_651_minor, context, GetContext_652_result, ctx, id
	return nil
}

var _ = replay.Replayer(&GlCreateShader{}) // interface compliance check
func (ϟa *GlCreateShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_653_major := uint32(2) // u32
	minRequiredVersion_653_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	case GLenum_GL_COMPUTE_SHADER:
		minRequiredVersion_654_major := uint32(3) // u32
		minRequiredVersion_654_minor := uint32(1) // u32
		_, _ = minRequiredVersion_654_major, minRequiredVersion_654_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_656_result := context             // Contextʳ
	ctx := GetContext_656_result                 // Contextʳ
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Call(funcInfoGlCreateShader)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		ptr, found := ϟb.Remappings[key]
		if !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Remappings[key] = ptr
		}
		ϟb.Clone(0)
		ϟb.Store(ptr)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	id := ShaderId(ϟa.Result) // ShaderId
	ctx.Instances.Shaders[id] = func() *Shader {
		s := &Shader{}
		s.Init()
		return s
	}()
	s := ctx.Instances.Shaders.Get(id) // Shaderʳ
	s.Type = ϟa.Type
	_, _, _, _, _, _, _ = minRequiredVersion_653_major, minRequiredVersion_653_minor, context, GetContext_656_result, ctx, id, s
	return nil
}

var _ = replay.Replayer(&GlCreateShaderProgramv{}) // interface compliance check
func (ϟa *GlCreateShaderProgramv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_657_major := uint32(3) // u32
	minRequiredVersion_657_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Strings.value())
	ϟb.Call(funcInfoGlCreateShaderProgramv)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		ptr, found := ϟb.Remappings[key]
		if !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Remappings[key] = ptr
		}
		ϟb.Clone(0)
		ϟb.Store(ptr)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_657_major, minRequiredVersion_657_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteProgram{}) // interface compliance check
func (ϟa *GlDeleteProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_659_major := uint32(2)    // u32
	minRequiredVersion_659_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_660_result := context             // Contextʳ
	ctx := GetContext_660_result                 // Contextʳ
	delete(ctx.Instances.Programs, ϟa.Program)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDeleteProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_659_major, minRequiredVersion_659_minor, context, GetContext_660_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDeleteProgramPipelines{}) // interface compliance check
func (ϟa *GlDeleteProgramPipelines) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_661_major := uint32(3) // u32
	minRequiredVersion_661_minor := uint32(1) // u32
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Pipelines.value())
	ϟb.Call(funcInfoGlDeleteProgramPipelines)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_661_major, minRequiredVersion_661_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteShader{}) // interface compliance check
func (ϟa *GlDeleteShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_662_major := uint32(2)    // u32
	minRequiredVersion_662_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_663_result := context             // Contextʳ
	ctx := GetContext_663_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	s.Deletable = true
	delete(ctx.Instances.Shaders, ϟa.Shader)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDeleteShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_662_major, minRequiredVersion_662_minor, context, GetContext_663_result, ctx, s
	return nil
}

var _ = replay.Replayer(&GlDetachShader{}) // interface compliance check
func (ϟa *GlDetachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_664_major := uint32(2)    // u32
	minRequiredVersion_664_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_665_result := context             // Contextʳ
	ctx := GetContext_665_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	delete(p.Shaders, s.Type)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDetachShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_664_major, minRequiredVersion_664_minor, context, GetContext_665_result, ctx, p, s
	return nil
}

var _ = replay.Replayer(&GlDispatchCompute{}) // interface compliance check
func (ϟa *GlDispatchCompute) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_666_major := uint32(3) // u32
	minRequiredVersion_666_minor := uint32(1) // u32
	ϟb.Push(ϟa.NumGroupsX.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumGroupsY.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.NumGroupsZ.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDispatchCompute)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_666_major, minRequiredVersion_666_minor
	return nil
}

var _ = replay.Replayer(&GlDispatchComputeIndirect{}) // interface compliance check
func (ϟa *GlDispatchComputeIndirect) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_667_major := uint32(3) // u32
	minRequiredVersion_667_minor := uint32(1) // u32
	ϟb.Push(ϟa.Indirect.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDispatchComputeIndirect)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_667_major, minRequiredVersion_667_minor
	return nil
}

var _ = replay.Replayer(&GlGenProgramPipelines{}) // interface compliance check
func (ϟa *GlGenProgramPipelines) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_668_major := uint32(3) // u32
	minRequiredVersion_668_minor := uint32(1) // u32
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Pipelines.value())
	ϟb.Call(funcInfoGlGenProgramPipelines)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_668_major, minRequiredVersion_668_minor
	return nil
}

var _ = replay.Replayer(&GlGetActiveAttrib{}) // interface compliance check
func (ϟa *GlGetActiveAttrib) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_669_major := uint32(2) // u32
	minRequiredVersion_669_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufferSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufferBytesWritten.value())
	ϟb.Push(ϟa.VectorCount.value())
	ϟb.Push(ϟa.Type.value())
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetActiveAttrib)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.BufferBytesWritten) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // GLsizei
		ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Name.Slice(uint64(0), uint64(256), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_669_major, minRequiredVersion_669_minor
	return nil
}

var _ = replay.Replayer(&GlGetActiveUniform{}) // interface compliance check
func (ϟa *GlGetActiveUniform) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_670_major := uint32(2) // u32
	minRequiredVersion_670_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufferSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufferBytesWritten.value())
	ϟb.Push(ϟa.VectorCount.value())
	ϟb.Push(ϟa.Type.value())
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetActiveUniform)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	if (ϟa.BufferBytesWritten) != (GLsizeiᵖ{}) {
		l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // GLsizei
		ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = l
	} else {
		ϟa.Name.Slice(uint64(0), uint64(256), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLint(ϟa.VectorCount.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLenum(ϟa.Type.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_670_major, minRequiredVersion_670_minor
	return nil
}

var _ = replay.Replayer(&GlGetActiveUniformBlockName{}) // interface compliance check
func (ϟa *GlGetActiveUniformBlockName) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_671_major := uint32(3) // u32
	minRequiredVersion_671_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.UniformBlockIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufferSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufferBytesWritten.value())
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetActiveUniformBlockName)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	l := GLsizei(ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // GLsizei
	ϟa.BufferBytesWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(l, ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Name.Slice(uint64(GLsizei(int32(0))), uint64(l), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _ = minRequiredVersion_671_major, minRequiredVersion_671_minor, l
	return nil
}

var _ = replay.Replayer(&GlGetActiveUniformBlockiv{}) // interface compliance check
func (ϟa *GlGetActiveUniformBlockiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_672_major := uint32(3) // u32
	minRequiredVersion_672_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS, GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES, GLenum_GL_UNIFORM_BLOCK_BINDING, GLenum_GL_UNIFORM_BLOCK_DATA_SIZE, GLenum_GL_UNIFORM_BLOCK_NAME_LENGTH, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER, GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER:
	default:
		v := ϟa.ParameterName
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.UniformBlockIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.ParameterName))
	ϟb.Push(ϟa.Parameters.value())
	ϟb.Call(funcInfoGlGetActiveUniformBlockiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Parameters.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_672_major, minRequiredVersion_672_minor
	return nil
}

var _ = replay.Replayer(&GlGetActiveUniformsiv{}) // interface compliance check
func (ϟa *GlGetActiveUniformsiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_674_major := uint32(3) // u32
	minRequiredVersion_674_minor := uint32(0) // u32
	switch ϟa.ParameterName {
	case GLenum_GL_UNIFORM_ARRAY_STRIDE, GLenum_GL_UNIFORM_BLOCK_INDEX, GLenum_GL_UNIFORM_IS_ROW_MAJOR, GLenum_GL_UNIFORM_MATRIX_STRIDE, GLenum_GL_UNIFORM_NAME_LENGTH, GLenum_GL_UNIFORM_OFFSET, GLenum_GL_UNIFORM_SIZE, GLenum_GL_UNIFORM_TYPE:
	default:
		v := ϟa.ParameterName
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟa.UniformIndices.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.UniformCount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.UniformIndices.value())
	ϟb.Push(value.U32(ϟa.ParameterName))
	ϟb.Push(ϟa.Parameters.value())
	ϟb.Call(funcInfoGlGetActiveUniformsiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Parameters.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.UniformCount), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_674_major, minRequiredVersion_674_minor
	return nil
}

var _ = replay.Replayer(&GlGetAttachedShaders{}) // interface compliance check
func (ϟa *GlGetAttachedShaders) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_676_major := uint32(2)    // u32
	minRequiredVersion_676_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_677_result := context             // Contextʳ
	ctx := GetContext_677_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_678_a := int32(ϟa.BufferLength)          // s32
	min_678_b := int32(len(p.Shaders))           // s32
	min_678_result := func() (result int32) {
		switch (min_678_a) < (min_678_b) {
		case true:
			return min_678_a
		case false:
			return min_678_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_678_a) < (min_678_b), ϟa))
			return result
		}
	}() // s32
	l := min_678_result // s32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufferLength.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.ShadersLengthWritten.value())
	ϟb.Push(ϟa.Shaders.value())
	ϟb.Call(funcInfoGlGetAttachedShaders)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.ShadersLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_676_major, minRequiredVersion_676_minor, context, GetContext_677_result, ctx, p, min_678_a, min_678_b, min_678_result, l
	return nil
}

var _ = replay.Replayer(&GlGetAttribLocation{}) // interface compliance check
func (ϟa *GlGetAttribLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_679_major := uint32(2) // u32
	minRequiredVersion_679_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.Call(funcInfoGlGetAttribLocation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_679_major, minRequiredVersion_679_minor
	return nil
}

var _ = replay.Replayer(&GlGetFragDataLocation{}) // interface compliance check
func (ϟa *GlGetFragDataLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_680_major := uint32(3) // u32
	minRequiredVersion_680_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetFragDataLocation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_680_major, minRequiredVersion_680_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramBinary{}) // interface compliance check
func (ϟa *GlGetProgramBinary) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_681_major := uint32(3) // u32
	minRequiredVersion_681_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.BinaryFormat.value())
	ϟb.Push(ϟa.Binary.value())
	ϟb.Call(funcInfoGlGetProgramBinary)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_681_major, minRequiredVersion_681_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramInfoLog{}) // interface compliance check
func (ϟa *GlGetProgramInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_682_major := uint32(2)    // u32
	minRequiredVersion_682_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_683_result := context             // Contextʳ
	ctx := GetContext_683_result                 // Contextʳ
	p := ctx.Instances.Programs.Get(ϟa.Program)  // Programʳ
	min_684_a := int32(ϟa.BufferLength)          // s32
	min_684_b := int32(p.InfoLog.Count)          // s32
	min_684_result := func() (result int32) {
		switch (min_684_a) < (min_684_b) {
		case true:
			return min_684_a
		case false:
			return min_684_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_684_a) < (min_684_b), ϟa))
			return result
		}
	}() // s32
	l := min_684_result // s32
	ϟdst, ϟsrc := ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(p.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufferLength.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.StringLengthWritten.value())
	ϟb.Push(ϟa.Info.value())
	ϟb.Call(funcInfoGlGetProgramInfoLog)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_682_major, minRequiredVersion_682_minor, context, GetContext_683_result, ctx, p, min_684_a, min_684_b, min_684_result, l
	return nil
}

var _ = replay.Replayer(&GlGetProgramInterfaceiv{}) // interface compliance check
func (ϟa *GlGetProgramInterfaceiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_685_major := uint32(3) // u32
	minRequiredVersion_685_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_RESOURCES, GLenum_GL_MAX_NAME_LENGTH, GLenum_GL_MAX_NUM_ACTIVE_VARIABLES:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.ProgramInterface))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetProgramInterfaceiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_685_major, minRequiredVersion_685_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramPipelineInfoLog{}) // interface compliance check
func (ϟa *GlGetProgramPipelineInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_688_major := uint32(3) // u32
	minRequiredVersion_688_minor := uint32(1) // u32
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.InfoLog.value())
	ϟb.Call(funcInfoGlGetProgramPipelineInfoLog)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_688_major, minRequiredVersion_688_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramPipelineiv{}) // interface compliance check
func (ϟa *GlGetProgramPipelineiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_689_major := uint32(3) // u32
	minRequiredVersion_689_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_PROGRAM, GLenum_GL_COMPUTE_SHADER, GLenum_GL_FRAGMENT_SHADER, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_VALIDATE_STATUS, GLenum_GL_VERTEX_SHADER:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetProgramPipelineiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_689_major, minRequiredVersion_689_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramResourceIndex{}) // interface compliance check
func (ϟa *GlGetProgramResourceIndex) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_691_major := uint32(3) // u32
	minRequiredVersion_691_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.ProgramInterface))
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetProgramResourceIndex)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_691_major, minRequiredVersion_691_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramResourceLocation{}) // interface compliance check
func (ϟa *GlGetProgramResourceLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_693_major := uint32(3) // u32
	minRequiredVersion_693_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_UNIFORM:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.ProgramInterface))
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetProgramResourceLocation)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_693_major, minRequiredVersion_693_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramResourceName{}) // interface compliance check
func (ϟa *GlGetProgramResourceName) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_695_major := uint32(3) // u32
	minRequiredVersion_695_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.ProgramInterface))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetProgramResourceName)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_695_major, minRequiredVersion_695_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramResourceiv{}) // interface compliance check
func (ϟa *GlGetProgramResourceiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_697_major := uint32(3) // u32
	minRequiredVersion_697_minor := uint32(1) // u32
	switch ϟa.ProgramInterface {
	case GLenum_GL_ATOMIC_COUNTER_BUFFER, GLenum_GL_BUFFER_VARIABLE, GLenum_GL_PROGRAM_INPUT, GLenum_GL_PROGRAM_OUTPUT, GLenum_GL_SHADER_STORAGE_BLOCK, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER, GLenum_GL_TRANSFORM_FEEDBACK_VARYING, GLenum_GL_UNIFORM, GLenum_GL_UNIFORM_BLOCK:
	default:
		v := ϟa.ProgramInterface
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.ProgramInterface))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.PropCount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Props.value())
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetProgramResourceiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_697_major, minRequiredVersion_697_minor
	return nil
}

var _ = replay.Replayer(&GlGetProgramiv{}) // interface compliance check
func (ϟa *GlGetProgramiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_699_major := uint32(2) // u32
	minRequiredVersion_699_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_ACTIVE_ATTRIBUTES, GLenum_GL_ACTIVE_ATTRIBUTE_MAX_LENGTH, GLenum_GL_ACTIVE_UNIFORMS, GLenum_GL_ACTIVE_UNIFORM_MAX_LENGTH, GLenum_GL_ATTACHED_SHADERS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_LINK_STATUS, GLenum_GL_VALIDATE_STATUS:
	case GLenum_GL_ACTIVE_UNIFORM_BLOCKS, GLenum_GL_ACTIVE_UNIFORM_BLOCK_MAX_NAME_LENGTH, GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_MODE, GLenum_GL_TRANSFORM_FEEDBACK_VARYINGS, GLenum_GL_TRANSFORM_FEEDBACK_VARYING_MAX_LENGTH:
		minRequiredVersion_700_major := uint32(3) // u32
		minRequiredVersion_700_minor := uint32(0) // u32
		_, _ = minRequiredVersion_700_major, minRequiredVersion_700_minor
	case GLenum_GL_ACTIVE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_701_major := uint32(3) // u32
		minRequiredVersion_701_minor := uint32(1) // u32
		_, _ = minRequiredVersion_701_major, minRequiredVersion_701_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetProgramiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_699_major, minRequiredVersion_699_minor
	return nil
}

var _ = replay.Replayer(&GlGetShaderInfoLog{}) // interface compliance check
func (ϟa *GlGetShaderInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_703_major := uint32(2)    // u32
	minRequiredVersion_703_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_704_result := context             // Contextʳ
	ctx := GetContext_704_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_705_a := int32(ϟa.BufferLength)          // s32
	min_705_b := int32(s.InfoLog.Count)          // s32
	min_705_result := func() (result int32) {
		switch (min_705_a) < (min_705_b) {
		case true:
			return min_705_a
		case false:
			return min_705_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_705_a) < (min_705_b), ϟa))
			return result
		}
	}() // s32
	l := min_705_result // s32
	ϟdst, ϟsrc := ϟa.Info.Slice(uint64(int32(0)), uint64(l), ϟs).Copy(s.InfoLog.Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufferLength.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.StringLengthWritten.value())
	ϟb.Push(ϟa.Info.value())
	ϟb.Call(funcInfoGlGetShaderInfoLog)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_703_major, minRequiredVersion_703_minor, context, GetContext_704_result, ctx, s, min_705_a, min_705_b, min_705_result, l
	return nil
}

var _ = replay.Replayer(&GlGetShaderPrecisionFormat{}) // interface compliance check
func (ϟa *GlGetShaderPrecisionFormat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_706_major := uint32(2) // u32
	minRequiredVersion_706_minor := uint32(0) // u32
	switch ϟa.ShaderType {
	case GLenum_GL_FRAGMENT_SHADER, GLenum_GL_VERTEX_SHADER:
	default:
		v := ϟa.ShaderType
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.PrecisionType {
	case GLenum_GL_HIGH_FLOAT, GLenum_GL_HIGH_INT, GLenum_GL_LOW_FLOAT, GLenum_GL_LOW_INT, GLenum_GL_MEDIUM_FLOAT, GLenum_GL_MEDIUM_INT:
	default:
		v := ϟa.PrecisionType
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.ShaderType))
	ϟb.Push(value.U32(ϟa.PrecisionType))
	ϟb.Push(ϟa.Range.value())
	ϟb.Push(ϟa.Precision.value())
	ϟb.Call(funcInfoGlGetShaderPrecisionFormat)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Range.Slice(uint64(0), uint64(2), ϟs).OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Precision.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _ = minRequiredVersion_706_major, minRequiredVersion_706_minor
	return nil
}

var _ = replay.Replayer(&GlGetShaderSource{}) // interface compliance check
func (ϟa *GlGetShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_709_major := uint32(2)    // u32
	minRequiredVersion_709_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_710_result := context             // Contextʳ
	ctx := GetContext_710_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	min_711_a := int32(ϟa.BufferLength)          // s32
	min_711_b := int32(len(s.Source))            // s32
	min_711_result := func() (result int32) {
		switch (min_711_a) < (min_711_b) {
		case true:
			return min_711_a
		case false:
			return min_711_b
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", (min_711_a) < (min_711_b), ϟa))
			return result
		}
	}() // s32
	l := min_711_result // s32
	ϟdst, ϟsrc := Charᵖ(ϟa.Source).Slice(uint64(int32(0)), uint64(l), ϟs).Copy(MakeCharˢFromString(s.Source, ϟs).Slice(uint64(int32(0)), uint64(l), ϟs), ϟs, ϟd, ϟl)
	ϟsrc.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.BufferLength.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.StringLengthWritten.value())
	ϟb.Push(ϟa.Source.value())
	ϟb.Call(funcInfoGlGetShaderSource)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟdst.OnWrite(ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟa.StringLengthWritten.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(GLsizei(l), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _, _, _, _ = minRequiredVersion_709_major, minRequiredVersion_709_minor, context, GetContext_710_result, ctx, s, min_711_a, min_711_b, min_711_result, l
	return nil
}

var _ = replay.Replayer(&GlGetShaderiv{}) // interface compliance check
func (ϟa *GlGetShaderiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_712_major := uint32(2) // u32
	minRequiredVersion_712_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_COMPILE_STATUS, GLenum_GL_DELETE_STATUS, GLenum_GL_INFO_LOG_LENGTH, GLenum_GL_SHADER_SOURCE_LENGTH, GLenum_GL_SHADER_TYPE:
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_714_result := context             // Contextʳ
	ctx := GetContext_714_result                 // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)    // Shaderʳ
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetShaderiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_SHADER_TYPE:
			return GLint(s.Type)
		case GLenum_GL_DELETE_STATUS:
			return func() (result GLint) {
				switch s.Deletable {
				case true:
					return GLint(int32(1))
				case false:
					return GLint(int32(0))
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", s.Deletable, ϟa))
					return result
				}
			}()
		case GLenum_GL_COMPILE_STATUS:
			return func() (result GLint) {
				switch s.Compiled {
				case true:
					return GLint(int32(1))
				case false:
					return GLint(int32(0))
				default:
					// TODO: better unmatched handling
					panic(fmt.Errorf("Unmatched switch(%v) in atom %T", s.Compiled, ϟa))
					return result
				}
			}()
		case GLenum_GL_INFO_LOG_LENGTH:
			return GLint(int32(s.InfoLog.Count))
		case GLenum_GL_SHADER_SOURCE_LENGTH:
			return GLint(int32(len(s.Source)))
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _ = minRequiredVersion_712_major, minRequiredVersion_712_minor, context, GetContext_714_result, ctx, s
	return nil
}

var _ = replay.Replayer(&GlGetUniformBlockIndex{}) // interface compliance check
func (ϟa *GlGetUniformBlockIndex) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_715_major := uint32(3) // u32
	minRequiredVersion_715_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.UniformBlockName.value())
	ϟb.Call(funcInfoGlGetUniformBlockIndex)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_715_major, minRequiredVersion_715_minor
	return nil
}

var _ = replay.Replayer(&GlGetUniformIndices{}) // interface compliance check
func (ϟa *GlGetUniformIndices) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_716_major := uint32(3) // u32
	minRequiredVersion_716_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.UniformCount.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.UniformNames.value())
	ϟb.Push(ϟa.UniformIndices.value())
	ϟb.Call(funcInfoGlGetUniformIndices)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_716_major, minRequiredVersion_716_minor
	return nil
}

var _ = replay.Replayer(&GlGetUniformLocation{}) // interface compliance check
func (ϟa *GlGetUniformLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_717_major := uint32(2) // u32
	minRequiredVersion_717_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟb.String(ϟa.Name))
	ϟb.Call(funcInfoGlGetUniformLocation)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		ptr, found := ϟb.Remappings[key]
		if !found {
			ptr = ϟb.AllocateMemory(uint64(4))
			ϟb.Remappings[key] = ptr
		}
		ϟb.Clone(0)
		ϟb.Store(ptr)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_717_major, minRequiredVersion_717_minor
	return nil
}

var _ = replay.Replayer(&GlGetUniformfv{}) // interface compliance check
func (ϟa *GlGetUniformfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_718_major := uint32(2) // u32
	minRequiredVersion_718_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetUniformfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_718_major, minRequiredVersion_718_minor
	return nil
}

var _ = replay.Replayer(&GlGetUniformiv{}) // interface compliance check
func (ϟa *GlGetUniformiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_719_major := uint32(2) // u32
	minRequiredVersion_719_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetUniformiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_719_major, minRequiredVersion_719_minor
	return nil
}

var _ = replay.Replayer(&GlGetUniformuiv{}) // interface compliance check
func (ϟa *GlGetUniformuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_720_major := uint32(3) // u32
	minRequiredVersion_720_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetUniformuiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_720_major, minRequiredVersion_720_minor
	return nil
}

var _ = replay.Replayer(&GlIsProgram{}) // interface compliance check
func (ϟa *GlIsProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_721_major := uint32(2)    // u32
	minRequiredVersion_721_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_722_result := context             // Contextʳ
	ctx := GetContext_722_result                 // Contextʳ
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_721_major, minRequiredVersion_721_minor, context, GetContext_722_result, ctx
	return nil
}

var _ = replay.Replayer(&GlIsProgramPipeline{}) // interface compliance check
func (ϟa *GlIsProgramPipeline) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_723_major := uint32(3) // u32
	minRequiredVersion_723_minor := uint32(1) // u32
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsProgramPipeline)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_723_major, minRequiredVersion_723_minor
	return nil
}

var _ = replay.Replayer(&GlIsShader{}) // interface compliance check
func (ϟa *GlIsShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_724_major := uint32(2)    // u32
	minRequiredVersion_724_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_725_result := context             // Contextʳ
	ctx := GetContext_725_result                 // Contextʳ
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsShader)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_724_major, minRequiredVersion_724_minor, context, GetContext_725_result, ctx
	return nil
}

var _ = replay.Replayer(&GlLinkProgram{}) // interface compliance check
func (ϟa *GlLinkProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_726_major := uint32(2) // u32
	minRequiredVersion_726_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlLinkProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_726_major, minRequiredVersion_726_minor
	return nil
}

var _ = replay.Replayer(&GlMemoryBarrier{}) // interface compliance check
func (ϟa *GlMemoryBarrier) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_727_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_727_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_728_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_728_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
	if (GLbitfield_GL_ALL_BARRIER_BITS)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_COMMAND_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_UNIFORM_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	ϟb.Push(value.U32(ϟa.Barriers))
	ϟb.Call(funcInfoGlMemoryBarrier)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_727_major, minRequiredVersion_727_minor, supportsBits_728_seenBits, supportsBits_728_validBits
	return nil
}

var _ = replay.Replayer(&GlMemoryBarrierByRegion{}) // interface compliance check
func (ϟa *GlMemoryBarrierByRegion) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_729_major := uint32(3)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	minRequiredVersion_729_minor := uint32(1)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // u32
	supportsBits_730_seenBits := ϟa.Barriers                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                // GLbitfield
	supportsBits_730_validBits := (GLbitfield_GL_ALL_BARRIER_BITS) | ((GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT) | ((GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_COMMAND_BARRIER_BIT) | ((GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT) | ((GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT) | ((GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT) | ((GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT) | ((GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT) | ((GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT) | ((GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT) | ((GLbitfield_GL_UNIFORM_BARRIER_BIT) | (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT))))))))))))) // GLbitfield
	if (GLbitfield_GL_ALL_BARRIER_BITS)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ATOMIC_COUNTER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_BUFFER_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_COMMAND_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_ELEMENT_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_FRAMEBUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_PIXEL_BUFFER_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_IMAGE_ACCESS_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_SHADER_STORAGE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_FETCH_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TEXTURE_UPDATE_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_TRANSFORM_FEEDBACK_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_UNIFORM_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	if (GLbitfield_GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT)&(ϟa.Barriers) != 0 {
	}
	ϟb.Push(value.U32(ϟa.Barriers))
	ϟb.Call(funcInfoGlMemoryBarrierByRegion)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_729_major, minRequiredVersion_729_minor, supportsBits_730_seenBits, supportsBits_730_validBits
	return nil
}

var _ = replay.Replayer(&GlProgramBinary{}) // interface compliance check
func (ϟa *GlProgramBinary) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_731_major := uint32(3) // u32
	minRequiredVersion_731_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		v := ϟa.BinaryFormat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.BinaryFormat))
	ϟb.Push(ϟa.Binary.value())
	ϟb.Push(ϟa.Length.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramBinary)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_731_major, minRequiredVersion_731_minor
	return nil
}

var _ = replay.Replayer(&GlProgramParameteri{}) // interface compliance check
func (ϟa *GlProgramParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_733_major := uint32(3) // u32
	minRequiredVersion_733_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_PROGRAM_BINARY_RETRIEVABLE_HINT:
	case GLenum_GL_PROGRAM_SEPARABLE:
		minRequiredVersion_734_major := uint32(3) // u32
		minRequiredVersion_734_minor := uint32(1) // u32
		_, _ = minRequiredVersion_734_major, minRequiredVersion_734_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramParameteri)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_733_major, minRequiredVersion_733_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1f{}) // interface compliance check
func (ϟa *GlProgramUniform1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_736_major := uint32(3) // u32
	minRequiredVersion_736_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform1f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_736_major, minRequiredVersion_736_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1fv{}) // interface compliance check
func (ϟa *GlProgramUniform1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_737_major := uint32(3) // u32
	minRequiredVersion_737_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform1fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_737_major, minRequiredVersion_737_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1i{}) // interface compliance check
func (ϟa *GlProgramUniform1i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_738_major := uint32(3) // u32
	minRequiredVersion_738_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform1i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_738_major, minRequiredVersion_738_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1iv{}) // interface compliance check
func (ϟa *GlProgramUniform1iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_739_major := uint32(3) // u32
	minRequiredVersion_739_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform1iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_739_major, minRequiredVersion_739_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1ui{}) // interface compliance check
func (ϟa *GlProgramUniform1ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_740_major := uint32(3) // u32
	minRequiredVersion_740_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform1ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_740_major, minRequiredVersion_740_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform1uiv{}) // interface compliance check
func (ϟa *GlProgramUniform1uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_741_major := uint32(3) // u32
	minRequiredVersion_741_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform1uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_741_major, minRequiredVersion_741_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2f{}) // interface compliance check
func (ϟa *GlProgramUniform2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_742_major := uint32(3) // u32
	minRequiredVersion_742_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform2f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_742_major, minRequiredVersion_742_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2fv{}) // interface compliance check
func (ϟa *GlProgramUniform2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_743_major := uint32(3) // u32
	minRequiredVersion_743_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_743_major, minRequiredVersion_743_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2i{}) // interface compliance check
func (ϟa *GlProgramUniform2i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_744_major := uint32(3) // u32
	minRequiredVersion_744_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform2i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_744_major, minRequiredVersion_744_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2iv{}) // interface compliance check
func (ϟa *GlProgramUniform2iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_745_major := uint32(3) // u32
	minRequiredVersion_745_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform2iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_745_major, minRequiredVersion_745_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2ui{}) // interface compliance check
func (ϟa *GlProgramUniform2ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_746_major := uint32(3) // u32
	minRequiredVersion_746_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform2ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_746_major, minRequiredVersion_746_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform2uiv{}) // interface compliance check
func (ϟa *GlProgramUniform2uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_747_major := uint32(3) // u32
	minRequiredVersion_747_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform2uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_747_major, minRequiredVersion_747_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3f{}) // interface compliance check
func (ϟa *GlProgramUniform3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_748_major := uint32(3) // u32
	minRequiredVersion_748_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform3f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_748_major, minRequiredVersion_748_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3fv{}) // interface compliance check
func (ϟa *GlProgramUniform3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_749_major := uint32(3) // u32
	minRequiredVersion_749_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_749_major, minRequiredVersion_749_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3i{}) // interface compliance check
func (ϟa *GlProgramUniform3i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_750_major := uint32(3) // u32
	minRequiredVersion_750_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform3i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_750_major, minRequiredVersion_750_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3iv{}) // interface compliance check
func (ϟa *GlProgramUniform3iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_751_major := uint32(3) // u32
	minRequiredVersion_751_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform3iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_751_major, minRequiredVersion_751_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3ui{}) // interface compliance check
func (ϟa *GlProgramUniform3ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_752_major := uint32(3) // u32
	minRequiredVersion_752_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform3ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_752_major, minRequiredVersion_752_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform3uiv{}) // interface compliance check
func (ϟa *GlProgramUniform3uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_753_major := uint32(3) // u32
	minRequiredVersion_753_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform3uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_753_major, minRequiredVersion_753_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4f{}) // interface compliance check
func (ϟa *GlProgramUniform4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_754_major := uint32(3) // u32
	minRequiredVersion_754_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform4f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_754_major, minRequiredVersion_754_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4fv{}) // interface compliance check
func (ϟa *GlProgramUniform4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_755_major := uint32(3) // u32
	minRequiredVersion_755_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_755_major, minRequiredVersion_755_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4i{}) // interface compliance check
func (ϟa *GlProgramUniform4i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_756_major := uint32(3) // u32
	minRequiredVersion_756_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform4i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_756_major, minRequiredVersion_756_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4iv{}) // interface compliance check
func (ϟa *GlProgramUniform4iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_757_major := uint32(3) // u32
	minRequiredVersion_757_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform4iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_757_major, minRequiredVersion_757_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4ui{}) // interface compliance check
func (ϟa *GlProgramUniform4ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_758_major := uint32(3) // u32
	minRequiredVersion_758_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlProgramUniform4ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_758_major, minRequiredVersion_758_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniform4uiv{}) // interface compliance check
func (ϟa *GlProgramUniform4uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_759_major := uint32(3) // u32
	minRequiredVersion_759_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniform4uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_759_major, minRequiredVersion_759_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix2fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_760_major := uint32(3) // u32
	minRequiredVersion_760_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_760_major, minRequiredVersion_760_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix2x3fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix2x3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_761_major := uint32(3) // u32
	minRequiredVersion_761_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix2x3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_761_major, minRequiredVersion_761_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix2x4fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix2x4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_762_major := uint32(3) // u32
	minRequiredVersion_762_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix2x4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_762_major, minRequiredVersion_762_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix3fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_763_major := uint32(3) // u32
	minRequiredVersion_763_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_763_major, minRequiredVersion_763_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix3x2fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix3x2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_764_major := uint32(3) // u32
	minRequiredVersion_764_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix3x2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_764_major, minRequiredVersion_764_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix3x4fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix3x4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_765_major := uint32(3) // u32
	minRequiredVersion_765_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix3x4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_765_major, minRequiredVersion_765_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix4fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_766_major := uint32(3) // u32
	minRequiredVersion_766_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_766_major, minRequiredVersion_766_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix4x2fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix4x2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_767_major := uint32(3) // u32
	minRequiredVersion_767_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix4x2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_767_major, minRequiredVersion_767_minor
	return nil
}

var _ = replay.Replayer(&GlProgramUniformMatrix4x3fv{}) // interface compliance check
func (ϟa *GlProgramUniformMatrix4x3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_768_major := uint32(3) // u32
	minRequiredVersion_768_minor := uint32(1) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlProgramUniformMatrix4x3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_768_major, minRequiredVersion_768_minor
	return nil
}

var _ = replay.Replayer(&GlReleaseShaderCompiler{}) // interface compliance check
func (ϟa *GlReleaseShaderCompiler) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_769_major := uint32(2) // u32
	minRequiredVersion_769_minor := uint32(0) // u32
	ϟb.Call(funcInfoGlReleaseShaderCompiler)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_769_major, minRequiredVersion_769_minor
	return nil
}

var _ = replay.Replayer(&GlShaderBinary{}) // interface compliance check
func (ϟa *GlShaderBinary) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_770_major := uint32(2) // u32
	minRequiredVersion_770_minor := uint32(0) // u32
	switch ϟa.BinaryFormat {
	default:
		v := ϟa.BinaryFormat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Shaders.value())
	ϟb.Push(value.U32(ϟa.BinaryFormat))
	ϟb.Push(ϟa.Binary.value())
	ϟb.Push(ϟa.BinarySize.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlShaderBinary)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_770_major, minRequiredVersion_770_minor
	return nil
}

var _ = replay.Replayer(&GlShaderSource{}) // interface compliance check
func (ϟa *GlShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_772_major := uint32(2)                                   // u32
	minRequiredVersion_772_minor := uint32(0)                                   // u32
	sources := ϟa.Source.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLcharᶜᵖˢ
	lengths := ϟa.Length.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                // Contextʳ
	GetContext_773_result := context                                            // Contextʳ
	ctx := GetContext_773_result                                                // Contextʳ
	s := ctx.Instances.Shaders.Get(ϟa.Shader)                                   // Shaderʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		str := func() (result string) {
			switch ((ϟa.Length) == (GLintᶜᵖ{})) || ((lengths.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)) < (GLint(int32(0)))) {
			case true:
				return strings.TrimRight(string(Charᵖ(sources.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)).StringSlice(ϟs, ϟd, ϟl, true).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), "\x00")
			case false:
				return string(Charᵖ(sources.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)).Slice(uint64(GLint(int32(0))), uint64(lengths.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
			default:
				// TODO: better unmatched handling
				panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ((ϟa.Length) == (GLintᶜᵖ{})) || ((lengths.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb)) < (GLint(int32(0)))), ϟa))
				return result
			}
		}() // string
		s.Source += str
		_ = str
	}
	if key, remap := ϟa.Shader.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Shader.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Shader.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Source.value())
	ϟb.Push(ϟa.Length.value())
	ϟb.Call(funcInfoGlShaderSource)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_772_major, minRequiredVersion_772_minor, sources, lengths, context, GetContext_773_result, ctx, s
	return nil
}

var _ = replay.Replayer(&GlUniform1f{}) // interface compliance check
func (ϟa *GlUniform1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_774_major := uint32(2)    // u32
	minRequiredVersion_774_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_775_result := context             // Contextʳ
	ctx := GetContext_775_result                 // Contextʳ
	v := MakeGLfloatˢ(uint64(1), ϟs)             // GLfloatˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform1f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_774_major, minRequiredVersion_774_minor, context, GetContext_775_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform1fv{}) // interface compliance check
func (ϟa *GlUniform1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_776_major := uint32(2)                             // u32
	minRequiredVersion_776_minor := uint32(0)                             // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_777_result := context                                      // Contextʳ
	ctx := GetContext_777_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLfloatˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_FLOAT
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform1fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_776_major, minRequiredVersion_776_minor, context, GetContext_777_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform1i{}) // interface compliance check
func (ϟa *GlUniform1i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_778_major := uint32(2)    // u32
	minRequiredVersion_778_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_779_result := context             // Contextʳ
	ctx := GetContext_779_result                 // Contextʳ
	v := MakeGLintˢ(uint64(1), ϟs)               // GLintˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform1i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(ϟa.Value, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_778_major, minRequiredVersion_778_minor, context, GetContext_779_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform1iv{}) // interface compliance check
func (ϟa *GlUniform1iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_780_major := uint32(2)                             // u32
	minRequiredVersion_780_minor := uint32(0)                             // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_781_result := context                                      // Contextʳ
	ctx := GetContext_781_result                                          // Contextʳ
	v := ϟa.Values.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // GLintˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)               // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                          // Uniform
	uniform.Type = GLenum_GL_INT
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform1iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_780_major, minRequiredVersion_780_minor, context, GetContext_781_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform1ui{}) // interface compliance check
func (ϟa *GlUniform1ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_782_major := uint32(3) // u32
	minRequiredVersion_782_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform1ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_782_major, minRequiredVersion_782_minor
	return nil
}

var _ = replay.Replayer(&GlUniform1uiv{}) // interface compliance check
func (ϟa *GlUniform1uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_783_major := uint32(3) // u32
	minRequiredVersion_783_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform1uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_783_major, minRequiredVersion_783_minor
	return nil
}

var _ = replay.Replayer(&GlUniform2f{}) // interface compliance check
func (ϟa *GlUniform2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_784_major := uint32(2)    // u32
	minRequiredVersion_784_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_785_result := context             // Contextʳ
	ctx := GetContext_785_result                 // Contextʳ
	v := MakeVec2fˢ(uint64(1), ϟs)               // Vec2fˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform2f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(Vec2f{Elements: [2]GLfloat{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_784_major, minRequiredVersion_784_minor, context, GetContext_785_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2fv{}) // interface compliance check
func (ϟa *GlUniform2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_786_major := uint32(2)                                     // u32
	minRequiredVersion_786_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_787_result := context                                              // Contextʳ
	ctx := GetContext_787_result                                                  // Contextʳ
	v := Vec2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_786_major, minRequiredVersion_786_minor, context, GetContext_787_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2i{}) // interface compliance check
func (ϟa *GlUniform2i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_788_major := uint32(2)    // u32
	minRequiredVersion_788_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_789_result := context             // Contextʳ
	ctx := GetContext_789_result                 // Contextʳ
	v := MakeVec2iˢ(uint64(1), ϟs)               // Vec2iˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform2i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(Vec2i{Elements: [2]GLint{ϟa.Value0, ϟa.Value1}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_788_major, minRequiredVersion_788_minor, context, GetContext_789_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2iv{}) // interface compliance check
func (ϟa *GlUniform2iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_790_major := uint32(2)                                     // u32
	minRequiredVersion_790_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_791_result := context                                              // Contextʳ
	ctx := GetContext_791_result                                                  // Contextʳ
	v := Vec2iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec2iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC2
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform2iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_790_major, minRequiredVersion_790_minor, context, GetContext_791_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform2ui{}) // interface compliance check
func (ϟa *GlUniform2ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_792_major := uint32(3) // u32
	minRequiredVersion_792_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform2ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_792_major, minRequiredVersion_792_minor
	return nil
}

var _ = replay.Replayer(&GlUniform2uiv{}) // interface compliance check
func (ϟa *GlUniform2uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_793_major := uint32(3) // u32
	minRequiredVersion_793_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform2uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_793_major, minRequiredVersion_793_minor
	return nil
}

var _ = replay.Replayer(&GlUniform3f{}) // interface compliance check
func (ϟa *GlUniform3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_794_major := uint32(2)    // u32
	minRequiredVersion_794_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_795_result := context             // Contextʳ
	ctx := GetContext_795_result                 // Contextʳ
	v := MakeVec3fˢ(uint64(1), ϟs)               // Vec3fˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform3f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(Vec3f{Elements: [3]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_794_major, minRequiredVersion_794_minor, context, GetContext_795_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3fv{}) // interface compliance check
func (ϟa *GlUniform3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_796_major := uint32(2)                                     // u32
	minRequiredVersion_796_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_797_result := context                                              // Contextʳ
	ctx := GetContext_797_result                                                  // Contextʳ
	v := Vec3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_796_major, minRequiredVersion_796_minor, context, GetContext_797_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3i{}) // interface compliance check
func (ϟa *GlUniform3i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_798_major := uint32(2)    // u32
	minRequiredVersion_798_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_799_result := context             // Contextʳ
	ctx := GetContext_799_result                 // Contextʳ
	v := MakeVec3iˢ(uint64(1), ϟs)               // Vec3iˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform3i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(Vec3i{Elements: [3]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_798_major, minRequiredVersion_798_minor, context, GetContext_799_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3iv{}) // interface compliance check
func (ϟa *GlUniform3iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_800_major := uint32(2)                                     // u32
	minRequiredVersion_800_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_801_result := context                                              // Contextʳ
	ctx := GetContext_801_result                                                  // Contextʳ
	v := Vec3iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec3iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC3
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform3iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_800_major, minRequiredVersion_800_minor, context, GetContext_801_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform3ui{}) // interface compliance check
func (ϟa *GlUniform3ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_802_major := uint32(3) // u32
	minRequiredVersion_802_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform3ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_802_major, minRequiredVersion_802_minor
	return nil
}

var _ = replay.Replayer(&GlUniform3uiv{}) // interface compliance check
func (ϟa *GlUniform3uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_803_major := uint32(3) // u32
	minRequiredVersion_803_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform3uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_803_major, minRequiredVersion_803_minor
	return nil
}

var _ = replay.Replayer(&GlUniform4f{}) // interface compliance check
func (ϟa *GlUniform4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_804_major := uint32(2)    // u32
	minRequiredVersion_804_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_805_result := context             // Contextʳ
	ctx := GetContext_805_result                 // Contextʳ
	v := MakeVec4fˢ(uint64(1), ϟs)               // Vec4fˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform4f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(Vec4f{Elements: [4]GLfloat{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_804_major, minRequiredVersion_804_minor, context, GetContext_805_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4fv{}) // interface compliance check
func (ϟa *GlUniform4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_806_major := uint32(2)                                     // u32
	minRequiredVersion_806_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_807_result := context                                              // Contextʳ
	ctx := GetContext_807_result                                                  // Contextʳ
	v := Vec4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_806_major, minRequiredVersion_806_minor, context, GetContext_807_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4i{}) // interface compliance check
func (ϟa *GlUniform4i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_808_major := uint32(2)    // u32
	minRequiredVersion_808_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_809_result := context             // Contextʳ
	ctx := GetContext_809_result                 // Contextʳ
	v := MakeVec4iˢ(uint64(1), ϟs)               // Vec4iˢ
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform4i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	v.Index(uint64(0), ϟs).replayWrite(Vec4i{Elements: [4]GLint{ϟa.Value0, ϟa.Value1, ϟa.Value2, ϟa.Value3}}, ϟa, ϟs, ϟd, ϟl, ϟb)
	program := ctx.Instances.Programs.Get(ctx.BoundProgram) // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)            // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs)
	program.Uniforms[ϟa.Location] = uniform
	_, _, _, _, _, _, _, _ = minRequiredVersion_808_major, minRequiredVersion_808_minor, context, GetContext_809_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4iv{}) // interface compliance check
func (ϟa *GlUniform4iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_810_major := uint32(2)                                     // u32
	minRequiredVersion_810_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_811_result := context                                              // Contextʳ
	ctx := GetContext_811_result                                                  // Contextʳ
	v := Vec4iᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Vec4iˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_INT_VEC4
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniform4iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_810_major, minRequiredVersion_810_minor, context, GetContext_811_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniform4ui{}) // interface compliance check
func (ϟa *GlUniform4ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_812_major := uint32(3) // u32
	minRequiredVersion_812_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.V0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniform4ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_812_major, minRequiredVersion_812_minor
	return nil
}

var _ = replay.Replayer(&GlUniform4uiv{}) // interface compliance check
func (ϟa *GlUniform4uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_813_major := uint32(3) // u32
	minRequiredVersion_813_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniform4uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_813_major, minRequiredVersion_813_minor
	return nil
}

var _ = replay.Replayer(&GlUniformBlockBinding{}) // interface compliance check
func (ϟa *GlUniformBlockBinding) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_814_major := uint32(3) // u32
	minRequiredVersion_814_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.UniformBlockIndex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.UniformBlockBinding.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlUniformBlockBinding)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_814_major, minRequiredVersion_814_minor
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix2fv{}) // interface compliance check
func (ϟa *GlUniformMatrix2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_815_major := uint32(2)                                     // u32
	minRequiredVersion_815_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_816_result := context                                              // Contextʳ
	ctx := GetContext_816_result                                                  // Contextʳ
	v := Mat2fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat2fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT2
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniformMatrix2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_815_major, minRequiredVersion_815_minor, context, GetContext_816_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix2x3fv{}) // interface compliance check
func (ϟa *GlUniformMatrix2x3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_817_major := uint32(3) // u32
	minRequiredVersion_817_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix2x3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_817_major, minRequiredVersion_817_minor
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix2x4fv{}) // interface compliance check
func (ϟa *GlUniformMatrix2x4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_818_major := uint32(3) // u32
	minRequiredVersion_818_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix2x4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_818_major, minRequiredVersion_818_minor
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix3fv{}) // interface compliance check
func (ϟa *GlUniformMatrix3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_819_major := uint32(2)                                     // u32
	minRequiredVersion_819_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_820_result := context                                              // Contextʳ
	ctx := GetContext_820_result                                                  // Contextʳ
	v := Mat3fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat3fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Type = GLenum_GL_FLOAT_MAT3
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniformMatrix3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_819_major, minRequiredVersion_819_minor, context, GetContext_820_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix3x2fv{}) // interface compliance check
func (ϟa *GlUniformMatrix3x2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_821_major := uint32(3) // u32
	minRequiredVersion_821_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix3x2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_821_major, minRequiredVersion_821_minor
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix3x4fv{}) // interface compliance check
func (ϟa *GlUniformMatrix3x4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_822_major := uint32(3) // u32
	minRequiredVersion_822_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix3x4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_822_major, minRequiredVersion_822_minor
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix4fv{}) // interface compliance check
func (ϟa *GlUniformMatrix4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_823_major := uint32(2)                                     // u32
	minRequiredVersion_823_minor := uint32(0)                                     // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                  // Contextʳ
	GetContext_824_result := context                                              // Contextʳ
	ctx := GetContext_824_result                                                  // Contextʳ
	v := Mat4fᵖ(ϟa.Values).Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // Mat4fˢ
	program := ctx.Instances.Programs.Get(ctx.BoundProgram)                       // Programʳ
	uniform := program.Uniforms.Get(ϟa.Location)                                  // Uniform
	uniform.Value = AsU8ˢ(v, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
	program.Uniforms[ϟa.Location] = uniform
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlUniformMatrix4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _ = minRequiredVersion_823_major, minRequiredVersion_823_minor, context, GetContext_824_result, ctx, v, program, uniform
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix4x2fv{}) // interface compliance check
func (ϟa *GlUniformMatrix4x2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_825_major := uint32(3) // u32
	minRequiredVersion_825_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix4x2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_825_major, minRequiredVersion_825_minor
	return nil
}

var _ = replay.Replayer(&GlUniformMatrix4x3fv{}) // interface compliance check
func (ϟa *GlUniformMatrix4x3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_826_major := uint32(3) // u32
	minRequiredVersion_826_minor := uint32(0) // u32
	if key, remap := ϟa.Location.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeInt32, ϟa.Location.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Transpose.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlUniformMatrix4x3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_826_major, minRequiredVersion_826_minor
	return nil
}

var _ = replay.Replayer(&GlUseProgram{}) // interface compliance check
func (ϟa *GlUseProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_827_major := uint32(2)    // u32
	minRequiredVersion_827_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_828_result := context             // Contextʳ
	ctx := GetContext_828_result                 // Contextʳ
	ctx.BoundProgram = ϟa.Program
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlUseProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_827_major, minRequiredVersion_827_minor, context, GetContext_828_result, ctx
	return nil
}

var _ = replay.Replayer(&GlUseProgramStages{}) // interface compliance check
func (ϟa *GlUseProgramStages) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_829_major := uint32(3)                                                                                                                                        // u32
	minRequiredVersion_829_minor := uint32(1)                                                                                                                                        // u32
	supportsBits_830_seenBits := ϟa.Stages                                                                                                                                           // GLbitfield
	supportsBits_830_validBits := (GLbitfield_GL_ALL_SHADER_BITS) | ((GLbitfield_GL_COMPUTE_SHADER_BIT) | ((GLbitfield_GL_FRAGMENT_SHADER_BIT) | (GLbitfield_GL_VERTEX_SHADER_BIT))) // GLbitfield
	if (GLbitfield_GL_ALL_SHADER_BITS)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_COMPUTE_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_FRAGMENT_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	if (GLbitfield_GL_VERTEX_SHADER_BIT)&(ϟa.Stages) != 0 {
	}
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Stages))
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlUseProgramStages)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_829_major, minRequiredVersion_829_minor, supportsBits_830_seenBits, supportsBits_830_validBits
	return nil
}

var _ = replay.Replayer(&GlValidateProgram{}) // interface compliance check
func (ϟa *GlValidateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_831_major := uint32(2) // u32
	minRequiredVersion_831_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlValidateProgram)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_831_major, minRequiredVersion_831_minor
	return nil
}

var _ = replay.Replayer(&GlValidateProgramPipeline{}) // interface compliance check
func (ϟa *GlValidateProgramPipeline) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_832_major := uint32(3) // u32
	minRequiredVersion_832_minor := uint32(1) // u32
	ϟb.Push(ϟa.Pipeline.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlValidateProgramPipeline)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_832_major, minRequiredVersion_832_minor
	return nil
}

var _ = replay.Replayer(&GlCullFace{}) // interface compliance check
func (ϟa *GlCullFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_833_major := uint32(2) // u32
	minRequiredVersion_833_minor := uint32(0) // u32
	switch ϟa.Mode {
	case GLenum_GL_BACK, GLenum_GL_FRONT, GLenum_GL_FRONT_AND_BACK:
	default:
		v := ϟa.Mode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_835_result := context             // Contextʳ
	ctx := GetContext_835_result                 // Contextʳ
	ctx.Rasterizing.CullFace = ϟa.Mode
	ϟb.Push(value.U32(ϟa.Mode))
	ϟb.Call(funcInfoGlCullFace)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_833_major, minRequiredVersion_833_minor, context, GetContext_835_result, ctx
	return nil
}

var _ = replay.Replayer(&GlDepthRangef{}) // interface compliance check
func (ϟa *GlDepthRangef) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_836_major := uint32(2)    // u32
	minRequiredVersion_836_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_837_result := context             // Contextʳ
	ctx := GetContext_837_result                 // Contextʳ
	ctx.Rasterizing.DepthNear = ϟa.Near
	ctx.Rasterizing.DepthFar = ϟa.Far
	ϟb.Push(ϟa.Near.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Far.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDepthRangef)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_836_major, minRequiredVersion_836_minor, context, GetContext_837_result, ctx
	return nil
}

var _ = replay.Replayer(&GlFrontFace{}) // interface compliance check
func (ϟa *GlFrontFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_838_major := uint32(2) // u32
	minRequiredVersion_838_minor := uint32(0) // u32
	switch ϟa.Orientation {
	case GLenum_GL_CCW, GLenum_GL_CW:
	default:
		v := ϟa.Orientation
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_840_result := context             // Contextʳ
	ctx := GetContext_840_result                 // Contextʳ
	ctx.Rasterizing.FrontFace = ϟa.Orientation
	ϟb.Push(value.U32(ϟa.Orientation))
	ϟb.Call(funcInfoGlFrontFace)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_838_major, minRequiredVersion_838_minor, context, GetContext_840_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetMultisamplefv{}) // interface compliance check
func (ϟa *GlGetMultisamplefv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_841_major := uint32(3) // u32
	minRequiredVersion_841_minor := uint32(1) // u32
	switch ϟa.Pname {
	case GLenum_GL_SAMPLE_POSITION:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Val.value())
	ϟb.Call(funcInfoGlGetMultisamplefv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_841_major, minRequiredVersion_841_minor
	return nil
}

var _ = replay.Replayer(&GlLineWidth{}) // interface compliance check
func (ϟa *GlLineWidth) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_843_major := uint32(2)    // u32
	minRequiredVersion_843_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_844_result := context             // Contextʳ
	ctx := GetContext_844_result                 // Contextʳ
	ctx.Rasterizing.LineWidth = ϟa.Width
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlLineWidth)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_843_major, minRequiredVersion_843_minor, context, GetContext_844_result, ctx
	return nil
}

var _ = replay.Replayer(&GlPolygonOffset{}) // interface compliance check
func (ϟa *GlPolygonOffset) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_845_major := uint32(2)    // u32
	minRequiredVersion_845_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_846_result := context             // Contextʳ
	ctx := GetContext_846_result                 // Contextʳ
	ctx.Rasterizing.PolygonOffsetUnits = ϟa.Units
	ctx.Rasterizing.PolygonOffsetFactor = ϟa.ScaleFactor
	ϟb.Push(ϟa.ScaleFactor.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Units.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPolygonOffset)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_845_major, minRequiredVersion_845_minor, context, GetContext_846_result, ctx
	return nil
}

var _ = replay.Replayer(&GlViewport{}) // interface compliance check
func (ϟa *GlViewport) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_847_major := uint32(2)    // u32
	minRequiredVersion_847_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_848_result := context             // Contextʳ
	ctx := GetContext_848_result                 // Contextʳ
	ctx.Rasterizing.Viewport = func() Rect {
		s := Rect{}
		s.Init()
		s.X = ϟa.X
		s.Y = ϟa.Y
		s.Width = ϟa.Width
		s.Height = ϟa.Height
		return s
	}()
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlViewport)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_847_major, minRequiredVersion_847_minor, context, GetContext_848_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetBooleani_v{}) // interface compliance check
func (ϟa *GlGetBooleani_v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_849_major := uint32(3) // u32
	minRequiredVersion_849_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE, GLenum_GL_VIEWPORT:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlGetBooleani_v)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_849_major, minRequiredVersion_849_minor
	return nil
}

var _ = replay.Replayer(&GlGetBooleanv{}) // interface compliance check
func (ϟa *GlGetBooleanv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_851_major := uint32(2) // u32
	minRequiredVersion_851_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_852_major := uint32(3) // u32
		minRequiredVersion_852_minor := uint32(0) // u32
		_, _ = minRequiredVersion_852_major, minRequiredVersion_852_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_853_major := uint32(3) // u32
		minRequiredVersion_853_minor := uint32(1) // u32
		_, _ = minRequiredVersion_853_major, minRequiredVersion_853_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // Boolˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_855_result := context                                                                    // Contextʳ
	ctx := GetContext_855_result                                                                        // Contextʳ
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetBooleanv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case GLenum_GL_BLEND:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_BLEND), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_CULL_FACE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_CULL_FACE), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_TEST:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_DEPTH_TEST), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DITHER:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_DITHER), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_POLYGON_OFFSET_FILL:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_POLYGON_OFFSET_FILL), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_COVERAGE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_SAMPLE_COVERAGE), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SCISSOR_TEST:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_SCISSOR_TEST), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_TEST:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Capabilities.Get(GLenum_GL_STENCIL_TEST), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.DepthMask, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_COLOR_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.ColorMaskRed, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.ColorMaskGreen, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(ctx.Rasterizing.ColorMaskBlue, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(ctx.Rasterizing.ColorMaskAlpha, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_COVERAGE_INVERT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.SampleCoverageInvert, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SHADER_COMPILER:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_851_major, minRequiredVersion_851_minor, v, context, GetContext_855_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetFloatv{}) // interface compliance check
func (ϟa *GlGetFloatv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_856_major := uint32(2) // u32
	minRequiredVersion_856_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_857_major := uint32(3) // u32
		minRequiredVersion_857_minor := uint32(0) // u32
		_, _ = minRequiredVersion_857_major, minRequiredVersion_857_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_858_major := uint32(3) // u32
		minRequiredVersion_858_minor := uint32(1) // u32
		_, _ = minRequiredVersion_858_major, minRequiredVersion_858_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // GLfloatˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_860_result := context                                                                    // Contextʳ
	ctx := GetContext_860_result                                                                        // Contextʳ
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetFloatv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case GLenum_GL_DEPTH_RANGE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.DepthNear, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.DepthFar, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_LINE_WIDTH:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.LineWidth, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_POLYGON_OFFSET_FACTOR:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.PolygonOffsetFactor, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_POLYGON_OFFSET_UNITS:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.PolygonOffsetUnits, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_COVERAGE_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.SampleCoverageValue, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_COLOR_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Clearing.ClearColor.Red, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Clearing.ClearColor.Green, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(ctx.Clearing.ClearColor.Blue, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(ctx.Clearing.ClearColor.Alpha, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Clearing.ClearDepth, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALIASED_LINE_WIDTH_RANGE:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALIASED_POINT_SIZE_RANGE:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_856_major, minRequiredVersion_856_minor, v, context, GetContext_860_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetInteger64i_v{}) // interface compliance check
func (ϟa *GlGetInteger64i_v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_861_major := uint32(3) // u32
	minRequiredVersion_861_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_862_major := uint32(3) // u32
		minRequiredVersion_862_minor := uint32(1) // u32
		_, _ = minRequiredVersion_862_major, minRequiredVersion_862_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlGetInteger64i_v)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_861_major, minRequiredVersion_861_minor
	return nil
}

var _ = replay.Replayer(&GlGetInteger64v{}) // interface compliance check
func (ϟa *GlGetInteger64v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_864_major := uint32(3) // u32
	minRequiredVersion_864_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_865_major := uint32(3) // u32
		minRequiredVersion_865_minor := uint32(1) // u32
		_, _ = minRequiredVersion_865_major, minRequiredVersion_865_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlGetInteger64v)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_864_major, minRequiredVersion_864_minor
	return nil
}

var _ = replay.Replayer(&GlGetIntegeri_v{}) // interface compliance check
func (ϟa *GlGetIntegeri_v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_867_major := uint32(3) // u32
	minRequiredVersion_867_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_DRAW_BUFFER, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING, GLenum_GL_VIEWPORT:
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_868_major := uint32(3) // u32
		minRequiredVersion_868_minor := uint32(1) // u32
		_, _ = minRequiredVersion_868_major, minRequiredVersion_868_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlGetIntegeri_v)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_867_major, minRequiredVersion_867_minor
	return nil
}

var _ = replay.Replayer(&GlGetIntegerv{}) // interface compliance check
func (ϟa *GlGetIntegerv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_870_major := uint32(2) // u32
	minRequiredVersion_870_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE, GLenum_GL_ALIASED_LINE_WIDTH_RANGE, GLenum_GL_ALIASED_POINT_SIZE_RANGE, GLenum_GL_ALPHA_BITS, GLenum_GL_ARRAY_BUFFER_BINDING, GLenum_GL_BLEND, GLenum_GL_BLEND_COLOR, GLenum_GL_BLEND_DST_ALPHA, GLenum_GL_BLEND_DST_RGB, GLenum_GL_BLEND_EQUATION_ALPHA, GLenum_GL_BLEND_EQUATION_RGB, GLenum_GL_BLEND_SRC_ALPHA, GLenum_GL_BLEND_SRC_RGB, GLenum_GL_BLUE_BITS, GLenum_GL_COLOR_CLEAR_VALUE, GLenum_GL_COLOR_WRITEMASK, GLenum_GL_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_CULL_FACE, GLenum_GL_CULL_FACE_MODE, GLenum_GL_CURRENT_PROGRAM, GLenum_GL_DEPTH_BITS, GLenum_GL_DEPTH_CLEAR_VALUE, GLenum_GL_DEPTH_FUNC, GLenum_GL_DEPTH_RANGE, GLenum_GL_DEPTH_TEST, GLenum_GL_DEPTH_WRITEMASK, GLenum_GL_DITHER, GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING, GLenum_GL_DRAW_FRAMEBUFFER_BINDING, GLenum_GL_FRONT_FACE, GLenum_GL_GENERATE_MIPMAP_HINT, GLenum_GL_GREEN_BITS, GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT, GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE, GLenum_GL_LINE_WIDTH, GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE, GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS, GLenum_GL_MAX_RENDERBUFFER_SIZE, GLenum_GL_MAX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_TEXTURE_SIZE, GLenum_GL_MAX_VARYING_VECTORS, GLenum_GL_MAX_VERTEX_ATTRIBS, GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS, GLenum_GL_MAX_VIEWPORT_DIMS, GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS, GLenum_GL_NUM_SHADER_BINARY_FORMATS, GLenum_GL_PACK_ALIGNMENT, GLenum_GL_POLYGON_OFFSET_FACTOR, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_POLYGON_OFFSET_UNITS, GLenum_GL_RED_BITS, GLenum_GL_RENDERBUFFER_BINDING, GLenum_GL_SAMPLES, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_BUFFERS, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SAMPLE_COVERAGE_INVERT, GLenum_GL_SAMPLE_COVERAGE_VALUE, GLenum_GL_SCISSOR_BOX, GLenum_GL_SCISSOR_TEST, GLenum_GL_SHADER_BINARY_FORMATS, GLenum_GL_SHADER_COMPILER, GLenum_GL_STENCIL_BACK_FAIL, GLenum_GL_STENCIL_BACK_FUNC, GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS, GLenum_GL_STENCIL_BACK_REF, GLenum_GL_STENCIL_BACK_VALUE_MASK, GLenum_GL_STENCIL_BACK_WRITEMASK, GLenum_GL_STENCIL_BITS, GLenum_GL_STENCIL_CLEAR_VALUE, GLenum_GL_STENCIL_FAIL, GLenum_GL_STENCIL_FUNC, GLenum_GL_STENCIL_PASS_DEPTH_FAIL, GLenum_GL_STENCIL_PASS_DEPTH_PASS, GLenum_GL_STENCIL_REF, GLenum_GL_STENCIL_TEST, GLenum_GL_STENCIL_VALUE_MASK, GLenum_GL_STENCIL_WRITEMASK, GLenum_GL_SUBPIXEL_BITS, GLenum_GL_TEXTURE_BINDING_2D, GLenum_GL_TEXTURE_BINDING_CUBE_MAP, GLenum_GL_UNPACK_ALIGNMENT, GLenum_GL_VIEWPORT:
	case GLenum_GL_COPY_READ_BUFFER_BINDING, GLenum_GL_COPY_WRITE_BUFFER_BINDING, GLenum_GL_DRAW_BUFFER, GLenum_GL_FRAGMENT_SHADER_DERIVATIVE_HINT, GLenum_GL_MAJOR_VERSION, GLenum_GL_MAX_3D_TEXTURE_SIZE, GLenum_GL_MAX_ARRAY_TEXTURE_LAYERS, GLenum_GL_MAX_COLOR_ATTACHMENTS, GLenum_GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_UNIFORM_BLOCKS, GLenum_GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MAX_DRAW_BUFFERS, GLenum_GL_MAX_ELEMENTS_INDICES, GLenum_GL_MAX_ELEMENTS_VERTICES, GLenum_GL_MAX_ELEMENT_INDEX, GLenum_GL_MAX_FRAGMENT_INPUT_COMPONENTS, GLenum_GL_MAX_FRAGMENT_UNIFORM_BLOCKS, GLenum_GL_MAX_FRAGMENT_UNIFORM_COMPONENTS, GLenum_GL_MAX_PROGRAM_TEXEL_OFFSET, GLenum_GL_MAX_SAMPLES, GLenum_GL_MAX_SERVER_WAIT_TIMEOUT, GLenum_GL_MAX_TEXTURE_LOD_BIAS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS, GLenum_GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS, GLenum_GL_MAX_UNIFORM_BLOCK_SIZE, GLenum_GL_MAX_UNIFORM_BUFFER_BINDINGS, GLenum_GL_MAX_VARYING_COMPONENTS, GLenum_GL_MAX_VERTEX_OUTPUT_COMPONENTS, GLenum_GL_MAX_VERTEX_UNIFORM_BLOCKS, GLenum_GL_MAX_VERTEX_UNIFORM_COMPONENTS, GLenum_GL_MINOR_VERSION, GLenum_GL_MIN_PROGRAM_TEXEL_OFFSET, GLenum_GL_NUM_EXTENSIONS, GLenum_GL_NUM_PROGRAM_BINARY_FORMATS, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_PIXEL_PACK_BUFFER_BINDING, GLenum_GL_PIXEL_UNPACK_BUFFER_BINDING, GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_PROGRAM_BINARY_FORMATS, GLenum_GL_RASTERIZER_DISCARD, GLenum_GL_READ_BUFFER, GLenum_GL_READ_FRAMEBUFFER_BINDING, GLenum_GL_SAMPLER_BINDING, GLenum_GL_TEXTURE_BINDING_2D_ARRAY, GLenum_GL_TEXTURE_BINDING_3D, GLenum_GL_TRANSFORM_FEEDBACK_ACTIVE, GLenum_GL_TRANSFORM_FEEDBACK_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_BINDING, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_SIZE, GLenum_GL_TRANSFORM_FEEDBACK_BUFFER_START, GLenum_GL_TRANSFORM_FEEDBACK_PAUSED, GLenum_GL_UNIFORM_BUFFER_BINDING, GLenum_GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_UNIFORM_BUFFER_SIZE, GLenum_GL_UNIFORM_BUFFER_START, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS, GLenum_GL_VERTEX_ARRAY_BINDING:
		minRequiredVersion_871_major := uint32(3) // u32
		minRequiredVersion_871_minor := uint32(0) // u32
		_, _ = minRequiredVersion_871_major, minRequiredVersion_871_minor
	case GLenum_GL_DISPATCH_INDIRECT_BUFFER_BINDING, GLenum_GL_IMAGE_BINDING_LAYERED, GLenum_GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS, GLenum_GL_MAX_COLOR_TEXTURE_SAMPLES, GLenum_GL_MAX_COMBINED_ATOMIC_COUNTERS, GLenum_GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTERS, GLenum_GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS, GLenum_GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS, GLenum_GL_MAX_COMPUTE_UNIFORM_BLOCKS, GLenum_GL_MAX_COMPUTE_UNIFORM_COMPONENTS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_COUNT, GLenum_GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS, GLenum_GL_MAX_COMPUTE_WORK_GROUP_SIZE, GLenum_GL_MAX_FRAGMENT_ATOMIC_COUNTERS, GLenum_GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS, GLenum_GL_MAX_FRAMEBUFFER_HEIGHT, GLenum_GL_MAX_FRAMEBUFFER_SAMPLES, GLenum_GL_MAX_FRAMEBUFFER_WIDTH, GLenum_GL_MAX_INTEGER_SAMPLES, GLenum_GL_MAX_SAMPLE_MASK_WORDS, GLenum_GL_MAX_SHADER_STORAGE_BLOCK_SIZE, GLenum_GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS, GLenum_GL_MAX_UNIFORM_LOCATIONS, GLenum_GL_MAX_VERTEX_ATOMIC_COUNTERS, GLenum_GL_MAX_VERTEX_ATTRIB_BINDINGS, GLenum_GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET, GLenum_GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS, GLenum_GL_PROGRAM_PIPELINE_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_BINDING, GLenum_GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT, GLenum_GL_SHADER_STORAGE_BUFFER_SIZE, GLenum_GL_SHADER_STORAGE_BUFFER_START, GLenum_GL_TEXTURE_BINDING_2D_MULTISAMPLE, GLenum_GL_VERTEX_BINDING_DIVISOR, GLenum_GL_VERTEX_BINDING_OFFSET, GLenum_GL_VERTEX_BINDING_STRIDE:
		minRequiredVersion_872_major := uint32(3) // u32
		minRequiredVersion_872_minor := uint32(1) // u32
		_, _ = minRequiredVersion_872_major, minRequiredVersion_872_minor
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	v := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs) // GLintˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                                                        // Contextʳ
	GetContext_874_result := context                                                                    // Contextʳ
	ctx := GetContext_874_result                                                                        // Contextʳ
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetIntegerv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	switch ϟa.Param {
	case GLenum_GL_ACTIVE_TEXTURE:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.ActiveTextureUnit), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.BoundBuffers.Get(GLenum_GL_ARRAY_BUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.BoundBuffers.Get(GLenum_GL_ELEMENT_ARRAY_BUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_SRC_ALPHA:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Blending.SrcAlphaBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_SRC_RGB:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Blending.SrcRgbBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_DST_ALPHA:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Blending.DstAlphaBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_DST_RGB:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Blending.DstRgbBlendFactor), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_EQUATION_RGB:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Blending.BlendEquationRgb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_EQUATION_ALPHA:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Blending.BlendEquationAlpha), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLEND_COLOR:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Blending.BlendColor.Red), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(GLint(ctx.Blending.BlendColor.Green), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(GLint(ctx.Blending.BlendColor.Blue), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(GLint(ctx.Blending.BlendColor.Alpha), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_FUNC:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Rasterizing.DepthTestFunction), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Clearing.ClearDepth), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Rasterizing.StencilMask.Get(GLenum_GL_FRONT)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_BACK_WRITEMASK:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Rasterizing.StencilMask.Get(GLenum_GL_BACK)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_VIEWPORT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.Viewport.X, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.Viewport.Y, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(GLint(ctx.Rasterizing.Viewport.Width), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(GLint(ctx.Rasterizing.Viewport.Height), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SCISSOR_BOX:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Rasterizing.Scissor.X, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(ctx.Rasterizing.Scissor.Y, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(2), ϟs).replayWrite(GLint(ctx.Rasterizing.Scissor.Width), ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(3), ϟs).replayWrite(GLint(ctx.Rasterizing.Scissor.Height), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_FRONT_FACE:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Rasterizing.FrontFace), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_CULL_FACE_MODE:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.Rasterizing.CullFace), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_CLEAR_VALUE:
		v.Index(uint64(0), ϟs).replayWrite(ctx.Clearing.ClearStencil, ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DRAW_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.BoundFramebuffers.Get(GLenum_GL_FRAMEBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_READ_FRAMEBUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.BoundFramebuffers.Get(GLenum_GL_READ_FRAMEBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_RENDERBUFFER_BINDING:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.BoundRenderbuffers.Get(GLenum_GL_RENDERBUFFER)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_CURRENT_PROGRAM:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.BoundProgram), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BINDING_2D:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_TEXTURE_BINDING_CUBE_MAP:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP)), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GENERATE_MIPMAP_HINT:
		v.Index(uint64(0), ϟs).replayWrite(GLint(ctx.GenerateMipmapHint), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_RENDERBUFFER_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_TEXTURE_IMAGE_UNITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_TEXTURE_SIZE:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VARYING_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VERTEX_ATTRIBS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_MAX_VIEWPORT_DIMS:
		max_width := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)  // any
		max_height := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(1), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(max_width, ϟa, ϟs, ϟd, ϟl, ϟb)
		v.Index(uint64(1), ϟs).replayWrite(max_height, ϟa, ϟs, ϟd, ϟl, ϟb)
		_, _ = max_width, max_height
	case GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_NUM_SHADER_BINARY_FORMATS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_PACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.PixelStorage.Get(GLenum_GL_PACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_UNPACK_ALIGNMENT:
		v.Index(uint64(0), ϟs).replayWrite(ctx.PixelStorage.Get(GLenum_GL_UNPACK_ALIGNMENT), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_ALPHA_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_BLUE_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GREEN_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_RED_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_DEPTH_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLE_BUFFERS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SAMPLES:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_STENCIL_BITS:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_SUBPIXEL_BITS:
		result := ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb) // any
		v.Index(uint64(0), ϟs).replayWrite(result, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = result
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	case GLenum_GL_GPU_DISJOINT_EXT:
		v.Index(uint64(0), ϟs).replayWrite(ϟa.Values.Slice(uint64(int32(0)), uint64(externs{ϟs, ϟd, ϟl}.stateVariableSize(ϟa.Param)), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	_, _, _, _, _, _ = minRequiredVersion_870_major, minRequiredVersion_870_minor, v, context, GetContext_874_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetInternalformativ{}) // interface compliance check
func (ϟa *GlGetInternalformativ) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_875_major := uint32(3) // u32
	minRequiredVersion_875_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_RENDERBUFFER:
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_876_major := uint32(3) // u32
		minRequiredVersion_876_minor := uint32(1) // u32
		_, _ = minRequiredVersion_876_major, minRequiredVersion_876_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_NUM_SAMPLE_COUNTS, GLenum_GL_SAMPLES:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetInternalformativ)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_875_major, minRequiredVersion_875_minor
	return nil
}

var _ = replay.Replayer(&GlGetString{}) // interface compliance check
func (ϟa *GlGetString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_880_major := uint32(2) // u32
	minRequiredVersion_880_minor := uint32(0) // u32
	switch ϟa.Param {
	case GLenum_GL_EXTENSIONS, GLenum_GL_RENDERER, GLenum_GL_SHADING_LANGUAGE_VERSION, GLenum_GL_VENDOR, GLenum_GL_VERSION:
	default:
		v := ϟa.Param
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Param))
	ϟb.Call(funcInfoGlGetString)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_880_major, minRequiredVersion_880_minor
	return nil
}

var _ = replay.Replayer(&GlGetStringi{}) // interface compliance check
func (ϟa *GlGetStringi) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_882_major := uint32(3) // u32
	minRequiredVersion_882_minor := uint32(0) // u32
	switch ϟa.Name {
	case GLenum_GL_EXTENSIONS:
	default:
		v := ϟa.Name
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Name))
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlGetStringi)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_882_major, minRequiredVersion_882_minor
	return nil
}

var _ = replay.Replayer(&GlIsEnabled{}) // interface compliance check
func (ϟa *GlIsEnabled) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_884_major := uint32(2) // u32
	minRequiredVersion_884_minor := uint32(0) // u32
	switch ϟa.Capability {
	case GLenum_GL_BLEND, GLenum_GL_CULL_FACE, GLenum_GL_DEPTH_TEST, GLenum_GL_DITHER, GLenum_GL_POLYGON_OFFSET_FILL, GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE, GLenum_GL_SAMPLE_COVERAGE, GLenum_GL_SCISSOR_TEST, GLenum_GL_STENCIL_TEST:
	case GLenum_GL_PRIMITIVE_RESTART_FIXED_INDEX, GLenum_GL_RASTERIZER_DISCARD:
		minRequiredVersion_885_major := uint32(3) // u32
		minRequiredVersion_885_minor := uint32(0) // u32
		_, _ = minRequiredVersion_885_major, minRequiredVersion_885_minor
	default:
		v := ϟa.Capability
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_887_result := context             // Contextʳ
	ctx := GetContext_887_result                 // Contextʳ
	ϟb.Push(value.U32(ϟa.Capability))
	ϟb.Call(funcInfoGlIsEnabled)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_884_major, minRequiredVersion_884_minor, context, GetContext_887_result, ctx
	return nil
}

var _ = replay.Replayer(&GlClientWaitSync{}) // interface compliance check
func (ϟa *GlClientWaitSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_888_major := uint32(3)                           // u32
	minRequiredVersion_888_minor := uint32(0)                           // u32
	supportsBits_889_seenBits := ϟa.SyncFlags                           // GLbitfield
	supportsBits_889_validBits := GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT // GLbitfield
	if (GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT)&(ϟa.SyncFlags) != 0 {
	}
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.SyncFlags))
	ϟb.Push(ϟa.Timeout.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlClientWaitSync)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _ = minRequiredVersion_888_major, minRequiredVersion_888_minor, supportsBits_889_seenBits, supportsBits_889_validBits
	return nil
}

var _ = replay.Replayer(&GlDeleteSync{}) // interface compliance check
func (ϟa *GlDeleteSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_890_major := uint32(3) // u32
	minRequiredVersion_890_minor := uint32(0) // u32
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlDeleteSync)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_890_major, minRequiredVersion_890_minor
	return nil
}

var _ = replay.Replayer(&GlFenceSync{}) // interface compliance check
func (ϟa *GlFenceSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_891_major := uint32(3) // u32
	minRequiredVersion_891_minor := uint32(0) // u32
	switch ϟa.Condition {
	case GLenum_GL_SYNC_GPU_COMMANDS_COMPLETE:
	default:
		v := ϟa.Condition
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Condition))
	ϟb.Push(value.U32(ϟa.SyncFlags))
	ϟb.Call(funcInfoGlFenceSync)
	if key, remap := ϟa.Result.remap(ϟa, ϟs); remap {
		ptr, found := ϟb.Remappings[key]
		if !found {
			ptr = ϟb.AllocateMemory(uint64(8))
			ϟb.Remappings[key] = ptr
		}
		ϟb.Clone(0)
		ϟb.Store(ptr)
	}
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_891_major, minRequiredVersion_891_minor
	return nil
}

var _ = replay.Replayer(&GlGetSynciv{}) // interface compliance check
func (ϟa *GlGetSynciv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_893_major := uint32(3) // u32
	minRequiredVersion_893_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_OBJECT_TYPE, GLenum_GL_SYNC_CONDITION, GLenum_GL_SYNC_FLAGS, GLenum_GL_SYNC_STATUS:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetSynciv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_893_major, minRequiredVersion_893_minor
	return nil
}

var _ = replay.Replayer(&GlIsSync{}) // interface compliance check
func (ϟa *GlIsSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_895_major := uint32(3) // u32
	minRequiredVersion_895_minor := uint32(0) // u32
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsSync)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_895_major, minRequiredVersion_895_minor
	return nil
}

var _ = replay.Replayer(&GlWaitSync{}) // interface compliance check
func (ϟa *GlWaitSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_896_major := uint32(3) // u32
	minRequiredVersion_896_minor := uint32(0) // u32
	if key, remap := ϟa.Sync.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint64, ϟa.Sync.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Sync.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.SyncFlags))
	ϟb.Push(ϟa.Timeout.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlWaitSync)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_896_major, minRequiredVersion_896_minor
	return nil
}

var _ = replay.Replayer(&GlActiveTexture{}) // interface compliance check
func (ϟa *GlActiveTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_897_major := uint32(2) // u32
	minRequiredVersion_897_minor := uint32(0) // u32
	switch ϟa.Unit {
	case GLenum_GL_TEXTURE0, GLenum_GL_TEXTURE1, GLenum_GL_TEXTURE10, GLenum_GL_TEXTURE11, GLenum_GL_TEXTURE12, GLenum_GL_TEXTURE13, GLenum_GL_TEXTURE14, GLenum_GL_TEXTURE15, GLenum_GL_TEXTURE16, GLenum_GL_TEXTURE17, GLenum_GL_TEXTURE18, GLenum_GL_TEXTURE19, GLenum_GL_TEXTURE2, GLenum_GL_TEXTURE20, GLenum_GL_TEXTURE21, GLenum_GL_TEXTURE22, GLenum_GL_TEXTURE23, GLenum_GL_TEXTURE24, GLenum_GL_TEXTURE25, GLenum_GL_TEXTURE26, GLenum_GL_TEXTURE27, GLenum_GL_TEXTURE28, GLenum_GL_TEXTURE29, GLenum_GL_TEXTURE3, GLenum_GL_TEXTURE30, GLenum_GL_TEXTURE31, GLenum_GL_TEXTURE4, GLenum_GL_TEXTURE5, GLenum_GL_TEXTURE6, GLenum_GL_TEXTURE7, GLenum_GL_TEXTURE8, GLenum_GL_TEXTURE9:
	default:
		v := ϟa.Unit
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_899_result := context             // Contextʳ
	ctx := GetContext_899_result                 // Contextʳ
	ctx.ActiveTextureUnit = ϟa.Unit
	if !(ctx.TextureUnits.Contains(ϟa.Unit)) {
		ctx.TextureUnits[ϟa.Unit] = ctx.TextureUnits.Get(ϟa.Unit)
	}
	ϟb.Push(value.U32(ϟa.Unit))
	ϟb.Call(funcInfoGlActiveTexture)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_897_major, minRequiredVersion_897_minor, context, GetContext_899_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindImageTexture{}) // interface compliance check
func (ϟa *GlBindImageTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_900_major := uint32(3) // u32
	minRequiredVersion_900_minor := uint32(1) // u32
	switch ϟa.Access {
	case GLenum_GL_READ_ONLY, GLenum_GL_READ_WRITE, GLenum_GL_WRITE_ONLY:
	default:
		v := ϟa.Access
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Unit.value(ϟb, ϟa, ϟs))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Layered.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Layer.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Access))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Call(funcInfoGlBindImageTexture)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_900_major, minRequiredVersion_900_minor
	return nil
}

var _ = replay.Replayer(&GlBindSampler{}) // interface compliance check
func (ϟa *GlBindSampler) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_903_major := uint32(3) // u32
	minRequiredVersion_903_minor := uint32(0) // u32
	ϟb.Push(ϟa.Unit.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBindSampler)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_903_major, minRequiredVersion_903_minor
	return nil
}

var _ = replay.Replayer(&GlBindTexture{}) // interface compliance check
func (ϟa *GlBindTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_904_major := uint32(2) // u32
	minRequiredVersion_904_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_905_major := uint32(3) // u32
		minRequiredVersion_905_minor := uint32(0) // u32
		_, _ = minRequiredVersion_905_major, minRequiredVersion_905_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_906_major := uint32(3) // u32
		minRequiredVersion_906_minor := uint32(1) // u32
		_, _ = minRequiredVersion_906_major, minRequiredVersion_906_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_908_result := context             // Contextʳ
	ctx := GetContext_908_result                 // Contextʳ
	if !(ctx.Instances.Textures.Contains(ϟa.Texture)) {
		ctx.Instances.Textures[ϟa.Texture] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
	}
	ctx.TextureUnits.Get(ctx.ActiveTextureUnit)[ϟa.Target] = ϟa.Texture
	ϟb.Push(value.U32(ϟa.Target))
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindTexture)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_904_major, minRequiredVersion_904_minor, context, GetContext_908_result, ctx
	return nil
}

var _ = replay.Replayer(&GlCompressedTexImage2D{}) // interface compliance check
func (ϟa *GlCompressedTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_909_major := uint32(2) // u32
	minRequiredVersion_909_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_911_major := uint32(3) // u32
		minRequiredVersion_911_minor := uint32(0) // u32
		_, _ = minRequiredVersion_911_major, minRequiredVersion_911_minor
	case GLenum_GL_ATC_RGB_AMD, GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD, GLenum_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD:
		requiresExtension_912_ext := ExtensionId_GL_AMD_compressed_ATC_texture // ExtensionId
		_ = requiresExtension_912_ext
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_914_result := context             // Contextʳ
	ctx := GetContext_914_result                 // Contextʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                         // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                               // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = uint32(ϟa.ImageSize)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Border.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.ImageSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCompressedTexImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_909_major, minRequiredVersion_909_minor, context, GetContext_914_result, ctx
	return nil
}

var _ = replay.Replayer(&GlCompressedTexImage3D{}) // interface compliance check
func (ϟa *GlCompressedTexImage3D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_915_major := uint32(3) // u32
	minRequiredVersion_915_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Border.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.ImageSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlCompressedTexImage3D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_915_major, minRequiredVersion_915_minor
	return nil
}

var _ = replay.Replayer(&GlCompressedTexSubImage2D{}) // interface compliance check
func (ϟa *GlCompressedTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_918_major := uint32(2) // u32
	minRequiredVersion_918_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
		minRequiredVersion_920_major := uint32(3) // u32
		minRequiredVersion_920_minor := uint32(0) // u32
		_, _ = minRequiredVersion_920_major, minRequiredVersion_920_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.ImageSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCompressedTexSubImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_918_major, minRequiredVersion_918_minor
	return nil
}

var _ = replay.Replayer(&GlCompressedTexSubImage3D{}) // interface compliance check
func (ϟa *GlCompressedTexSubImage3D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_922_major := uint32(3) // u32
	minRequiredVersion_922_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.ImageSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value())
	ϟb.Call(funcInfoGlCompressedTexSubImage3D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_922_major, minRequiredVersion_922_minor
	return nil
}

var _ = replay.Replayer(&GlCopyTexImage2D{}) // interface compliance check
func (ϟa *GlCopyTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_925_major := uint32(2) // u32
	minRequiredVersion_925_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
		minRequiredVersion_927_major := uint32(3) // u32
		minRequiredVersion_927_minor := uint32(0) // u32
		_, _ = minRequiredVersion_927_major, minRequiredVersion_927_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Border.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyTexImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_925_major, minRequiredVersion_925_minor
	return nil
}

var _ = replay.Replayer(&GlCopyTexSubImage2D{}) // interface compliance check
func (ϟa *GlCopyTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_929_major := uint32(2) // u32
	minRequiredVersion_929_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyTexSubImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_929_major, minRequiredVersion_929_minor
	return nil
}

var _ = replay.Replayer(&GlCopyTexSubImage3D{}) // interface compliance check
func (ϟa *GlCopyTexSubImage3D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_931_major := uint32(3) // u32
	minRequiredVersion_931_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlCopyTexSubImage3D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_931_major, minRequiredVersion_931_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteSamplers{}) // interface compliance check
func (ϟa *GlDeleteSamplers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_933_major := uint32(3) // u32
	minRequiredVersion_933_minor := uint32(0) // u32
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Samplers.value())
	ϟb.Call(funcInfoGlDeleteSamplers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_933_major, minRequiredVersion_933_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteTextures{}) // interface compliance check
func (ϟa *GlDeleteTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_934_major := uint32(2)                               // u32
	minRequiredVersion_934_minor := uint32(0)                               // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_935_result := context                                        // Contextʳ
	ctx := GetContext_935_result                                            // Contextʳ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.Textures, t.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Textures.value())
	ϟb.Call(funcInfoGlDeleteTextures)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_934_major, minRequiredVersion_934_minor, t, context, GetContext_935_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenSamplers{}) // interface compliance check
func (ϟa *GlGenSamplers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_936_major := uint32(3) // u32
	minRequiredVersion_936_minor := uint32(0) // u32
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Samplers.value())
	ϟb.Call(funcInfoGlGenSamplers)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_936_major, minRequiredVersion_936_minor
	return nil
}

var _ = replay.Replayer(&GlGenTextures{}) // interface compliance check
func (ϟa *GlGenTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_937_major := uint32(2)                               // u32
	minRequiredVersion_937_minor := uint32(0)                               // u32
	t := ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // TextureIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                            // Contextʳ
	GetContext_938_result := context                                        // Contextʳ
	ctx := GetContext_938_result                                            // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Textures.value())
	ϟb.Call(funcInfoGlGenTextures)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := TextureId(ϟa.Textures.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // TextureId
		ctx.Instances.Textures[id] = func() *Texture {
			s := &Texture{}
			s.Init()
			return s
		}()
		t.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_937_major, minRequiredVersion_937_minor, t, context, GetContext_938_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenerateMipmap{}) // interface compliance check
func (ϟa *GlGenerateMipmap) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_939_major := uint32(2) // u32
	minRequiredVersion_939_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_940_major := uint32(3) // u32
		minRequiredVersion_940_minor := uint32(0) // u32
		_, _ = minRequiredVersion_940_major, minRequiredVersion_940_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Call(funcInfoGlGenerateMipmap)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_939_major, minRequiredVersion_939_minor
	return nil
}

var _ = replay.Replayer(&GlGetSamplerParameterfv{}) // interface compliance check
func (ϟa *GlGetSamplerParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_942_major := uint32(3) // u32
	minRequiredVersion_942_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetSamplerParameterfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_942_major, minRequiredVersion_942_minor
	return nil
}

var _ = replay.Replayer(&GlGetSamplerParameteriv{}) // interface compliance check
func (ϟa *GlGetSamplerParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_944_major := uint32(3) // u32
	minRequiredVersion_944_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetSamplerParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_944_major, minRequiredVersion_944_minor
	return nil
}

var _ = replay.Replayer(&GlGetTexLevelParameterfv{}) // interface compliance check
func (ϟa *GlGetTexLevelParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_946_major := uint32(3) // u32
	minRequiredVersion_946_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetTexLevelParameterfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_946_major, minRequiredVersion_946_minor
	return nil
}

var _ = replay.Replayer(&GlGetTexLevelParameteriv{}) // interface compliance check
func (ϟa *GlGetTexLevelParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_949_major := uint32(3) // u32
	minRequiredVersion_949_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_2D_MULTISAMPLE, GLenum_GL_TEXTURE_3D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_ALPHA_SIZE, GLenum_GL_TEXTURE_ALPHA_TYPE, GLenum_GL_TEXTURE_BLUE_SIZE, GLenum_GL_TEXTURE_BLUE_TYPE, GLenum_GL_TEXTURE_COMPRESSED, GLenum_GL_TEXTURE_DEPTH, GLenum_GL_TEXTURE_DEPTH_SIZE, GLenum_GL_TEXTURE_DEPTH_TYPE, GLenum_GL_TEXTURE_FIXED_SAMPLE_LOCATIONS, GLenum_GL_TEXTURE_GREEN_SIZE, GLenum_GL_TEXTURE_GREEN_TYPE, GLenum_GL_TEXTURE_HEIGHT, GLenum_GL_TEXTURE_INTERNAL_FORMAT, GLenum_GL_TEXTURE_RED_SIZE, GLenum_GL_TEXTURE_RED_TYPE, GLenum_GL_TEXTURE_SAMPLES, GLenum_GL_TEXTURE_SHARED_SIZE, GLenum_GL_TEXTURE_STENCIL_SIZE, GLenum_GL_TEXTURE_WIDTH:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetTexLevelParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_949_major, minRequiredVersion_949_minor
	return nil
}

var _ = replay.Replayer(&GlGetTexParameterfv{}) // interface compliance check
func (ϟa *GlGetTexParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_952_major := uint32(2) // u32
	minRequiredVersion_952_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_953_major := uint32(3) // u32
		minRequiredVersion_953_minor := uint32(0) // u32
		_, _ = minRequiredVersion_953_major, minRequiredVersion_953_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_954_major := uint32(3) // u32
		minRequiredVersion_954_minor := uint32(1) // u32
		_, _ = minRequiredVersion_954_major, minRequiredVersion_954_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_956_major := uint32(3) // u32
		minRequiredVersion_956_minor := uint32(0) // u32
		_, _ = minRequiredVersion_956_major, minRequiredVersion_956_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_957_major := uint32(3) // u32
		minRequiredVersion_957_minor := uint32(1) // u32
		_, _ = minRequiredVersion_957_major, minRequiredVersion_957_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_959_result := context                                 // Contextʳ
	ctx := GetContext_959_result                                     // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetTexParameterfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result GLfloat) {
		switch ϟa.Parameter {
		case GLenum_GL_TEXTURE_MAG_FILTER:
			return GLfloat(t.MagFilter)
		case GLenum_GL_TEXTURE_MIN_FILTER:
			return GLfloat(t.MinFilter)
		case GLenum_GL_TEXTURE_WRAP_S:
			return GLfloat(t.WrapS)
		case GLenum_GL_TEXTURE_WRAP_T:
			return GLfloat(t.WrapT)
		case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return GLfloat(t.MaxAnisotropy)
		case GLenum_GL_TEXTURE_SWIZZLE_R:
			return GLfloat(t.SwizzleR)
		case GLenum_GL_TEXTURE_SWIZZLE_G:
			return GLfloat(t.SwizzleG)
		case GLenum_GL_TEXTURE_SWIZZLE_B:
			return GLfloat(t.SwizzleB)
		case GLenum_GL_TEXTURE_SWIZZLE_A:
			return GLfloat(t.SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_952_major, minRequiredVersion_952_minor, context, GetContext_959_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlGetTexParameteriv{}) // interface compliance check
func (ϟa *GlGetTexParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_960_major := uint32(2) // u32
	minRequiredVersion_960_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_961_major := uint32(3) // u32
		minRequiredVersion_961_minor := uint32(0) // u32
		_, _ = minRequiredVersion_961_major, minRequiredVersion_961_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_962_major := uint32(3) // u32
		minRequiredVersion_962_minor := uint32(1) // u32
		_, _ = minRequiredVersion_962_major, minRequiredVersion_962_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_IMMUTABLE_FORMAT, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_964_major := uint32(3) // u32
		minRequiredVersion_964_minor := uint32(0) // u32
		_, _ = minRequiredVersion_964_major, minRequiredVersion_964_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE, GLenum_GL_IMAGE_FORMAT_COMPATIBILITY_TYPE, GLenum_GL_TEXTURE_IMMUTABLE_LEVELS:
		minRequiredVersion_965_major := uint32(3) // u32
		minRequiredVersion_965_minor := uint32(1) // u32
		_, _ = minRequiredVersion_965_major, minRequiredVersion_965_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_967_result := context                                 // Contextʳ
	ctx := GetContext_967_result                                     // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Values.value())
	ϟb.Call(funcInfoGlGetTexParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Values.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(func() (result GLint) {
		switch ϟa.Parameter {
		case GLenum_GL_TEXTURE_MAG_FILTER:
			return GLint(t.MagFilter)
		case GLenum_GL_TEXTURE_MIN_FILTER:
			return GLint(t.MinFilter)
		case GLenum_GL_TEXTURE_WRAP_S:
			return GLint(t.WrapS)
		case GLenum_GL_TEXTURE_WRAP_T:
			return GLint(t.WrapT)
		case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
			return GLint(t.MaxAnisotropy)
		case GLenum_GL_TEXTURE_SWIZZLE_R:
			return GLint(t.SwizzleR)
		case GLenum_GL_TEXTURE_SWIZZLE_G:
			return GLint(t.SwizzleG)
		case GLenum_GL_TEXTURE_SWIZZLE_B:
			return GLint(t.SwizzleB)
		case GLenum_GL_TEXTURE_SWIZZLE_A:
			return GLint(t.SwizzleA)
		default:
			// TODO: better unmatched handling
			panic(fmt.Errorf("Unmatched switch(%v) in atom %T", ϟa.Parameter, ϟa))
			return result
		}
	}(), ϟa, ϟs, ϟd, ϟl, ϟb)
	_, _, _, _, _, _, _ = minRequiredVersion_960_major, minRequiredVersion_960_minor, context, GetContext_967_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlIsSampler{}) // interface compliance check
func (ϟa *GlIsSampler) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_968_major := uint32(3) // u32
	minRequiredVersion_968_minor := uint32(0) // u32
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsSampler)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_968_major, minRequiredVersion_968_minor
	return nil
}

var _ = replay.Replayer(&GlIsTexture{}) // interface compliance check
func (ϟa *GlIsTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_969_major := uint32(2)    // u32
	minRequiredVersion_969_minor := uint32(0)    // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_970_result := context             // Contextʳ
	ctx := GetContext_970_result                 // Contextʳ
	if key, remap := ϟa.Texture.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Texture.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Texture.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsTexture)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_969_major, minRequiredVersion_969_minor, context, GetContext_970_result, ctx
	return nil
}

var _ = replay.Replayer(&GlPixelStorei{}) // interface compliance check
func (ϟa *GlPixelStorei) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_971_major := uint32(2) // u32
	minRequiredVersion_971_minor := uint32(0) // u32
	switch ϟa.Parameter {
	case GLenum_GL_PACK_ALIGNMENT, GLenum_GL_UNPACK_ALIGNMENT:
	case GLenum_GL_PACK_IMAGE_HEIGHT, GLenum_GL_PACK_ROW_LENGTH, GLenum_GL_PACK_SKIP_IMAGES, GLenum_GL_PACK_SKIP_PIXELS, GLenum_GL_PACK_SKIP_ROWS, GLenum_GL_UNPACK_IMAGE_HEIGHT, GLenum_GL_UNPACK_ROW_LENGTH, GLenum_GL_UNPACK_SKIP_IMAGES, GLenum_GL_UNPACK_SKIP_PIXELS, GLenum_GL_UNPACK_SKIP_ROWS:
		minRequiredVersion_972_major := uint32(3) // u32
		minRequiredVersion_972_minor := uint32(0) // u32
		_, _ = minRequiredVersion_972_major, minRequiredVersion_972_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_974_result := context             // Contextʳ
	ctx := GetContext_974_result                 // Contextʳ
	ctx.PixelStorage[ϟa.Parameter] = ϟa.Value
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlPixelStorei)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_971_major, minRequiredVersion_971_minor, context, GetContext_974_result, ctx
	return nil
}

var _ = replay.Replayer(&GlSamplerParameterf{}) // interface compliance check
func (ϟa *GlSamplerParameterf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_975_major := uint32(3) // u32
	minRequiredVersion_975_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlSamplerParameterf)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_975_major, minRequiredVersion_975_minor
	return nil
}

var _ = replay.Replayer(&GlSamplerParameterfv{}) // interface compliance check
func (ϟa *GlSamplerParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_977_major := uint32(3) // u32
	minRequiredVersion_977_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value())
	ϟb.Call(funcInfoGlSamplerParameterfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_977_major, minRequiredVersion_977_minor
	return nil
}

var _ = replay.Replayer(&GlSamplerParameteri{}) // interface compliance check
func (ϟa *GlSamplerParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_979_major := uint32(3) // u32
	minRequiredVersion_979_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlSamplerParameteri)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_979_major, minRequiredVersion_979_minor
	return nil
}

var _ = replay.Replayer(&GlSamplerParameteriv{}) // interface compliance check
func (ϟa *GlSamplerParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_981_major := uint32(3) // u32
	minRequiredVersion_981_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_WRAP_R, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Sampler.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Param.value())
	ϟb.Call(funcInfoGlSamplerParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_981_major, minRequiredVersion_981_minor
	return nil
}

var _ = replay.Replayer(&GlTexImage2D{}) // interface compliance check
func (ϟa *GlTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_983_major := uint32(2) // u32
	minRequiredVersion_983_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_985_major := uint32(3) // u32
		minRequiredVersion_985_minor := uint32(0) // u32
		_, _ = minRequiredVersion_985_major, minRequiredVersion_985_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_987_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_987_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_988_major := uint32(3) // u32
		minRequiredVersion_988_minor := uint32(0) // u32
		_, _ = minRequiredVersion_988_major, minRequiredVersion_988_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_990_result := context             // Contextʳ
	ctx := GetContext_990_result                 // Contextʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                         // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(Voidᶜᵖ{})) {
			if (ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                               // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if (ϟa.Data) != (TexturePointer(Voidᶜᵖ{})) {
			if (ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0))) {
				l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
			}
		} else {
			l.Data = MakeU8ˢ(uint64(l.Size), ϟs)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.InternalFormat.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Border.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_983_major, minRequiredVersion_983_minor, context, GetContext_990_result, ctx
	return nil
}

var _ = replay.Replayer(&GlTexImage3D{}) // interface compliance check
func (ϟa *GlTexImage3D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_991_major := uint32(3) // u32
	minRequiredVersion_991_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_994_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_994_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Internalformat.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Border.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Pixels.value())
	ϟb.Call(funcInfoGlTexImage3D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_991_major, minRequiredVersion_991_minor
	return nil
}

var _ = replay.Replayer(&GlTexParameterf{}) // interface compliance check
func (ϟa *GlTexParameterf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_996_major := uint32(2) // u32
	minRequiredVersion_996_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_997_major := uint32(3) // u32
		minRequiredVersion_997_minor := uint32(0) // u32
		_, _ = minRequiredVersion_997_major, minRequiredVersion_997_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_998_major := uint32(3) // u32
		minRequiredVersion_998_minor := uint32(1) // u32
		_, _ = minRequiredVersion_998_major, minRequiredVersion_998_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1000_major := uint32(3) // u32
		minRequiredVersion_1000_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1000_major, minRequiredVersion_1000_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1001_major := uint32(3) // u32
		minRequiredVersion_1001_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1001_major, minRequiredVersion_1001_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_1003_result := context                                // Contextʳ
	ctx := GetContext_1003_result                                    // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_S:
		t.WrapS = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_T:
		t.WrapT = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = float32(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = GLenum(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexParameterf)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_996_major, minRequiredVersion_996_minor, context, GetContext_1003_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlTexParameterfv{}) // interface compliance check
func (ϟa *GlTexParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1004_major := uint32(2) // u32
	minRequiredVersion_1004_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1005_major := uint32(3) // u32
		minRequiredVersion_1005_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1005_major, minRequiredVersion_1005_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1006_major := uint32(3) // u32
		minRequiredVersion_1006_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1006_major, minRequiredVersion_1006_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1008_major := uint32(3) // u32
		minRequiredVersion_1008_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1008_major, minRequiredVersion_1008_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1009_major := uint32(3) // u32
		minRequiredVersion_1009_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1009_major, minRequiredVersion_1009_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlTexParameterfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1004_major, minRequiredVersion_1004_minor
	return nil
}

var _ = replay.Replayer(&GlTexParameteri{}) // interface compliance check
func (ϟa *GlTexParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1011_major := uint32(2) // u32
	minRequiredVersion_1011_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1012_major := uint32(3) // u32
		minRequiredVersion_1012_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1012_major, minRequiredVersion_1012_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1013_major := uint32(3) // u32
		minRequiredVersion_1013_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1013_major, minRequiredVersion_1013_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1015_major := uint32(3) // u32
		minRequiredVersion_1015_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1015_major, minRequiredVersion_1015_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1016_major := uint32(3) // u32
		minRequiredVersion_1016_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1016_major, minRequiredVersion_1016_minor
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                     // Contextʳ
	GetContext_1018_result := context                                // Contextʳ
	ctx := GetContext_1018_result                                    // Contextʳ
	id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(ϟa.Target) // TextureId
	t := ctx.Instances.Textures.Get(id)                              // Textureʳ
	switch ϟa.Parameter {
	case GLenum_GL_TEXTURE_MAG_FILTER:
		t.MagFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MIN_FILTER:
		t.MinFilter = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_S:
		t.WrapS = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_WRAP_T:
		t.WrapT = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT:
		t.MaxAnisotropy = float32(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_R:
		t.SwizzleR = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_G:
		t.SwizzleG = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_B:
		t.SwizzleB = GLenum(ϟa.Value)
	case GLenum_GL_TEXTURE_SWIZZLE_A:
		t.SwizzleA = GLenum(ϟa.Value)
	default:
		v := ϟa.Parameter
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexParameteri)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _ = minRequiredVersion_1011_major, minRequiredVersion_1011_minor, context, GetContext_1018_result, ctx, id, t
	return nil
}

var _ = replay.Replayer(&GlTexParameteriv{}) // interface compliance check
func (ϟa *GlTexParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1019_major := uint32(2) // u32
	minRequiredVersion_1019_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
		minRequiredVersion_1020_major := uint32(3) // u32
		minRequiredVersion_1020_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1020_major, minRequiredVersion_1020_minor
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
		minRequiredVersion_1021_major := uint32(3) // u32
		minRequiredVersion_1021_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1021_major, minRequiredVersion_1021_minor
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Pname {
	case GLenum_GL_TEXTURE_MAG_FILTER, GLenum_GL_TEXTURE_MIN_FILTER, GLenum_GL_TEXTURE_WRAP_S, GLenum_GL_TEXTURE_WRAP_T:
	case GLenum_GL_TEXTURE_BASE_LEVEL, GLenum_GL_TEXTURE_COMPARE_FUNC, GLenum_GL_TEXTURE_COMPARE_MODE, GLenum_GL_TEXTURE_MAX_LEVEL, GLenum_GL_TEXTURE_MAX_LOD, GLenum_GL_TEXTURE_MIN_LOD, GLenum_GL_TEXTURE_SWIZZLE_A, GLenum_GL_TEXTURE_SWIZZLE_B, GLenum_GL_TEXTURE_SWIZZLE_G, GLenum_GL_TEXTURE_SWIZZLE_R, GLenum_GL_TEXTURE_WRAP_R:
		minRequiredVersion_1023_major := uint32(3) // u32
		minRequiredVersion_1023_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1023_major, minRequiredVersion_1023_minor
	case GLenum_GL_DEPTH_STENCIL_TEXTURE_MODE:
		minRequiredVersion_1024_major := uint32(3) // u32
		minRequiredVersion_1024_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1024_major, minRequiredVersion_1024_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlTexParameteriv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1019_major, minRequiredVersion_1019_minor
	return nil
}

var _ = replay.Replayer(&GlTexStorage2D{}) // interface compliance check
func (ϟa *GlTexStorage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1026_major := uint32(3) // u32
	minRequiredVersion_1026_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexStorage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1026_major, minRequiredVersion_1026_minor
	return nil
}

var _ = replay.Replayer(&GlTexStorage2DMultisample{}) // interface compliance check
func (ϟa *GlTexStorage2DMultisample) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1029_major := uint32(3) // u32
	minRequiredVersion_1029_minor := uint32(1) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_MULTISAMPLE:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Samples.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Fixedsamplelocations.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexStorage2DMultisample)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1029_major, minRequiredVersion_1029_minor
	return nil
}

var _ = replay.Replayer(&GlTexStorage3D{}) // interface compliance check
func (ϟa *GlTexStorage3D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1032_major := uint32(3) // u32
	minRequiredVersion_1032_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Internalformat {
	case GLenum_GL_COMPRESSED_R11_EAC, GLenum_GL_COMPRESSED_RG11_EAC, GLenum_GL_COMPRESSED_RGB8_ETC2, GLenum_GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_COMPRESSED_RGBA8_ETC2_EAC, GLenum_GL_COMPRESSED_SIGNED_R11_EAC, GLenum_GL_COMPRESSED_SIGNED_RG11_EAC, GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC, GLenum_GL_COMPRESSED_SRGB8_ETC2, GLenum_GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2, GLenum_GL_DEPTH24_STENCIL8, GLenum_GL_DEPTH32F_STENCIL8, GLenum_GL_DEPTH_COMPONENT16, GLenum_GL_DEPTH_COMPONENT24, GLenum_GL_DEPTH_COMPONENT32F, GLenum_GL_R11F_G11F_B10F, GLenum_GL_R16F, GLenum_GL_R16I, GLenum_GL_R16UI, GLenum_GL_R32F, GLenum_GL_R32I, GLenum_GL_R32UI, GLenum_GL_R8, GLenum_GL_R8I, GLenum_GL_R8UI, GLenum_GL_R8_SNORM, GLenum_GL_RG16F, GLenum_GL_RG16I, GLenum_GL_RG16UI, GLenum_GL_RG32F, GLenum_GL_RG32I, GLenum_GL_RG32UI, GLenum_GL_RG8, GLenum_GL_RG8I, GLenum_GL_RG8UI, GLenum_GL_RG8_SNORM, GLenum_GL_RGB10_A2, GLenum_GL_RGB10_A2UI, GLenum_GL_RGB16F, GLenum_GL_RGB16I, GLenum_GL_RGB16UI, GLenum_GL_RGB32F, GLenum_GL_RGB32I, GLenum_GL_RGB32UI, GLenum_GL_RGB565, GLenum_GL_RGB5_A1, GLenum_GL_RGB8, GLenum_GL_RGB8I, GLenum_GL_RGB8UI, GLenum_GL_RGB8_SNORM, GLenum_GL_RGB9_E5, GLenum_GL_RGBA16F, GLenum_GL_RGBA16I, GLenum_GL_RGBA16UI, GLenum_GL_RGBA32F, GLenum_GL_RGBA32I, GLenum_GL_RGBA32UI, GLenum_GL_RGBA4, GLenum_GL_RGBA8, GLenum_GL_RGBA8I, GLenum_GL_RGBA8UI, GLenum_GL_RGBA8_SNORM, GLenum_GL_SRGB8, GLenum_GL_SRGB8_ALPHA8:
	default:
		v := ϟa.Internalformat
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Levels.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Internalformat))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexStorage3D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1032_major, minRequiredVersion_1032_minor
	return nil
}

var _ = replay.Replayer(&GlTexSubImage2D{}) // interface compliance check
func (ϟa *GlTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1035_major := uint32(2) // u32
	minRequiredVersion_1035_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RGB, GLenum_GL_RGBA:
	case GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
		minRequiredVersion_1037_major := uint32(3) // u32
		minRequiredVersion_1037_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1037_major, minRequiredVersion_1037_minor
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1039_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1039_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT:
		minRequiredVersion_1040_major := uint32(3) // u32
		minRequiredVersion_1040_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1040_major, minRequiredVersion_1040_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1042_result := context            // Contextʳ
	ctx := GetContext_1042_result                // Contextʳ
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_2D) // TextureId
		t := ctx.Instances.Textures.Get(id)                                         // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
		}
		t.Texture2D[ϟa.Level] = l
		t.Kind = TextureKind_TEXTURE2D
		t.Format = ϟa.Format
		_, _, _ = id, t, l
	case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y, GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
		id := ctx.TextureUnits.Get(ctx.ActiveTextureUnit).Get(GLenum_GL_TEXTURE_CUBE_MAP) // TextureId
		t := ctx.Instances.Textures.Get(id)                                               // Textureʳ
		l := func() Image {
			s := Image{}
			s.Init()
			s.Width = ϟa.Width
			s.Height = ϟa.Height
			s.Size = externs{ϟs, ϟd, ϟl}.imageSize(uint32(ϟa.Width), uint32(ϟa.Height), ϟa.Format, ϟa.Type)
			s.Format = ϟa.Format
			return s
		}() // Image
		if ((ctx.BoundBuffers.Get(GLenum_GL_PIXEL_UNPACK_BUFFER)) == (BufferId(uint32(0)))) && ((ϟa.Data) != (TexturePointer(Voidᶜᵖ{}))) {
			l.Data = U8ᵖ(ϟa.Data).Slice(uint64(uint32(0)), uint64(l.Size), ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb).Clone(ϟs)
		}
		cube := t.Cubemap.Get(ϟa.Level) // CubemapLevel
		cube.Faces[ϟa.Target] = l
		t.Cubemap[ϟa.Level] = cube
		t.Kind = TextureKind_CUBEMAP
		t.Format = ϟa.Format
		_, _, _, _ = id, t, l, cube
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlTexSubImage2D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1035_major, minRequiredVersion_1035_minor, context, GetContext_1042_result, ctx
	return nil
}

var _ = replay.Replayer(&GlTexSubImage3D{}) // interface compliance check
func (ϟa *GlTexSubImage3D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1043_major := uint32(3) // u32
	minRequiredVersion_1043_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TEXTURE_2D_ARRAY, GLenum_GL_TEXTURE_3D:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Format {
	case GLenum_GL_ALPHA, GLenum_GL_DEPTH_COMPONENT, GLenum_GL_DEPTH_STENCIL, GLenum_GL_LUMINANCE, GLenum_GL_LUMINANCE_ALPHA, GLenum_GL_RED, GLenum_GL_RED_INTEGER, GLenum_GL_RG, GLenum_GL_RGB, GLenum_GL_RGBA, GLenum_GL_RGBA_INTEGER, GLenum_GL_RGB_INTEGER, GLenum_GL_RG_INTEGER:
	default:
		v := ϟa.Format
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1046_ext := ExtensionId_GL_OES_texture_half_float // ExtensionId
		_ = requiresExtension_1046_ext
	case GLenum_GL_BYTE, GLenum_GL_FLOAT, GLenum_GL_FLOAT_32_UNSIGNED_INT_24_8_REV, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_10F_11F_11F_REV, GLenum_GL_UNSIGNED_INT_24_8, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT_5_9_9_9_REV, GLenum_GL_UNSIGNED_SHORT, GLenum_GL_UNSIGNED_SHORT_4_4_4_4, GLenum_GL_UNSIGNED_SHORT_5_5_5_1, GLenum_GL_UNSIGNED_SHORT_5_6_5:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Level.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Xoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Yoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Zoffset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Depth.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Format))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Pixels.value())
	ϟb.Call(funcInfoGlTexSubImage3D)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1043_major, minRequiredVersion_1043_minor
	return nil
}

var _ = replay.Replayer(&GlBeginTransformFeedback{}) // interface compliance check
func (ϟa *GlBeginTransformFeedback) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1048_major := uint32(3) // u32
	minRequiredVersion_1048_minor := uint32(0) // u32
	switch ϟa.PrimitiveMode {
	case GLenum_GL_LINES, GLenum_GL_POINTS, GLenum_GL_TRIANGLES:
	default:
		v := ϟa.PrimitiveMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.PrimitiveMode))
	ϟb.Call(funcInfoGlBeginTransformFeedback)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1048_major, minRequiredVersion_1048_minor
	return nil
}

var _ = replay.Replayer(&GlBindTransformFeedback{}) // interface compliance check
func (ϟa *GlBindTransformFeedback) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1050_major := uint32(3) // u32
	minRequiredVersion_1050_minor := uint32(0) // u32
	switch ϟa.Target {
	case GLenum_GL_TRANSFORM_FEEDBACK:
	default:
		v := ϟa.Target
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(value.U32(ϟa.Target))
	ϟb.Push(ϟa.Id.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBindTransformFeedback)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1050_major, minRequiredVersion_1050_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteTransformFeedbacks{}) // interface compliance check
func (ϟa *GlDeleteTransformFeedbacks) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1052_major := uint32(3) // u32
	minRequiredVersion_1052_minor := uint32(0) // u32
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Ids.value())
	ϟb.Call(funcInfoGlDeleteTransformFeedbacks)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1052_major, minRequiredVersion_1052_minor
	return nil
}

var _ = replay.Replayer(&GlEndTransformFeedback{}) // interface compliance check
func (ϟa *GlEndTransformFeedback) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1053_major := uint32(3) // u32
	minRequiredVersion_1053_minor := uint32(0) // u32
	ϟb.Call(funcInfoGlEndTransformFeedback)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1053_major, minRequiredVersion_1053_minor
	return nil
}

var _ = replay.Replayer(&GlGenTransformFeedbacks{}) // interface compliance check
func (ϟa *GlGenTransformFeedbacks) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1054_major := uint32(3) // u32
	minRequiredVersion_1054_minor := uint32(0) // u32
	ϟb.Push(ϟa.N.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Ids.value())
	ϟb.Call(funcInfoGlGenTransformFeedbacks)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1054_major, minRequiredVersion_1054_minor
	return nil
}

var _ = replay.Replayer(&GlGetTransformFeedbackVarying{}) // interface compliance check
func (ϟa *GlGetTransformFeedbackVarying) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1055_major := uint32(3) // u32
	minRequiredVersion_1055_minor := uint32(0) // u32
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.BufSize.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Length.value())
	ϟb.Push(ϟa.Size.value())
	ϟb.Push(ϟa.Type.value())
	ϟb.Push(ϟa.Name.value())
	ϟb.Call(funcInfoGlGetTransformFeedbackVarying)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1055_major, minRequiredVersion_1055_minor
	return nil
}

var _ = replay.Replayer(&GlIsTransformFeedback{}) // interface compliance check
func (ϟa *GlIsTransformFeedback) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1056_major := uint32(3) // u32
	minRequiredVersion_1056_minor := uint32(0) // u32
	ϟb.Push(ϟa.Id.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlIsTransformFeedback)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1056_major, minRequiredVersion_1056_minor
	return nil
}

var _ = replay.Replayer(&GlPauseTransformFeedback{}) // interface compliance check
func (ϟa *GlPauseTransformFeedback) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1057_major := uint32(3) // u32
	minRequiredVersion_1057_minor := uint32(0) // u32
	ϟb.Call(funcInfoGlPauseTransformFeedback)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1057_major, minRequiredVersion_1057_minor
	return nil
}

var _ = replay.Replayer(&GlResumeTransformFeedback{}) // interface compliance check
func (ϟa *GlResumeTransformFeedback) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1058_major := uint32(3) // u32
	minRequiredVersion_1058_minor := uint32(0) // u32
	ϟb.Call(funcInfoGlResumeTransformFeedback)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1058_major, minRequiredVersion_1058_minor
	return nil
}

var _ = replay.Replayer(&GlTransformFeedbackVaryings{}) // interface compliance check
func (ϟa *GlTransformFeedbackVaryings) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1059_major := uint32(3) // u32
	minRequiredVersion_1059_minor := uint32(0) // u32
	switch ϟa.BufferMode {
	case GLenum_GL_INTERLEAVED_ATTRIBS, GLenum_GL_SEPARATE_ATTRIBS:
	default:
		v := ϟa.BufferMode
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	if key, remap := ϟa.Program.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Program.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Program.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Varyings.value())
	ϟb.Push(value.U32(ϟa.BufferMode))
	ϟb.Call(funcInfoGlTransformFeedbackVaryings)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1059_major, minRequiredVersion_1059_minor
	return nil
}

var _ = replay.Replayer(&GlBindVertexArray{}) // interface compliance check
func (ϟa *GlBindVertexArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1061_major := uint32(3)   // u32
	minRequiredVersion_1061_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1062_result := context            // Contextʳ
	ctx := GetContext_1062_result                // Contextʳ
	if !(ctx.Instances.VertexArrays.Contains(ϟa.Array)) {
		ctx.Instances.VertexArrays[ϟa.Array] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
	}
	ctx.BoundVertexArray = ϟa.Array
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlBindVertexArray)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1061_major, minRequiredVersion_1061_minor, context, GetContext_1062_result, ctx
	return nil
}

var _ = replay.Replayer(&GlBindVertexBuffer{}) // interface compliance check
func (ϟa *GlBindVertexBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1063_major := uint32(3) // u32
	minRequiredVersion_1063_minor := uint32(1) // u32
	ϟb.Push(ϟa.Bindingindex.value(ϟb, ϟa, ϟs))
	if key, remap := ϟa.Buffer.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Buffer.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Buffer.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(ϟa.Offset.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Stride.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlBindVertexBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1063_major, minRequiredVersion_1063_minor
	return nil
}

var _ = replay.Replayer(&GlDeleteVertexArrays{}) // interface compliance check
func (ϟa *GlDeleteVertexArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1064_major := uint32(3)                            // u32
	minRequiredVersion_1064_minor := uint32(0)                            // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_1065_result := context                                     // Contextʳ
	ctx := GetContext_1065_result                                         // Contextʳ
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		delete(ctx.Instances.VertexArrays, a.Index(uint64(i), ϟs).replayRead(ϟa, ϟs, ϟd, ϟl, ϟb))
	}
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Arrays.value())
	ϟb.Call(funcInfoGlDeleteVertexArrays)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1064_major, minRequiredVersion_1064_minor, context, GetContext_1065_result, ctx, a
	return nil
}

var _ = replay.Replayer(&GlDisableVertexAttribArray{}) // interface compliance check
func (ϟa *GlDisableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1066_major := uint32(2)   // u32
	minRequiredVersion_1066_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1067_result := context            // Contextʳ
	ctx := GetContext_1067_result                // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = false
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlDisableVertexAttribArray)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1066_major, minRequiredVersion_1066_minor, context, GetContext_1067_result, ctx
	return nil
}

var _ = replay.Replayer(&GlEnableVertexAttribArray{}) // interface compliance check
func (ϟa *GlEnableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1068_major := uint32(2)   // u32
	minRequiredVersion_1068_minor := uint32(0)   // u32
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1069_result := context            // Contextʳ
	ctx := GetContext_1069_result                // Contextʳ
	ctx.VertexAttributeArrays.Get(ϟa.Location).Enabled = true
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlEnableVertexAttribArray)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _ = minRequiredVersion_1068_major, minRequiredVersion_1068_minor, context, GetContext_1069_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGenVertexArrays{}) // interface compliance check
func (ϟa *GlGenVertexArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1070_major := uint32(3)                            // u32
	minRequiredVersion_1070_minor := uint32(0)                            // u32
	a := ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs) // VertexArrayIdˢ
	context := ϟc.Contexts.Get(ϟc.CurrentThread)                          // Contextʳ
	GetContext_1071_result := context                                     // Contextʳ
	ctx := GetContext_1071_result                                         // Contextʳ
	ϟb.Push(ϟa.Count.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Arrays.value())
	ϟb.Call(funcInfoGlGenVertexArrays)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	for i := GLsizei(GLsizei(int32(0))); i < ϟa.Count; i++ {
		id := VertexArrayId(ϟa.Arrays.Slice(uint64(GLsizei(int32(0))), uint64(ϟa.Count), ϟs).Index(uint64(i), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)) // VertexArrayId
		ctx.Instances.VertexArrays[id] = func() *VertexArray {
			s := &VertexArray{}
			s.Init()
			return s
		}()
		a.Index(uint64(i), ϟs).replayWrite(id, ϟa, ϟs, ϟd, ϟl, ϟb)
		_ = id
	}
	_, _, _, _, _, _ = minRequiredVersion_1070_major, minRequiredVersion_1070_minor, a, context, GetContext_1071_result, ctx
	return nil
}

var _ = replay.Replayer(&GlGetVertexAttribIiv{}) // interface compliance check
func (ϟa *GlGetVertexAttribIiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1072_major := uint32(3) // u32
	minRequiredVersion_1072_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1073_major := uint32(3) // u32
		minRequiredVersion_1073_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1073_major, minRequiredVersion_1073_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetVertexAttribIiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1072_major, minRequiredVersion_1072_minor
	return nil
}

var _ = replay.Replayer(&GlGetVertexAttribIuiv{}) // interface compliance check
func (ϟa *GlGetVertexAttribIuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1075_major := uint32(3) // u32
	minRequiredVersion_1075_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1076_major := uint32(3) // u32
		minRequiredVersion_1076_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1076_major, minRequiredVersion_1076_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetVertexAttribIuiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1075_major, minRequiredVersion_1075_minor
	return nil
}

var _ = replay.Replayer(&GlGetVertexAttribPointerv{}) // interface compliance check
func (ϟa *GlGetVertexAttribPointerv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1078_major := uint32(2) // u32
	minRequiredVersion_1078_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_POINTER:
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Pointer.value())
	ϟb.Call(funcInfoGlGetVertexAttribPointerv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1078_major, minRequiredVersion_1078_minor
	return nil
}

var _ = replay.Replayer(&GlGetVertexAttribfv{}) // interface compliance check
func (ϟa *GlGetVertexAttribfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1080_major := uint32(2) // u32
	minRequiredVersion_1080_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1081_major := uint32(3) // u32
		minRequiredVersion_1081_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1081_major, minRequiredVersion_1081_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1082_major := uint32(3) // u32
		minRequiredVersion_1082_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1082_major, minRequiredVersion_1082_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetVertexAttribfv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1080_major, minRequiredVersion_1080_minor
	return nil
}

var _ = replay.Replayer(&GlGetVertexAttribiv{}) // interface compliance check
func (ϟa *GlGetVertexAttribiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1084_major := uint32(2) // u32
	minRequiredVersion_1084_minor := uint32(0) // u32
	switch ϟa.Pname {
	case GLenum_GL_CURRENT_VERTEX_ATTRIB, GLenum_GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING, GLenum_GL_VERTEX_ATTRIB_ARRAY_ENABLED, GLenum_GL_VERTEX_ATTRIB_ARRAY_NORMALIZED, GLenum_GL_VERTEX_ATTRIB_ARRAY_SIZE, GLenum_GL_VERTEX_ATTRIB_ARRAY_STRIDE, GLenum_GL_VERTEX_ATTRIB_ARRAY_TYPE:
	case GLenum_GL_VERTEX_ATTRIB_ARRAY_DIVISOR, GLenum_GL_VERTEX_ATTRIB_ARRAY_INTEGER:
		minRequiredVersion_1085_major := uint32(3) // u32
		minRequiredVersion_1085_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1085_major, minRequiredVersion_1085_minor
	case GLenum_GL_VERTEX_ATTRIB_BINDING:
		minRequiredVersion_1086_major := uint32(3) // u32
		minRequiredVersion_1086_minor := uint32(1) // u32
		_, _ = minRequiredVersion_1086_major, minRequiredVersion_1086_minor
	default:
		v := ϟa.Pname
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Pname))
	ϟb.Push(ϟa.Params.value())
	ϟb.Call(funcInfoGlGetVertexAttribiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1084_major, minRequiredVersion_1084_minor
	return nil
}

var _ = replay.Replayer(&GlIsVertexArray{}) // interface compliance check
func (ϟa *GlIsVertexArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1088_major := uint32(3) // u32
	minRequiredVersion_1088_minor := uint32(0) // u32
	if key, remap := ϟa.Array.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Array.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Array.value(ϟb, ϟa, ϟs))
	}
	ϟb.Call(funcInfoGlIsVertexArray)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1088_major, minRequiredVersion_1088_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib1f{}) // interface compliance check
func (ϟa *GlVertexAttrib1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1089_major := uint32(2) // u32
	minRequiredVersion_1089_minor := uint32(0) // u32
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttrib1f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1089_major, minRequiredVersion_1089_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib1fv{}) // interface compliance check
func (ϟa *GlVertexAttrib1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1090_major := uint32(2) // u32
	minRequiredVersion_1090_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib1fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1090_major, minRequiredVersion_1090_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib2f{}) // interface compliance check
func (ϟa *GlVertexAttrib2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1091_major := uint32(2) // u32
	minRequiredVersion_1091_minor := uint32(0) // u32
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttrib2f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1091_major, minRequiredVersion_1091_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib2fv{}) // interface compliance check
func (ϟa *GlVertexAttrib2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1092_major := uint32(2) // u32
	minRequiredVersion_1092_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(2), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib2fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1092_major, minRequiredVersion_1092_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib3f{}) // interface compliance check
func (ϟa *GlVertexAttrib3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1093_major := uint32(2) // u32
	minRequiredVersion_1093_minor := uint32(0) // u32
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value2.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttrib3f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1093_major, minRequiredVersion_1093_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib3fv{}) // interface compliance check
func (ϟa *GlVertexAttrib3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1094_major := uint32(2) // u32
	minRequiredVersion_1094_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(3), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib3fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1094_major, minRequiredVersion_1094_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib4f{}) // interface compliance check
func (ϟa *GlVertexAttrib4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1095_major := uint32(2) // u32
	minRequiredVersion_1095_minor := uint32(0) // u32
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value0.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value1.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value2.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value3.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttrib4f)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1095_major, minRequiredVersion_1095_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttrib4fv{}) // interface compliance check
func (ϟa *GlVertexAttrib4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1096_major := uint32(2) // u32
	minRequiredVersion_1096_minor := uint32(0) // u32
	ϟa.Value.Slice(uint64(0), uint64(4), ϟs).OnRead(ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlVertexAttrib4fv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1096_major, minRequiredVersion_1096_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribBinding{}) // interface compliance check
func (ϟa *GlVertexAttribBinding) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1097_major := uint32(3) // u32
	minRequiredVersion_1097_minor := uint32(1) // u32
	ϟb.Push(ϟa.Attribindex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Bindingindex.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribBinding)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1097_major, minRequiredVersion_1097_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribDivisor{}) // interface compliance check
func (ϟa *GlVertexAttribDivisor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1098_major := uint32(3) // u32
	minRequiredVersion_1098_minor := uint32(0) // u32
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Divisor.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribDivisor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1098_major, minRequiredVersion_1098_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribFormat{}) // interface compliance check
func (ϟa *GlVertexAttribFormat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1099_major := uint32(3) // u32
	minRequiredVersion_1099_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Attribindex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Normalized.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Relativeoffset.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribFormat)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1099_major, minRequiredVersion_1099_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribI4i{}) // interface compliance check
func (ϟa *GlVertexAttribI4i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1101_major := uint32(3) // u32
	minRequiredVersion_1101_minor := uint32(0) // u32
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Z.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.W.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribI4i)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1101_major, minRequiredVersion_1101_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribI4iv{}) // interface compliance check
func (ϟa *GlVertexAttribI4iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1102_major := uint32(3) // u32
	minRequiredVersion_1102_minor := uint32(0) // u32
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlVertexAttribI4iv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1102_major, minRequiredVersion_1102_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribI4ui{}) // interface compliance check
func (ϟa *GlVertexAttribI4ui) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1103_major := uint32(3) // u32
	minRequiredVersion_1103_minor := uint32(0) // u32
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.X.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Y.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Z.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.W.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribI4ui)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1103_major, minRequiredVersion_1103_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribI4uiv{}) // interface compliance check
func (ϟa *GlVertexAttribI4uiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1104_major := uint32(3) // u32
	minRequiredVersion_1104_minor := uint32(0) // u32
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.V.value())
	ϟb.Call(funcInfoGlVertexAttribI4uiv)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1104_major, minRequiredVersion_1104_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribIFormat{}) // interface compliance check
func (ϟa *GlVertexAttribIFormat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1105_major := uint32(3) // u32
	minRequiredVersion_1105_minor := uint32(1) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Attribindex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Relativeoffset.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribIFormat)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1105_major, minRequiredVersion_1105_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribIPointer{}) // interface compliance check
func (ϟa *GlVertexAttribIPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1107_major := uint32(3) // u32
	minRequiredVersion_1107_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1108_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1108_ext
	case GLenum_GL_BYTE, GLenum_GL_INT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_SHORT:
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	ϟb.Push(ϟa.Index.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(ϟa.Stride.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Pointer.value())
	ϟb.Call(funcInfoGlVertexAttribIPointer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1107_major, minRequiredVersion_1107_minor
	return nil
}

var _ = replay.Replayer(&GlVertexAttribPointer{}) // interface compliance check
func (ϟa *GlVertexAttribPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1110_major := uint32(2) // u32
	minRequiredVersion_1110_minor := uint32(0) // u32
	switch ϟa.Type {
	case GLenum_GL_BYTE, GLenum_GL_FIXED, GLenum_GL_FLOAT, GLenum_GL_SHORT, GLenum_GL_UNSIGNED_BYTE, GLenum_GL_UNSIGNED_SHORT:
	case GLenum_GL_HALF_FLOAT_OES:
		requiresExtension_1111_ext := ExtensionId_GL_OES_vertex_half_float // ExtensionId
		_ = requiresExtension_1111_ext
	case GLenum_GL_HALF_FLOAT, GLenum_GL_INT, GLenum_GL_INT_2_10_10_10_REV, GLenum_GL_UNSIGNED_INT, GLenum_GL_UNSIGNED_INT_2_10_10_10_REV:
		minRequiredVersion_1112_major := uint32(3) // u32
		minRequiredVersion_1112_minor := uint32(0) // u32
		_, _ = minRequiredVersion_1112_major, minRequiredVersion_1112_minor
	default:
		v := ϟa.Type
		return fmt.Errorf("Missing switch case handler for value %T %v", v, v)
	}
	context := ϟc.Contexts.Get(ϟc.CurrentThread)    // Contextʳ
	GetContext_1114_result := context               // Contextʳ
	ctx := GetContext_1114_result                   // Contextʳ
	a := ctx.VertexAttributeArrays.Get(ϟa.Location) // VertexAttributeArrayʳ
	a.Size = uint32(ϟa.Size)
	a.Type = ϟa.Type
	a.Normalized = ϟa.Normalized
	a.Stride = ϟa.Stride
	a.Pointer = ϟa.Data
	a.Buffer = ctx.BoundBuffers.Get(GLenum_GL_ARRAY_BUFFER)
	ϟb.Push(ϟa.Location.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Size.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.Type))
	ϟb.Push(value.Bool(ϟa.Normalized))
	ϟb.Push(ϟa.Stride.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Data.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexAttribPointer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _ = minRequiredVersion_1110_major, minRequiredVersion_1110_minor, context, GetContext_1114_result, ctx, a
	return nil
}

var _ = replay.Replayer(&GlVertexBindingDivisor{}) // interface compliance check
func (ϟa *GlVertexBindingDivisor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	minRequiredVersion_1115_major := uint32(3) // u32
	minRequiredVersion_1115_minor := uint32(1) // u32
	ϟb.Push(ϟa.Bindingindex.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Divisor.value(ϟb, ϟa, ϟs))
	ϟb.Call(funcInfoGlVertexBindingDivisor)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _ = minRequiredVersion_1115_major, minRequiredVersion_1115_minor
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjecti64v{}) // interface compliance check
func (ϟa *GlGetQueryObjecti64v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjecti64v)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&GlGetQueryObjectui64v{}) // interface compliance check
func (ϟa *GlGetQueryObjectui64v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	if key, remap := ϟa.Query.remap(ϟa, ϟs); remap {
		loadRemap(ϟb, key, protocol.TypeUint32, ϟa.Query.value(ϟb, ϟa, ϟs))
	} else {
		ϟb.Push(ϟa.Query.value(ϟb, ϟa, ϟs))
	}
	ϟb.Push(value.U32(ϟa.Parameter))
	ϟb.Push(ϟa.Value.value())
	ϟb.Call(funcInfoGlGetQueryObjectui64v)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayWrite(ϟa.Value.Slice(uint64(0), uint64(1), ϟs).Index(uint64(0), ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb)
	return nil
}

var _ = replay.Replayer(&ReplayCreateRenderer{}) // interface compliance check
func (ϟa *ReplayCreateRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Id))
	ϟb.Call(funcInfoReplayCreateRenderer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&ReplayBindRenderer{}) // interface compliance check
func (ϟa *ReplayBindRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U32(ϟa.Id))
	ϟb.Call(funcInfoReplayBindRenderer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&BackbufferInfo{}) // interface compliance check
func (ϟa *BackbufferInfo) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	context := ϟc.Contexts.Get(ϟc.CurrentThread) // Contextʳ
	GetContext_1127_result := context            // Contextʳ
	ctx := GetContext_1127_result                // Contextʳ
	ctx.PreserveBuffersOnSwap = ϟa.PreserveBuffersOnSwap
	backbuffer := ctx.Instances.Framebuffers.Get(FramebufferId(uint32(0)))                        // Framebufferʳ
	color_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_COLOR_ATTACHMENT0).Object)    // RenderbufferId
	color_buffer := ctx.Instances.Renderbuffers.Get(color_id)                                     // Renderbufferʳ
	depth_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_DEPTH_ATTACHMENT).Object)     // RenderbufferId
	depth_buffer := ctx.Instances.Renderbuffers.Get(depth_id)                                     // Renderbufferʳ
	stencil_id := RenderbufferId(backbuffer.Attachments.Get(GLenum_GL_STENCIL_ATTACHMENT).Object) // RenderbufferId
	stencil_buffer := ctx.Instances.Renderbuffers.Get(stencil_id)                                 // Renderbufferʳ
	color_buffer.Width = ϟa.Width
	color_buffer.Height = ϟa.Height
	color_buffer.Format = ϟa.ColorFmt
	depth_buffer.Width = ϟa.Width
	depth_buffer.Height = ϟa.Height
	depth_buffer.Format = ϟa.DepthFmt
	stencil_buffer.Width = ϟa.Width
	stencil_buffer.Height = ϟa.Height
	stencil_buffer.Format = ϟa.StencilFmt
	if ϟa.ResetViewportScissor {
		ctx.Rasterizing.Scissor.Width = ϟa.Width
		ctx.Rasterizing.Scissor.Height = ϟa.Height
		ctx.Rasterizing.Viewport.Width = ϟa.Width
		ctx.Rasterizing.Viewport.Height = ϟa.Height
	}
	ϟb.Push(ϟa.Width.value(ϟb, ϟa, ϟs))
	ϟb.Push(ϟa.Height.value(ϟb, ϟa, ϟs))
	ϟb.Push(value.U32(ϟa.ColorFmt))
	ϟb.Push(value.U32(ϟa.DepthFmt))
	ϟb.Push(value.U32(ϟa.StencilFmt))
	ϟb.Push(value.Bool(ϟa.ResetViewportScissor))
	ϟb.Push(value.Bool(ϟa.PreserveBuffersOnSwap))
	ϟb.Call(funcInfoBackbufferInfo)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	_, _, _, _, _, _, _, _, _, _ = context, GetContext_1127_result, ctx, backbuffer, color_id, color_buffer, depth_id, depth_buffer, stencil_id, stencil_buffer
	return nil
}

var _ = replay.Replayer(&StartTimer{}) // interface compliance check
func (ϟa *StartTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U8(ϟa.Index))
	ϟb.Call(funcInfoStartTimer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&StopTimer{}) // interface compliance check
func (ϟa *StopTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Push(value.U8(ϟa.Index))
	ϟb.Call(funcInfoStopTimer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}

var _ = replay.Replayer(&FlushPostBuffer{}) // interface compliance check
func (ϟa *FlushPostBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error {
	ϟc := getState(ϟs)
	_ = ϟc
	ϟa.observations.ApplyReads(ϟs.Memory[memory.ApplicationPool])
	ϟb.Call(funcInfoFlushPostBuffer)
	ϟa.observations.ApplyWrites(ϟs.Memory[memory.ApplicationPool])
	return nil
}
func (p Voidᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Voidᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U8ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint8 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint8 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U8ᵖ) replayWrite(value uint8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U8ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLcharᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLchar {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLchar {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᵖ) replayWrite(value GLchar, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLcharᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Charᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) byte {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) byte {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Charᵖ) replayWrite(value byte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Charᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLuintᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuintᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuintᶜᵖ) replayWrite(value GLuint, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLuintᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLuintᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuintᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuintᵖ) replayWrite(value GLuint, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLuintᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLcharᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLchar {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLchar {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᶜᵖ) replayWrite(value GLchar, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLcharᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLenumᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLenum {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLenumᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLenum {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLenumᵖ) replayWrite(value GLenum, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLenumᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLsizeiᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLsizei {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLsizeiᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLsizei {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLsizeiᵖ) replayWrite(value GLsizei, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLsizeiᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Voidᵖᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Voidᵖᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Voidᵖᵖ) replayWrite(value Voidᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Voidᵖᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLintᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLintᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLintᵖ) replayWrite(value GLint, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLintᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLintᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLintᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLintᶜᵖ) replayWrite(value GLint, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLintᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p QueryIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p QueryIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p QueryIdᶜᵖ) replayWrite(value QueryId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p QueryIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p QueryIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p QueryIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p QueryIdᵖ) replayWrite(value QueryId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p QueryIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p BufferIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p BufferIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p BufferIdᶜᵖ) replayWrite(value BufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p BufferIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p BufferIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p BufferIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p BufferIdᵖ) replayWrite(value BufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p BufferIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLint64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLint64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLint64ᵖ) replayWrite(value GLint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLint64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLenumᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLenum {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLenumᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLenum {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLenumᶜᵖ) replayWrite(value GLenum, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLenumᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLfloatᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLfloat {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLfloatᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLfloat {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLfloatᶜᵖ) replayWrite(value GLfloat, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLfloatᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLfloatᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLfloat {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLfloatᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLfloat {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLfloatᵖ) replayWrite(value GLfloat, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLfloatᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLcharᶜᵖᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharᶜᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᶜᵖᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharᶜᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᶜᵖᵖ) replayWrite(value GLcharᶜᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLcharᶜᵖᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p PipelineIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) PipelineId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p PipelineIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) PipelineId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p PipelineIdᶜᵖ) replayWrite(value PipelineId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p PipelineIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p PipelineIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) PipelineId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p PipelineIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) PipelineId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p PipelineIdᵖ) replayWrite(value PipelineId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p PipelineIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p VertexArrayIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p VertexArrayIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p VertexArrayIdᶜᵖ) replayWrite(value VertexArrayId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p VertexArrayIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p VertexArrayIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p VertexArrayIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p VertexArrayIdᵖ) replayWrite(value VertexArrayId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p VertexArrayIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p FramebufferIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferIdᵖ) replayWrite(value FramebufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p FramebufferIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p ProgramIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ProgramId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ProgramIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ProgramId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ProgramIdᵖ) replayWrite(value ProgramId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p ProgramIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p RenderbufferIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p RenderbufferIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p RenderbufferIdᵖ) replayWrite(value RenderbufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p RenderbufferIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p ShaderIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderIdᵖ) replayWrite(value ShaderId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p ShaderIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p TextureIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TextureIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TextureIdᵖ) replayWrite(value TextureId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p TextureIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLubyteᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLubyte {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLubyteᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLubyte {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLubyteᵖ) replayWrite(value GLubyte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLubyteᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLuint64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuint64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuint64ᵖ) replayWrite(value GLuint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLuint64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLvoidᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLsizeiᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLsizei {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLsizeiᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLsizei {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLsizeiᶜᵖ) replayWrite(value GLsizei, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLsizeiᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Voidᶜᵖᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᶜᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Voidᶜᵖᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᶜᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Voidᶜᵖᶜᵖ) replayWrite(value Voidᶜᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Voidᶜᵖᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Voidᶜᵖᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᶜᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Voidᶜᵖᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᶜᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Voidᶜᵖᵖ) replayWrite(value Voidᶜᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Voidᶜᵖᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLubyteᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLubyte {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLubyteᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLubyte {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLubyteᶜᵖ) replayWrite(value GLubyte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLubyteᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLuint64ᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuint64ᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLuint64ᶜᵖ) replayWrite(value GLuint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLuint64ᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p FramebufferIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p FramebufferIdᶜᵖ) replayWrite(value FramebufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p FramebufferIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p RenderbufferIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p RenderbufferIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p RenderbufferIdᶜᵖ) replayWrite(value RenderbufferId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p RenderbufferIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLcharᶜᵖᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharᶜᵖ {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᶜᵖᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharᶜᵖ {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLcharᶜᵖᶜᵖ) replayWrite(value GLcharᶜᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLcharᶜᵖᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p ShaderIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p ShaderIdᶜᵖ) replayWrite(value ShaderId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p ShaderIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Vec2fᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2f {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec2fᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2f {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec2fᵖ) replayWrite(value Vec2f, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Vec2fᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Vec2iᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2i {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec2iᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2i {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec2iᵖ) replayWrite(value Vec2i, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Vec2iᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Vec3fᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3f {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec3fᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3f {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec3fᵖ) replayWrite(value Vec3f, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Vec3fᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Vec3iᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3i {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec3iᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3i {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec3iᵖ) replayWrite(value Vec3i, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Vec3iᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Vec4fᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4f {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec4fᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4f {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec4fᵖ) replayWrite(value Vec4f, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Vec4fᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Vec4iᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4i {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec4iᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4i {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Vec4iᵖ) replayWrite(value Vec4i, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Vec4iᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Mat2fᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat2f {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Mat2fᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat2f {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Mat2fᵖ) replayWrite(value Mat2f, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Mat2fᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Mat3fᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat3f {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Mat3fᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat3f {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Mat3fᵖ) replayWrite(value Mat3f, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Mat3fᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Mat4fᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat4f {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Mat4fᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat4f {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Mat4fᵖ) replayWrite(value Mat4f, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Mat4fᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p GLbooleanᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLboolean {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLbooleanᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLboolean {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p GLbooleanᵖ) replayWrite(value GLboolean, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p GLbooleanᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Boolᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) bool {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Boolᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) bool {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Boolᵖ) replayWrite(value bool, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Boolᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p SamplerIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) SamplerId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p SamplerIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) SamplerId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p SamplerIdᶜᵖ) replayWrite(value SamplerId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p SamplerIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p SamplerIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) SamplerId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p SamplerIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) SamplerId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p SamplerIdᵖ) replayWrite(value SamplerId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p SamplerIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p TextureIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TextureIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TextureIdᶜᵖ) replayWrite(value TextureId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p TextureIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p TransformFeedbackIdᶜᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TransformFeedbackId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TransformFeedbackIdᶜᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TransformFeedbackId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TransformFeedbackIdᶜᵖ) replayWrite(value TransformFeedbackId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p TransformFeedbackIdᶜᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p TransformFeedbackIdᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TransformFeedbackId {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TransformFeedbackIdᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TransformFeedbackId {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p TransformFeedbackIdᵖ) replayWrite(value TransformFeedbackId, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p TransformFeedbackIdᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p EGLintᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLint {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p EGLintᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLint {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p EGLintᵖ) replayWrite(value EGLint, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p EGLintᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p Intᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Intᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p Intᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p Intᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p CGLContextObjᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObj {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGLContextObjᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObj {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGLContextObjᵖ) replayWrite(value CGLContextObj, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p CGLContextObjᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p CGSConnectionIDᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSConnectionID {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGSConnectionIDᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSConnectionID {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGSConnectionIDᵖ) replayWrite(value CGSConnectionID, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p CGSConnectionIDᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p CGSWindowIDᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSWindowID {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGSWindowIDᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSWindowID {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGSWindowIDᵖ) replayWrite(value CGSWindowID, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p CGSWindowIDᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p CGSSurfaceIDᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSSurfaceID {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGSSurfaceIDᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSSurfaceID {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p CGSSurfaceIDᵖ) replayWrite(value CGSSurfaceID, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p CGSSurfaceIDᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p F64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p F64ᵖ) replayWrite(value float64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p F64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p S64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p S64ᵖ) replayWrite(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p S64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (p U64ᵖ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	p.Slice(0, 1, ϟs).replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U64ᵖ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	p.Slice(0, 1, ϟs).onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p.Read(ϟs, ϟd, ϟl)
}
func (p U64ᵖ) replayWrite(value uint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Write(value, ϟs)
	p.Slice(0, 1, ϟs).onReplayWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
}
func (p U64ᵖ) value() value.Pointer {
	if p.Address != 0 {
		return value.RemappedPointer(p.Address)
	} else {
		return value.AbsolutePointer(0)
	}
}
func (s Boolˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Boolˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Boolˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Boolˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []bool {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s BufferIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s BufferIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) BufferIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s BufferIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s BufferIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []BufferId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s CGLContextObjˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObjˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value(ϟb, ϟa, ϟs))
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s CGLContextObjˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGLContextObjˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s CGLContextObjˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s CGLContextObjˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []CGLContextObj {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s CGSConnectionIDˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSConnectionIDˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value(ϟb, ϟa, ϟs))
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s CGSConnectionIDˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSConnectionIDˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s CGSConnectionIDˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s CGSConnectionIDˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []CGSConnectionID {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s CGSSurfaceIDˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSSurfaceIDˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s CGSSurfaceIDˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSSurfaceIDˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s CGSSurfaceIDˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s CGSSurfaceIDˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []CGSSurfaceID {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s CGSWindowIDˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSWindowIDˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s CGSWindowIDˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) CGSWindowIDˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s CGSWindowIDˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s CGSWindowIDˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []CGSWindowID {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Charˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Charˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Charˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Charˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []byte {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s EGLintˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLintˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s EGLintˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) EGLintˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s EGLintˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s EGLintˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []EGLint {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s F64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s F64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s F64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s F64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []float64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s FramebufferIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s FramebufferIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) FramebufferIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s FramebufferIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s FramebufferIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []FramebufferId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLbooleanˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLbooleanˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLbooleanˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLbooleanˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLbooleanˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLbooleanˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLboolean {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLcharˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLcharˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLcharˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLcharˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLchar {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLcharᶜᵖˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharᶜᵖˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value())
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s GLcharᶜᵖˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLcharᶜᵖˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLcharᶜᵖˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLcharᶜᵖˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLcharᶜᵖ {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLenumˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLenumˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLenumˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLenumˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLenumˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLenumˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLenum {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLfloatˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLfloatˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLfloatˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLfloatˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLfloatˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLfloatˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLfloat {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLint64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLint64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLint64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLint64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLint64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLint64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLintˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLintˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLintˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLintˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLintˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLintˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLint {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLsizeiˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLsizeiˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLsizeiˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLsizeiˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLsizeiˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLsizeiˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLsizei {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLubyteˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLubyteˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLubyteˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLubyteˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLubyteˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLubyteˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLubyte {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLuint64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLuint64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuint64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLuint64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLuint64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLuint64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLuintˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuintˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLuintˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLuintˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLuintˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s GLuintˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []GLuint {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s GLvoidˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLvoidˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s GLvoidˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) GLvoidˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s GLvoidˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Intˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Intˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Intˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Intˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Mat2fˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat2fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Mat2fˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat2fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Mat2fˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Mat2fˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Mat2f {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Mat3fˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat3fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Mat3fˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat3fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Mat3fˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Mat3fˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Mat3f {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Mat4fˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat4fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Mat4fˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Mat4fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Mat4fˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Mat4fˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Mat4f {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s PipelineIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) PipelineIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s PipelineIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) PipelineIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s PipelineIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s PipelineIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []PipelineId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s ProgramIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ProgramIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s ProgramIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ProgramIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s ProgramIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s ProgramIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []ProgramId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s QueryIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s QueryIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) QueryIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s QueryIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s QueryIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []QueryId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s RenderbufferIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s RenderbufferIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) RenderbufferIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s RenderbufferIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s RenderbufferIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []RenderbufferId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s S64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s S64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s S64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s S64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s SamplerIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) SamplerIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s SamplerIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) SamplerIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s SamplerIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s SamplerIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []SamplerId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s ShaderIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s ShaderIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) ShaderIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s ShaderIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s ShaderIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []ShaderId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s TextureIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s TextureIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TextureIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s TextureIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s TextureIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []TextureId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s TransformFeedbackIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TransformFeedbackIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s TransformFeedbackIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) TransformFeedbackIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s TransformFeedbackIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s TransformFeedbackIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []TransformFeedbackId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s U64ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U64ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s U64ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s U64ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint64 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s U8ˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s U8ˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s U8ˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s U8ˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint8 {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Vec2fˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Vec2fˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Vec2fˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Vec2fˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Vec2f {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Vec2iˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2iˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Vec2iˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec2iˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Vec2iˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Vec2iˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Vec2i {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Vec3fˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Vec3fˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Vec3fˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Vec3fˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Vec3f {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Vec3iˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3iˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Vec3iˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec3iˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Vec3iˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Vec3iˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Vec3i {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Vec4fˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Vec4fˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4fˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Vec4fˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Vec4fˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Vec4f {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Vec4iˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4iˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Vec4iˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Vec4iˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Vec4iˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Vec4iˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Vec4i {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s VertexArrayIdˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s VertexArrayIdˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) VertexArrayIdˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(size)
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
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
func (s VertexArrayIdˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s VertexArrayIdˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []VertexArrayId {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Voidˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}
func (s Voidˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Voidˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Voidᵖˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᵖˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value())
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s Voidᵖˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᵖˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Voidᵖˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Voidᵖˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Voidᵖ {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
func (s Voidᶜᵖˢ) onReplayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᶜᵖˢ {
	if s.Root.Pool == memory.ApplicationPool {
		s.replayMap(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs))
		for _, v := range s.Read(ϟs, ϟd, ϟl) {
			ϟb.Push(v.value())
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}
func (s Voidᶜᵖˢ) onReplayWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᶜᵖˢ {
	if s.Root.Pool == memory.ApplicationPool {
		ϟb.MapMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Voidᶜᵖˢ) replayMap(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	if s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.MapMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
}
func (s Voidᶜᵖˢ) replayRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Voidᶜᵖ {
	s.onReplayRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return s.Read(ϟs, ϟd, ϟl)
}
