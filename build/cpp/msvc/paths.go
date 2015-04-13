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

// +build windows

package msvc

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"android.googlesource.com/platform/tools/gpu/build"
)

const msvcVersion = "12"
const sdkVersion = "8.1"

var paths = Paths{}

// ResolvePaths uses the system environment variables and Windows registry to
// the MSVC toolchain directories.
func ResolvePaths() (Paths, error) {
	if paths.resolved {
		return paths, nil
	}

	env := fmt.Sprintf("VS%s0COMNTOOLS", msvcVersion)
	vstools := build.File(os.Getenv(env))
	if vstools == "" {
		return Paths{}, fmt.Errorf("Environment variable '%s' not set", env)
	}
	if !vstools.Exists() {
		return Paths{}, fmt.Errorf("Environment variable '%s' directory '%s' does not exist", env, vstools)
	}

	sdk, err := getWinSdk()
	if err != nil {
		return Paths{}, fmt.Errorf("Failed to find Windows SDK directory from registry: %v", err)
	}

	vc := vstools.Join("..", "..", "VC")

	paths.Cl = vc.Join("bin", "amd64", "cl.exe")
	paths.Lib = vc.Join("bin", "amd64", "lib.exe")
	paths.Link = vc.Join("bin", "amd64", "link.exe")
	paths.IncludeSearchPaths = build.Files(
		vc.Join("include"),
		sdk.Join("Include", "shared"),
		sdk.Join("Include", "um"),
	)
	paths.LibrarySearchPaths = build.Files(
		vc.Join("Lib", "amd64"),
		sdk.Join("Lib", "winv6.3", "um", "x64"),
	)

	if !paths.Cl.Exists() {
		return Paths{}, fmt.Errorf("MSVC tool '%s' was not found", paths.Cl)
	}

	if !paths.Link.Exists() {
		return Paths{}, fmt.Errorf("MSVC tool '%s' was not found", paths.Link)
	}

	paths.resolved = true
	return paths, nil
}

func getWinSdk() (build.File, error) {
	var handle syscall.Handle
	sdkkey := "SOFTWARE\\Wow6432Node\\Microsoft\\Microsoft SDKs\\Windows\\v" + sdkVersion
	if err := syscall.RegOpenKeyEx(
		syscall.HKEY_LOCAL_MACHINE,
		syscall.StringToUTF16Ptr(sdkkey),
		0, syscall.KEY_READ, &handle); err != nil {
		return "", err
	}
	defer syscall.RegCloseKey(handle)

	var str [1024]uint16
	p := (*byte)(unsafe.Pointer(&str[0]))
	n := uint32(len(str) * 2) // in bytes
	if err := syscall.RegQueryValueEx(
		handle,
		syscall.StringToUTF16Ptr("InstallationFolder"),
		nil, nil, p, &n); err != nil {
		return "", err
	}

	path := syscall.UTF16ToString(str[:])
	return build.File(path), nil
}
