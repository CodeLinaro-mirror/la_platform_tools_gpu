# glsl
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl"

Package glsl contains routines for manipulation of OpenGL ES Shading Language
programs.

It exposes functions for parsing, serializing and evaluating GLES Shading
Language programs. While this package contains a number of sub-packages, the
only sub-package which is expected to be imported directly is the ast package,
which contains the definitions of the AST of the parsed program. The main
functionality of the other packages is exposed through the functions of this
package.

## Usage

#### func  Analyze

```go
func Analyze(program interface{}) (err []error)
```
Analyze performs semantic analysis on the parsed program AST. It computes the
types of all program expression, array sizes and values of constant variables.
Any encountered errors are returned as a result.

#### func  Formatter

```go
func Formatter(node interface{}) fmt.Formatter
```
Formatter is a helper function which turns any AST node into something that can
be printed with %v. The returned object's default format will print the tree
under the ast node in a reindented form. The alternate format flag (%#v) will
print the node while preserving original whitespace, if this is present in the
***Cst nodes of the tree.

#### func  Parse

```go
func Parse(src string, lang ast.Language) (program interface{}, err []error)
```
Parse preprocesses and parses an OpenGL ES Shading language program present in
the first argument. The second argument specifies the language, whose syntax to
employ during parsing. The parsed AST is returned in the first result. If any
parsing errors are encountered, they are returned in the second result.
