# rpc
--
    import "android.googlesource.com/platform/tools/gpu/rpc"

Package rpc implements a remote procedure call system.

RPC uses the binary package for serialization and the multiplexer package to
merge in-flight requests onto a single stream.

## Usage

```go
var ErrInvalidHeader = NewError("Invalid RPC header")
```
ErrInvalidHeader is returned when either client or server detects an incorrectly
formed rpc header.

```go
var Namespace = registry.NewNamespace()
```

#### func  Serve

```go
func Serve(r io.Reader, w io.Writer, mtu int, l log.Logger, handler Handler)
```
Server implements the receiving side of a client server rpc pair. It listens on
the reader for calls, and dispatches them to the supplied handler. Any result
returned from the handler is then sent back down the writer.

#### type Client

```go
type Client struct {
}
```

Client implements the sending side of a client-server rpc pair.

#### func  NewClient

```go
func NewClient(m *multiplexer.Multiplexer, n *registry.Namespace) Client
```
NewClient creates a new rpc client object that uses the multiplexer m for
communication the namespace n for decoding objects. If n is nil then the global
namespace is used.

#### func (Client) Multiplexer

```go
func (c Client) Multiplexer() *multiplexer.Multiplexer
```
Multiplexer returns the multiplexer used for communication to the server.

#### func (Client) Namespace

```go
func (c Client) Namespace() *registry.Namespace
```
Namespace returns the custom namespace used for decoding responses from the
server, or nil if no custom namespace has been specified.

#### func (Client) Send

```go
func (c Client) Send(call binary.Object) (interface{}, error)
```
Send encodes an rpc call and sends it to the server. It blocks until a reply is
received or an error indicating there will be no reply occurs. This method is
safe for concurrent use.

#### type Error

```go
type Error struct {
	binary.Generate
}
```

Error is an implementation of error that can be sent over the wire.

#### func  NewError

```go
func NewError(msg string, args ...interface{}) *Error
```
NewError is used to create new rpc error objects with the specified human
readable message.

#### func (*Error) Class

```go
func (*Error) Class() binary.Class
```

#### func (*Error) Error

```go
func (e *Error) Error() string
```

#### type Handler

```go
type Handler func(interface{}) binary.Object
```

Handler is the signature for a function that handles incoming rpc calls.
