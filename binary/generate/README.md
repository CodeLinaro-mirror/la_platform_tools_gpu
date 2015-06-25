# generate
--
    import "android.googlesource.com/platform/tools/gpu/binary/generate"

Package generate has support for generating encode and decode methods for the
binary package automatically.

Package generate has support for generating encode and decode methods for the
binary package automatically.

## Usage

#### func  Sort

```go
func Sort(structs []*Struct)
```
Sort is used to ensure stable ordering of Struct slices. This is to ensure
automatically generated code has minimum diffs. The sort order is by Struct
name, but guarantees dependencies occur first.

#### type Directory

```go
type Directory struct {
	Name       string // The package name (as used in package declarations)
	ImportPath string // The full import path (as used in import statements)
	Dir        string // The actual directory in which the files live
	Scan       bool   // Whether to scan this directory for structs
	Module     Module // The main module data for this directory
	Test       Module // The test module data for this directory
}
```


#### type File

```go
type File struct {
	Copyright  string
	Package    string
	Import     string
	IsTest     bool
	Path       string
	Directives map[string]string
	Structs    []*Struct
	Constants  schema.Constants
	Imports    Imports
	Style
}
```


#### type Generator

```go
type Generator struct {
}
```


#### func  NewGenerator

```go
func NewGenerator() *Generator
```

#### func (*Generator) CppFile

```go
func (g *Generator) CppFile(file *File) ([]byte, error)
```
CppFile generates the all the cpp code for a file with a set of structs.

#### func (*Generator) GoFile

```go
func (g *Generator) GoFile(file *File) ([]byte, error)
```
GoFile generates the all the go code for a file with a set of structs.

#### func (*Generator) JavaFile

```go
func (g *Generator) JavaFile(file *File) ([]byte, error)
```
JavaFile generates the all the java code for a file with a set of structs.

#### type Imports

```go
type Imports map[string]struct{}
```


#### type Loader

```go
type Loader struct {
	Path        string                // The base path of file scanning
	ForceSource bool                  // If true, forces source only imports
	Directories map[string]*Directory // The set of directories considered
	FileSet     *token.FileSet        // The parser file set
}
```


#### func  NewLoader

```go
func NewLoader(path string, forceSource bool) *Loader
```
NewLoader creates a new go source loader.

#### func (*Loader) GetDir

```go
func (l *Loader) GetDir(importPath string) *Directory
```
GetDir returns the Directory that matches the supplied import path. It will add
a new one if needed.

#### func (*Loader) Process

```go
func (l *Loader) Process() error
```
Process generates output data for all packages that have been marked as needing
to be scanned.

#### func (*Loader) ScanFile

```go
func (l *Loader) ScanFile(filename, source string)
```
ScanFile adds a fake package with the file as it's only source.

#### func (*Loader) ScanPackage

```go
func (l *Loader) ScanPackage(importPath string)
```
ScanPackage marks the directory specified by the import path as needing to be
scanned for binary structures.

#### type Module

```go
type Module struct {
	Sources []Source       // The set of sources to parse
	Files   []*ast.File    // The parsed files included
	Types   *types.Package // The resolved type information
	Output  *File          // The prepared structures to generate for
}
```

Module represents a resolvable module. Under normal go layout conditions a
directory has one module that represents the files that are considered when the
directory is imported, and a second one that also includes the test files. They
must be considered separately because otherwise you can get import cycles.

#### type Source

```go
type Source struct {
	Filename   string            // The filename for this source
	Content    interface{}       // The content of this source, see ParseFiles for details.
	AST        *ast.File         // The parsed syntax tree
	Directives map[string]string // the set of comment overrides
}
```

Source holds a file a filename content pair as consumed by go/parser.ParseFile.

#### type Struct

```go
type Struct struct {
	schema.Class
	IDName    string // The name to give the ID of the type.
	Signature string // The full string type signature of the Struct.
}
```

Struct is a description of an encodable struct. Signature includes the package,
name and name and type of all the fields. Any change to the Signature will cause
the ID to change.

#### func  FromTypename

```go
func FromTypename(pkg *types.Package, n *types.TypeName, imports Imports) *Struct
```
FromTypename creates and initializes a Struct from a types.Typename. It assumes
that the typename will map to a types.Struct, and adds all the fields of that
struct to the Struct information.

#### func (*Struct) UpdateID

```go
func (s *Struct) UpdateID()
```
UpdateID recalculates the struct ID from the current signature.

#### type Style

```go
type Style struct {
	ClassPrefix  string
	MemberPrefix string
	Indent       string
}
```
