# msvc
--
    import "android.googlesource.com/platform/tools/gpu/build/cpp/msvc"

Package msvc contains C++ toolchains for building with Microsoft Visual Studio.

## Usage

```go
var MSVC = &cpp.Toolchain{
	Compiler:  compile,
	Archiver:  archive,
	DllLinker: linkDll,
	ExeLinker: linkExe,
	LibName:   func(cfg cpp.Config) string { return cfg.Name + ".lib" },
	DllName:   func(cfg cpp.Config) string { return cfg.Name + ".dll" },
	ExeName:   func(cfg cpp.Config) string { return cfg.Name + ".exe" },
	ObjExt:    func(cpp.Config) string { return ".obj" },
}
```

#### type Paths

```go
type Paths struct {
	Cl   build.File // The MSVC compiler.
	Ml   build.File // The MSVC macro assembler.
	Lib  build.File // The MSVC static library tool.
	Link build.File // The MSVC linker tool.

	IncludeSearchPaths build.FileSet // MSVC additional include paths.
	LibrarySearchPaths build.FileSet // MSVC additional library paths.
}
```

Paths contains the list of directories required by the MSVC toolchain. It is
acquired by calling ResolvePaths.

#### func  ResolvePaths

```go
func ResolvePaths() (Paths, error)
```
ResolvePaths uses the system environment variables and Windows registry to the
MSVC toolchain directories.
