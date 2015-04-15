; Copyright (C) 2015 The Android Open Source Project
;
; Licensed under the Apache License, Version 2.0 (the "License");
; you may not use this file except in compliance with the License.
; You may obtain a copy of the License at
;
;      http://www.apache.org/licenses/LICENSE-2.0
;
; Unless required by applicable law or agreed to in writing, software
; distributed under the License is distributed on an "AS IS" BASIS,
; WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
; See the License for the specific language governing permissions and
; limitations under the License.

; This masm exposes trampoline functions for OpenGL 1.1 and WGL functions that
; are expected to be exposed by OpenGL32.dll but not declared by the API files.

; Lazy resolve all functions, then trampoline to the resolved address
TRAMPOLINE macro FUNC, REAL
    local skip
    public FUNC
    extern REAL:proc

FUNC PROC
    mov eax, resolved
    cmp eax, 0
    jnz skip
    ; Function are not resolved, resolve now...

    ; Preserve arguments stored in registers (ecx, edx, r8 and r9)
    push rcx
    push rdx
    push r8
    push r9

    ; give resolve() the necessary 32 byte 'shadow space', and call resolve
    sub rsp, 28h
    call resolve
    add rsp, 28h

    ; Reload argument registers
    pop r9
    pop r8
    pop rdx
    pop rcx
skip:
    mov rax, [REAL]
    mov rax, [rax]
    jmp rax
FUNC ENDP
    endm

extern resolve:proc
extern resolved:SDWORD

