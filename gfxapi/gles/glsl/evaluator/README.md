# evaluator
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/evaluator"

Package evaluator is responsible for evaluating expressions in OpenGL ES Shading
Language programs.

## Usage

#### func  Evaluate

```go
func Evaluate(expr ast.Expression, res func(sym ast.ValueSymbol) ast.Value,
	lang ast.Language) (val ast.Value, err []error)
```
Evaluate evaluates an expression in a GLES Shading Language program. It assumes
the expression has already been typed and statically checked for correctness by
the sema package. This function is suitable to be passed as the evaluator
argument to the sema.Analyze function. Its arguments are:

- expr the expression to evaluate

- res a function resolving Symbols occuring in the expressions to values

- lang the language whose semantics to employ during evaluation

The function returns the value obtained by evaluating the expression. If the
evaluation fails, the second result shall contain a list of errors.

#### func  EvaluatePreprocessorExpression

```go
func EvaluatePreprocessorExpression(e ast.Expression) (val ast.IntValue, err []error)
```
EvaluatePreprocessorExpression is a wrapper around Evaluate, which adapts it to
processing preprocessor expressions. This function is suitable to be passed as
the evaluator argument to parse.Parse function.
