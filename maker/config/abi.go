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

package config

import (
	"fmt"
	"strings"
)

var (
	ABIs = []*ABI{
		&ABI{
			Name:         "linux_x64",
			OS:           Linux,
			Architecture: X86_64,
			Version:      "",
			Toolchain:    "",
		},
		&ABI{
			Name:         "osx_x64",
			OS:           OSX,
			Architecture: X86_64,
			Version:      "",
			Toolchain:    "",
		},
		&ABI{
			Name:         "windows_x64",
			OS:           Windows,
			Architecture: X86_64,
			Version:      "",
			Toolchain:    "",
		},
		&ABI{
			Name:         "armeabi-v7a",
			OS:           Android,
			Architecture: Arm,
			Version:      "4.9",
			Toolchain:    "arm-linux-androideabi",
		},
		&ABI{
			Name:         "arm64-v8a",
			OS:           Android,
			Architecture: Arm64,
			Version:      "4.9",
			Toolchain:    "aarch64-linux-android",
		},
		&ABI{
			Name:         "mips",
			OS:           Android,
			Architecture: Mips,
			Version:      "4.9",
			Toolchain:    "mipsel-linux-android",
		},
		&ABI{
			Name:         "mips64",
			OS:           Android,
			Architecture: Mips64,
			Version:      "4.9",
			Toolchain:    "mips64el-linux-android",
		},
		&ABI{
			Name:         "x86",
			OS:           Android,
			Architecture: X86,
			Version:      "4.9",
			Toolchain:    "x86",
		},
		&ABI{
			Name:         "x86_64",
			OS:           Android,
			Architecture: X86_64,
			Version:      "4.9",
			Toolchain:    "x86_64",
		},
	}
)

type ABI struct {
	Name         string
	OS           *OS
	Architecture *Architecture
	Version      string
	Toolchain    string
	Disabled     bool
}

func (a *ABI) String() string {
	return a.Name
}

func (a *ABI) ToolchainRoot() string {
	i := strings.Index(a.Toolchain, "-")
	if i < 0 {
		return a.Toolchain
	}
	return a.Toolchain[:i]
}

func GetABI(os *OS, arch *Architecture) *ABI {
	for _, abi := range ABIs {
		if abi.OS == os && abi.Architecture == arch {
			return abi
		}
	}
	panic(fmt.Errorf("OS %s Architecture %s is not a valid abi pair", os, arch))
}
