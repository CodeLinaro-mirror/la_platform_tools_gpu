# test
--
    import "android.googlesource.com/platform/tools/gpu/rpc/test"


## Usage

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

#### type Derived

```go
type Derived struct {
	binary.Generate
	Name string
	Enum Enum
}
```

Class Derived

#### func  CreateDerived

```go
func CreateDerived(
	Name string,
	Enum Enum,
) *Derived
```

#### func (*Derived) Class

```go
func (*Derived) Class() binary.Class
```

#### func (*Derived) GetEnum

```go
func (c *Derived) GetEnum() Enum
```

#### func (*Derived) GetName

```go
func (c *Derived) GetName() string
```

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

#### func (Enum) String

```go
func (i Enum) String() string
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

#### type ListNodeArray

```go
type ListNodeArray []*ListNode
```

Array ListNodeRefArray

#### func (ListNodeArray) Format

```go
func (a ListNodeArray) Format(f fmt.State, c rune)
```

#### type RPC

```go
type RPC interface {
	Add(l log.Logger, a uint32, b uint32) (uint32, error)
	EnumToString(l log.Logger, e Enum) (string, error)
	GetStruct(l log.Logger) (Struct, error)
	SetStruct(l log.Logger, s Struct) error
	GetResource(l log.Logger) (ResourceId, error)
	UseResource(l log.Logger, r ResourceId) error
	ResolveResource(l log.Logger, r ResourceId) (Resource, error)
	GetSingleListNode(l log.Logger) (*ListNode, error)
	GetListNodeChain(l log.Logger) (*ListNode, error)
	GetListNodeChainArray(l log.Logger) (ListNodeArray, error)
	GetBase(l log.Logger) (Base, error)
	GetDerived(l log.Logger) (Base, error)
}
```


#### func  CreateClient

```go
func CreateClient(r io.Reader, w io.Writer, mtu int) RPC
```

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
