# parse
--
    import "android.googlesource.com/platform/tools/gpu/parse"

Package parse provides support functionality for writing scannerless parsers.

The main entry point is the Parse function. It works by setting up a parser on
the supplied content, and then invoking the supplied root parsing function. The
CST is built for you inside the ParseLeaf and ParseBranch methods of the parser,
but it is up to the supplied parsing functions to hold on to the CST if you want
it, and also to build the AST.

## Usage

```go
const RuneEOL = '\n'
```
RuneEOL is the rune that marks the end of a line.

```go
var (
	// ParseErrorLimit is the maximum number of errors before a parse is aborted.
	ParseErrorLimit = 10
	// AbortParse is paniced when a parse cannot continue. It is recovered at the
	// top level, to allow the errors to be cleanly returned to the caller.
	AbortParse = errors.New("abort")
)
```

```go
var ConstantValues schema.Constants
```

#### func  Parse

```go
func Parse(root BranchParser, data string, skip Skip) []Error
```
Parse is the main entry point to the parse library. Given a root parse function,
the input string and the Skip controller, it builds a and initializes a Parser,
runs the root using it, verifies it worked correctly and then returns the errors
generated if any.

#### type Branch

```go
type Branch struct {

	// Children is the slice of child nodes for this Branch.
	Children []Node
}
```

Branch is a CST node that can have children.

#### func (*Branch) AddPrefix

```go
func (n *Branch) AddPrefix(s Separator)
```

#### func (*Branch) AddSuffix

```go
func (n *Branch) AddSuffix(s Separator)
```

#### func (*Branch) Prefix

```go
func (n *Branch) Prefix() Separator
```

#### func (*Branch) Suffix

```go
func (n *Branch) Suffix() Separator
```

#### func (*Branch) Token

```go
func (n *Branch) Token() Token
```

#### func (*Branch) WriteTo

```go
func (n *Branch) WriteTo(w io.Writer) error
```

#### type BranchParser

```go
type BranchParser func(p *Parser, cst *Branch)
```

BranchParser is a function that is passed to ParseBranch. It is handed the
Branch to fill in and the Parser to fill it from, and must either succeed or add
an error to the parser.

#### type Error

```go
type Error struct {
	// At is the parse fragment that was being processed when the error was encountered.
	At Fragment
	// Message is the message associated with the error.
	Message string
	// Stack is the captured stack trace at the point the error was noticed.
	Stack []byte
}
```

Error represents the information that us useful in debugging a parse failure.

#### func (Error) Error

```go
func (err Error) Error() string
```

#### func (Error) Format

```go
func (err Error) Format(f fmt.State, c rune)
```

#### type ErrorList

```go
type ErrorList []Error
```

ErrorList is a convenience type for managing lists of errors.

#### func (*ErrorList) Add

```go
func (l *ErrorList) Add(r *Reader, at Fragment, message string, args ...interface{})
```

#### func (ErrorList) Error

```go
func (errs ErrorList) Error() string
```

#### type Fragment

```go
type Fragment interface {
	// Token returns the underlying token of this node.
	Token() Token
	// Write is used to write the underlying token out to the writer.
	WriteTo(io.Writer) error
}
```

Fragment is a component of a cst that is backed by a token. This includes Nodes
and all forms of space and comment.

#### func  NewFragment

```go
func NewFragment(token Token) Fragment
```

#### type Leaf

```go
type Leaf struct {
}
```

Leaf nodes are part of the cst that cannot have child nodes, they represent a
single token from the input.

#### func (*Leaf) AddPrefix

```go
func (n *Leaf) AddPrefix(s Separator)
```

#### func (*Leaf) AddSuffix

```go
func (n *Leaf) AddSuffix(s Separator)
```

#### func (*Leaf) Prefix

```go
func (n *Leaf) Prefix() Separator
```

#### func (*Leaf) SetToken

```go
func (n *Leaf) SetToken(token Token)
```

#### func (*Leaf) Suffix

```go
func (n *Leaf) Suffix() Separator
```

#### func (*Leaf) Token

```go
func (n *Leaf) Token() Token
```

