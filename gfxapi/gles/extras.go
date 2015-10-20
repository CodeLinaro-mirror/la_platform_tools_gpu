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
	"android.googlesource.com/platform/tools/gpu/binary"
)

// ProgramInfo is an atom extra used to describe the uniform and attribute
// layout of a linked shader program.
type ProgramInfo struct {
	binary.Generate

	// Uniforms describes the layout of all uniforms of a shader program.
	Uniforms map[UniformLocation]UniformInfo

	// Attributes describes the layout of all attributes of a shader program.
	Attributes map[AttributeLocation]AttributeInfo
}

// UniformInfo describes the layout of a single shader uniform.
type UniformInfo struct {
	binary.Generate

	Name        string // Name of the uniform.
	VectorCount GLint  // Number of vectors that make up the uniform.
	Type        GLenum // Type of the uniform.
}

// AttributeInfo describes the layout of a single shader uniform.
type AttributeInfo struct {
	binary.Generate

	Name        string // Name of the attribute.
	VectorCount GLint  // Number of vectors that make up the attribute.
	Type        GLenum // Type of the attribute.
}
