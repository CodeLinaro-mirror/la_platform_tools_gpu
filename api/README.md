# api
--
    import "android.googlesource.com/platform/tools/gpu/api"

Package api holds the main interface to the api language libraries. It provides
functions for going from api files to abstract syntax trees and processed
semantic trees.

## Usage

```go
var DefaultProcessor = Processor{
	Parsed:   map[string]*ast.API{},
	Resolved: map[string]*semantic.API{},
}
```
DefaultProcessor is the Processor used in the package level functions. Most
applications will not need multiple instances of a Processor, and can just use
this one.

#### func  Parse

```go
func Parse(apiname string) (*ast.API, parse.ErrorList)
```
Parse parses the api file with the DefaultProcessor. See Processor.Parse for
details.

#### func  Resolve

```go
func Resolve(apiname string, mappings resolver.ASTToSemantic) (*semantic.API, parse.ErrorList)
```
Resolve resolves the api file with the DefaultProcessor. See Processor.Resolve
for details.

#### type Processor

```go
type Processor struct {
	Parsed   map[string]*ast.API
	Resolved map[string]*semantic.API
}
```

Processor holds the state when resolving multiple api files.

#### func (*Processor) Parse

```go
func (p *Processor) Parse(path string) (*ast.API, parse.ErrorList)
```
Parse returns an ast that represents the supplied filename. It if the file has
already been parsed, the cached ast will be returned, otherwise it invokes
parser.Parse on the content of the supplied file name.

#### func (*Processor) Resolve

```go
func (p *Processor) Resolve(apiname string, mappings resolver.ASTToSemantic) (*semantic.API, parse.ErrorList)
```
Resolve returns a semantic.API that represents the supplied api file name. If
the file has already been resolved, the cached semantic tree is returned,
otherwise the file and all dependant files are parsed using Processor.Parse.
Recursive calls are made to Resolve for all named imports, and then finally the
ast and all included ast's are handed to resolver.Resolve to do semantic
processing.
