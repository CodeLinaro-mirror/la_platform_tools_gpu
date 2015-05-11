# generate
--
    import "android.googlesource.com/platform/tools/gpu/rpc/generate"

Package generate contains the rpc code generation functionality.

## Usage

#### func  Go

```go
func Go(f *template.Functions) error
```
Go invokes the main go code generation template.

#### func  Init

```go
func Init(apiFile string, api *semantic.API) *template.Functions
```
Init prepares a new template processor that layers the apic one with the
functions from the binary codec generate package.

#### func  Java

```go
func Java(f *template.Functions) error
```
Java invokes the main go code generation template.
