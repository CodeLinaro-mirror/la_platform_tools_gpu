# test
--
    import "android.googlesource.com/platform/tools/gpu/rpc/test"


## Usage

```go
var ConstantValues schema.Constants
```

```go
var Namespace = registry.NewNamespace()
```

#### func  BindServer

```go
func BindServer(r io.Reader, w io.Writer, mtu int, l log.Logger, server RPC)
```

#### type Base

```go
type Base interface {
	binary.Object
	GetName() string
}
```

Interface Base

#### type Client

```go
type Client interface {
	// Client exposes all the RPC interface methods.
	RPC
	// Multiplexer returns the multiplexer used for communication to the server.
	Multiplexer() *multiplexer.Multiplexer
	// Namespace returns the custom namespace used for decoding responses from the
	// server, or nil if no custom namespace has been specified.
	Namespace() *registry.Namespace
}
```

Client is the client interface for RPC calls.

#### func  NewClient

```go
func NewClient(m *multiplexer.Multiplexer, n *registry.Namespace) Client
```
NewClient creates a new rpc client object that uses the multiplexer m for
communication the namespace n for decoding objects. If n is nil then the global
namespace is used.

#### type Enum

```go
type Enum int
```

Enum Enum

```go
const (
	EnumOne   Enum = 1
	EnumTwo   Enum = 2
	EnumThree Enum = 3
)
```

#### func (Enum) IsOne

```go
func (i Enum) IsOne() bool
```

#### func (Enum) IsThree

```go
func (i Enum) IsThree() bool
```

#### func (Enum) IsTwo

```go
func (i Enum) IsTwo() bool
```

#### func (*Enum) Parse

```go
func (v *Enum) Parse(s string) error
```

#### func (Enum) String

```go
func (v Enum) String() string
```

#### type ListNode

```go
type ListNode struct {
	binary.Generate
	Name string
	Next *ListNode
}
```

Class ListNode

#### func  CreateListNode

```go
func CreateListNode(
	Name string,
	Next *ListNode,
) *ListNode
```

#### func (*ListNode) Class

```go
func (*ListNode) Class() binary.Class
```

#### func (*ListNode) GetName

```go
func (c *ListNode) GetName() string
```

#### func (*ListNode) GetNext

```go
func (c *ListNode) GetNext() *ListNode
```

#### type RPC

```go
type RPC interface {
	Add(a uint32, b uint32, l log.Logger) (uint32, error)
	EnumToString(e Enum, l log.Logger) (string, error)
	GetStruct(l log.Logger) (Struct, error)
	SetStruct(s Struct, l log.Logger) error
	GetResource(l log.Logger) (ResourceId, error)
	UseResource(r ResourceId, l log.Logger) error
	ResolveResource(r ResourceId, l log.Logger) (Resource, error)
	GetSingleListNode(l log.Logger) (*ListNode, error)
	GetListNodeChain(l log.Logger) (*ListNode, error)
	GetListNodeChainArray(l log.Logger) ([]*ListNode, error)
}
```


#### type Resolver

```go
type Resolver struct {
	Database database.Database
}
```


#### func (Resolver) ResolveResource

```go
func (r Resolver) ResolveResource(id ResourceId, l log.Logger) (Resource, error)
```
ResolveResource loads and returns the Resource stored in the resolver's
database, using id.

#### type Resource

```go
type Resource struct {
	binary.Generate
	Int    uint32
	Float  float32
	String string
}
```

Class Resource

#### func  CreateResource

```go
func CreateResource(
	Int uint32,
	Float float32,
	String string,
) *Resource
```

#### func  ResolveResource

```go
func ResolveResource(id ResourceId, d database.Database, l log.Logger) (res Resource, err error)
```
ResolveResource loads and returns the Resource stored in the database d, using
id.

#### func (*Resource) Class

```go
func (*Resource) Class() binary.Class
```

#### func (*Resource) GetFloat

```go
func (c *Resource) GetFloat() float32
```

#### func (*Resource) GetInt

```go
func (c *Resource) GetInt() uint32
```

#### func (*Resource) GetString

```go
func (c *Resource) GetString() string
```

#### type ResourceId

```go
type ResourceId struct {
	binary.Generate
	ID binary.ID
}
```

Handle ResourceId

#### func  StoreResource

```go
func StoreResource(v *Resource, d database.Database, l log.Logger) (ResourceId, error)
```
StoreResource stores v into the database d, returning the ResourceId.

#### func (*ResourceId) Class

```go
func (*ResourceId) Class() binary.Class
```

#### func (ResourceId) Valid

```go
func (h ResourceId) Valid() bool
```

#### type Struct

```go
type Struct struct {
	binary.Generate
	String string
	U32    uint32
	Enum   Enum
}
```

Class Struct

#### func  CreateStruct

```go
func CreateStruct(
	String string,
	U32 uint32,
	Enum Enum,
) *Struct
```

#### func (*Struct) Class

```go
func (*Struct) Class() binary.Class
```

#### func (*Struct) GetEnum

```go
func (c *Struct) GetEnum() Enum
```

#### func (*Struct) GetString

```go
func (c *Struct) GetString() string
```

#### func (*Struct) GetU32

```go
func (c *Struct) GetU32() uint32
```
