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

var (
	Arm = &Architecture{
		Name: "arm",
	}
	Arm64 = &Architecture{
		Name: "arm64",
	}
	Mips = &Architecture{
		Name: "mips",
	}
	Mips64 = &Architecture{
		Name: "mips64",
	}
	X86 = &Architecture{
		Name: "x86",
	}
	X86_64 = &Architecture{
		Name: "X86_64",
	}
)

type Architecture struct {
	Name string
}

func (a *Architecture) String() string {
	return a.Name
}
