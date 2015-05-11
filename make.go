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

// The fab command is used to build the gpu project.
package main

import (
	"fmt"
	"os"
	"strings"

	"android.googlesource.com/platform/tools/gpu/build"
	. "android.googlesource.com/platform/tools/gpu/maker"
)

func main() { Run() }

const (
	GPURoot = "android.googlesource.com/platform/tools/gpu"
)

var (
	glespath   = Path(gpusrc, "gfxapi/gles")
	gapirpath  = Path(gpusrc, "cc/gapir")
	spypath    = Path(gpusrc, "cc/gfxspy2/src")
	spywinpath = Path(gpusrc, "cc/gfxspy2/src/windows")
	spyosxpath = Path(gpusrc, "cc/gfxspy2/src/osx")
	testpath   = Path(gpusrc, "gfxapi/test")
	glesapi    = Path(gpusrc, "gfxapi/gles/gles.api")
	testapi    = Path(gpusrc, "gfxapi/test/gfxapi_test.api")
	javacore   = Path(Paths.Root, "../base/rpclib/src/main/java/com/android/tools/rpclib/rpccore/")
	javarpc    = Path(Paths.Root, "../adt/idea/android/src/com/android/tools/idea/editors/gfxtrace/rpc")
	gpusrc     = GoSrcPath(GPURoot)

	Tools struct {
		Embed    Entity
		Rpcapi   Entity
		Apic     Entity
		Codergen Entity
		Stringer Entity
	}

	Apps struct {
		Gapis Entity
		Gapir Entity
	}
)

