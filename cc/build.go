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

package cc

import (
	"flag"
	"strings"

	"android.googlesource.com/platform/tools/gpu/build"
	"android.googlesource.com/platform/tools/gpu/build/cpp"
	"android.googlesource.com/platform/tools/gpu/build/cpp/gcc"
	"android.googlesource.com/platform/tools/gpu/build/cpp/msvc"
	"android.googlesource.com/platform/tools/gpu/build/cpp/ndk"
	"android.googlesource.com/platform/tools/gpu/log"
)

var (
	keystore   = flag.String("keystore", GPURoot.Join("build", "keystore", "debug.keystore").Absolute(), "The keystore used to sign APKs")
	storepass  = flag.String("storepass", "android", "The password to the keystore")
	keypass    = flag.String("keypass", "android", "The password to the keystore's key")
	keyalias   = flag.String("alias", "androiddebugkey", "The alias of the key used to sign APKs")
	logfile    = flag.String("logfile", "", "Writes logging to a file instead of stdout")
	forcebuild = flag.Bool("f", false, "All build steps will be forced")
	verbose    = flag.Bool("verbose", false, "Enable verbose logging")
	debug      = flag.Bool("d", false, "Generate debug binaries")
)

var (
	ExternalRoot = build.RepoRoot.Path.Join("external")
	GPURoot      = build.RepoRoot.Path.Join("tools", "gpu", "src", "android.googlesource.com", "platform", "tools", "gpu")
	BinRoot      = build.RepoRoot.Path.Join("tools", "gpu", "bin")
	CCRoot       = GPURoot.Join("cc")
	GapicRoot    = CCRoot.Join("gapic")
	GapiiRoot    = CCRoot.Join("gapii")
	GapirRoot    = CCRoot.Join("gapir")
	ReplaydRoot  = CCRoot.Join("replayd")
	GmockRoot    = ExternalRoot.Join("gmock")
	GtestRoot    = ExternalRoot.Join("gtest")
)

func getEnvironment(logger log.Logger) build.Environment {
	return build.Environment{
		Intermediates: build.RepoRoot.Path.Join("tools", "gpu", "pkg"),
		Roots: build.RootList{
			build.Root{Name: "external", Path: ExternalRoot},
			build.Root{Name: "gpu", Path: GPURoot},
		},
		Keystore:   build.File(*keystore),
		Storepass:  *storepass,
		Keypass:    *keypass,
		Keyalias:   *keyalias,
		Logger:     logger,
		ForceBuild: *forcebuild,
		Verbose:    *verbose,
	}
}

// Generate "maker" graph for building the cpp code.
func Graph(targetNames []string) {
	// Note this logger is captured in the build graph. There is no good place
	// to call logger.Close()
	logger := getLogger()
	if run(targetNames, logger) == 1 {
		panic("There were errors making the build graph")
	}
}

func getLogger() log.Logger {
	if len(*logfile) > 0 {
		logger, err := log.File(*logfile)
		if err != nil {
			panic(err)
		}
		return logger
	} else {
		return log.Std()
	}
}

func run(targetNames []string, logger log.Logger) int {
	env := getEnvironment(logger)
	buildTargets := getBuildTargets()

	targets := make([]Target, len(targetNames))
	for i, targetName := range targetNames {
		if target, found := buildTargets[targetName]; found {
			targets[i] = target
		} else {
			available := []string{}
			for t := range buildTargets {
				available = append(available, t)
			}
			logger.Errorf("Unknown target '%s'. Available targets: %v", targetName, available)
			return 1
		}
	}

	for i := range targets {
		env.Logger = logger.Fork().Enter(targetNames[i])
		targets[i].Build(env)
	}

	return 0
}

type Target struct {
	SourceFiles []string
	Gtest       cpp.Config
	Gmock       cpp.Config
	Gapic       cpp.Config
	GapicTests  cpp.Config
	Gapii       cpp.Config
	Gapir       cpp.Config
	GapirTests  cpp.Config
	Spy         cpp.Config
	Replayd     cpp.Config
}

func (t Target) Extend(n Target) Target {
	return Target{
		SourceFiles: append(append([]string{}, t.SourceFiles...), n.SourceFiles...),
		Gtest:       t.Gtest.Extend(n.Gtest),
		Gmock:       t.Gmock.Extend(n.Gmock),
		Gapic:       t.Gapic.Extend(n.Gapic),
		GapicTests:  t.GapicTests.Extend(n.GapicTests),
		Gapii:       t.Gapii.Extend(n.Gapii),
		Gapir:       t.Gapir.Extend(n.Gapir),
		GapirTests:  t.GapirTests.Extend(n.GapirTests),
		Spy:         t.Spy.Extend(n.Spy),
		Replayd:     t.Replayd.Extend(n.Replayd),
	}
}

