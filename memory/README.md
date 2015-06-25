# memory
--
    import "android.googlesource.com/platform/tools/gpu/memory"

Package memory contains types used for representing and simulating memory
observed in the capture.

## Usage

```go
var Tmp = Range{
	Base: 0x00000000ff000000,
	Size: 0x0000000000ffffff,
}
```

#### func  Reader

```go
func Reader(s Slice, d database.Database, l log.Logger) io.Reader
```
Reader returns a binary reader for the specified Slice.

#### func  Write

```go
func Write(w binary.Writer, arch device.Architecture, v interface{}) error
```
Write writes the value v to the writer w. If v is an array or slice, then each
of the elements will be written, sequentially.

#### func  Writer

```go
func Writer(p *Pool, rng Range) io.Writer
```
Writer returns a binary writer for the specified memory pool and range.

#### type Pointer

```go
type Pointer uint64
```

Pointer is the type representing a memory pointer.

#### func (Pointer) Offset

```go
func (p Pointer) Offset(n uint64) Pointer
```
Offset returns the pointer offset by n bytes.

#### func (Pointer) Range

```go
func (p Pointer) Range(s uint64) Range
```
Range returns a Range of size s with the base of this pointer.

#### func (Pointer) String

```go
func (p Pointer) String() string
```

#### type Pool

```go
type Pool struct {
	binary.Generate `disable:"true"`
}
```

Pool represents an unbounded and isolated memory space. Pool can be used to
represent the application address space, or hidden GPU Pool.

Pool can be sliced into smaller regions which can be read or written to. All
writes to Pool or its slices do not actually perform binary data copies, but
instead all writes are stored as lightweight records. Only when a Pool slice has
Get called will any resolving, loading or copying of binary data occur.

#### func (*Pool) At

```go
func (m *Pool) At(p Pointer) Slice
```
At returns an unbounded Slice starting at p.

#### func (*Pool) Slice

```go
func (m *Pool) Slice(rng Range) Slice
```
Slice returns a Slice referencing the subset of the Pool range.

#### func (*Pool) String

```go
func (m *Pool) String() string
```
String returns the full history of writes performed to this pool.

#### func (*Pool) Write

```go
func (m *Pool) Write(dst Pointer, src Slice)
```
Write copies the slice src to dst.

#### type PoolID

```go
type PoolID uint32
```

PoolID is an indentifier of a Pool.

```go
const ApplicationPool PoolID = 0
```
ApplicationPool is the PoolID of Pool representing the application's memory
address space.

#### func (*PoolID) Parse

```go
func (v *PoolID) Parse(s string) error
```

#### func (PoolID) String

```go
func (v PoolID) String() string
```

#### type Range

```go
type Range struct {
	binary.Generate
	Base Pointer // A pointer to the first byte in the memory range.
	Size uint64  // The size in bytes of the memory range.
}
```

Range represents a region of memory.

#### func (*Range) Class

```go
func (*Range) Class() binary.Class
```

#### func (Range) Contains

```go
func (i Range) Contains(p Pointer) bool
```
Contains returns true if the pointer p is within the Range.

#### func (Range) End

```go
func (i Range) End() Pointer
```
End returns a Pointer to one byte beyond the end of the Range.

#### func (Range) Expand

```go
func (i Range) Expand(p Pointer) Range
```
Expand returns a new Range that is grown to include the pointer p.

#### func (Range) First

```go
func (i Range) First() Pointer
```
First returns a Pointer to the first byte in the Range.

#### func (Range) Intersect

```go
func (i Range) Intersect(other Range) Range
```
Intersect returns the Range that is common between this Range and other. If the
two memory ranges do not intersect, then this function panics.

#### func (Range) Last

```go
func (i Range) Last() Pointer
```
Last returns a Pointer to the last byte in the Range.

#### func (Range) Overlaps

```go
func (i Range) Overlaps(other Range) bool
```
Overlaps returns true if other overlaps this memory range.

#### func (Range) Span

```go
func (i Range) Span() interval.U64Span
```
Span returns the Range as a U64Span.

#### func (Range) String

```go
func (i Range) String() string
```

#### type RangeList

```go
type RangeList []Range
```


#### func (*RangeList) Copy

```go
func (l *RangeList) Copy(to, from, count int)
```
Copy performs a copy of ranges within the RangeList.

#### func (*RangeList) GetSpan

```go
func (l *RangeList) GetSpan(index int) interval.U64Span
```
GetSpan returns the span of the range with the specified index in the RangeList.

#### func (*RangeList) Length

```go
func (l *RangeList) Length() int
```
Length returns the number of ranges in the RangeList.

#### func (*RangeList) Resize

```go
func (l *RangeList) Resize(length int)
```
Resize resizes the RangeList to the specified length.

#### func (*RangeList) SetSpan

```go
func (l *RangeList) SetSpan(index int, span interval.U64Span)
```
SetSpan adjusts the range of the span with the specified index in the RangeList.

#### type Slice

```go
type Slice interface {
	// Get resolves all the bytes representing the slice.
	Get(d database.Database, l log.Logger) ([]byte, error)

	// ResourceID returns the identifier of the resource representing the slice,
	// creating a new resource if it isn't already backed by one.
	ResourceID(d database.Database, l log.Logger) (binary.ID, error)

	// Size returns the number of bytes that would be returned by calling Get.
	Size() uint64

	// Slice returns a new Slice referencing a subset range of the data.
	// The range r is relative to the base of the Slice. For example a slice of
	// [0, 4] would return a Slice referencing the first 5 bytes of this Slice.
	// Attempting to slice outside the range of this Slice will result in a
	// panic.
	Slice(r Range) Slice
}
```

Slice is the interface for a data source that can be resolved to a byte slice
with Get, or 'sliced' to a subset of the data source.

#### func  Blob

```go
func Blob(data []byte) Slice
```
Blob returns a read-only Slice that wraps data.

#### func  Data

```go
func Data(arch device.Architecture, data ...interface{}) Slice
```
Data returns a read-only Slice that contains the encoding of data.

#### func  Resource

```go
func Resource(resId binary.ID, size uint64) Slice
```
Resource returns a Slice that wraps a resource stored in the database. resId is
the identifier of the data and size is the size in bytes of the data.
