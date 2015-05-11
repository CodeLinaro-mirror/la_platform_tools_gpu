# gcc
--
    import "android.googlesource.com/platform/tools/gpu/build/cpp/gcc"

Package gcc contains C++ toolchains for building with gcc.

## Usage

```go
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
```
