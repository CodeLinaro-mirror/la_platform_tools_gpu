# parser
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/parser"

Package parser implements parsing and serializing an OpenGL ES Shading Language
programs.

## Usage

#### func  Formatter

```go
func Formatter(e interface{}) fmt.Formatter
```
Formatter is a helper function which turns any AST node into something that can
be printed with %v. The returned object's default format will print the tree
under the ast node in a reindented form. The alternate format flag (%#v) will
print the node while preserving original whitespace, if this is present in the
***Cst nodes of the tree.

#### func  Parse

```go
func Parse(in string, language ast.Language,
	eval PreprocessorExpressionEvaluator) (program interface{}, err []error)
```
Parse is the main entry point into the parser. It parses GLES Shading Language
programs and returns their AST representations. With suitable arguments it can
also parse preprocessor #if expressions. It's arguments are:

- in: the string to parse. It is automatically preprocessed.

- language: which language are we parsing. In case we are parsing #if
expressions (ast.LangPreprocessor), the eval function can be null.

- eval: An evaluator function which evaluates preprocessor #if expressions. It
is only used when parsing full GLES Shading Language programs
(ast.LangVertexShader, ast.LangFragmentShader).

The result is an object of type *ast.Ast in case of vertex and fragment shaders.
In case of preprocessor expressions the result is an ast.Expression interface.
The function also returns any errors it encounters during processing.

#### type PreprocessorExpressionEvaluator

```go
type PreprocessorExpressionEvaluator func(expr ast.Expression) (val ast.IntValue, err []error)
```

PreprocessorExpressionEvaluator is the type of the function arugment of the
Parse function. These functions are supposed to evaluate preprocessor
expressions, given an ast.Expression and return an IntValue and possibly a list
of errors.