func (t Target) Build(env build.Environment) {
	// Build gtest into a library
	gtestSource := GtestRoot.Join("src").Glob("gtest-all.cc", "gtest_main.cc")
	gtestLib := cpp.MakeStaticLibrary(gtestSource, t.Gtest, env)

	// Build gmock into a library
	gmockSource := GmockRoot.Join("src").Glob("gmock-all.cc")
	gmockLib := cpp.MakeStaticLibrary(gmockSource, t.Gmock, env)

	// Gather the source files for gapic
	gapicSource := GapicRoot.Glob(t.SourceFiles...).
		Append(GapicRoot.Join(t.Gapic.OS).Glob(t.SourceFiles...)...).
		Exclude("*_test.cpp")

	// Build the gapic static library.
	gapicLib := cpp.MakeStaticLibrary(gapicSource, t.Gapic, env)

	// Gather the source files for gapii
	gapiiSource := GapiiRoot.Glob(t.SourceFiles...).
		Append(GapiiRoot.Join(t.Gapic.OS).Glob(t.SourceFiles...)...).
		Exclude("*_test.cpp")

	// Build the gapii static library.
	gapiiLib := cpp.MakeStaticLibrary(gapiiSource, t.Gapii, env)
	_ = gapiiLib

	// Gather the source files for gapir
	gapirSource := GapirRoot.Glob(t.SourceFiles...).
		Append(GapirRoot.Join(t.Gapir.OS).Glob(t.SourceFiles...)...).
		Exclude("*_test.cpp")

	// Build the gapir static library.
	gapirLib := cpp.MakeStaticLibrary(gapirSource, t.Gapir, env)

	// Build the spy from the gapii static library.
	// TODO: using the static-lib strips symbol visibility from the
	// dynamic-library.
	// Investigate linker flags to use build.Files(gapiiLib, gapicLib)
	spyInputs := gapiiSource.Append(gapicLib)
	cpp.MakeDynamicLibrary(spyInputs, t.Spy, env)

	// Build replayd from the gapir static library and Main.cpp.
	replaydSource := build.Files(ReplaydRoot.Join("main.cpp"))
	replaydInputs := replaydSource.Append(gapirLib, gapicLib)
	replayd := cpp.MakeExecutable(replaydInputs, t.Replayd, env)

	// Build tests.
	gapicTestSource := GapicRoot.Glob(t.SourceFiles...).
		Append(GapicRoot.Join(t.Gapic.OS).Glob(t.SourceFiles...)...).
		Filter("*_test.cpp")
	gapicTestInputs := gapicTestSource.Append(gtestLib, gmockLib, gapicLib)
	gapicTest := cpp.MakeExecutable(gapicTestInputs, t.GapicTests, env)

	gapirTestSource := GapirRoot.Glob(t.SourceFiles...).
		Append(GapirRoot.Join(t.Gapir.OS).Glob(t.SourceFiles...)...).
		Filter("*_test.cpp")
	gapirTestInputs := gapirTestSource.Append(gtestLib, gmockLib, gapirLib, gapicLib)
	gapirTest := cpp.MakeExecutable(gapirTestInputs, t.GapirTests, env)

	if t.Replayd.OS == build.HostOS {
		cpp.MakeRunTest(gapicTest, t.GapicTests)

		cpp.MakeRunTest(gapirTest, t.GapirTests)

		src, dst := replayd, BinRoot.Join(replayd.Name())
		cpp.MakeCopy(src, dst, t.Replayd, env)
	}
}

func base(toolchain *cpp.Toolchain, os, architecture string) Target {
	base := cpp.Config{
		Toolchain:    toolchain,
		OS:           os,
		Architecture: architecture,
		Defines: map[string]string{
			"TARGET_OS_" + strings.ToUpper(os): "1",
		},
	}

	if *debug {
		base.Flavor = "debug"
		base.OptimizationLevel = cpp.NoOptimization
		base.Defines["LOG_LEVEL"] = "3" // LOG_LEVEL_DEBUG
	} else {
		base.Flavor = "release"
		base.OptimizationLevel = cpp.FullOptimization
	}

	base.OutputDir = BinRoot.Join(os+"-"+architecture, base.Flavor)

	if toolchain == gcc.GCC {
		base.LibrarySearchPaths = build.Files("/usr/local/lib")
	}

	if toolchain != msvc.MSVC {
		base.Libraries = build.Files("stdc++")
		base.CompilerArgs = []string{"-std=c++11"}
	} else {
		base.Flavor += "-msvc"
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
		Gapic: base.Extend(cpp.Config{
			Name:               "gapic",
			IncludeSearchPaths: build.FileSet{CCRoot},
		}),
		GapicTests: base.Extend(cpp.Config{
			Name: "gapic-tests",
			IncludeSearchPaths: build.FileSet{
				CCRoot,
				GmockRoot.Join("include"),
				GtestRoot.Join("include"),
			},
		}),
		Gapii: base.Extend(cpp.Config{
			Name:               "gapii",
			IncludeSearchPaths: build.FileSet{CCRoot},
		}),
		Gapir: base.Extend(cpp.Config{
			Name:               "gapir",
			IncludeSearchPaths: build.FileSet{CCRoot},
		}),
		GapirTests: base.Extend(cpp.Config{
			Name: "gapir-tests",
			IncludeSearchPaths: build.FileSet{
				CCRoot,
				GmockRoot.Join("include"),
				GtestRoot.Join("include"),
			},
		}),
		Spy: base.Extend(cpp.Config{
			Name:               "spy",
			IncludeSearchPaths: build.FileSet{CCRoot},
		}),
		Replayd: base.Extend(cpp.Config{
			Name: "replayd",
			IncludeSearchPaths: build.FileSet{
				CCRoot,
			},
			Permissions: []string{
				"android.permission.INTERNET",
				"android.permission.READ_EXTERNAL_STORAGE",
				"android.permission.WRITE_EXTERNAL_STORAGE",
			},
		}),
	}
}

