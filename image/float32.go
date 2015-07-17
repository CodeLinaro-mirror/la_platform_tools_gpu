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

package image

import "android.googlesource.com/platform/tools/gpu/binary"

type fmtFloat32 struct{ binary.Generate }

func (*fmtFloat32) String() string                 { return "Float32" }
func (*fmtFloat32) Check(d []byte, w, h int) error { return checkSize(d, w, h, 32) }

// RGBA returns a format containing an 8-bit red, green, blue and alpha channel
// per pixel.
func Float32() Format { return &fmtFloat32{} }
