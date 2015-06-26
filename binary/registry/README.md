# registry
--
    import "android.googlesource.com/platform/tools/gpu/binary/registry"


## Usage

```go
var (
	// Global is the default global Namespace object.
	Global = NewNamespace()
)
```

#### type Namespace

```go
type Namespace struct {
}
```

Namespace represents a mapping of type identifiers to their Class.

#### func  NewNamespace

```go
func NewNamespace(fallbacks ...*Namespace) *Namespace
```
NewNamespace creates a new namespace layered on top of the specified fallback.

#### func (*Namespace) Add

```go
func (n *Namespace) Add(class binary.Class)
```
Add a new type to the Namespace.

#### func (*Namespace) AddFallbacks

```go
func (n *Namespace) AddFallbacks(fallbacks ...*Namespace)
```
AddFallbacks appends new Namespaces to the fallback list of this Namespace.

#### func (*Namespace) Count

```go
func (n *Namespace) Count() int
```
Count returns the number of entries reachable through this namespace. Because it
sums the counts of the namespaces it depends on, this may be more than the
number of unique keys.

#### func (*Namespace) Lookup

```go
func (n *Namespace) Lookup(id binary.ID) binary.Class
```
Lookup looks up a Class by the given type id in the Namespace. If there is no
match, it will return nil.

#### func (Namespace) Visit

```go
func (n Namespace) Visit(visitor func(binary.Class))
```
Visit invokes the visitor for every class object reachable through this
namespace. The visitor maybe be called with the same id more than once if it is
present in multiple namespaces.
