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
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/build"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/maker"
)

// Add steps to the "maker" build system for compiling, linking of c++.

// Target for generating generated code
const code = "code"

var sourcePatterns = []string{"*.cpp", "*.c", "*.cc", "*.mm", "*.asm"}

func makeEntity(f build.File) maker.Entity {
	return maker.File(f.Absolute())
}

func makeEntities(fs build.FileSet) []interface{} {
	var es []interface{}
	for _, f := range fs {
		es = append(es, makeEntity(f))
	}
	return es
}

func makeStep(name string, output build.File, source build.FileSet, always bool, a func(*maker.Step) error) *maker.Step {
	s := maker.NewStep(a)
	e := makeEntity(output)
	s.Creates(e)
	s.DependsOn(maker.DirOf(e))
	s.DependsOn(makeEntities(source)...)
	// This is a bit sad, but codergen does not know what its outputs are.
	s.DependsOn(code)
	if always {
		s.AlwaysRun()
	}
	if name != "" {
		maker.List("cc:" + name).DependsOn(e)
	}
	maker.List("cc").DependsOn(e)
	return s
}

func logger(env build.Environment, name string) log.Logger {
	rootLogger := env.Logger
	if env.Verbose > 0 {
		log.Infof(rootLogger, "Building %s", name)
	}
	return log.Enter(rootLogger, name)
}

// Add make steps to compile all the specified source files.
// Returns fileset of the resulting object files.
// Note we do not get a dependency on the compiler.
func MakeCompile(sources build.FileSet, cfg Config, env build.Environment) build.FileSet {
	ext := cfg.Toolchain.ObjExt(cfg)
	if cfg.OutputExt != nil {
		ext = *cfg.OutputExt
	}

	sources = sources.Append(cfg.AdditionalSources...)
	objects := make([]build.File, len(sources))
	for i, source := range sources {
		i, source, env := i, source, env
		object := IntermediatePath(source, ext, cfg, env)
		objects[i] = object
		prev := maker.Creator(makeEntity(object))
		if prev != nil {
			// skip any objects which have existing build steps
			continue
		}
		s := makeStep("", object, build.Files(source), env.ForceBuild,
			func(*maker.Step) error {
				env.Logger = log.Enter(env.Logger, object.Name())
				env.Logger = log.Enter(env.Logger, "C++.Compile")
				return cfg.Toolchain.Compiler(source, object, cfg, env)
			})
		depsFor := cfg.Toolchain.DepsFor
		if depsFor == nil {
			// If the toolchain can't check dependencies, then we have to build
			s.DependsOn(code)
			s.AlwaysRun()
			continue
		}
		deps, valid := depsFor(object, cfg, env)
		if !valid {
			depFileFor := cfg.Toolchain.DepFileFor
			if depFileFor != nil {
				s.DependsOn(maker.DirOf(depFileFor(object, cfg, env).Absolute()))
			}
			s.DependsOn(code)
			s.AlwaysRun()
			continue
		}
		entities := makeEntities(deps)
		for _, dep := range entities {
			e := maker.EntityOf(dep)
			if e != nil && e.Timestamp().IsZero() {
				// If any dependencies are missing then do code generation
				s.DependsOn(code)
				break
			}
		}
		s.DependsOn(entities...)
	}
	return objects
}

// Add make steps to construct a static library. The inputs may be a
// combination of source and / or object files. Source file inputs will
// generate compilation steps. Returns the output library name.
// If intermediate is true, then the output is placed in the intermediate
// directory instead of the output directory.
// Note we do not get a dependency on the archiver.
func MakeStaticLibrary(inputs build.FileSet, cfg Config, env build.Environment, intermediate bool) build.File {
	objects := MakeCompile(inputs.Filter(sourcePatterns...), cfg, env)
	objects = objects.Append(inputs.Filter("*" + cfg.Toolchain.ObjExt(cfg))...)

	name := cfg.Toolchain.LibName(cfg)
	output := cfg.OutputDir
	if intermediate {
		output = env.Intermediates.Join(Triplet(cfg))
	}
	output = output.Join(name).ChangeExt(cfg.Toolchain.LibExt(cfg))
	if cfg.OutputExt != nil {
		output = output.ChangeExt(*cfg.OutputExt)
	}

	makeStep(cfg.Name, output, objects, env.ForceBuild,
		func(*maker.Step) error {
			env.Logger = log.Enter(logger(env, cfg.Name), "C++.StaticLibrary")
			return cfg.Toolchain.Archiver(objects, output, cfg, env)
		})
	return output
}

