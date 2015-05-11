# preprocessor
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/preprocessor"

Package preprocessor defines the preprocessor for the OpenGL ES Shading
Language.

It can be used in two ways. The function Preprocess takes an input strings and
returns the sequence of preprocessed tokens as a result. The function
PreprocessStream returns a Preprocessor object, which can be used to access the
result in a streaming fashion, using the Next() and Peek() methods, returning
TokenInfo objects. The end of input is signalled by a nil Token field of the
object.

The preprocessor does not attempt to distinguish between prefix and postfix
versions of the increment and decrement operators even though it defines the
corresponding tokens. This job is left to the parser, which should not assume
that both UoPreinc and UoPostinc can be returned for a "++" token, etc.

Functions Preprocess and PreprocessStream are parameterized by an
ExpressionEvaluator function, which is used for evaluating #if expressions. The
third parameter is the file identifier, while will be returned by the __FILE__
macro.

## Usage

```go
var (
	KwCentroid  = appendQualifier("centroid")
	KwConst     = appendQualifier("const")
	KwFlat      = appendQualifier("flat")
	KwIn        = appendQualifier("in")
	KwInvariant = appendQualifier("invariant")
	KwLayout    = appendQualifier("layout")
	KwOut       = appendQualifier("out")
	KwSmooth    = appendQualifier("smooth")
	KwUniform   = appendQualifier("uniform")
	KwAttribute = appendQualifier("attribute") // 1.0
	KwVarying   = appendQualifier("varying")   // 1.0

	KwBreak     = appendKeyword("break")
	KwCase      = appendKeyword("case")
	KwContinue  = appendKeyword("continue")
	KwDefault   = appendKeyword("default")
	KwDiscard   = appendKeyword("discard")
	KwDo        = appendKeyword("do")
	KwElse      = appendKeyword("else")
	KwFalse     = appendKeyword("false")
	KwFor       = appendKeyword("for")
	KwIf        = appendKeyword("if")
	KwInout     = appendKeyword("inout")
	KwLowp      = appendKeyword("lowp")
	KwMediump   = appendKeyword("mediump")
	KwHighp     = appendKeyword("highp")
	KwPrecision = appendKeyword("precision")
	KwReturn    = appendKeyword("return")
	KwStruct    = appendKeyword("struct")
	KwSwitch    = appendKeyword("switch")
	KwTrue      = appendKeyword("true")
	KwWhile     = appendKeyword("while")
)
```
All the non-type keywords in the language, one variable per keyword.

```go
var (
	OpColon     = appendOperator(":")
	OpDot       = appendOperator(".")
	OpLBrace    = appendOperator("{")
	OpLBracket  = appendOperator("[")
	OpLParen    = appendOperator("(")
	OpQuestion  = appendOperator("?")
	OpRBrace    = appendOperator("}")
	OpRBracket  = appendOperator("]")
	OpRParen    = appendOperator(")")
	OpSemicolon = appendOperator(";")
)
```
Operator tokens not defined in the ast package.

#### func  IsPrecision

```go
func IsPrecision(token Token) bool
```
Is token a precision keyword?

#### func  IsQualifier

```go
func IsQualifier(token Token) bool
```
Is token a qualifier keyword?

#### func  Preprocess

```go
func Preprocess(data string, eval ExpressionEvaluator, file int) (tokens []TokenInfo, err []error)
```
Preprocess preprocesses a GLES Shading language program string into a sequence
of tokens. Any preprocessing errors are returned in the second return value.

#### type ExpressionEvaluator

```go
type ExpressionEvaluator func(*Preprocessor) (val ast.IntValue, err []error)
```

ExpressionEvaluator is a function type. These functions are used to process #if
expressions. If you want to implement your own expression evaluator, pass a
function which given a Preprocessor, returns a value and a list of errors.
Otherwise, pass ast.EvaluatePreprocessorExpression.

This function is passed a preprocessor object, which reads the tokens following
the #if directive. The tokens are passed as-is, with the exception of
defined(MACRO) expressions, which are evaluated to "0" or "1".

#### type Identifier

```go
type Identifier string
```

Identifier represents a language identifier returned by the lexer.

#### func (Identifier) String

```go
func (id Identifier) String() string
```

#### type Keyword

```go
type Keyword struct {
}
```

Keyword represents a language keyword token returned by the preprocessor.

#### func (Keyword) String

```go
func (key Keyword) String() string
```

#### type Operator

```go
type Operator struct {
}
```

Operator represents an operator returned by the preprocessor.

#### func (Operator) String

```go
func (op Operator) String() string
```

#### type Preprocessor

```go
type Preprocessor struct {
}
```

Preprocessor is the streaming interface to the GLES Shading Language
preprocessor. It supports the usual Peek/Next operations. At any point it can
return the list of detected preprocessor errors.

#### func  PreprocessStream

```go
func PreprocessStream(data string, eval ExpressionEvaluator, file int) *Preprocessor
```

#### func (*Preprocessor) GetErrors

```go
func (p *Preprocessor) GetErrors() []error
```
Return the list of detected preprocessor errors.

#### func (*Preprocessor) Next

```go
func (p *Preprocessor) Next() TokenInfo
```
Return the next token and advance the stream.

#### func (*Preprocessor) Peek

```go
func (p *Preprocessor) Peek() TokenInfo
```
Return the next token.

#### func (*Preprocessor) PeekN

```go
func (p *Preprocessor) PeekN(n int) TokenInfo
```
Return the n-th token (for the next token, set n=0).

#### type Token

```go
type Token interface {
	fmt.Stringer
}
```

An interface for all the tokens returned by the preprocessor.

#### type TokenInfo

```go
type TokenInfo struct {
	Token      Token       // The token.
	Whitespace bool        // Whether this token was preceeded by any whitespace.
	Newline    bool        // Whether this token was preceeded by a newline character.
	Cst        *parse.Leaf // Structure containing the preceeding and following whitespace.
}
```

TokenInfo contains whitespace information about the contained token.