func getBuildTargets() map[string]Target {
	noext := ""

	linux := base(gcc.GCC, "linux", "x64").Extend(Target{
		GapicTests: cpp.Config{
			Libraries: build.FileSet{"pthread"},
		},
		GapirTests: cpp.Config{
			Libraries: build.FileSet{"dl", "GL", "m", "pthread", "X11", "rt"},
		},
		Replayd: cpp.Config{
			Libraries: build.FileSet{"dl", "GL", "m", "pthread", "X11", "rt"},
		},
		Spy: cpp.Config{
			Libraries: build.FileSet{"pthread"},
		},
	})

	osx := base(gcc.GCC, "osx", "x64").Extend(Target{
		SourceFiles: []string{"*.mm"},
		Spy: cpp.Config{
			Name:      "OpenGL",
			OutputExt: &noext,
		},
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
	})
	osx.Spy.OutputDir = osx.Spy.OutputDir.Join("OpenGL.framework", "Versions", "A")

	windows := base(gcc.GCC, "windows", "x64").Extend(Target{
		Spy: cpp.Config{
			Libraries: build.FileSet{"ws2_32", "opengl32", "gdi32", "user32"},
		},
		GapirTests: cpp.Config{
			Libraries: build.FileSet{"ws2_32", "opengl32", "gdi32", "user32"},
		},
		Replayd: cpp.Config{
			Libraries: build.FileSet{"ws2_32", "opengl32", "gdi32", "user32"},
		},
	})

	windows_msvc := base(msvc.MSVC, "windows", "x64").Extend(Target{
		Spy: cpp.Config{
			Name:              "opengl32",
			AdditionalSources: build.FileSet{GapiiRoot.Join("windows", "opengl32_x64.asm")},
			ModuleDefinition:  GapiiRoot.Join("windows", "opengl32_exports.def"),
			Libraries:         build.FileSet{"ws2_32", "opengl32", "gdi32", "user32"},
		},
		GapirTests: cpp.Config{
			Libraries: build.FileSet{"ws2_32", "opengl32", "gdi32", "user32"},
		},
		Replayd: cpp.Config{
			Libraries: build.FileSet{"ws2_32", "opengl32", "gdi32", "user32"},
		},
	})

	android_target := Target{
		GapicTests: cpp.Config{
			Toolchain: ndk.EXE,
			Libraries: build.FileSet{"log", "android", "z", "m"},
		},
		GapirTests: cpp.Config{
			Toolchain: ndk.EXE,
			Libraries: build.FileSet{"EGL", "log", "android", "z", "m"},
		},
		Replayd: cpp.Config{
			Libraries:          build.FileSet{"EGL", "log", "android", "z", "m"},
			IncludeSearchPaths: build.FileSet{ndkRoot().Join("sources", "android", "native_app_glue")},
			AdditionalSources:  build.FileSet{ndkRoot().Join("sources", "android", "native_app_glue", "android_native_app_glue.c")},
		},
		Spy: cpp.Config{
			Libraries: build.FileSet{"log", "z", "m", "dl"},
		},
	}
	android_arm := base(ndk.APK, "android", "arm").Extend(android_target)
	android_arm64 := base(ndk.APK, "android", "arm64").Extend(android_target)

	return map[string]Target{
		"linux":         linux,
		"osx":           osx,
		"windows":       windows,
		"windows-msvc":  windows_msvc,
		"android-arm":   android_arm,
		"android-arm64": android_arm64,
	}
}

func ndkRoot() build.File {
	paths, _ := ndk.ResolvePaths()
	return paths.NDK
}
