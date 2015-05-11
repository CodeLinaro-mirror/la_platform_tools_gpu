# flat
--
    import "android.googlesource.com/platform/tools/gpu/binary/flat"

Objects are encoded as:

    type [20]byte // A unique identifier for the type of the object
    ...data...    // The object's data (length dependent on the object type)

If the object is nil, then the object is encoded as a null type.

## Usage

#### func  Decoder

```go
func Decoder(reader binary.Reader) binary.Decoder
```
Decoder creates a binary.Decoder that reads from the provided binary.Reader.

#### func  Encoder

```go
func Encoder(writer binary.Writer) binary.Encoder
```
Encoder creates a binary.Encoder that writes to the supplied binary.Writer.
