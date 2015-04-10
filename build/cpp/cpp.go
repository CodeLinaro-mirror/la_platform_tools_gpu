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

package cpp

import (
	"errors"
	"path/filepath"
	"strings"
	"sync"

	"android.googlesource.com/platform/tools/gpu/build"
)

type tool func(inputs build.FileSet, output build.File, cfg Config, env build.Environment) error

var sourcePatterns = []string{"*.cpp", "*.c", "*.cc", "*.mm"}

// Toolchain is a collection of tools used to build objects, libraries and programs.
type Toolchain struct {
	Compiler tool                // Tool used to compile source to object files.
	Archiver tool                // Tool used to package object files into archives.
	Linker   tool                // Tool used to link objects and packages into executables.
	ExeName  func(Config) string // Returns the name of the emitted executable.
	LibName  func(Config) string // Returns the name of the emitted static library file.
	ObjExt   func(Config) string // Extension used for object files.
}

func (t Toolchain) LibExt(cfg Config) string {
	return build.File(t.LibName(cfg)).Ext()
}

// Config is the configuration for a C++, C or Objective-C++ build.
type Config struct {
	Name               string            // The name that may be mangled to produce the output file.
	Toolchain          *Toolchain        // The toolchain used to build.
	OS                 string            // The target operating system, e.g. "windows".
	Architecture       string            // The target architecture, e.g. "x64".
	Flavor             string            // Flavor is used to separate different build configurations, e.g. "release".
	Defines            map[string]string // The list of defines used for compilation.
	CompilerArgs       []string          // Custom complier flags.
	ArchiverArgs       []string          // Custom archiver flags.
	LinkerArgs         []string          // Custom linker flags.
	Libraries          build.FileSet     // The list of libraries used by the linker.
	LibrarySearchPaths build.FileSet     // The list of library search paths.
	IncludeSearchPaths build.FileSet     // The list of include search paths.
	AdditionalSources  build.FileSet     // Additional list of source files to compile.
	Permissions        []string          // APK required permissions (NDK only).
}

// Extend returns a new Config based on this Config, but with field added or
// replaced by any non-default fields of n.
func (c Config) Extend(n Config) Config {
	if n.Name != "" {
		c.Name = n.Name
	}
	if n.Toolchain != nil {
		c.Toolchain = n.Toolchain
	}
	if n.OS != "" {
		c.OS = n.OS
	}
	if n.Architecture != "" {
		c.Architecture = n.Architecture
	}
	if n.Flavor != "" {
		c.Flavor = n.Flavor
	}
	for n, v := range n.Defines {
		c.Defines[n] = v
	}
	c.CompilerArgs = append(append([]string{}, c.CompilerArgs...), n.CompilerArgs...)
	c.ArchiverArgs = append(append([]string{}, c.ArchiverArgs...), n.ArchiverArgs...)
	c.LinkerArgs = append(append([]string{}, c.LinkerArgs...), n.LinkerArgs...)
	c.Libraries = c.Libraries.Append(n.Libraries...)
	c.LibrarySearchPaths = c.LibrarySearchPaths.Append(n.LibrarySearchPaths...)
	c.IncludeSearchPaths = c.IncludeSearchPaths.Append(n.IncludeSearchPaths...)
	c.AdditionalSources = c.AdditionalSources.Append(n.AdditionalSources...)
	c.Permissions = append(append([]string{}, c.Permissions...), n.Permissions...)
	return c
}

func combineErrors(errs []error) error {
	msgs := []string{}
	for _, err := range errs {
		if err != nil {
			msgs = append(msgs, err.Error())
		}
	}
	if len(msgs) > 0 {
		return errors.New(strings.Join(msgs, "\n"))
	} else {
		return nil
	}
}

// Compile compiles the list of source files into object files using the Config
// and build Environment. Compile returns the list of object files.
func Compile(sources build.FileSet, cfg Config, env build.Environment) (build.FileSet, error) {
	env.Logger = env.Logger.Enter("C++.Compile")

	wg := sync.WaitGroup{}
	sources = sources.Append(cfg.AdditionalSources...)
	objects := make([]build.File, len(sources))
	errors := make([]error, len(sources))
	wg.Add(len(sources))
	for i, source := range sources {
		i, source, env := i, source, env
		env.Logger = env.Logger.Fork() // Give each go-routine a unique logger context id.
		go func() {
			defer wg.Done()
			object := IntermediatePath(source, cfg.Toolchain.ObjExt(cfg), cfg, env)
			objects[i] = object
			errors[i] = cfg.Toolchain.Compiler(build.FileSet{source}, object, cfg, env)
		}()
	}
	wg.Wait()
	return objects, combineErrors(errors)
}

// Executable links the list of input files into an executable using the Config
// and build Environment. The inputs can be a combination of source files,
// object files and / or library files. Compile returns the output executable
// file.
func Executable(inputs build.FileSet, cfg Config, env build.Environment) (build.File, error) {
	env.Logger = env.Logger.Enter("C++.Executable")

	objects, err := Compile(inputs.Filter(sourcePatterns...), cfg, env)
	if err != nil {
		return "", err
	}

	objects = objects.Append(inputs.Filter("*" + cfg.Toolchain.ObjExt(cfg))...)

	sourceLibraries := build.FileSet{}
	for _, library := range inputs.Filter("*" + cfg.Toolchain.LibExt(cfg)) {
		dir, base := filepath.Split(library.Absolute())
		sourceLibraries = sourceLibraries.Append(build.File(base))
		cfg.LibrarySearchPaths = cfg.LibrarySearchPaths.Append(build.File(dir))
	}

	cfg.Libraries = append(sourceLibraries, cfg.Libraries...)

	output := env.Output.Join(cfg.Toolchain.ExeName(cfg))
	output.MkdirAll()

	return output, cfg.Toolchain.Linker(objects, output, cfg, env)
}

// StaticLibrary archives the list of input files into an static library using
// the Config and build Environment. The inputs can be a combination of source
// files and / or object files. StaticLibrary returns the output static library
// file.
func StaticLibrary(inputs build.FileSet, cfg Config, env build.Environment) (build.File, error) {
	env.Logger = env.Logger.Enter("C++.StaticLibrary")

	objects, err := Compile(inputs.Filter(sourcePatterns...), cfg, env)
	if err != nil {
		return "", err
	}

	objects = objects.Append(inputs.Filter("*" + cfg.Toolchain.ObjExt(cfg))...)

	name := cfg.Toolchain.LibName(cfg)
	output := env.Intermediates.Join(Triplet(cfg), name).ChangeExt(cfg.Toolchain.LibExt(cfg))
	output.MkdirAll()

	return output, cfg.Toolchain.Archiver(objects, output, cfg, env)
}

// Triplet returns a string combining the os, architecture and flavour of cfg.
func Triplet(cfg Config) string {
	return strings.Join([]string{cfg.OS, cfg.Architecture, cfg.Flavor}, "-")
}

// IntermediatePath returns a File in the intermediate directory for generating
// a file built from source using cfg and env.
func IntermediatePath(source build.File, ext string, cfg Config, env build.Environment) build.File {
	rel := source.RelativeTo(build.Root)
	rel = strings.Replace(rel, "..", "__", -1) // prevent leaking outside of the Intermediates directory
	out := env.Intermediates.Join(Triplet(cfg), rel).ChangeExt(ext)
	out.MkdirAll()
	return out
}
