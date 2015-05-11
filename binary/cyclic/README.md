# cyclic
--
    import "android.googlesource.com/platform/tools/gpu/binary/cyclic"

Object encoding details

Objects are encoded in three different forms, depending on whether the object
was nil, or was already encoded to the stream. If the object is non-nil and is
encoded for the first time for a given Encoder then the object is encoded as:

    key  uint32   // A unique identifier for the object instance
    type [20]byte // A unique identifier for the type of the object
    ...data...    // The object's data (length dependent on the object type)

All subsequent encodings of the object are encoded as the 16 bit key identifier
without any additional data:

    key uint32 // An identifier of a previously encoded object instance

If the object is nil, then the object is encoded as a uint32 of 0.

## Usage

#### func  Decoder

```go
func Decoder(reader binary.Reader) *decoder
```
Decoder creates a binary.Decoder that reads from the provided binary.Reader.

#### func  Encoder

```go
func Encoder(writer binary.Writer) binary.Encoder
```
Encoder creates a binary.Encoder that writes to the supplied binary.Writer.
