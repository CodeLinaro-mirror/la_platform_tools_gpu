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

// Package ndk contains C++ toolchains for building Android applications and
// libraries.
package ndk

import (
	"fmt"
	"io/ioutil"
	"strings"

	"android.googlesource.com/platform/tools/gpu/build"
	"android.googlesource.com/platform/tools/gpu/build/cpp"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/maker"
)

type ndkTarget struct {
	name    string
	version string
	abi     string
}

var ndkArchToTarget = map[string]ndkTarget{
	"arm":    ndkTarget{"arm-linux-androideabi", "4.9", "armeabi-v7a"},
	"arm64":  ndkTarget{"aarch64-linux-android", "4.9", "arm64-v8a"},
	"mips":   ndkTarget{"mipsel-linux-android", "4.9", "mips"},
	"mips64": ndkTarget{"mips64el-linux-android", "4.9", "mips64"},
	"x86":    ndkTarget{"x86", "4.9", "x86"},
	"x86_64": ndkTarget{"x86_64", "4.9", "x86_64"},
}

var osToSystem = map[string]string{
	"windows": "windows-x86_64",
	"linux":   "linux-x86_64",
	"osx":     "darwin-x86_64",
}

// Toolchain for building an Android executable env, without being packaged into an
// APK (user-debug only).
var EXE = &cpp.Toolchain{
	Compiler:   compile,
	Archiver:   archive,
	DllLinker:  linkSo,
	ExeLinker:  linkExe,
	DepsFor:    depsFor,
	DepFileFor: depFileFor,
	LibName:    func(cfg cpp.Config) string { return "lib" + cfg.Name + ".a" },
	DllName:    func(cfg cpp.Config) string { return "lib" + cfg.Name + ".so" },
	ExeName:    func(cfg cpp.Config) string { return cfg.Name },
	ObjExt:     func(cfg cpp.Config) string { return ".o" },
}

// Toolchain for building an Android APK.
var APK = &cpp.Toolchain{
	Compiler:   compile,
	Archiver:   archive,
	DllLinker:  linkSo,
	ExeLinker:  linkApk,
	DepsFor:    depsFor,
	DepFileFor: depFileFor,
	LibName:    func(cfg cpp.Config) string { return "lib" + cfg.Name + ".a" },
	DllName:    func(cfg cpp.Config) string { return "lib" + cfg.Name + ".so" },
	ExeName:    func(cfg cpp.Config) string { return cfg.Name + ".apk" },
	ObjExt:     func(cfg cpp.Config) string { return ".o" },
}

type tools struct {
	as      build.File
	cc      build.File
	ar      build.File
	sysroot build.File
	incdirs build.FileSet
	libdirs build.FileSet
	libs    build.FileSet
}

func getTools(cfg cpp.Config) (*tools, error) {
	paths, err := ResolvePaths()
	if err != nil {
		return nil, err
	}

	target, ok := ndkArchToTarget[cfg.Architecture]
	if !ok {
		return nil, fmt.Errorf("NDK architecture '%s' not supported", cfg.Architecture)
	}

	system, ok := osToSystem[maker.HostOS]
	if !ok {
		return nil, fmt.Errorf("NDK host OS '%s' not supported", maker.HostOS)
	}

	stlBase := paths.NDK.Join("sources", "cxx-stl", "gnu-libstdc++", target.version)
	if !stlBase.Exists() {
		return nil, fmt.Errorf("NDK gnu-libstdc++ for version %s not found", target.version)
	}
	arch := fmt.Sprintf("arch-%s", cfg.Architecture)
	bin := paths.NDK.Join("toolchains", target.name+"-"+target.version, "prebuilt", system, "bin")
	if !bin.Exists() {
		return nil, fmt.Errorf("NDK toolchain for %s version %s not found", target.name, target.version)
	}

	ndkPlatform := fmt.Sprintf("android-%d", ndkAndroidVersion)

	return &tools{
		as: bin.Join(target.name + "-as" + maker.HostExecutableExtension),
		cc: bin.Join(target.name + "-gcc" + maker.HostExecutableExtension),
		ar: bin.Join(target.name + "-ar" + maker.HostExecutableExtension),
		incdirs: build.FileSet{
			stlBase.Join("include"),
			stlBase.Join("libs", target.abi, "include"),
		},
		libdirs: build.FileSet{
			stlBase.Join("libs", target.abi),
		},
		libs: build.FileSet{
			"libgnustl_static.a",
		},
		sysroot: paths.NDK.Join("platforms", ndkPlatform, arch),
	}, nil
}