#### func (*Leaf) WriteTo

```go
func (n *Leaf) WriteTo(w io.Writer) error
```

#### type LeafParser

```go
type LeafParser func(p *Parser, cst *Leaf)
```

LeafParser is a function that is passed to ParseLeaf. It is handed the Leaf to
fill in and the Parser to fill it from, and must either succeed or add an error
to the parser.

#### type Node

```go
type Node interface {
	Fragment
	// Prefix returns the set of skippable fragments associated with this Node
	// that precede it in the stream. Association is defined by the Skip function
	// in use.
	Prefix() Separator
	// AddPrefix adds more fragments to the Prefix list.
	AddPrefix(Separator)
	// Suffix returns the set of skippable fragments associated with this Node
	// that follow it in the stream. Association is defined by the Skip function
	// in use, the default is until the end of the line.
	Suffix() Separator
	// AddSuffix adds more fragments to the Suffix list.
	AddSuffix(Separator)
}
```

Node is a Fragment in a cst that represents unskipped tokens.

#### type NumberKind

```go
type NumberKind uint8
```

NumberKind is a type used by Reader.Numeric for identifying various kinds of
numbers.

```go
const (
	// No number was found.
	NotNumeric NumberKind = iota
	// A decimal number.
	Decimal
	// An octal number, starting with "0". PS: A lone "0" is classified as octal.
	Octal
	// A hexadecimal number, starting with "0x".
	Hexadecimal
	// A floating point number: "123.456". Whole and the fractional parts are optional (but
	// not both at the same time).
	Floating
	// A floating point number in scientific notation: "123.456e±789". The fractional part,
	// the dot and the exponent sign are all optional.
	Scientific
)
```

#### func (*NumberKind) Parse

```go
func (v *NumberKind) Parse(s string) error
```

#### func (NumberKind) String

```go
func (v NumberKind) String() string
```

#### type Parser

```go
type Parser struct {
	Reader           // The token reader for this parser.
	Errors ErrorList // The set of errors generated during the parse.
}
```

Parser contains all the context needed while parsing. They are built for you by
the Parse function.

#### func (*Parser) Error

```go
func (p *Parser) Error(message string, args ...interface{})
```
Error adds a new error to the parser error list. It will attempt to consume a
token from the reader to act as a place holder, and also to ensure progress is
made in the presence of errors.

#### func (*Parser) ErrorAt

```go
func (p *Parser) ErrorAt(at Fragment, message string, args ...interface{})
```
ErrorAt is like Error, except because it is handed a fragment, it will not try
to consume anything itself.

#### func (*Parser) Expected

```go
func (p *Parser) Expected(value string)
```
Expected is a wrapper around p.ErrorAt for the very common case of an unexpected
input. It uses value as the expected input, and parses a token of the stream for
the unexpected actual input.

#### func (*Parser) ParseBranch

```go
func (p *Parser) ParseBranch(cst *Branch, do BranchParser)
```
ParseLeaf adds a new Branch to cst and then calls the do function to parse the
branch. This is called recursively to build the node tree.

#### func (*Parser) ParseLeaf

```go
func (p *Parser) ParseLeaf(cst *Branch, do LeafParser)
```
ParseLeaf adds a new Leaf to cst and then calls the do function to parse the
Leaf. If do is nil, a leaf will be built with the current unconsumed input.

#### type Reader

```go
type Reader struct {
}
```

Reader is the interface to an object that converts a rune array into tokens.

#### func  NewReader

```go
func NewReader(data string) *Reader
```
NewReader creates a new reader which reads from the supplied string.

#### func (*Reader) Advance

```go
func (r *Reader) Advance()
```
Advance moves the cursor one rune forward.

#### func (*Reader) AlphaNumeric

```go
func (r *Reader) AlphaNumeric() bool
```
AlphaNumeric moves past anything that starts with a letter or underscore, and
consists of letters, numbers or underscores. It returns true if the pattern was
matched, false otherwise.

#### func (*Reader) Consume

```go
func (r *Reader) Consume() Token
```
Consume consumes the current token.

