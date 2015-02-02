/*
 * Copyright 2014, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef ANDROID_CAZE_GL_INCLUDE_H
#define ANDROID_CAZE_GL_INCLUDE_H

#include "Target.h"

// Platform specific include files for the Gl functions

#if TARGET_OS == CAZE_OS_ANDROID
#   include <GLES/gl.h>
#   include <GLES/glext.h>
#   include <GLES3/gl3.h>
#   include <GLES3/gl3ext.h>
#   include <EGL/egl.h>
#elif TARGET_OS == CAZE_OS_OSX
#   include <OpenGL/gl.h>
#   include <OpenGL/glext.h>
#   include <GLFW/glfw3.h>
#elif TARGET_OS == CAZE_OS_LINUX || TARGET_OS == CAZE_OS_WINDOWS
#   include <GL/glew.h>
#   include <GLFW/glfw3.h>
#endif

#if ((!defined EGL_VERSION_1_0) && (!defined GLFW_VERSION_MAJOR))
#   error "No supported GL native interface (EGL, GLFW) found!"
#endif

#endif  // ANDROID_CAZE_GL_INCLUDE_H