func depFileFor(output build.File, cfg cpp.Config, env build.Environment) build.File {
	return cpp.IntermediatePath(output, ".dep", cfg, env)
}

func compile(input build.File, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = log.Enter(env.Logger, "NDK.Compile")

	tools, err := getTools(cfg)
	if err != nil {
		return err
	}

	switch input.Ext() {
	case ".asm":
		return tools.as.Exec(env, "-o", output.Absolute(), input.Absolute())

	default:
		depfile := depFileFor(output, cfg, env)

		a := []string{
			"--sysroot=" + tools.sysroot.Absolute(),
			"-c", // Compile to .o
			optFlags(cfg),
			"-MMD", "-MF", depfile.Absolute(), // Generate dependency file
			"-fPIC", // TODO: Not required for exes
		}
		a = append(a, cfg.CompilerArgs...)
		for _, isp := range cfg.IncludeSearchPaths.Append(tools.incdirs...) {
			a = append(a, fmt.Sprintf("-I%s", isp))
		}
		for n, v := range cfg.Defines {
			a = append(a, fmt.Sprintf("-D%s=%s", n, v))
		}
		a = append(a, input.Name(), "-o", output.Absolute())
		return tools.cc.ExecAt(env, build.File(input.Dir()), a...)
	}
}

func archive(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = log.Enter(env.Logger, "NDK.Archive")

	tools, err := getTools(cfg)
	if err != nil {
		return err
	}

	a := []string{"-rcs"}
	a = append(a, output.Absolute())
	a = append(a, cfg.ArchiverArgs...)
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	return tools.ar.Exec(env, a...)
}

func linkSo(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = log.Enter(env.Logger, "NDK.LinkSo")

	tools, err := getTools(cfg)
	if err != nil {
		return err
	}

	a := []string{
		optFlags(cfg),
		"--sysroot=" + tools.sysroot.Absolute(),
		"--shared",
	}
	a = append(a, cfg.LinkerArgs...)
	for _, lsp := range cfg.LibrarySearchPaths.Append(tools.libdirs...) {
		a = append(a, fmt.Sprintf("-L%s", lsp))
	}
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	for _, library := range cfg.Libraries.Append(tools.libs...) {
		name := strings.TrimPrefix(strings.TrimSuffix(string(library), ".a"), "lib")
		a = append(a, fmt.Sprintf("-l%s", name))
	}

	a = append(a, "-o", string(output))
	return tools.cc.Exec(env, a...)
}

func linkExe(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = log.Enter(env.Logger, "NDK.LinkExe")

	tools, err := getTools(cfg)
	if err != nil {
		return err
	}

	a := []string{
		"--sysroot=" + tools.sysroot.Absolute(),
		optFlags(cfg),
	}
	a = append(a, cfg.LinkerArgs...)
	for _, lsp := range cfg.LibrarySearchPaths.Append(tools.libdirs...) {
		a = append(a, fmt.Sprintf("-L%s", lsp))
	}
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	for _, library := range cfg.Libraries.Append(tools.libs...) {
		name := strings.TrimPrefix(strings.TrimSuffix(string(library), ".a"), "lib")
		a = append(a, fmt.Sprintf("-l%s", name))
	}

	a = append(a, "-o", string(output))
	return tools.cc.Exec(env, a...)
}

