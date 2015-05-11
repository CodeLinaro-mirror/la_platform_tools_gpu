# memory
--
    import "android.googlesource.com/platform/tools/gpu/memory"

Package memory contains types used for representing and simulating memory
observed in the capture.

## Usage

#### type Data

```go
type Data interface {
	Get(db database.Database, logger log.Logger) ([]byte, error)
	Size() uint64
}
```

Data is the interface for a data source that can be resolved to a byte slice
with Get.

Size returns the number of bytes that would be returned by calling Get.

#### type DataSliceWriter

```go
type DataSliceWriter interface {
	Data
	Slice(r Range) DataSliceWriter
	Write(d Data)
}
```

DataSliceWriter is similar to DataSlicer, but also has the Write method for
modifying the data that the DataSliceWriter refers to.

Write will replace the data in this DataSliceWriter with d. If d is shorter in
length than the slice, then only the range [0, d.Size()-1] bytes will be
replaced. A sliced DataSliceWriter shares the same data from which it was sliced
- i.e. any mutations to a DataSliceWriter will also affect the Data it was
sliced from.

#### type DataSlicer

```go
type DataSlicer interface {
	Data
	Slice(r Range) DataSlicer
}
```

DataSlicer extends the Data interface with aditional support for slicing.

Slice returns a new DataSlicer referencing a subset range of the data. The range
r is relative to the base of the DataSlicer. For example a slice of [0, 4] would
return a DataSlicer referencing the first 5 bytes of this DataSlicer. Attempting
to slice outside the range of this DataSlicer will result in a panic.

#### func  ResourceData

```go
func ResourceData(resId binary.ID, size uint64) DataSlicer
```
ResourceData returns a DataSlicer that wraps a resource. resId is the identifier
of the resource and size is the size in bytes of the resource.

#### type Memory

```go
type Memory struct {
	binary.Generate `disable:"true"`
}
```

Memory represents an unbounded and isolated memory space. Memory can be used to
represent the application address space, or hidden GPU memory.

Memory can be sliced into smaller regions which can be read or written to. All
writes to Memory or its slices do not actually perform binary data copies, but
instead all writes are stored as lightweight records. Only when a Memory slice
has Get called will any resolving, loading or copying of binary data occur.

#### func (*Memory) Slice

```go
func (m *Memory) Slice(rng Range) DataSliceWriter
```
Slice returns a DataSliceWriter referencing the subset of the Memory range.

#### func (*Memory) Write

```go
func (m *Memory) Write(d Data)
```
Write copies d to the Memory slice [0, d.Size()-1].

#### type Pointer

```go
type Pointer uint64
```

Pointer is the type representing a memory pointer.

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
