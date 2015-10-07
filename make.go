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
	"fmt"
	"net"
	"strings"
	"time"

	"android.googlesource.com/platform/tools/gpu/cc"
	"android.googlesource.com/platform/tools/gpu/maker/config"
	"android.googlesource.com/platform/tools/gpu/maker/do"
	"android.googlesource.com/platform/tools/gpu/maker/graph"
	"android.googlesource.com/platform/tools/gpu/maker/run"
)

func main() { run.Run() }

const (
	GPURoot = "android.googlesource.com/platform/tools/gpu"
	// Request to send to gapir to have it shutdown. Note we do not use the
	// value in replay/protocol/connection_type.go to avoid introducing a
	// dependency on the code we are trying to build.
	ReplaydShutdownRequest = 2
)

var (
	gpusrc       = do.GoSrcPath(GPURoot)
	gapirpath    = gpusrc.Child("cc/gapir")
	gapiipath    = gpusrc.Child("cc/gapii")
	gapiiwinpath = gpusrc.Child("cc/gapii/windows")
	gapiiosxpath = gpusrc.Child("cc/gapii/osx")
	cppcoder     = gpusrc.Child("cc/gapic/coder")
	javabase     = config.Paths.Repo.Child("tools")
	javapaths    = []interface{}{
		javabase.Child("base/rpclib/src/main/java/com/android/tools/rpclib"),
		javabase.Child("base/rpclib/src/test/java/com/android/tools/rpclib"),
		javabase.Child("adt/idea/android/src/com/android/tools/idea/editors/gfxtrace"),
	}
	pkginfobase = gpusrc.Child("tools/pkginfo")

	Tools struct {
		Embed    graph.Entity
		Apic     graph.Entity
		Codergen graph.Entity
	}

	Apps struct {
		Gapis graph.Entity
		Gapir graph.Entity
		Gapid graph.Entity
		Gapit graph.Entity
	}
)

func init() {
	run.Register(func() {
		// Install rules for the build tools
		Tools.Embed = do.GoInstall(GPURoot, "tools/embed")
		Tools.Apic = do.GoInstall(GPURoot, "api/apic")
		Tools.Codergen = do.GoInstall(GPURoot, "tools/codergen")
		graph.List("tools").DependsStruct(Tools)
		// All the embed rules
		embedCopyright := Embed(gpusrc.Child("tools/copyright"))
		embedCodergen := Embed(gpusrc.Child("tools/codergen/template"))
		graph.Creator(Tools.Apic).DependsOn(embedCopyright)
		graph.Creator(Tools.Codergen).DependsOn(embedCopyright, embedCodergen)
		// All the apic rules
		GfxApi("test", "gfxapi_test.api")
		GfxApi("gles", "gles.api")
		// The codergen rule
		Codergen("codergen",
			"--signatures", gpusrc.Child("signatures.txt").Name(),
			"--go",
			"--java", javabase.Name(),
			"-cpp", cppcoder.Name(),
			GPURoot+"/...")
		//
		graph.List("code").DependsOn("embed", "apic", "codergen")
		// The native code rules
		cc.Graph()
		Apps.Gapir = graph.Virtual("cc:gapir")
		graph.Creator(Apps.Gapir).DependsOn(ShutdownReplayd(), "code")
		graph.Creator("cc:spy").DependsOn("code")
		graph.Creator("cc:gapii").DependsOn("code")
		// The testing rules
		gotest := do.GoTest(GPURoot + "/...")
		// Runtime dependencies
		graph.List("runtime").DependsOn(Apps.Gapir, "cc:spy", "cc:gapii")
		graph.Creator(gotest).DependsOn("code", "runtime")
		if config.TargetOS == config.HostOS {
			graph.List("test").DependsOn("go_test", "cc_test")
		} else {
			graph.List("test").DependsOn("go_test")
		}
		// The main binary rules
		Apps.Gapis = do.GoInstall(GPURoot, "server/gapis")
		graph.Creator(Apps.Gapis).DependsOn("code")
		Apps.Gapid = do.GoInstall(GPURoot, "_experimental/client/gapid")
		graph.Creator(Apps.Gapid).DependsOn("code")
		Apps.Gapit = do.GoInstall(GPURoot, "tools/gapit")
		graph.Creator(Apps.Gapit).DependsOn("code")
		graph.List("apps").DependsStruct(Apps)
		// Application launchers
		do.Exec(Apps.Gapis).Creates(graph.Virtual("gapis")).DependsOn(Apps.Gapir)
		do.Exec(Apps.Gapid, "--gxuidebug").Creates(graph.Virtual("gapid")).DependsOn(Apps.Gapis, "runtime")
		// Utilties
		do.GoRun(gpusrc.File("tools/clean_generated/main.go"), gpusrc).Creates(graph.Virtual("clean_gpu"))
		do.GoRun(gpusrc.File("tools/clean_generated/main.go"), javapaths...).Creates(graph.Virtual("clean_java"))
		do.GoRun(gpusrc.File("tools/copyright/copyright/main.go"), "-o", gpusrc).Creates(graph.Virtual("copyright")).DependsOn(embedCopyright)
		// pkginfo helper APK
		PkgInfo()
		// The default rules
		graph.List(graph.Default).DependsOn("apps", "test")
	})
}

