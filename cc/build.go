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

package main

import (
	"flag"
	"strings"

	"android.googlesource.com/platform/tools/gpu/build"
	"android.googlesource.com/platform/tools/gpu/build/cpp"
	"android.googlesource.com/platform/tools/gpu/build/cpp/gcc"
	"android.googlesource.com/platform/tools/gpu/build/cpp/ndk"
	"android.googlesource.com/platform/tools/gpu/log"
)

var (
	target    = flag.String("target", build.HostOS, "The target to build")
	runtests  = flag.Bool("runtests", false, "Run the tests after building")
	keystore  = flag.String("keystore", GPURoot.Join("build", "keystore", "debug.keystore").Absolute(), "The keystore used to sign APKs")
	storepass = flag.String("storepass", "android", "The password to the keystore")
	keypass   = flag.String("keypass", "android", "The password to the keystore's key")
	keyalias  = flag.String("alias", "androiddebugkey", "The alias of the key used to sign APKs")
	logfile   = flag.String("logfile", "", "Writes logging to a file instead of stdout")
	verbose   = flag.Bool("v", false, "Enable verbose logging")
)

var (
	GPURoot     = build.Root.Join("tools", "gpu", "src", "android.googlesource.com", "platform", "tools", "gpu")
	CCRoot      = GPURoot.Join("cc")
	ReplaydRoot = CCRoot.Join("replayd")
	GmockRoot   = build.Root.Join("external", "gmock")
	GtestRoot   = build.Root.Join("external", "gtest")
)

func main() {
	flag.Parse()

	var logger log.Logger
	if len(*logfile) > 0 {
		var err error
		logger, err = log.File(*logfile)
		if err != nil {
			panic(err)
		}
	} else {
		logger = log.Std()
	}
	defer logger.Close()

	env := build.Environment{
		Output:        build.Root.Join("tools", "gpu", "bin"),
		Intermediates: build.Root.Join("tools", "gpu", "pkg"),
		Keystore:      build.File(*keystore),
		Storepass:     *storepass,
		Keypass:       *keypass,
		Keyalias:      *keyalias,
		Logger:        logger,
		Verbose:       *verbose,
	}
	targets[*target].Build(env)
}

type Target struct {
	SourceFiles []string
	Gtest       cpp.Config
	Gmock       cpp.Config
	Gapir       cpp.Config
	GapirTests  cpp.Config
	Replayd     cpp.Config
}

func (t Target) Extend(n Target) Target {
	return Target{
		SourceFiles: append(append([]string{}, t.SourceFiles...), n.SourceFiles...),
		Gtest:       t.Gtest.Extend(n.Gtest),
		Gmock:       t.Gmock.Extend(n.Gmock),
		Gapir:       t.Gapir.Extend(n.Gapir),
		GapirTests:  t.GapirTests.Extend(n.GapirTests),
		Replayd:     t.Replayd.Extend(n.Replayd),
	}
}

func (t Target) Build(env build.Environment) {
	rootLogger := env.Logger
	begin := func(cfg cpp.Config) log.Logger {
		if env.Verbose {
			rootLogger.Info("Building %s", cfg.Name)
		}
		return rootLogger.Enter(cfg.Name)
	}

	// Build gtest into a library
	env.Logger = begin(t.Gtest)
	gtestSource := GtestRoot.Join("src").Glob("gtest-all.cc", "gtest_main.cc")
	gtestLib, err := cpp.StaticLibrary(gtestSource, t.Gtest, env)
	if err != nil {
		env.Logger.Error("%v", err)
		return
	}

	// Build gmock into a library
	env.Logger = begin(t.Gmock)
	gmockSource := GmockRoot.Join("src").Glob("gmock-all.cc")
	gmockLib, err := cpp.StaticLibrary(gmockSource, t.Gmock, env)
	if err != nil {
		env.Logger.Error("%v", err)
		return
	}

	// Gather the source files for gapir
	gapirSource := ReplaydRoot.Join("src").Glob(t.SourceFiles...).Exclude("Main.cpp")

	// Build the gapir static library.
	env.Logger = begin(t.Gapir)
	gapirLib, err := cpp.StaticLibrary(gapirSource, t.Gapir, env)
	if err != nil {
		env.Logger.Error("%v", err)
		return
	}

	// Build gapir tests.
	env.Logger = begin(t.GapirTests)
	gapirTestSource := ReplaydRoot.Join("test").Glob(t.SourceFiles...).Append(gapirLib, gtestLib, gmockLib)
	gapirTest, err := cpp.Executable(gapirTestSource, t.GapirTests, env)
	if err != nil {
		env.Logger.Error("%v", err)
		return
	}

	// Build replayd from the gapir static library and Main.cpp.
	env.Logger = begin(t.Replayd)
	replaydSource := build.Files(gapirLib, ReplaydRoot.Join("src", "Main.cpp"))
	if _, err := cpp.Executable(replaydSource, t.Replayd, env); err != nil {
		env.Logger.Error("%v", err)
		return
	}

	if *runtests {
		env.Logger.Info("Running gapir-tests:")
		if err := gapirTest.Exec(env); err != nil {
			env.Logger.Error("%v", err)
			return
		}
	}
}

