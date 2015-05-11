# endian
--
    import "android.googlesource.com/platform/tools/gpu/binary/endian"

Package endian implements binary.Reader and binary.Writer for writing .

Boolean values are encoded as single bytes, where 0 represents false and non-
zero represents true.

Numeric types are all encoded as the simple native representation, but no
attempt is made to align them.

Strings are encoded in C style null terminated form.

## Usage

```go
var (
	Little = ByteOrder(eb.LittleEndian)
	Big    = ByteOrder(eb.BigEndian)
)
```

#### func  Reader

```go
func Reader(r io.Reader, byteOrder ByteOrder) binary.Reader
```
Reader creates a binary.Reader that reads from the provided io.Reader, with the
specified byte order.

#### func  Writer

```go
func Writer(w io.Writer, byteOrder ByteOrder) binary.Writer
```
Writer creates a binary.Writer that writes to the supplied stream, with the
specified byte order.

#### type ByteOrder

```go
type ByteOrder eb.ByteOrder
```