func Embed(path *graph.Path) *graph.Path {
	out := path.File("embed.go")
	args := []string{"--out", out.Name()}
	files := path.Files().Exclude("*.go")
	for _, f := range files {
		args = append(args, f.Name())
	}
	s := do.Exec(Tools.Embed, args...).Creates(out)
	for _, f := range files {
		s.DependsOn(f)
	}
	graph.List("embed").DependsOn(out)
	return out
}

func Apic(dst *graph.Path, api *graph.Path, template *graph.Path) {
	deps := config.Paths.Deps.File(fmt.Sprintf("%v_%v.deps", api.Basename(), template.Basename()))
	s := do.Exec(Tools.Apic,
		"template",
		"--dir", dst.Name(),
		"--deps", deps.Name(),
		api.Name(), template.Name()).DependsOn(api, template, dst, config.Paths.Deps)
	s.UseDepsFile(deps)
	graph.List("apic").DependsOn(deps)
}

func Codergen(name string, args ...string) {
	do.Exec(Tools.Codergen, args...).Creates(graph.Virtual(name)).DependsOn("apic").Access(do.GoPkgResources)
}

func PkgInfo() {
	gradleAPK := pkginfobase.File("app/build/outputs/apk/app-debug.apk")
	binAPK := config.Paths.Bin.File("pkginfo.apk")
	graph.NewStep(func(step *graph.Step) error {
		return do.ExecAt(pkginfobase.Name(), config.Verbose, pkginfobase.File("gradlew").Name(), "build")
	}).Creates(gradleAPK)
	do.CopyFile(binAPK, gradleAPK)
	graph.List("pkginfo").DependsOn(binAPK)
}

func ShutdownReplayd() graph.Entity {
	e := graph.Virtual("shutdowngapir")
	graph.NewStep(func(*graph.Step) error {
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
	out := gpusrc.Child("gfxapi", pkg)
	in := out.File(api)
	templates := gpusrc.Child("gfxapi", "templates")
	istest := strings.HasSuffix(pkg, "test")
	// go code
	Apic(out, in, templates.File("api.go.tmpl"))
	Apic(out, in, templates.File("replay_writer.go.tmpl"))
	Apic(out, in, templates.File("schema.go.tmpl"))
	Apic(out, in, templates.File("state_mutator.go.tmpl"))
	if istest {
		return
	}
	// gapir code
	Apic(gapirpath, in, templates.File("gfx_api.cpp.tmpl"))
	Apic(gapirpath, in, templates.File("gfx_api.h.tmpl"))
	// gapii code
	Apic(gapiipath, in, templates.File("api_exports.cpp.tmpl"))
	Apic(gapiipath, in, templates.File("api_imports.cpp.tmpl"))
	Apic(gapiipath, in, templates.File("api_imports.h.tmpl"))
	Apic(gapiipath, in, templates.File("api_spy.h.tmpl"))
	Apic(gapiipath, in, templates.File("api_types.h.tmpl"))
	Apic(gapiiwinpath, in, templates.File("opengl32_exports.def.tmpl"))
	Apic(gapiiwinpath, in, templates.File("opengl32_resolve.cpp.tmpl"))
	Apic(gapiiwinpath, in, templates.File("opengl32_x64.asm.tmpl"))
	Apic(gapiiosxpath, in, templates.File("opengl_framework_exports.cpp.tmpl"))
}
