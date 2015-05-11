# sema
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/sema"

Package sema performs semantic checking of OpenGL ES Shading Language programs.

## Usage

#### func  Analyze

```go
func Analyze(program interface{}, evaluator Evaluator) (err []error)
```
Analyze is the main entry point of the package. It performs a semantic check of
a GLES Shading Language program. Its arguments are the AST representation of the
program, and a constant expression evaluating function. This function is used
for evaluating array size expressions and the values of constant variables. An
implementation of such a function can be found in the evaluator package.

The Analyze function is a work in progress. Currently it just computes the types
of all the expressions and evaluates all constant expressions.

#### type Evaluator

```go
type Evaluator func(expr ast.Expression, resolver func(symbol ast.ValueSymbol) ast.Value, lang ast.Language) (val ast.Value, err []error)
```

Evaluator is the type of the function used for evaluating constant expressions.
