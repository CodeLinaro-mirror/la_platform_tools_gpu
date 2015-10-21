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

package ndk

import (
	"fmt"
	"os"

	"android.googlesource.com/platform/tools/gpu/maker/build"
	"android.googlesource.com/platform/tools/gpu/maker/config"
)

// Settings for the particular SDK/NDK build/platform we're using.
const (
	javaEnvVar = "JAVA_HOME"
)

// Paths contains the list of directories and files required by the NDK
// toolchain. It is acquired by calling ResolvePaths.
var Paths = struct {
	NDK        build.File // The Android NDK root directory.
	SDK        build.File // The Android SDK root directory.
	GCC        build.File // The Android NDK gcc directory.
	Jarsigner  build.File // The jarsigner executable path.
	AAPT       build.File // The Android SDK aapt executable path.
	Zipalign   build.File // The Android SDK zipalign executable path.
	AndroidJar build.File // The Android SDK android.jar file.
}{}

func resolveEnvVarDir(env string) (build.File, error) {
	f := build.File(os.Getenv(env))
	switch {
	case f == "":
		return "", fmt.Errorf("Environment variable '%s' not set", env)

	case !f.Exists():
		return "", fmt.Errorf("Environment variable '%s' directory '%s' does not exist", env, f)
	}
	return f, nil
}

func init() {
	// Uses the system environment variable to find the Java directory and the
	// prebuilts to find the NDK/SDK toolchain.

	// Resolve the prebuilt SDK & NDK directories.
	Paths.SDK = build.RepoRoot.Path.Join("prebuilts", "fullsdk")
	Paths.NDK = build.RepoRoot.Path.Join("prebuilts", "ndk", "current")
	ndkBuildTools := ""
	switch config.HostOS {
	case config.Linux:
		Paths.GCC = build.RepoRoot.Path.Join("prebuilts", "gcc", "linux-x86")
		ndkBuildTools = fmt.Sprintf("%v-linux", ndkAndroidVersion)

	case config.OSX:
		Paths.GCC = build.RepoRoot.Path.Join("prebuilts", "gcc", "darwin-x86")
		ndkBuildTools = fmt.Sprintf("%v-darwin", ndkAndroidVersion)
	}

	if !Paths.NDK.Exists() || !Paths.SDK.Exists() || !Paths.GCC.Exists() {
		// NDK not found, disable android ta
		for _, abi := range config.ABIs {
			if abi.OS == config.Android {
				abi.Disabled = true
			}
		}
		return
	}

	// Resolve Java root directory from environment variable.
	java, err := resolveEnvVarDir(javaEnvVar)
	if err != nil {
		panic(err)
	}

	Paths.Jarsigner = java.Join("bin", "jarsigner"+config.HostOS.ExecutableExtension)
	if !Paths.Jarsigner.Exists() {
		panic(fmt.Errorf("Java SDK does not contain jarsigner"))
	}

	buildtools := Paths.SDK.Join("build-tools", ndkBuildTools)
	if !buildtools.Exists() {
		panic(fmt.Errorf("Android SDK does not contain required build-tools: %s", ndkBuildTools))
	}

	Paths.AAPT = buildtools.Join("aapt" + config.HostOS.ExecutableExtension)
	if !Paths.AAPT.Exists() {
		panic(fmt.Errorf("Android SDK does not contain aapt tool"))
	}

	Paths.Zipalign = buildtools.Join("zipalign" + config.HostOS.ExecutableExtension)
	if !Paths.Zipalign.Exists() {
		panic(fmt.Errorf("Android SDK does not contain zipalign tool"))
	}

	sdkPlatform := fmt.Sprintf("android-%d", ndkAndroidVersion)
	Paths.AndroidJar = Paths.SDK.Join("platforms", sdkPlatform, "android.jar")
	if !Paths.AndroidJar.Exists() {
		panic(fmt.Errorf("Android SDK does not contain required platform: %s", sdkPlatform))
	}
}
