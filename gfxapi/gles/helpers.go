// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// NewProgram returns the atoms to create a shader program with compiled vertex
// and fragment shaders. The returned program is not linked.
func NewProgram(a device.Architecture, d database.Database, l log.Logger,
	vertexShaderID, fragmentShaderID ShaderId, programID ProgramId,
	vertexShaderSource, fragmentShaderSource string) []atom.Atom {
	tmp := memory.Tmp.Base
	return []atom.Atom{
		NewGlCreateShader(ShaderType_GL_VERTEX_SHADER, vertexShaderID),
		NewGlShaderSource(vertexShaderID, 1, tmp, 0).
			AddRead(atom.Data(a, d, l, tmp, memory.Pointer(tmp+8))).
			AddRead(atom.Data(a, d, l, tmp+8, vertexShaderSource)),
		NewGlCompileShader(vertexShaderID),
		NewGlCreateShader(ShaderType_GL_FRAGMENT_SHADER, fragmentShaderID),
		NewGlShaderSource(fragmentShaderID, 1, tmp, 0).
			AddRead(atom.Data(a, d, l, tmp, memory.Pointer(tmp+8))).
			AddRead(atom.Data(a, d, l, tmp+8, fragmentShaderSource)),
		NewGlCompileShader(fragmentShaderID),
		NewGlCreateProgram(programID),
		NewGlAttachShader(programID, vertexShaderID),
		NewGlAttachShader(programID, fragmentShaderID),
	}
}
