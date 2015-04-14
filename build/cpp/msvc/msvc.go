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

// Package msvc contains C++ toolchains for building with Microsoft Visual
// Studio.
package msvc

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/build"
	"android.googlesource.com/platform/tools/gpu/build/cpp"
)

var MSVC = &cpp.Toolchain{
	Compiler: compile,
	Archiver: archive,
	Linker:   link,
	ExeName:  func(cfg cpp.Config) string { return cfg.Name + ".exe" },
	LibName:  func(cfg cpp.Config) string { return cfg.Name + ".lib" },
	ObjExt:   func(cpp.Config) string { return ".obj" },
}

// Paths contains the list of directories required by the MSVC toolchain.
// It is acquired by calling ResolvePaths.
type Paths struct {
	resolved bool

	Cl   build.File // The MSVC compiler.
	Lib  build.File // The MSVC static library tool.
	Link build.File // The MSVC linker tool.

	IncludeSearchPaths build.FileSet // MSVC additional include paths.
	LibrarySearchPaths build.FileSet // MSVC additional library paths.
}

func compile(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = env.Logger.Enter("MSVC.Compile")

	paths, err := ResolvePaths()
	if err != nil {
		return err
	}

	a := append([]string{
		"/c",    // Compile, don't link
		"/EHsc", // Enable exceptions
		"/nologo",
	}, cfg.CompilerArgs...)

	for _, isp := range cfg.IncludeSearchPaths.Append(paths.IncludeSearchPaths...) {
		a = append(a, fmt.Sprintf("/I%s", isp))
	}
	for n, v := range cfg.Defines {
		a = append(a, fmt.Sprintf("/D%s=%s", n, v))
	}
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	a = append(a, "/Fo:"+string(output))
	return paths.Cl.ExecAt(env, build.File(paths.Cl.Dir()), a...)
}

func archive(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = env.Logger.Enter("MSVC.Archive")

	paths, err := ResolvePaths()
	if err != nil {
		return err
	}

	a := []string{
		"/nologo",
	}
	a = append(a, cfg.ArchiverArgs...)
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	a = append(a, "/OUT:"+output.Absolute())
	return paths.Lib.Exec(env, a...)
}

func link(inputs build.FileSet, output build.File, cfg cpp.Config, env build.Environment) error {
	env.Logger = env.Logger.Enter("MSVC.Link")

	paths, err := ResolvePaths()
	if err != nil {
		return err
	}

	a := append([]string{
		"/MACHINE:X64",
		"/SUBSYSTEM:CONSOLE",
		"/nologo",
	}, cfg.LinkerArgs...)
	for _, lsp := range cfg.LibrarySearchPaths.Append(paths.LibrarySearchPaths...) {
		a = append(a, fmt.Sprintf("/LIBPATH:%s", lsp))
	}
	for _, input := range inputs {
		a = append(a, input.Absolute())
	}
	for _, library := range cfg.Libraries {
		a = append(a, string(library+".lib"))
	}
	a = append(a, "/OUT:"+string(output))
	return paths.Link.Exec(env, a...)
}