#### func (*Reader) GuessNextToken

```go
func (r *Reader) GuessNextToken() Token
```
GuessNextToken attempts to do a general purpose consume of a single arbitrary
token from the stream. It is used by error handlers to indicate where the error
occurred. It guarantees that if the stream is not finished, it will consume at
least one character.

#### func (*Reader) IsEOF

```go
func (r *Reader) IsEOF() bool
```
IsEOF returns true when the cursor is at the end of the input.

#### func (*Reader) NotSpace

```go
func (r *Reader) NotSpace() bool
```
Space skips over any non whitespace, returning true if it advanced the cursor.

#### func (*Reader) Numeric

```go
func (r *Reader) Numeric() NumberKind
```
Numeric tries to move past the common number pattern. It returns a constant of
type NumberKind describing the kind of number it found.

#### func (*Reader) Peek

```go
func (r *Reader) Peek() rune
```
Peek returns the next rune without advancing the cursor.

#### func (*Reader) Rollback

```go
func (r *Reader) Rollback()
```
Rollback sets the cursor back to the last consume point.

#### func (*Reader) Rune

```go
func (r *Reader) Rune(value rune) bool
```
Rune advances and returns true if the next rune after the cursor matches value.

#### func (*Reader) SeekRune

```go
func (r *Reader) SeekRune(value rune) bool
```
SeekRune advances the cursor until either the value is found or the end of
stream is reached. It returns true if it found value, false otherwise.

#### func (*Reader) Space

```go
func (r *Reader) Space() bool
```
Space skips over any whitespace, returning true if it advanced the cursor.

#### func (*Reader) String

```go
func (r *Reader) String(value string) bool
```
String checks to see if value occurs at cursor, if it does, it advances the
cursor past it and returns true.

#### func (*Reader) Token

```go
func (r *Reader) Token() Token
```
Token peeks at the current scanned token value. It does not consume anything.

#### type Separator

```go
type Separator []Fragment
```

Separator is a list type to manage fragments that were skipped.

#### func (Separator) WriteTo

```go
func (s Separator) WriteTo(w io.Writer) error
```

#### type Skip

```go
type Skip func(parser *Parser, mode SkipMode) Separator
```

Skip is the function used to skip separating tokens. A separating token is one
where, as far as the parser is concerned, the tokens do not exist, even though
the tokens may have been necessary to separate the lexical tokens (whitespace),
or carry useful information (comments).

#### func  NewSkip

```go
func NewSkip(line, blockstart, blockend string) Skip
```
NewSkip builds a Skip function for the common case of a parser that has one type
of line comment, one type of block comment, and want to treat all unicode space
characters as skippable.

#### type SkipMode

```go
type SkipMode int
```


```go
const (
	// SkipPrefix is the skip mode that skips tokens that are associated with
	// the following lexically relevant token. This is mostly important for
	// comment association.
	SkipPrefix SkipMode = iota
	// SkipSuffix is the skip mode that skips tokens that are associated with
	// the preceding lexically relevant token. This is mostly important for
	// comment association.
	SkipSuffix
)
```

#### func (*SkipMode) Parse

```go
func (v *SkipMode) Parse(s string) error
```

#### func (SkipMode) String

```go
func (v SkipMode) String() string
```

#### type Token

```go
type Token struct {
	Runes []rune // The full rune array for the string this token is from.
	Start int    // The start of the token in the full rune array.
	End   int    // One past the end of the token.
}
```

A Token represents the smallest consumed unit input.

#### func (Token) Cursor

```go
func (t Token) Cursor() (line int, column int)
```
Cursor is used to calculate the line and column of the start of the token. It
may be very expensive to call, and is intended to be used sparingly in producing
human readable error messages only.

#### func (Token) Format

```go
func (t Token) Format(f fmt.State, c rune)
```
Format implements fmt.Formatter writing the start end and value of the token.

#### func (Token) Len

```go
func (t Token) Len() int
```
Len returns the length of the token in runes.

#### func (Token) String

```go
func (t Token) String() string
```
String returns the string form of the rune range the token represents.
