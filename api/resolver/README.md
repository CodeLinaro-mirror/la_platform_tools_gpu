# resolver
--
    import "android.googlesource.com/platform/tools/gpu/api/resolver"

Package resolver implements a semantic resolving for the api language. It is
responsible for converting from an abstract syntax tree to a typed semantic
graph ready for code generation.

## Usage

```go
const (
	RefSuffix     = "ʳ"
	SliceSuffix   = "ˢ"
	ConstSuffix   = "ᶜ"
	PointerSuffix = "ᵖ"
	ArraySuffix   = "ᵃ"
	MapSuffix     = "ᵐ"
	TypeInfix     = "ː"
)
```

#### func  Resolve

```go
func Resolve(includes []*ast.API, symbols *semantic.Symbols, mappings ASTToSemantic) (*semantic.API, parse.ErrorList)
```
Resolve takes valid asts as produced by the parser and converts them to the
semantic graph form. If the asts are not fully valid (ie there were parse
errors) then the results are undefined. If there are semantic problems with the
ast, Resolve will return the set of errors it finds, and the returned graph may
be incomplete/invalid.

#### type ASTToSemantic

```go
type ASTToSemantic map[ast.Node]semantic.Node
```

ASTToSemantic is a relational map of AST nodes to semantic nodes.
