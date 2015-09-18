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
	"strings"

	"android.googlesource.com/platform/tools/gpu/maker/build"
)

type misotool func(inputs build.FileSet, output build.File, cfg Config, env build.Environment) error
type sisotool func(input build.File, output build.File, cfg Config, env build.Environment) error
type depsFor func(output build.File, cfg Config, env build.Environment) (deps build.FileSet, valid bool)
type depFileFor func(output build.File, cfg Config, env build.Environment) build.File

// OptimisationLevel is an enumerator of optimisation levels to use by a toolchain.
type OptimizationLevel int

const (
	NoOptimization   = iota // Use to disable optimization. Useful for debugging.
	FullOptimization        // Full optimizations enabled. Fastest option.
)

// Toolchain is a collection of tools used to build objects, libraries and programs.
type Toolchain struct {
	Compiler   sisotool            // Tool used to compile source to object files.
	Archiver   misotool            // Tool used to package object files into archives.
	DllLinker  misotool            // Tool used to link objects and packages into dynamic libraries.
	ExeLinker  misotool            // Tool used to link objects and packages into executables.
	DepFileFor depFileFor          // Returns the name of the dependency file
	DepsFor    depsFor             // Returns the list of dependencies for the given file.
	LibName    func(Config) string // Returns the name of the emitted static library file.
	DllName    func(Config) string // Returns the name of the emitted dynamic library file.
	ExeName    func(Config) string // Returns the name of the emitted executable.
	ObjExt     func(Config) string // Extension used for object files.
}

func (t Toolchain) LibExt(cfg Config) string {
	return build.File(t.LibName(cfg)).Ext()
}

// Config is the configuration for a C++, C or Objective-C++ build.
type Config struct {
	Name               string            // The name that may be mangled to produce the output file.
	OutputExt          *string           // The output file extension. If nil, a toolchain default is used.
	OutputDir          build.File        // The output directory
	Toolchain          *Toolchain        // The toolchain used to build.
	OptimizationLevel  OptimizationLevel // The optimization level to use.
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
	ModuleDefinition   build.File        // An optional module definition file for dlls.
}

// Extend returns a new Config based on this Config, but with field added or
// replaced by any non-default fields of n.
func (c Config) Extend(n Config) Config {
	if n.Name != "" {
		c.Name = n.Name
	}
	if n.OutputExt != nil {
		c.OutputExt = n.OutputExt
	}
	if n.OutputDir != "" {
		c.OutputDir = n.OutputDir
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
	if n.ModuleDefinition != "" {
		c.ModuleDefinition = n.ModuleDefinition
	}

	return c
}

// Triplet returns a string combining the os, architecture and flavour of cfg.
func Triplet(cfg Config) string {
	return strings.Join([]string{cfg.OS, cfg.Architecture, cfg.Flavor}, "-")
}

// IntermediatePath returns a File in the intermediate directory for generating
// a file built from source using cfg and env.
func IntermediatePath(source build.File, ext string, cfg Config, env build.Environment) build.File {
	root := env.Roots.Find(source)
	if root == nil {
		root = &build.RepoRoot
	}
	rel := source.RelativeTo(root.Path)
	rel = strings.Replace(rel, "..", "__", -1) // prevent leaking outside of the Intermediates directory
	out := env.Intermediates.Join(Triplet(cfg), root.Name, rel).ChangeExt(ext)
	return out
}
