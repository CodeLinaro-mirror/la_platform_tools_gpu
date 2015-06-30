# build
--
    import "android.googlesource.com/platform/tools/gpu/build"

Package build contains types and variables used for creating builds.

## Usage

```go
var RepoRoot = func() Root {
	gopaths := filepath.SplitList(os.Getenv("GOPATH"))
	for _, gopath := range gopaths {
		if File(gopath).Join("src", "android.googlesource.com", "platform", "tools", "gpu", "build", "root.go").Exists() {
			return Root{Name: "repo", Path: File(File(gopath).Join("..", "..").Absolute())}
		}
	}
	panic(fmt.Errorf("Project could not be found in any GOPATH: %v", gopaths))
}()
```
Root returns the repo-root of the project.

#### type Environment

```go
type Environment struct {
	Intermediates File       // The intermediates directory
	Roots         RootList   // List of root directories.
	Keystore      File       // The path to the keystore used to sign APKs.
	Storepass     string     // The password to the keystore.
	Keypass       string     // The password to the key in the keystore.
	Keyalias      string     // The alias of the key used to sign APKs.
	Logger        log.Logger // The logger to emit log messages to.
	ForceBuild    bool       // If true, all build steps will be forced.
	Verbose       bool       // If true, logging should be verbose.
}
```

Environment holds the global environment configuration for the build.

#### type File

```go
type File string
```

File represents the path to a file or directory.

#### func  Path

```go
func Path(segments ...string) File
```
Path returns a File formed from the joined path segments.

#### func (File) Absolute

```go
func (f File) Absolute() string
```
Absolute returns the absolute path of this File.

#### func (File) ChangeExt

```go
func (f File) ChangeExt(ext string) File
```
ChangeExt returns a new File with the extension changed to ext.

#### func (File) Contains

```go
func (d File) Contains(f File) bool
```
Contains returns true if this directory contains the file f.

#### func (File) CopyTo

```go
func (f File) CopyTo(dst File) error
```
CopyTo copied this File to dst, replacing any existing file at dst.

#### func (File) Dir

```go
func (f File) Dir() string
```
Dir returns the directory part of the File (without filename).

#### func (File) Exec

```go
func (f File) Exec(env Environment, args ...string) error
```
Exec executes this File with the specified arguments.

#### func (File) ExecAt

```go
func (f File) ExecAt(env Environment, wd File, args ...string) error
```
ExecAt executes this File with the specified arguments with the working
directory set to wd.

#### func (File) Exists

```go
func (f File) Exists() bool
```
Exists returns true if this File exists.

#### func (File) Ext

```go
func (f File) Ext() string
```
Ext returns the extension of the file, including the '.', or an empty string if
the file has no extension.

#### func (File) Glob

```go
func (f File) Glob(patterns ...string) FileSet
```
Glob returns a FileSet of all files matching any of the specified patterns.

#### func (File) Join

```go
func (f File) Join(ext ...string) File
```
Join returns a File formed from joining this File with ext.

#### func (File) LookPath

```go
func (f File) LookPath() (File, error)
```
LookPath returns the path to f, searching the system PATHs.

#### func (File) Matches

```go
func (f File) Matches(patterns ...string) bool
```
Matches returns true if the File matches any of the patterns.

#### func (File) MkdirAll

```go
func (f File) MkdirAll()
```
MkdirAll creates the directory hierarchy to hold this File.

#### func (File) Name

```go
func (f File) Name() string
```
Name returns the name part of the File (without directories).

#### func (File) RelativeTo

```go
func (f File) RelativeTo(base File) string
```
RelativeTo returns the path of this File relative to base.

#### type FileSet

```go
type FileSet []File
```

FileSet is a list of Files with no duplicates.

#### func  Files

```go
func Files(files ...File) FileSet
```
Files returns a FileSet formed from the list of files.

#### func (FileSet) Append

```go
func (fs FileSet) Append(files ...File) FileSet
```
Append returns a new FileSet with files appended to fs. Duplicates will not be
included in the returned FileSet.

#### func (FileSet) Exclude

```go
func (fs FileSet) Exclude(patterns ...string) FileSet
```
Exclude returns the list of files in this FileSet that does not match any
pattern in patterns.

#### func (FileSet) Filter

```go
func (fs FileSet) Filter(patterns ...string) FileSet
```
Filter returns the list of files in this FileSet that matches any pattern in
patterns.

#### type Root

```go
type Root struct {
	Name string // Name of this root.
	Path File   // Root path
}
```

Root describes a named root directory.

#### type RootList

```go
type RootList []Root
```

RootList is a list of roots.

#### func (RootList) Find

```go
func (l RootList) Find(f File) *Root
```
Find returns the first Root that contains the file f, or nil if no roots contain
the file.
