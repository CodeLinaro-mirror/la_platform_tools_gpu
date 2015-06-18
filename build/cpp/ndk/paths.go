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

	"android.googlesource.com/platform/tools/gpu/build"
	"android.googlesource.com/platform/tools/gpu/maker"
)

// Settings for the particular SDK/NDK build/platform we're using.
const (
	ndkEnvVar  = "ANDROID_NDK_ROOT"
	sdkEnvVar  = "ANDROID_HOME"
	javaEnvVar = "JAVA_HOME"
)

var paths = Paths{}

// Paths contains the list of directories and files required by the NDK
// toolchain. It is acquired by calling ResolvePaths.
type Paths struct {
	resolved   bool
	NDK        build.File // The Android NDK root directory.
	SDK        build.File // The Android SDK root directory.
	Jarsigner  build.File // The jarsigner executable path.
	AAPT       build.File // The Android SDK aapt executable path.
	Zipalign   build.File // The Android SDK zipalign executable path.
	AndroidJar build.File // The Android SDK android.jar file.
}

// ResolvePaths uses the system environment variables to find all the NDK, SDK
// and Java directories and executables required by the NDK toolchain.
func ResolvePaths() (Paths, error) {
	if paths.resolved {
		return paths, nil
	}

	resolveEnvVarDir := func(dir *build.File, env string) error {
		f := build.File(os.Getenv(env))
		switch {
		case f == "":
			return fmt.Errorf("Environment variable '%s' not set", env)

		case !f.Exists():
			return fmt.Errorf("Environment variable '%s' directory '%s' does not exist", env, f)
		}
		*dir = f
		return nil
	}

	var java build.File

	// Resolve root directories from environment variables
	if err := resolveEnvVarDir(&paths.NDK, ndkEnvVar); err != nil {
		return Paths{}, err
	}
	if err := resolveEnvVarDir(&paths.SDK, sdkEnvVar); err != nil {
		return Paths{}, err
	}
	if err := resolveEnvVarDir(&java, javaEnvVar); err != nil {
		return Paths{}, err
	}

	paths.Jarsigner = java.Join("bin", "jarsigner"+maker.HostExecutableExtension)
	if !paths.Jarsigner.Exists() {
		return Paths{}, fmt.Errorf("Java SDK does not contain jarsigner")
	}

	buildtools := paths.SDK.Join("build-tools", ndkBuildTools)
	if !buildtools.Exists() {
		return Paths{}, fmt.Errorf("Android SDK does not contain required build-tools: %s", ndkBuildTools)
	}

	paths.AAPT = buildtools.Join("aapt" + maker.HostExecutableExtension)
	if !paths.AAPT.Exists() {
		return Paths{}, fmt.Errorf("Android SDK does not contain aapt tool")
	}

	paths.Zipalign = buildtools.Join("zipalign" + maker.HostExecutableExtension)
	if !paths.Zipalign.Exists() {
		return Paths{}, fmt.Errorf("Android SDK does not contain zipalign tool")
	}

	sdkPlatform := fmt.Sprintf("android-%d", ndkAndroidVersion)
	paths.AndroidJar = paths.SDK.Join("platforms", sdkPlatform, "android.jar")
	if !paths.AndroidJar.Exists() {
		return Paths{}, fmt.Errorf("Android SDK does not contain required platform: %s", sdkPlatform)
	}

	paths.resolved = true
	return paths, nil
}
