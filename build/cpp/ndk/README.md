# ndk
--
    import "android.googlesource.com/platform/tools/gpu/build/cpp/ndk"

Package ndk contains C++ toolchains for building Android applications and
libraries.

## Usage

```go
var APK = &cpp.Toolchain{
	Compiler:   compile,
	Archiver:   archive,
	DllLinker:  linkSo,
	ExeLinker:  linkApk,
	DepsFor:    depsFor,
	DepFileFor: depFileFor,
	LibName:    func(cfg cpp.Config) string { return "lib" + cfg.Name + ".a" },
	DllName:    func(cfg cpp.Config) string { return cfg.Name + ".so" },
	ExeName:    func(cfg cpp.Config) string { return cfg.Name + ".apk" },
	ObjExt:     func(cfg cpp.Config) string { return ".o" },
}
```
Toolchain for building an Android APK.

```go
var EXE = &cpp.Toolchain{
	Compiler:   compile,
	Archiver:   archive,
	DllLinker:  linkSo,
	ExeLinker:  linkExe,
	DepsFor:    depsFor,
	DepFileFor: depFileFor,
	LibName:    func(cfg cpp.Config) string { return "lib" + cfg.Name + ".a" },
	DllName:    func(cfg cpp.Config) string { return cfg.Name + ".so" },
	ExeName:    func(cfg cpp.Config) string { return cfg.Name },
	ObjExt:     func(cfg cpp.Config) string { return ".o" },
}
```
Toolchain for building an Android executable env, without being packaged into an
APK (user-debug only).

#### type Paths

```go
type Paths struct {
	NDK        build.File // The Android NDK root directory.
	SDK        build.File // The Android SDK root directory.
	Jarsigner  build.File // The jarsigner executable path.
	AAPT       build.File // The Android SDK aapt executable path.
	Zipalign   build.File // The Android SDK zipalign executable path.
	AndroidJar build.File // The Android SDK android.jar file.
}
```

Paths contains the list of directories and files required by the NDK toolchain.
It is acquired by calling ResolvePaths.

#### func  ResolvePaths

```go
func ResolvePaths() (Paths, error)
```
ResolvePaths uses the system environment variables to find all the NDK, SDK and
Java directories and executables required by the NDK toolchain.
