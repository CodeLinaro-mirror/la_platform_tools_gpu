# parser
--
    import "android.googlesource.com/platform/tools/gpu/api/parser"

Package parser implements a parser for converting the api language into abstract
syntax trees.

## Usage

#### func  Parse

```go
func Parse(data string) (*ast.API, parse.ErrorList)
```
Parse takes a string containing a complete api description and returns the
abstract syntax tree representation of it. If the string is not syntactically
valid, it will also return the errors encountered. If errors are returned, the
ast returned will be the incomplete tree so far, and may not be structurally
valid.