func init() {
	Register(func() {
		// Install rules for the build tools
		Tools.Embed = GoInstall(GPURoot + "/tools/embed")
		Tools.Rpcapi = GoInstall(GPURoot + "/rpc/rpcapi")
		Tools.Apic = GoInstall(GPURoot + "/api/apic")
		Tools.Codergen = GoInstall(GPURoot + "/binary/codergen")
		Tools.Stringer = GoInstall("golang.org/x/tools/cmd/stringer")
		List("tools").DependsStruct(Tools)
		// All the embed rules
		embedRPC := Embed(Path(gpusrc, "rpc/generate"))
		embedCopyright := Embed(Path(gpusrc, "tools/copyright"))
		embedBinary := Embed(Path(gpusrc, "binary/generate"))
		Creator(Tools.Rpcapi).DependsOn(embedCopyright, embedRPC)
		Creator(Tools.Apic).DependsOn(embedCopyright)
		Creator(Tools.Codergen).DependsOn(embedCopyright, embedBinary)
		// All the rpc rules
		servicerpc := File(gpusrc, "service/service.api")
		RpcApiGo(File(gpusrc, "rpc/test/rpc_test.api"))
		RpcApiGo(servicerpc)
		// All the apic rules
		Apic(glespath, glesapi, Path(gpusrc, "gfxapi/templates/api.go.tmpl"))
		Apic(glespath, glesapi, Path(gpusrc, "gfxapi/templates/replay_writer.go.tmpl"))
		Apic(glespath, glesapi, Path(gpusrc, "gfxapi/templates/schema.go.tmpl"))
		Apic(glespath, glesapi, Path(gpusrc, "gfxapi/templates/state_mutator.go.tmpl"))
		Apic(gapirpath, glesapi, Path(gpusrc, "gfxapi/templates/gfx_api.cpp.tmpl"))
		Apic(gapirpath, glesapi, Path(gpusrc, "gfxapi/templates/gfx_api.h.tmpl"))
		Apic(spypath, glesapi, Path(gpusrc, "gfxapi/templates/api_exports.cpp.tmpl"))
		Apic(spypath, glesapi, Path(gpusrc, "gfxapi/templates/api_imports.cpp.tmpl"))
		Apic(spypath, glesapi, Path(gpusrc, "gfxapi/templates/api_imports.h.tmpl"))
		Apic(spypath, glesapi, Path(gpusrc, "gfxapi/templates/api_spy.h.tmpl"))
		Apic(spypath, glesapi, Path(gpusrc, "gfxapi/templates/api_state.h.tmpl"))
		Apic(spypath, glesapi, Path(gpusrc, "gfxapi/templates/api_types.h.tmpl"))
		Apic(spywinpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl32_exports.def.tmpl"))
		Apic(spywinpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl32_resolve.cpp.tmpl"))
		Apic(spywinpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl32_x64.asm.tmpl"))
		Apic(spyosxpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl_framework_exports.cpp.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/api.go.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/replay_writer.go.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/schema.go.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/state_mutator.go.tmpl"))
		// The codergen rule
		Codergen("codergen", "--go", GPURoot+"/...")
		// Enum string rules
		Stringer(Path(gpusrc, "binary/generate"), "Kind")
		Stringer(Path(gpusrc, "log"), "Kind")
		// The java code generation rules
		Codergen("javacoders", "--java", javacore, GPURoot+"/rpc/...")
		RpcApi("--java", javarpc, servicerpc).Creates(Virtual("javarpc"))
		List("java").DependsOn("javacoders", "javarpc")
		//
		List("code").DependsOn("embed", "rpcapi", "apic", "codergen", "stringer")
		// The native code rules
		Apps.Gapir = Virtual("gapir")
		GoRun(Path(gpusrc, "cc/build.go"),
			"--runtests",
			"--targets="+build.HostOS+",android-arm",
		).Creates(Apps.Gapir).DependsOn("code")
		// The testing rules
		gotest := GoTest(GPURoot + "/...")
		Creator(gotest).DependsOn("code", Apps.Gapir)
		List("test").DependsOn(gotest)
		// The main binary rules
		Apps.Gapis = GoInstall(GPURoot + "/server/gapis")
		Creator(Apps.Gapis).DependsOn("code")
		List("apps").DependsStruct(Apps)
		// Application launchers
		Command(Apps.Gapis).Creates(Virtual("gapis")).DependsOn(Apps.Gapir)
		// Utilties
		GoRun(Path(gpusrc, "tools/clean_generated/main.go"), gpusrc).Creates(Virtual("clean_gpu"))
		// The default rules
		List(Default).DependsOn("apps", "test")
	})
}

func Embed(path string) Entity {
	out := File(path, "embed.go")
	args := []string{"--out", out.Name()}
	files := FilesOf(path, func(i os.FileInfo) bool {
		return !i.IsDir() && !strings.HasSuffix(i.Name(), ".go") && !strings.HasSuffix(i.Name(), ".md")
	})
	for _, f := range files {
		args = append(args, f.Name())
	}
	s := Command(Tools.Embed, args...).Creates(out)
	for _, f := range files {
		s.DependsOn(f)
	}
	List("embed").DependsOn(out)
	return out
}

func RpcApiGo(api Entity) {
	out := File(strings.TrimSuffix(api.Name(), ".api") + "_rpc.go")
	RpcApi("--go", DirOf(api).Name(), api).Creates(out)
	List("rpcapi").DependsOn(out)
}

func RpcApi(language string, path string, api Entity) *Step {
	return Command(Tools.Rpcapi, "--dir", path, language, api.Name()).DependsOn(api)
}

func Apic(path string, api string, template string) {
	dst := Dir(path)
	a := File(api)
	t := File(template)
	_, apiname := PathSplit(api)
	_, templatename := PathSplit(template)
	deps := File(Paths.Deps, fmt.Sprintf("%v_%v.deps", apiname, templatename))
	s := Command(Tools.Apic,
		"template",
		"--dir", dst.Name(),
		"--deps", deps.Name(),
		a.Name(), t.Name()).DependsOn(a, t, dst)
	s.UseDepsFile(deps)
	List("apic").DependsOn(deps)
}

func Codergen(name string, args ...string) {
	Command(Tools.Codergen, args...).Creates(Virtual(name)).DependsOn("rpcapi", "apic")
}

func Stringer(path, name string) {
	r := Paths.Root
	defer func() { Paths.Root = r }()
	Paths.Root = path
	e := Virtual("")
	Command(Tools.Stringer, "--type", name).Creates(e)
	List("stringer").DependsOn(e)
}