func base(toolchain *cpp.Toolchain, os, architecture string) Target {
	base := cpp.Config{
		Toolchain:    toolchain,
		OS:           os,
		Architecture: architecture,
		Flavor:       "release",
		Defines: map[string]string{
			"TARGET_OS_" + strings.ToUpper(os): "1",
		},
		CompilerArgs: []string{"-std=c++11"},
	}
	if toolchain == gcc.GCC {
		base.LibrarySearchPaths = build.FileSet{"/usr/local/lib"}
	}
	return Target{
		SourceFiles: []string{"*.cpp", "*.cc"},
		Gtest: base.Extend(cpp.Config{
			Name: "gtest",
			IncludeSearchPaths: build.FileSet{
				GtestRoot,
				GtestRoot.Join("include"),
			},
		}),
		Gmock: base.Extend(cpp.Config{
			Name: "gmock",
			IncludeSearchPaths: build.FileSet{
				GmockRoot,
				GmockRoot.Join("include"),
				GtestRoot,
				GtestRoot.Join("include"),
			},
		}),
		Gapir: base.Extend(cpp.Config{
			Name: "gapir",
			IncludeSearchPaths: build.FileSet{
				CCRoot.Join("common"),
			},
		}),
		GapirTests: base.Extend(cpp.Config{
			Name:      "gapir-tests",
			Libraries: build.FileSet{"stdc++"},
			IncludeSearchPaths: build.FileSet{
				ReplaydRoot.Join("src"),
				CCRoot.Join("common"),
				GmockRoot.Join("include"),
				GtestRoot.Join("include"),
			},
		}),
		Replayd: base.Extend(cpp.Config{
			Name:      "replayd",
			Libraries: build.FileSet{"stdc++"},
			IncludeSearchPaths: build.FileSet{
				CCRoot.Join("common"),
			},
			Permissions: []string{
				"android.permission.INTERNET",
				"android.permission.READ_EXTERNAL_STORAGE",
				"android.permission.WRITE_EXTERNAL_STORAGE",
			},
		}),
	}
}

var targets = map[string]Target{
	"linux": base(gcc.GCC, "linux", "x64").Extend(Target{
		GapirTests: cpp.Config{
			Libraries: build.FileSet{"dl", "GL", "stdc++", "m", "pthread", "X11", "rt"},
		},
		Replayd: cpp.Config{
			Libraries: build.FileSet{"dl", "GL", "stdc++", "m", "pthread", "X11", "rt"},
		},
	}),

	"osx": base(gcc.GCC, "osx", "x64").Extend(Target{
		SourceFiles: []string{"*.mm"},
		GapirTests: cpp.Config{
			LinkerArgs: []string{
				"-framework", "Cocoa",
				"-framework", "OpenGL",
			},
			Libraries: build.FileSet{"pthread"},
		},
		Replayd: cpp.Config{
			LinkerArgs: []string{
				"-framework", "Cocoa",
				"-framework", "OpenGL",
			},
			Libraries: build.FileSet{"pthread"},
		},
	}),

	"windows": base(gcc.GCC, "windows", "x64").Extend(Target{
		GapirTests: cpp.Config{
			Libraries: build.FileSet{"ws2_32", "opengl32", "gdi32"},
		},
		Replayd: cpp.Config{
			Libraries: build.FileSet{"ws2_32", "opengl32", "gdi32"},
		},
	}),

	"android-arm": base(ndk.APK, "android", "arm").Extend(Target{
		GapirTests: cpp.Config{
			Toolchain: ndk.EXE,
			Libraries: build.FileSet{"EGL", "log", "android", "z", "m"},
		},
		Replayd: cpp.Config{
			Libraries:          build.FileSet{"EGL", "log", "android", "z", "m"},
			IncludeSearchPaths: build.FileSet{ndkRoot().Join("sources", "android", "native_app_glue")},
			AdditionalSources:  build.FileSet{ndkRoot().Join("sources", "android", "native_app_glue", "android_native_app_glue.c")},
		},
	}),
}

func ndkRoot() build.File {
	paths, _ := ndk.ResolvePaths()
	return paths.NDK
}
