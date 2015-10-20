/*
 * Copyright 2015, The Android Open Source Project
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

template<typename T>
inline void addExtras(const T&, int reduce_overload_precedence = 0) {}

// getProgramInfo returns a scratch-allocated ProgramInfo, populated with the details of all the
// attributes and uniforms exposed by program.
inline gapic::coder::gles::ProgramInfo* getProgramInfo(int32_t program) {
    using namespace gapic::coder::gles;

    const int kNameSize = 1024;
    char name[kNameSize];

    auto pi = mScratch.create<ProgramInfo>();
    int32_t uniformCount = 0;
    mImports.glGetProgramiv(program, GLenum::GL_ACTIVE_UNIFORMS, &uniformCount);
    pi->mUniforms = mScratch.map<int32_t, UniformInfo>(uniformCount);
    for (int32_t i = 0; i < uniformCount; i++) {
        UniformInfo ui;
        int32_t nameLen = 0;
        mImports.glGetActiveUniform(
            program, i, kNameSize, &nameLen, &ui.mVectorCount, &ui.mType, name);
        ui.mName = mScratch.create<char>(nameLen + 1);
        strncpy(ui.mName, name, nameLen + 1);
        int32_t location = mImports.glGetUniformLocation(program, name);
        pi->mUniforms.set(location, ui);
    }

    int32_t attributeCount = 0;
    mImports.glGetProgramiv(program, GLenum::GL_ACTIVE_ATTRIBUTES, &attributeCount);
    pi->mAttributes = mScratch.map<uint32_t, AttributeInfo>(attributeCount);
    for (int32_t i = 0; i < attributeCount; i++) {
        AttributeInfo ai;
        int32_t nameLen = 0;
        mImports.glGetActiveAttrib(
            program, i, kNameSize, &nameLen, &ai.mVectorCount, &ai.mType, name);
        ai.mName = mScratch.create<char>(nameLen + 1);
        strncpy(ai.mName, name, nameLen + 1);
        int32_t location = mImports.glGetAttribLocation(program, name);
        pi->mAttributes.set(location, ai);
    }

    return pi;
}

inline void addExtras(gapic::coder::gles::GlLinkProgram& atom) {
    atom.mextras.append(getProgramInfo(atom.mProgram));
}

inline void addExtras(gapic::coder::gles::GlProgramBinary& atom) {
    atom.mextras.append(getProgramInfo(atom.mProgram));
}