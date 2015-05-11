# multiplexer
--
    import "android.googlesource.com/platform/tools/gpu/multiplexer"

Package multiplexer provides multiple data-stream multiplexing over a single
binary data-stream.

## Usage

```go
var ErrChannelClosed = errors.New("Channel closed")
```

#### type Multiplexer

```go
type Multiplexer struct {
}
```

Multiplexer provides multiple data-stream multiplexing over a single binary
data-stream. Channels can be opened and closed by either endpoint, and data can
be sent on an open channel in either direction. The muliplexer attempts to
evenly distribute bandwidth between channels.

#### func  New

```go
func New(in io.Reader, out io.Writer, mtu int, channelOpenedCallback func(io.ReadWriteCloser)) *Multiplexer
```
New creates and returns a new Multiplexer using the specified reader and writer
for communication. mtu defines the maximum size of each packet of data.
channelOpenedCallback will be called for each channel that was opened by the
remote endpoint.

#### func (*Multiplexer) MTU

```go
func (m *Multiplexer) MTU() int
```
MTU returns the maximum transmission unit size the multiplexer was created with.

#### func (*Multiplexer) OpenChannel

```go
func (m *Multiplexer) OpenChannel() (io.ReadWriteCloser, error)
```
OpenChannel opens a new channel for communication. This will invoke the the
channel-opened callback on the remote endpoint.
