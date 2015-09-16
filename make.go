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
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"android.googlesource.com/platform/tools/gpu/cc"
	. "android.googlesource.com/platform/tools/gpu/maker"
)

var targetOS = flag.String("targetos", HostOS, "target OS to build")

func main() { Run() }

const (
	GPURoot = "android.googlesource.com/platform/tools/gpu"
	// Request to send to gapir to have it shutdown. Note we do not use the
	// value in replay/protocol/connection_type.go to avoid introducing a
	// dependency on the code we are trying to build.
	ReplaydShutdownRequest = 2
)

var (
	gapirpath    = Path(gpusrc, "cc/gapir")
	gapiipath    = Path(gpusrc, "cc/gapii")
	gapiiwinpath = Path(gpusrc, "cc/gapii/windows")
	gapiiosxpath = Path(gpusrc, "cc/gapii/osx")
	cppcoder     = Path(gpusrc, "cc/gapic/coder")
	javabase     = Path(Paths.Root, "../")
	javapaths    = []string{
		Path(javabase, "base/rpclib/src/main/java/com/android/tools/rpclib"),
		Path(javabase, "base/rpclib/src/test/java/com/android/tools/rpclib"),
		Path(javabase, "adt/idea/android/src/com/android/tools/idea/editors/gfxtrace"),
	}
	gpusrc = GoSrcPath(GPURoot)

	Tools struct {
		Embed    Entity
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
		Tools.Embed = GoInstall(GPURoot, "tools/embed")
		Tools.Apic = GoInstall(GPURoot, "api/apic")
		Tools.Codergen = GoInstall(GPURoot, "tools/codergen")
		Tools.Gapit = GoInstall(GPURoot, "tools/gapit")
		List("gapit").DependsOn(Tools.Gapit)
		List("tools").DependsStruct(Tools)
		// All the embed rules
		embedCopyright := Embed(Path(gpusrc, "tools/copyright"))
		embedCodergen := Embed(Path(gpusrc, "tools/codergen/template"))
		Creator(Tools.Apic).DependsOn(embedCopyright)
		Creator(Tools.Codergen).DependsOn(embedCopyright, embedCodergen)
		// All the apic rules
		GfxApi("test", "gfxapi_test.api")
		GfxApi("gles", "gles.api")
		// The codergen rule
		Codergen("codergen", "--go", "--java", javabase, "-cpp", cppcoder, GPURoot+"/...")
		//
		List("code").DependsOn("embed", "apic", "codergen")
		// The native code rules
		cctargets := []string{*targetOS}
		if os.Getenv("ANDROID_NDK_ROOT") != "" {
			cctargets = append(cctargets, []string{"android-arm", "android-arm64"}...)
		}
		cc.Graph(cctargets)
		Apps.Gapir = Virtual("cc:gapir")
		Creator(Apps.Gapir).DependsOn(ShutdownReplayd(), "code")
		Creator("cc:spy").DependsOn("code")
		Creator("cc:gapii").DependsOn("code")
		// The testing rules
		gotest := GoTest(GPURoot + "/...")
		// Runtime dependencies
		Creator(Tools.Gapit).DependsOn("cc:spy")
		List("runtime").DependsOn(Apps.Gapir, "cc:spy", "cc:gapii")
		Creator(gotest).DependsOn("code", "runtime")
		if *targetOS == HostOS {
			List("test").DependsOn("go_test", "cc_test")
		} else {
			List("test").DependsOn("go_test")
		}
		// The main binary rules
		Apps.Gapis = GoInstall(GPURoot, "server/gapis")
		Creator(Apps.Gapis).DependsOn("code")
		Apps.Gapid = GoInstall(GPURoot, "_experimental/client/gapid")
		Creator(Apps.Gapid).DependsOn("code")
		List("apps").DependsStruct(Apps)
		// Application launchers
		Command(Apps.Gapis).Creates(Virtual("gapis")).DependsOn(Apps.Gapir)
		Command(Apps.Gapid, "--gxuidebug").Creates(Virtual("gapid")).DependsOn(Apps.Gapis, "runtime")
		// Utilties
		GoRun(Path(gpusrc, "tools/clean_generated/main.go"), gpusrc).Creates(Virtual("clean_gpu"))
		GoRun(Path(gpusrc, "tools/clean_generated/main.go"), javapaths...).Creates(Virtual("clean_java"))
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
	Command(Tools.Codergen, args...).Creates(Virtual(name)).DependsOn("apic").Access(GoPkgResources)
}

func ShutdownReplayd() Entity {
	e := Virtual("shutdowngapir")
	NewStep(func(*Step) error {
		for _, endpoint := range []string{"localhost:9283", "localhost:9284"} {
			const maxRetries = 10
			for i := 0; i < maxRetries; i++ {
				conn, err := net.Dial("tcp", endpoint)
				if err != nil {
					// Assume this means there is no gapir
					break
				}
				defer conn.Close()
				msg := []byte{ReplaydShutdownRequest}
				n, err := conn.Write(msg)
				if err != nil {
					return fmt.Errorf("Failed to send shutdown request to Replayd %v", err)
				}
				if n != len(msg) {
					return fmt.Errorf("Failed to send shutdown request to Replayd (only sent %v bytes", n)
				}
				if i == maxRetries-1 {
					return fmt.Errorf("Replayd at %v did not die", endpoint)
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
		return nil
	}).Creates(e)
	return e
}

func GfxApi(pkg string, api string) {
	out := Path(gpusrc, "gfxapi", pkg)
	in := Path(out, api)
	templates := Path(gpusrc, "gfxapi", "templates")
	istest := strings.HasSuffix(pkg, "test")
	// go code
	Apic(out, in, Path(templates, "api.go.tmpl"))
	Apic(out, in, Path(templates, "replay_writer.go.tmpl"))
	Apic(out, in, Path(templates, "schema.go.tmpl"))
	Apic(out, in, Path(templates, "state_mutator.go.tmpl"))
	if istest {
		return
	}
	// gapir code
	Apic(gapirpath, in, Path(templates, "gfx_api.cpp.tmpl"))
	Apic(gapirpath, in, Path(templates, "gfx_api.h.tmpl"))
	// gapii code
	Apic(gapiipath, in, Path(templates, "api_exports.cpp.tmpl"))
	Apic(gapiipath, in, Path(templates, "api_imports.cpp.tmpl"))
	Apic(gapiipath, in, Path(templates, "api_imports.h.tmpl"))
	Apic(gapiipath, in, Path(templates, "api_spy.h.tmpl"))
	Apic(gapiipath, in, Path(templates, "api_types.h.tmpl"))
	Apic(gapiiwinpath, in, Path(templates, "opengl32_exports.def.tmpl"))
	Apic(gapiiwinpath, in, Path(templates, "opengl32_resolve.cpp.tmpl"))
	Apic(gapiiwinpath, in, Path(templates, "opengl32_x64.asm.tmpl"))
	Apic(gapiiosxpath, in, Path(templates, "opengl_framework_exports.cpp.tmpl"))
}