.CODE
TRAMPOLINE glAccum, real__glAccum
TRAMPOLINE glAlphaFunc, real__glAlphaFunc
TRAMPOLINE glAreTexturesResident, real__glAreTexturesResident
TRAMPOLINE glArrayElement, real__glArrayElement
TRAMPOLINE glBegin, real__glBegin
TRAMPOLINE glBitmap, real__glBitmap
TRAMPOLINE glCallList, real__glCallList
TRAMPOLINE glCallLists, real__glCallLists
TRAMPOLINE glClearAccum, real__glClearAccum
TRAMPOLINE glClearDepth, real__glClearDepth
TRAMPOLINE glClearIndex, real__glClearIndex
TRAMPOLINE glClipPlane, real__glClipPlane
TRAMPOLINE glColor3b, real__glColor3b
TRAMPOLINE glColor3bv, real__glColor3bv
TRAMPOLINE glColor3d, real__glColor3d
TRAMPOLINE glColor3dv, real__glColor3dv
TRAMPOLINE glColor3f, real__glColor3f
TRAMPOLINE glColor3fv, real__glColor3fv
TRAMPOLINE glColor3i, real__glColor3i
TRAMPOLINE glColor3iv, real__glColor3iv
TRAMPOLINE glColor3s, real__glColor3s
TRAMPOLINE glColor3sv, real__glColor3sv
TRAMPOLINE glColor3ub, real__glColor3ub
TRAMPOLINE glColor3ubv, real__glColor3ubv
TRAMPOLINE glColor3ui, real__glColor3ui
TRAMPOLINE glColor3uiv, real__glColor3uiv
TRAMPOLINE glColor3us, real__glColor3us
TRAMPOLINE glColor3usv, real__glColor3usv
TRAMPOLINE glColor4b, real__glColor4b
TRAMPOLINE glColor4bv, real__glColor4bv
TRAMPOLINE glColor4d, real__glColor4d
TRAMPOLINE glColor4dv, real__glColor4dv
TRAMPOLINE glColor4f, real__glColor4f
TRAMPOLINE glColor4fv, real__glColor4fv
TRAMPOLINE glColor4i, real__glColor4i
TRAMPOLINE glColor4iv, real__glColor4iv
TRAMPOLINE glColor4s, real__glColor4s
TRAMPOLINE glColor4sv, real__glColor4sv
TRAMPOLINE glColor4ub, real__glColor4ub
TRAMPOLINE glColor4ubv, real__glColor4ubv
TRAMPOLINE glColor4ui, real__glColor4ui
TRAMPOLINE glColor4uiv, real__glColor4uiv
TRAMPOLINE glColor4us, real__glColor4us
TRAMPOLINE glColor4usv, real__glColor4usv
TRAMPOLINE glColorMaterial, real__glColorMaterial
TRAMPOLINE glColorPointer, real__glColorPointer
TRAMPOLINE glCopyPixels, real__glCopyPixels
TRAMPOLINE glCopyTexImage1D, real__glCopyTexImage1D
TRAMPOLINE glCopyTexSubImage1D, real__glCopyTexSubImage1D
TRAMPOLINE glDebugEntry, real__glDebugEntry
TRAMPOLINE glDeleteLists, real__glDeleteLists
TRAMPOLINE glDepthRange, real__glDepthRange
TRAMPOLINE glDrawBuffer, real__glDrawBuffer
TRAMPOLINE glDrawPixels, real__glDrawPixels
TRAMPOLINE glEdgeFlag, real__glEdgeFlag
TRAMPOLINE glEdgeFlagPointer, real__glEdgeFlagPointer
TRAMPOLINE glEdgeFlagv, real__glEdgeFlagv
TRAMPOLINE glEnd, real__glEnd
TRAMPOLINE glEndList, real__glEndList
TRAMPOLINE glEvalCoord1d, real__glEvalCoord1d
TRAMPOLINE glEvalCoord1dv, real__glEvalCoord1dv
TRAMPOLINE glEvalCoord1f, real__glEvalCoord1f
TRAMPOLINE glEvalCoord1fv, real__glEvalCoord1fv
TRAMPOLINE glEvalCoord2d, real__glEvalCoord2d
TRAMPOLINE glEvalCoord2dv, real__glEvalCoord2dv
TRAMPOLINE glEvalCoord2f, real__glEvalCoord2f
TRAMPOLINE glEvalCoord2fv, real__glEvalCoord2fv
TRAMPOLINE glEvalMesh1, real__glEvalMesh1
TRAMPOLINE glEvalMesh2, real__glEvalMesh2
TRAMPOLINE glEvalPoint1, real__glEvalPoint1
TRAMPOLINE glEvalPoint2, real__glEvalPoint2
TRAMPOLINE glFeedbackBuffer, real__glFeedbackBuffer
TRAMPOLINE glFogf, real__glFogf
TRAMPOLINE glFogfv, real__glFogfv
TRAMPOLINE glFogi, real__glFogi
TRAMPOLINE glFogiv, real__glFogiv
TRAMPOLINE glFrustum, real__glFrustum
TRAMPOLINE glGenLists, real__glGenLists
TRAMPOLINE glGetClipPlane, real__glGetClipPlane
TRAMPOLINE glGetDoublev, real__glGetDoublev
TRAMPOLINE glGetLightfv, real__glGetLightfv
TRAMPOLINE glGetLightiv, real__glGetLightiv
TRAMPOLINE glGetMapdv, real__glGetMapdv
TRAMPOLINE glGetMapfv, real__glGetMapfv
TRAMPOLINE glGetMapiv, real__glGetMapiv
TRAMPOLINE glGetMaterialfv, real__glGetMaterialfv
TRAMPOLINE glGetMaterialiv, real__glGetMaterialiv
TRAMPOLINE glGetPixelMapfv, real__glGetPixelMapfv
TRAMPOLINE glGetPixelMapuiv, real__glGetPixelMapuiv
TRAMPOLINE glGetPixelMapusv, real__glGetPixelMapusv
TRAMPOLINE glGetPointerv, real__glGetPointerv
TRAMPOLINE glGetPolygonStipple, real__glGetPolygonStipple
TRAMPOLINE glGetTexEnvfv, real__glGetTexEnvfv
TRAMPOLINE glGetTexEnviv, real__glGetTexEnviv
TRAMPOLINE glGetTexGendv, real__glGetTexGendv
TRAMPOLINE glGetTexGenfv, real__glGetTexGenfv
TRAMPOLINE glGetTexGeniv, real__glGetTexGeniv
TRAMPOLINE glGetTexImage, real__glGetTexImage
TRAMPOLINE glGetTexLevelParameterfv, real__glGetTexLevelParameterfv
TRAMPOLINE glGetTexLevelParameteriv, real__glGetTexLevelParameteriv
TRAMPOLINE glIndexd, real__glIndexd
TRAMPOLINE glIndexdv, real__glIndexdv
TRAMPOLINE glIndexf, real__glIndexf
TRAMPOLINE glIndexfv, real__glIndexfv
TRAMPOLINE glIndexi, real__glIndexi
TRAMPOLINE glIndexiv, real__glIndexiv
TRAMPOLINE glIndexMask, real__glIndexMask
TRAMPOLINE glIndexPointer, real__glIndexPointer
TRAMPOLINE glIndexs, real__glIndexs
TRAMPOLINE glIndexsv, real__glIndexsv
TRAMPOLINE glIndexub, real__glIndexub
TRAMPOLINE glIndexubv, real__glIndexubv
TRAMPOLINE glInitNames, real__glInitNames
TRAMPOLINE glInterleavedArrays, real__glInterleavedArrays
TRAMPOLINE glIsList, real__glIsList
TRAMPOLINE glLightf, real__glLightf
TRAMPOLINE glLightfv, real__glLightfv
TRAMPOLINE glLighti, real__glLighti
TRAMPOLINE glLightiv, real__glLightiv
TRAMPOLINE glLightModelf, real__glLightModelf
TRAMPOLINE glLightModelfv, real__glLightModelfv
TRAMPOLINE glLightModeli, real__glLightModeli
TRAMPOLINE glLightModeliv, real__glLightModeliv
TRAMPOLINE glLineStipple, real__glLineStipple
TRAMPOLINE glListBase, real__glListBase
TRAMPOLINE glLoadIdentity, real__glLoadIdentity
TRAMPOLINE glLoadMatrixd, real__glLoadMatrixd
TRAMPOLINE glLoadMatrixf, real__glLoadMatrixf
TRAMPOLINE glLoadName, real__glLoadName
TRAMPOLINE glLogicOp, real__glLogicOp
TRAMPOLINE glMap1d, real__glMap1d
TRAMPOLINE glMap1f, real__glMap1f
TRAMPOLINE glMap2d, real__glMap2d
TRAMPOLINE glMap2f, real__glMap2f
TRAMPOLINE glMapGrid1d, real__glMapGrid1d
TRAMPOLINE glMapGrid1f, real__glMapGrid1f
TRAMPOLINE glMapGrid2d, real__glMapGrid2d
TRAMPOLINE glMapGrid2f, real__glMapGrid2f
TRAMPOLINE glMaterialf, real__glMaterialf
TRAMPOLINE glMaterialfv, real__glMaterialfv
TRAMPOLINE glMateriali, real__glMateriali
TRAMPOLINE glMaterialiv, real__glMaterialiv
TRAMPOLINE glMatrixMode, real__glMatrixMode
TRAMPOLINE GlmfBeginGlsBlock, real__GlmfBeginGlsBlock
TRAMPOLINE GlmfCloseMetaFile, real__GlmfCloseMetaFile
TRAMPOLINE GlmfEndGlsBlock, real__GlmfEndGlsBlock
TRAMPOLINE GlmfEndPlayback, real__GlmfEndPlayback
TRAMPOLINE GlmfInitPlayback, real__GlmfInitPlayback
TRAMPOLINE GlmfPlayGlsRecord, real__GlmfPlayGlsRecord
TRAMPOLINE glMultMatrixd, real__glMultMatrixd
TRAMPOLINE glMultMatrixf, real__glMultMatrixf
TRAMPOLINE glNewList, real__glNewList
TRAMPOLINE glNormal3b, real__glNormal3b
TRAMPOLINE glNormal3bv, real__glNormal3bv
TRAMPOLINE glNormal3d, real__glNormal3d
TRAMPOLINE glNormal3dv, real__glNormal3dv
TRAMPOLINE glNormal3f, real__glNormal3f
TRAMPOLINE glNormal3fv, real__glNormal3fv
TRAMPOLINE glNormal3i, real__glNormal3i
TRAMPOLINE glNormal3iv, real__glNormal3iv
TRAMPOLINE glNormal3s, real__glNormal3s
TRAMPOLINE glNormal3sv, real__glNormal3sv
TRAMPOLINE glNormalPointer, real__glNormalPointer
TRAMPOLINE glOrtho, real__glOrtho
TRAMPOLINE glPassThrough, real__glPassThrough
TRAMPOLINE glPixelMapfv, real__glPixelMapfv
TRAMPOLINE glPixelMapuiv, real__glPixelMapuiv
TRAMPOLINE glPixelMapusv, real__glPixelMapusv
TRAMPOLINE glPixelStoref, real__glPixelStoref
TRAMPOLINE glPixelTransferf, real__glPixelTransferf
TRAMPOLINE glPixelTransferi, real__glPixelTransferi
TRAMPOLINE glPixelZoom, real__glPixelZoom
TRAMPOLINE glPointSize, real__glPointSize
TRAMPOLINE glPolygonMode, real__glPolygonMode
TRAMPOLINE glPolygonStipple, real__glPolygonStipple
TRAMPOLINE glPopAttrib, real__glPopAttrib
TRAMPOLINE glPopClientAttrib, real__glPopClientAttrib
TRAMPOLINE glPopMatrix, real__glPopMatrix
TRAMPOLINE glPopName, real__glPopName
TRAMPOLINE glPrioritizeTextures, real__glPrioritizeTextures
TRAMPOLINE glPushAttrib, real__glPushAttrib
TRAMPOLINE glPushClientAttrib, real__glPushClientAttrib
TRAMPOLINE glPushMatrix, real__glPushMatrix
TRAMPOLINE glPushName, real__glPushName
TRAMPOLINE glRasterPos2d, real__glRasterPos2d
TRAMPOLINE glRasterPos2dv, real__glRasterPos2dv
TRAMPOLINE glRasterPos2f, real__glRasterPos2f
TRAMPOLINE glRasterPos2fv, real__glRasterPos2fv
TRAMPOLINE glRasterPos2i, real__glRasterPos2i
TRAMPOLINE glRasterPos2iv, real__glRasterPos2iv
TRAMPOLINE glRasterPos2s, real__glRasterPos2s
TRAMPOLINE glRasterPos2sv, real__glRasterPos2sv
TRAMPOLINE glRasterPos3d, real__glRasterPos3d
TRAMPOLINE glRasterPos3dv, real__glRasterPos3dv
TRAMPOLINE glRasterPos3f, real__glRasterPos3f
TRAMPOLINE glRasterPos3fv, real__glRasterPos3fv
TRAMPOLINE glRasterPos3i, real__glRasterPos3i
TRAMPOLINE glRasterPos3iv, real__glRasterPos3iv
TRAMPOLINE glRasterPos3s, real__glRasterPos3s
TRAMPOLINE glRasterPos3sv, real__glRasterPos3sv
TRAMPOLINE glRasterPos4d, real__glRasterPos4d
TRAMPOLINE glRasterPos4dv, real__glRasterPos4dv
TRAMPOLINE glRasterPos4f, real__glRasterPos4f
TRAMPOLINE glRasterPos4fv, real__glRasterPos4fv
TRAMPOLINE glRasterPos4i, real__glRasterPos4i
TRAMPOLINE glRasterPos4iv, real__glRasterPos4iv
TRAMPOLINE glRasterPos4s, real__glRasterPos4s
TRAMPOLINE glRasterPos4sv, real__glRasterPos4sv
TRAMPOLINE glReadBuffer, real__glReadBuffer
TRAMPOLINE glRectd, real__glRectd
TRAMPOLINE glRectdv, real__glRectdv
TRAMPOLINE glRectf, real__glRectf
TRAMPOLINE glRectfv, real__glRectfv
TRAMPOLINE glRecti, real__glRecti
TRAMPOLINE glRectiv, real__glRectiv
TRAMPOLINE glRects, real__glRects
TRAMPOLINE glRectsv, real__glRectsv
TRAMPOLINE glRenderMode, real__glRenderMode
TRAMPOLINE glRotated, real__glRotated
TRAMPOLINE glRotatef, real__glRotatef
TRAMPOLINE glScaled, real__glScaled
TRAMPOLINE glScalef, real__glScalef
TRAMPOLINE glSelectBuffer, real__glSelectBuffer
TRAMPOLINE glShadeModel, real__glShadeModel
TRAMPOLINE glStencilFunc, real__glStencilFunc
TRAMPOLINE glStencilOp, real__glStencilOp
TRAMPOLINE glTexCoord1d, real__glTexCoord1d
TRAMPOLINE glTexCoord1dv, real__glTexCoord1dv
TRAMPOLINE glTexCoord1f, real__glTexCoord1f
TRAMPOLINE glTexCoord1fv, real__glTexCoord1fv
TRAMPOLINE glTexCoord1i, real__glTexCoord1i
TRAMPOLINE glTexCoord1iv, real__glTexCoord1iv
TRAMPOLINE glTexCoord1s, real__glTexCoord1s
TRAMPOLINE glTexCoord1sv, real__glTexCoord1sv
TRAMPOLINE glTexCoord2d, real__glTexCoord2d
TRAMPOLINE glTexCoord2dv, real__glTexCoord2dv
TRAMPOLINE glTexCoord2f, real__glTexCoord2f
TRAMPOLINE glTexCoord2fv, real__glTexCoord2fv
TRAMPOLINE glTexCoord2i, real__glTexCoord2i
TRAMPOLINE glTexCoord2iv, real__glTexCoord2iv
TRAMPOLINE glTexCoord2s, real__glTexCoord2s
TRAMPOLINE glTexCoord2sv, real__glTexCoord2sv
TRAMPOLINE glTexCoord3d, real__glTexCoord3d
TRAMPOLINE glTexCoord3dv, real__glTexCoord3dv
TRAMPOLINE glTexCoord3f, real__glTexCoord3f
TRAMPOLINE glTexCoord3fv, real__glTexCoord3fv
TRAMPOLINE glTexCoord3i, real__glTexCoord3i
TRAMPOLINE glTexCoord3iv, real__glTexCoord3iv
TRAMPOLINE glTexCoord3s, real__glTexCoord3s
TRAMPOLINE glTexCoord3sv, real__glTexCoord3sv
TRAMPOLINE glTexCoord4d, real__glTexCoord4d
TRAMPOLINE glTexCoord4dv, real__glTexCoord4dv
TRAMPOLINE glTexCoord4f, real__glTexCoord4f
TRAMPOLINE glTexCoord4fv, real__glTexCoord4fv
TRAMPOLINE glTexCoord4i, real__glTexCoord4i
TRAMPOLINE glTexCoord4iv, real__glTexCoord4iv
TRAMPOLINE glTexCoord4s, real__glTexCoord4s
TRAMPOLINE glTexCoord4sv, real__glTexCoord4sv
TRAMPOLINE glTexCoordPointer, real__glTexCoordPointer
TRAMPOLINE glTexEnvf, real__glTexEnvf
TRAMPOLINE glTexEnvfv, real__glTexEnvfv
TRAMPOLINE glTexEnvi, real__glTexEnvi
TRAMPOLINE glTexEnviv, real__glTexEnviv
TRAMPOLINE glTexGend, real__glTexGend
TRAMPOLINE glTexGendv, real__glTexGendv
TRAMPOLINE glTexGenf, real__glTexGenf
TRAMPOLINE glTexGenfv, real__glTexGenfv
TRAMPOLINE glTexGeni, real__glTexGeni
TRAMPOLINE glTexGeniv, real__glTexGeniv
TRAMPOLINE glTexImage1D, real__glTexImage1D
TRAMPOLINE glTexParameterfv, real__glTexParameterfv
TRAMPOLINE glTexParameteriv, real__glTexParameteriv
TRAMPOLINE glTexSubImage1D, real__glTexSubImage1D
TRAMPOLINE glTranslated, real__glTranslated
TRAMPOLINE glTranslatef, real__glTranslatef
TRAMPOLINE glVertex2d, real__glVertex2d
TRAMPOLINE glVertex2dv, real__glVertex2dv
TRAMPOLINE glVertex2f, real__glVertex2f
TRAMPOLINE glVertex2fv, real__glVertex2fv
TRAMPOLINE glVertex2i, real__glVertex2i
TRAMPOLINE glVertex2iv, real__glVertex2iv
TRAMPOLINE glVertex2s, real__glVertex2s
TRAMPOLINE glVertex2sv, real__glVertex2sv
TRAMPOLINE glVertex3d, real__glVertex3d
TRAMPOLINE glVertex3dv, real__glVertex3dv
TRAMPOLINE glVertex3f, real__glVertex3f
TRAMPOLINE glVertex3fv, real__glVertex3fv
TRAMPOLINE glVertex3i, real__glVertex3i
TRAMPOLINE glVertex3iv, real__glVertex3iv
TRAMPOLINE glVertex3s, real__glVertex3s
TRAMPOLINE glVertex3sv, real__glVertex3sv
TRAMPOLINE glVertex4d, real__glVertex4d
TRAMPOLINE glVertex4dv, real__glVertex4dv
TRAMPOLINE glVertex4f, real__glVertex4f
TRAMPOLINE glVertex4fv, real__glVertex4fv
TRAMPOLINE glVertex4i, real__glVertex4i
TRAMPOLINE glVertex4iv, real__glVertex4iv
TRAMPOLINE glVertex4s, real__glVertex4s
TRAMPOLINE glVertex4sv, real__glVertex4sv
TRAMPOLINE glVertexPointer, real__glVertexPointer
TRAMPOLINE wglChoosePixelFormat, real__wglChoosePixelFormat
TRAMPOLINE wglCopyContext, real__wglCopyContext
TRAMPOLINE wglCreateLayerContext, real__wglCreateLayerContext
TRAMPOLINE wglDeleteContext, real__wglDeleteContext
TRAMPOLINE wglDescribeLayerPlane, real__wglDescribeLayerPlane
TRAMPOLINE wglDescribePixelFormat, real__wglDescribePixelFormat
TRAMPOLINE wglGetCurrentContext, real__wglGetCurrentContext
TRAMPOLINE wglGetCurrentDC, real__wglGetCurrentDC
TRAMPOLINE wglGetDefaultProcAddress, real__wglGetDefaultProcAddress
TRAMPOLINE wglGetLayerPaletteEntries, real__wglGetLayerPaletteEntries
TRAMPOLINE wglGetPixelFormat, real__wglGetPixelFormat
TRAMPOLINE wglRealizeLayerPalette, real__wglRealizeLayerPalette
TRAMPOLINE wglSetLayerPaletteEntries, real__wglSetLayerPaletteEntries
TRAMPOLINE wglSetPixelFormat, real__wglSetPixelFormat
TRAMPOLINE wglShareLists, real__wglShareLists
TRAMPOLINE wglSwapLayerBuffers, real__wglSwapLayerBuffers
TRAMPOLINE wglSwapMultipleBuffers, real__wglSwapMultipleBuffers
TRAMPOLINE wglUseFontBitmapsA, real__wglUseFontBitmapsA
TRAMPOLINE wglUseFontBitmapsW, real__wglUseFontBitmapsW
TRAMPOLINE wglUseFontOutlinesA, real__wglUseFontOutlinesA
TRAMPOLINE wglUseFontOutlinesW, real__wglUseFontOutlinesW
end
