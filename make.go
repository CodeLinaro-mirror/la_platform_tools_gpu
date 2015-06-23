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
	"flag"
	"fmt"
	"os"
	"strings"

	"android.googlesource.com/platform/tools/gpu/cc"
	. "android.googlesource.com/platform/tools/gpu/maker"
)

var targetOS = flag.String("targetos", HostOS, "target OS to build")

func main() { Run() }

const (
	GPURoot = "android.googlesource.com/platform/tools/gpu"
)

var (
	glespath     = Path(gpusrc, "gfxapi/gles")
	gapirpath    = Path(gpusrc, "cc/gapir")
	gapiipath    = Path(gpusrc, "cc/gapii")
	gapiiwinpath = Path(gpusrc, "cc/gapii/windows")
	gapiiosxpath = Path(gpusrc, "cc/gapii/osx")
	cppcoder     = Path(gpusrc, "cc/gapic/coder")
	testpath     = Path(gpusrc, "gfxapi/test")
	glesapi      = Path(gpusrc, "gfxapi/gles/gles.api")
	testapi      = Path(gpusrc, "gfxapi/test/gfxapi_test.api")
	javabase     = Path(Paths.Root, "../")
	javarpc      = Path(javabase, "adt/idea/android/src/com/android/tools/idea/editors/gfxtrace/rpc")
	gpusrc       = GoSrcPath(GPURoot)

	Tools struct {
		Embed    Entity
		Rpcapi   Entity
		Apic     Entity
		Codergen Entity
		Gapit    Entity
	}

	Apps struct {
		Gapis Entity
		Gapir Entity
		Gapid Entity
	}
)

func init() {
	Register(func() {
		// Install rules for the build tools
		Tools.Embed = GoInstall(GPURoot + "/tools/embed")
		Tools.Rpcapi = GoInstall(GPURoot + "/rpc/rpcapi")
		Tools.Apic = GoInstall(GPURoot + "/api/apic")
		Tools.Codergen = GoInstall(GPURoot + "/binary/codergen")
		Tools.Gapit = GoInstall(GPURoot + "/tools/gapit")
		List("gapit").DependsOn(Tools.Gapit)
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
		Apic(gapiipath, glesapi, Path(gpusrc, "gfxapi/templates/api_exports.cpp.tmpl"))
		Apic(gapiipath, glesapi, Path(gpusrc, "gfxapi/templates/api_imports.cpp.tmpl"))
		Apic(gapiipath, glesapi, Path(gpusrc, "gfxapi/templates/api_imports.h.tmpl"))
		Apic(gapiipath, glesapi, Path(gpusrc, "gfxapi/templates/api_spy.h.tmpl"))
		Apic(gapiipath, glesapi, Path(gpusrc, "gfxapi/templates/api_types.h.tmpl"))
		Apic(gapiiwinpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl32_exports.def.tmpl"))
		Apic(gapiiwinpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl32_resolve.cpp.tmpl"))
		Apic(gapiiwinpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl32_x64.asm.tmpl"))
		Apic(gapiiosxpath, glesapi, Path(gpusrc, "gfxapi/templates/opengl_framework_exports.cpp.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/api.go.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/replay_writer.go.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/schema.go.tmpl"))
		Apic(testpath, testapi, Path(gpusrc, "gfxapi/templates/state_mutator.go.tmpl"))
		// The codergen rule
		Codergen("codergen", "--go", "--java", javabase, "-cpp", cppcoder, GPURoot+"/...")
		// The java code generation rules
		RpcApi("--java", javarpc, servicerpc).Creates(Virtual("javarpc"))
		List("java").DependsOn("codergen", "javarpc")
		//
		List("code").DependsOn("embed", "rpcapi", "apic", "codergen")
		// The native code rules
		Apps.Gapir = Virtual("cc:replayd")
		cctargets := []string{*targetOS}
		if os.Getenv("ANDROID_NDK_ROOT") != "" {
			cctargets = append(cctargets, []string{"android-arm", "android-arm64"}...)
		}
		cc.Graph(cctargets)
		// The testing rules
		gotest := GoTest(GPURoot + "/...")
		// Runtime dependencies
		Creator(Tools.Gapit).DependsOn("cc:spy")
		List("runtime").DependsOn(Apps.Gapir, "cc:spy")
		Creator(gotest).DependsOn("code", "runtime")
		if *targetOS == HostOS {
			List("test").DependsOn("go_test", "cc_test")
		} else {
			List("test").DependsOn("go_test")
		}
		// The main binary rules
		Apps.Gapis = GoInstall(GPURoot + "/server/gapis")
		Creator(Apps.Gapis).DependsOn("code")
		Apps.Gapid = GoInstall(GPURoot + "/_experimental/client/gapid")
		Creator(Apps.Gapid).DependsOn("code")
		List("apps").DependsStruct(Apps)
		// Application launchers
		Command(Apps.Gapis).Creates(Virtual("gapis")).DependsOn(Apps.Gapir)
		Command(Apps.Gapid, "--gxuidebug").Creates(Virtual("gapid")).DependsOn(Apps.Gapis, "runtime")
		// Utilties
		GoRun(Path(gpusrc, "tools/clean_generated/main.go"), gpusrc).Creates(Virtual("clean_gpu"))
		GoRun(Path(gpusrc, "tools/copyright/copyright/main.go"), "-o", gpusrc).Creates(Virtual("copyright")).DependsOn(embedCopyright)
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
		a.Name(), t.Name()).DependsOn(a, t, dst, Dir(Paths.Deps))
	s.UseDepsFile(deps)
	List("apic").DependsOn(deps)
}

func Codergen(name string, args ...string) {
	Command(Tools.Codergen, args...).Creates(Virtual(name)).DependsOn("rpcapi", "apic").Access(GoPkgResources)
}