// Add make steps to construct a dynamic library. The inputs can be a
// combination of source files, object files and / or library files.
// Source file inputs will generate compilation steps.
// Returns the output library name.
// Note we do not get a dependency on the linker.
func MakeDynamicLibrary(name string, inputs build.FileSet, cfg Config, env build.Environment) build.File {
	objects := MakeCompile(inputs.Filter(sourcePatterns...), cfg, env)
	objects = objects.Append(inputs.Filter("*" + cfg.Toolchain.ObjExt(cfg))...)
	libs := inputs.Filter("*" + cfg.Toolchain.LibExt(cfg))

	libraries := build.FileSet{}
	for _, library := range libs {
		dir, base := filepath.Split(library.Absolute())
		libraries = libraries.Append(build.File(base))
		cfg.LibrarySearchPaths = cfg.LibrarySearchPaths.Append(build.File(dir))
	}

	cfg.Libraries = append(libraries, cfg.Libraries...)

	output := cfg.OutputDir.Join(cfg.Toolchain.DllName(cfg))
	if cfg.OutputExt != nil {
		output = output.ChangeExt(*cfg.OutputExt)
	}

	deps := libs.Append(objects...)
	makeStep(name, output, deps, env.ForceBuild,
		func(*maker.Step) error {
			env.Logger = log.Enter(logger(env, cfg.Name), "C++.DynamicLibrary")
			return cfg.Toolchain.DllLinker(objects, output, cfg, env)
		})
	return output
}

// Add make steps to construct an executable. The inputs can be a combination
// of source files, object files and / or library files. Source file inputs
// will generate compilation steps. Returns the output library name.
// Note we do not get a dependency on the linker.
func MakeExecutable(name string, inputs build.FileSet, cfg Config, env build.Environment) build.File {
	objects := MakeCompile(inputs.Filter(sourcePatterns...), cfg, env)
	objects = objects.Append(inputs.Filter("*" + cfg.Toolchain.ObjExt(cfg))...)
	libs := inputs.Filter("*" + cfg.Toolchain.LibExt(cfg))

	libraries := build.FileSet{}
	for _, library := range libs {
		dir, base := filepath.Split(library.Absolute())
		libraries = libraries.Append(build.File(base))
		cfg.LibrarySearchPaths = cfg.LibrarySearchPaths.Append(build.File(dir))
	}

	cfg.Libraries = append(libraries, cfg.Libraries...)

	output := cfg.OutputDir.Join(cfg.Toolchain.ExeName(cfg))
	if cfg.OutputExt != nil {
		output = output.ChangeExt(*cfg.OutputExt)
	}

	deps := libs.Append(objects...)
	makeStep(name, output, deps, env.ForceBuild,
		func(*maker.Step) error {
			env.Logger = log.Enter(logger(env, cfg.Name), "C++.Executable")
			return cfg.Toolchain.ExeLinker(objects, output, cfg, env)
		})
	return output
}

// Add a make step to run a test. The test only runs on the host OS.
// The test is automatically included in the "test" target.
func MakeRunTest(test build.File, cfg Config) {
	// We only run tests on the host OS.
	if cfg.OS == maker.HostOS {
		// cfg.Defines["LOG_LEVEL"] = "0" // Disable all log messages except for GAPID_FATAL.
		phony := maker.Virtual("cc:" + test.Name() + ":run")
		maker.Command(makeEntity(test)).Creates(phony)
		maker.List("cc_test").DependsOn(phony)
	}
}

// Add a make step to copy a file from "src" to "dest". The copy only
// happens on the host OS.
func MakeCopy(src build.File, dest build.File, cfg Config, env build.Environment) {
	logger := env.Logger
	if cfg.OS == maker.HostOS {
		makeStep(cfg.Name, dest, build.Files(src), env.ForceBuild,
			func(*maker.Step) error {
				log.Infof(logger, "Copying %s -> %s", src, dest)
				if err := src.CopyTo(dest); err != nil {
					log.Errorf(logger, "Copy failed: %v", err)
					// Error delibrately suppressed
				}
				return nil
			}).AlwaysRun()
	}
}
