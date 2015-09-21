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
	Linux = &OS{
		Name:                "linux",
		ExecutableExtension: "",
		DllExtension:        ".so",
		System:              "linux-x86_64",
	}
	OSX = &OS{
		Name:                "osx",
		ExecutableExtension: "",
		DllExtension:        ".dylib",
		System:              "darwin-x86_64",
	}
	Windows = &OS{
		Name:                "windows",
		ExecutableExtension: ".exe",
		DllExtension:        ".dll",
		System:              "windows-x86_64",
	}
	Android = &OS{
		Name:                "android",
		ExecutableExtension: "",
		DllExtension:        ".so",
		System:              "",
	}
)

var (
	osByName = map[string]*OS{
		Linux.Name:   Linux,
		OSX.Name:     OSX,
		Windows.Name: Windows,
		Android.Name: Android,
	}
)

type OS struct {
	Name                string
	ExecutableExtension string
	DllExtension        string
	System              string
}

func (os *OS) String() string {
	return os.Name
}

func FindOS(name string) *OS {
	return osByName[name]
}
