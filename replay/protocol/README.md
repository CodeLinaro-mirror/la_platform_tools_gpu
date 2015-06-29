# protocol
--
    import "android.googlesource.com/platform/tools/gpu/replay/protocol"

Package protocol contains the constants and types used to communicate with the
replay system and its virtual-machine interpreter.

Any changes to the values in this package must also be made to the replay
source.

## Usage

```go
const (
	// ConnectionTypeDeviceInfo is the type of connection used to query device information.
	ConnectionTypeDeviceInfo = ConnectionType(0)
	// ConnectionTypeReplay is the type of connection used to issue a replay.
	ConnectionTypeReplay = ConnectionType(1)
	// ConnectionTypeShutdown is used to request a shutdown of replayd.
	ConnectionTypeShutdown = ConnectionType(2)
)
```

```go
const (
	// MessageTypeGet is sent for a packet that requests a resource.
	MessageTypeGet = MessageType(0)
	// MessageTypePost is sent for a packet containing postback data.
	MessageTypePost = MessageType(1)
)
```

```go
const (
	OpCall     = 0
	OpPushI    = 1
	OpLoadC    = 2
	OpLoadV    = 3
	OpLoad     = 4
	OpPop      = 5
	OpStoreV   = 6
	OpStore    = 7
	OpResource = 8
	OpPost     = 9
	OpCopy     = 10
	OpClone    = 11
	OpStrcpy   = 12
	OpExtend   = 13
	OpLabel    = 14
)
```

```go
var Namespace = registry.NewNamespace()
```

#### type ConnectionType

```go
type ConnectionType uint8
```

ConnectionType is sent from the server to the replay system to define the type
of connection.

#### func (*ConnectionType) Parse

```go
func (v *ConnectionType) Parse(s string) error
```

#### func (ConnectionType) String

```go
func (v ConnectionType) String() string
```

#### type MessageType

```go
type MessageType uint8
```

MessageType defines the packet type sent from the replay system to the server.

#### func (*MessageType) Parse

```go
func (v *MessageType) Parse(s string) error
```

#### func (MessageType) String

```go
func (v MessageType) String() string
```

#### type Opcode

```go
type Opcode int
```

Opcode is one of the opcodes supported by the replay virtual machine.

#### func (Opcode) String

```go
func (t Opcode) String() string
```
String returns the human-readable name of the opcode.

#### type Payload

```go
type Payload struct {
	binary.Generate
	StackSize          uint32         // Maximum number of values.
	VolatileMemorySize uint32         // In bytes.
	Constants          []byte         // The constant buffer.
	Resources          []ResourceInfo // Resources used by this replay payload.
	Opcodes            []byte         // The encoded list of opcodes.
}
```

Payload contains all the information to perform a replay. The encoded form is
what is passed to the replay system.

#### func (*Payload) Class

```go
func (*Payload) Class() binary.Class
```

#### type ResourceInfo

```go
type ResourceInfo struct {
	binary.Generate
	ID   string // The resource identifier as a string.
	Size uint32 // The size in bytes of the resource.
}
```

ResourceInfo describes a resource used by a Payload.

#### func (*ResourceInfo) Class

```go
func (*ResourceInfo) Class() binary.Class
```

#### type Type

```go
type Type uint32
```

Type is one of the primitive types supported by the replay virtual machine.

```go
const (
	TypeBool            Type = 0  // A boolean type.
	TypeInt8            Type = 1  // A signed 8-bit integer type.
	TypeInt16           Type = 2  // A signed 16-bit integer type.
	TypeInt32           Type = 3  // A signed 32-bit integer type.
	TypeInt64           Type = 4  // A signed 64-bit integer type.
	TypeUint8           Type = 5  // An unsigned 8-bit integer type.
	TypeUint16          Type = 6  // An unsigned 16-bit integer type.
	TypeUint32          Type = 7  // An unsigned 32-bit integer type.
	TypeUint64          Type = 8  // An unsigned 64-bit integer type.
	TypeFloat           Type = 9  // A 32-bit floating-point number type.
	TypeDouble          Type = 10 // A 64-bit floating-point number type.
	TypeAbsolutePointer Type = 11 // A pointer type that is not remapped by the protocol.
	TypeConstantPointer Type = 12 // A pointer into the constant buffer space.
	TypeVolatilePointer Type = 13 // A pointer into the volatile buffer space.

	TypeVoid Type = 0xffffffff // A non-existant type. Not handled by the protocol.
)
```

#### func (*Type) Parse

```go
func (v *Type) Parse(s string) error
```

#### func (Type) Size

```go
func (t Type) Size(pointerSize int) int
```
Size returns the size in bytes of the type. pointerSize is the size in bytes of
a pointer for the target architecture.

#### func (Type) String

```go
func (v Type) String() string
```
