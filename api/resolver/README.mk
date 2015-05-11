# resolver
--
    import "android.googlesource.com/platform/tools/gpu/api/resolver"

Package resolver implements a semantic resolving for the api language. It is
responsible for converting from an abstract syntax tree to a typed semantic
graph ready for code generation.

## Usage

#### func  Resolve

```go
func Resolve(compiled *ast.API) (*semantic.API, parse.ErrorList, ASTToSemantic)
```
Resolve takes a valid ast as produced by the parser and converts it to the
semantic graph form. If the ast is not fully valid (ie there were parse errors)
then the results are undefined, and may include null pointer access. If there
are semantic problems with the ast, Resolve will return the set of errors it
finds, and the returned graph may be incomplete/invalid.

#### type ASTToSemantic

```go
type ASTToSemantic map[ast.Node]semantic.Node
```

ASTToSemantic is a relational map of AST nodes to semantic nodes.

#### type Alias

```go
type Alias struct {
	AST  *ast.Alias
	Name string
	To   semantic.Type
}
```

Alias is used as a temporary type holder during type resolution. It is not
present in the final semantic tree returned, but may be present in the AST ->
semantic map.

#### func (Alias) Member

```go
func (t Alias) Member(name string) semantic.Node
```

#### func (Alias) Typename

```go
func (t Alias) Typename() string
```
