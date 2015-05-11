# interval
--
    import "android.googlesource.com/platform/tools/gpu/interval"

Package interval implements algorithms that operate on lists of intervals.

## Usage

#### func  Contains

```go
func Contains(l List, value uint64) bool
```
Contains returns true if the value is found inside on of the intervals.

#### func  IndexOf

```go
func IndexOf(l List, value uint64) int
```
IndexOf returns the index of the span the value is a part of, or -1 if not found

#### func  Intersect

```go
func Intersect(l List, span U64Span) (first, count int)
```
Intersect finds the intervals from the list that overlap with the specified
span.

#### func  Merge

```go
func Merge(l List, span U64Span, joinAdj bool) int
```
Merge adds a span to the list, merging it with existing spans if it overlaps
them, and returns the index of that span. If the joinAdj parameter is true, then
any intervals that are immediately adjacent to span will be merged with span.
For example, consider the merging of intervals [0, 2] and [3, 5]:

When joinAdj == false:

    ╭       ╮       ╭       ╮   ╭       ╮╭       ╮
    │0  1  2│ merge │3  4  5│ = │0  1  2││3  4  5│
    ╰       ╯       ╰       ╯   ╰       ╯╰       ╯

When join == true:

    ╭       ╮       ╭       ╮   ╭                ╮
    │0  1  2│ merge │3  4  5│ = │0  1  2  3  4  5│
    ╰       ╯       ╰       ╯   ╰                ╯

#### func  Remove

```go
func Remove(l List, span U64Span)
```
Remove strips the specified span from the list, cutting it out of any
overlapping intervals

#### func  Replace

```go
func Replace(l List, span U64Span) int
```
Replace cuts the span out of any existing intervals, and then adds a new
interval, and returns it's index.

#### func  Search

```go
func Search(l List, t Predicate) int
```
Search finds the first interval in the list that the supplied predicate returns
true for. If no interval matches the predicate, it returns the length of the
list.

#### type List

```go
type List interface {
	// Length returns the number of elements in the list
	Length() int
	// GetSpan returns the span for the element at index in the list
	GetSpan(index int) U64Span
	// SetSpan sets the span for the element at index in the list
	SetSpan(index int, span U64Span)
	// Copy count list entries
	Copy(to, from, count int)
	// Resize adjusts the length of the array
	Resize(length int)
}
```

List is the interface to an object that can be used as an interval list by the
algorithms in this package.

#### type Predicate

```go
type Predicate func(test U64Span) bool
```

Predicate is used as the condition for a Search

#### type U64Range

```go
type U64Range struct {
	First uint64 // the first value in the interval
	Count uint64 // the count of values in the interval
}
```

U64Range is an interval specified by a beginning and size.

#### func (U64Range) Span

```go
func (r U64Range) Span() U64Span
```
Span converts a U64Range to a U64Span

#### type U64RangeList

```go
type U64RangeList []U64Range
```

U64RangeList implements List for an array of U64Range intervals

#### func (U64RangeList) Copy

```go
func (l U64RangeList) Copy(to, from, count int)
```

#### func (U64RangeList) GetSpan

```go
func (l U64RangeList) GetSpan(index int) U64Span
```

#### func (U64RangeList) Length

```go
func (l U64RangeList) Length() int
```

#### func (*U64RangeList) Resize

```go
func (l *U64RangeList) Resize(length int)
```

#### func (U64RangeList) SetSpan

```go
func (l U64RangeList) SetSpan(index int, span U64Span)
```

#### type U64Span

```go
type U64Span struct {
	Start uint64 // the value at which the interval begins
	End   uint64 // the next value not included in the interval.
}
```

U64Span is the base interval type understood by the algorithms in this package.
It is a half open interval that includes the lower bound, but not the upper.

#### func (U64Span) Range

```go
func (s U64Span) Range() U64Range
```
Range converts a U64Span to a U64Range

#### type U64SpanList

```go
type U64SpanList []U64Span
```

U64SpanList implements List for an array of U64Span intervals

#### func (U64SpanList) Copy

```go
func (l U64SpanList) Copy(to, from, count int)
```

#### func (U64SpanList) GetSpan

```go
func (l U64SpanList) GetSpan(index int) U64Span
```

#### func (U64SpanList) Length

```go
func (l U64SpanList) Length() int
```

#### func (*U64SpanList) Resize

```go
func (l *U64SpanList) Resize(length int)
```

#### func (U64SpanList) SetSpan

```go
func (l U64SpanList) SetSpan(index int, span U64Span)
```
