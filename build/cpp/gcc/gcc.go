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

// Package gcc contains C++ toolchains for building with gcc.
package gcc

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/build"
	"android.googlesource.com/platform/tools/gpu/build/cpp"
)

var GCC = &cpp.Toolchain{
	Compiler:  compile,
	Archiver:  archive,
	DllLinker: linkDll,
	ExeLinker: linkExe,
	DepsFor:   depsFor,
	LibName:   func(cfg cpp.Config) string { return "lib" + cfg.Name + ".a" },
	DllName: func(cfg cpp.Config) string {
		switch cfg.OS {
		case "windows":
			return cfg.Name + ".dll"
		case "osx":
			return cfg.Name + ".dylib"
		default:
			return cfg.Name + ".so"
		}
	},
	ExeName: func(cfg cpp.Config) string {
		switch cfg.OS {
		case "windows":
			return cfg.Name + ".exe"
		default:
			return cfg.Name
		}
	},
	ObjExt: func(cpp.Config) string { return ".o" },
}

type tools struct {
	cc build.File
	ar build.File
}

func getTools(cfg cpp.Config) (*tools, error) {
	var t tools

	if build.HostOS == "linux" {
		switch cfg.OS {
		case "linux":
			bin := build.RepoRoot.Path.Join("prebuilts", "gcc", "linux-x86", "host", "x86_64-linux-glibc2.11-4.8", "bin")
			t = tools{
				cc: bin.Join("x86_64-linux-gcc"),
				ar: bin.Join("x86_64-linux-ar"),
			}

		case "windows":
			bin := build.RepoRoot.Path.Join("prebuilts", "gcc", "linux-x86", "host", "x86_64-w64-mingw32-4.8", "bin")
			t = tools{
				cc: bin.Join("x86_64-w64-mingw32-gcc"),
				ar: bin.Join("x86_64-w64-mingw32-ar"),
			}

		default:
			return nil, fmt.Errorf("Cross compiling to '%s' is currently not avaliable", cfg.OS)
		}
	} else {
		t = tools{
			cc: build.File("gcc").LookPath(),
			ar: build.File("ar").LookPath(),
		}
	}

	if !t.cc.Exists() {
		return nil, fmt.Errorf("GCC tool '%s' was not found", t.cc)
	}

	if !t.ar.Exists() {
		return nil, fmt.Errorf("GCC tool '%s' was not found", t.ar)
	}

	return &t, nil
}

func depFileFor(output build.File, cfg cpp.Config, env build.Environment) build.File {
	return cpp.IntermediatePath(output, ".dep", cfg, env)
}

func compile(input build.File, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = env.Logger.Enter("GCC.Compile")

	tools, err := getTools(cfg)
	if err != nil {
		return err
	}

	depfile := depFileFor(output, cfg, env)

	a := append([]string{
		"-c", // Compile to .o
		optFlags(cfg),
		"-g", // Generate debug info
		"-fPIC",
		"-fvisibility=hidden",
		"-fvisibility-inlines-hidden",
		// "-fcolor-diagnostics", clang-only
		"-MMD", "-MF", depfile.Absolute(), // Generate dependency file
	}, cfg.CompilerArgs...)
	for _, isp := range cfg.IncludeSearchPaths {
		a = append(a, fmt.Sprintf("-I%s", isp))
	}
	for n, v := range cfg.Defines {
		a = append(a, fmt.Sprintf("-D%s=%s", n, v))
	}
	a = append(a, input.Absolute(), "-o", string(output))
	return tools.cc.Exec(env, a...)
}

func archive(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = env.Logger.Enter("GCC.Archive")

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

func linkDll(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = env.Logger.Enter("GCC.LinkDll")

	tools, err := getTools(cfg)
	if err != nil {
		return err
	}

	a := append([]string{
		optFlags(cfg),
		"-fPIC",
	}, cfg.LinkerArgs...)

	switch cfg.OS {
	case "osx":
		a = append(a, "-dynamiclib")
		a = append(a, "-compatibility_version", "1.0.0")
		a = append(a, "-current_version", "1.0.0")
	default:
		a = append(a, "-shared")
	}

	for _, lsp := range cfg.LibrarySearchPaths {
		a = append(a, fmt.Sprintf("-L%s", lsp))
	}
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	for _, library := range cfg.Libraries {
		name := strings.TrimPrefix(strings.TrimSuffix(string(library), ".a"), "lib")
		a = append(a, fmt.Sprintf("-l%s", name))
	}
	a = append(a, "-o", string(output))
	return tools.cc.Exec(env, a...)
}

func linkExe(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = env.Logger.Enter("GCC.LinkExe")

	tools, err := getTools(cfg)
	if err != nil {
		return err
	}

	a := append([]string{
		optFlags(cfg),
	}, cfg.LinkerArgs...)
	for _, lsp := range cfg.LibrarySearchPaths {
		a = append(a, fmt.Sprintf("-L%s", lsp))
	}
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	for _, library := range cfg.Libraries {
		name := strings.TrimPrefix(strings.TrimSuffix(string(library), ".a"), "lib")
		a = append(a, fmt.Sprintf("-l%s", name))
	}
	a = append(a, "-o", string(output))
	return tools.cc.Exec(env, a...)
}

func depsFor(output build.File, cfg cpp.Config, env build.Environment) (deps build.FileSet, valid bool) {
	env.Logger = env.Logger.Enter("GCC.Deps")

	depfile := depFileFor(output, cfg, env)

	return cpp.ParseDepFile(depfile, env)
}

func optFlags(cfg cpp.Config) string {
	switch cfg.OptimizationLevel {
	case cpp.NoOptimization:
		return "-O0"
	default:
		return "-O2"
	}
}
