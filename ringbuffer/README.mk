# ringbuffer
--
    import "android.googlesource.com/platform/tools/gpu/ringbuffer"

Package ringbuffer implements an in-memory circular buffer conforming to the
io.ReadWriteCloser interface.

## Usage

#### func  New

```go
func New(capacity int) io.ReadWriteCloser
```
New constructs a new ring-buffer with the specified capacity in bytes.

Writes to the ring-buffer will block until all the bytes are written into the
buffer or the buffer is closed (whichever comes first.) Reads from the
ring-buffer will block until a single byte is read or the buffer is closed
(whichever comes first.) It is safe to call Read and Write in parallel with each
other or with Close. If the ring-buffer is closed while a read or write is in
progress then io.ErrClosedPipe will be returned by the read / write function.