func linkApk(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = log.Enter(env.Logger, "NDK.LinkApk")

	paths, err := ResolvePaths()
	if err != nil {
		return err
	}

	// Create an intermediate directory to hold the files going into the apk.
	log.Debugf(env.Logger, "Creating APK files")
	root := build.File(cpp.IntermediatePath(output, "", cfg, env).Dir())

	// Emit the intermetiate files and .so:
	//    root/AndroidManifest.xml
	//    root/values/strings.xml
	//    root/lib/<abi>/libxxx.so
	manifest := root.Join("AndroidManifest.xml")
	manifest.MkdirAll()

	if err := ioutil.WriteFile(manifest.Absolute(), androidManifest(cfg), 0666); err != nil {
		return fmt.Errorf("Failed to build AndroidManifest.xml: %v", err)
	}
	res := root.Join("res")
	strings := res.Join("values", "strings.xml")
	strings.MkdirAll()
	if err := ioutil.WriteFile(strings.Absolute(), stringsXml(cfg.Name), 0666); err != nil {
		return fmt.Errorf("Failed to build strings.xml: %v", err)
	}
	so := root.Join("lib", ndkArchToTarget[cfg.Architecture].abi, "lib"+cfg.Name+".so")
	so.MkdirAll()
	if err := linkSo(inputs, so, cfg, env); err != nil {
		return err
	}

	unaligned := root.Join(cfg.Name + "-unaligned.apk")

	// Build the apk, unsigned. This pulls in AndroidManifest.xml and strings.xml
	log.Debugf(env.Logger, "Building unsigned APK")
	if err := paths.AAPT.Exec(env,
		"package", "-v", "-f",
		"-M", manifest.Absolute(),
		"-S", res.Absolute(),
		"-I", paths.AndroidJar.Absolute(),
		"-F", unaligned.Absolute()); err != nil {
		return err
	}

	log.Debugf(env.Logger, "Adding .so to APK")
	if err := paths.AAPT.ExecAt(env, root,
		"add", "-f", unaligned.Absolute(), so.RelativeTo(root)); err != nil {
		return err
	}

	log.Debugf(env.Logger, "Signing APK")
	if err := paths.Jarsigner.Exec(env,
		"-verbose",
		"-keystore", env.Keystore.Absolute(),
		"-storepass", env.Storepass,
		"-keypass", env.Keypass,
		unaligned.Absolute(), env.Keyalias); err != nil {
		return err
	}

	log.Debugf(env.Logger, "Zip-aligning APK")
	if err := paths.Zipalign.Exec(env,
		"-v", "-f", "4",
		unaligned.Absolute(), output.Absolute()); err != nil {
		return err
	}

	log.Debugf(env.Logger, "Done")
	return nil
}

func depsFor(output build.File, cfg cpp.Config, env build.Environment) (deps build.FileSet, valid bool) {
	env.Logger = log.Enter(env.Logger, "NDK.Deps")

	depfile := depFileFor(output, cfg, env)

	return cpp.ParseDepFile(depfile, env)
}

func androidManifest(cfg cpp.Config) []byte {
	permissions := make([]string, len(cfg.Permissions))
	for i := range cfg.Permissions {
		permissions[i] = fmt.Sprintf(`    <uses-permission android:name="%s" />`, cfg.Permissions[i])
	}
	return []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<manifest xmlns:android="http://schemas.android.com/apk/res/android" package="com.android.gapid.%s">
    <uses-sdk android:minSdkVersion="%d" />
    %s
    <application android:label="@string/app_name" android:hasCode="false" android:debuggable="true">
        <activity android:name="android.app.NativeActivity"
                  android:label="@string/app_name"
                  android:configChanges="orientation|keyboardHidden">
            <meta-data android:name="android.app.lib_name" android:value="%[1]s" />
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
    </application>
</manifest>`, cfg.Name, minSdkVersion, strings.Join(permissions, "\n")))
}

func stringsXml(name string) []byte {
	return []byte(fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<resources>
    <string name="app_name">%s</string>
</resources>`, name))
}

func optFlags(cfg cpp.Config) string {
	switch cfg.OptimizationLevel {
	case cpp.NoOptimization:
		return "-O0"
	default:
		return "-O2"
	}
}
