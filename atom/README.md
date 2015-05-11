# atom
--
    import "android.googlesource.com/platform/tools/gpu/atom"

Package atom provides the fundamental types used to describe a capture stream.

## Usage

```go
const NoID = ^ID(0)
```
NoID is used when you have to pass an ID, but don't have one to use.

#### func  Register

```go
func Register(ty TypeInfo)
```
Register registers the atom type ty with the atom registry. If another atom is
already registered with the same type identifer then Register will panic.

#### type Atom

```go
type Atom interface {
	binary.Object

	// API returns the graphics API this atom belongs to.
	API() gfxapi.API

	// TypeID returns the identifier of this atom's type.
	TypeID() TypeID

	// Flags returns the flags of the atom.
	Flags() Flags

	// Mutate mutates the State using the atom.
	Mutate(*gfxapi.State) error
}
```

Atom is the interface implemented by all objects that describe an single event
in a capture stream. Typical implementations of Atom describe an application's
call to a graphics API function or provide meta-data describing observed memory
or state at the time of capture.

Each implementation of Atom should have a unique and stable TypeID to ensure
binary compatibility with old capture formats. Any change to the Atom's binary
format should also result in a new TypeID.

#### func  New

```go
func New(id TypeID) (Atom, error)
```
New builds a new instance of the atom with type identifier id. The type must
have previously been registered with Register.

#### type Flags

```go
type Flags uint32
```

Flags is a bitfield describing characteristics of an atom.

```go
const (
	DrawCall Flags = 1 << iota
	EndOfFrame
)
```

#### func (Flags) IsDrawCall

```go
func (f Flags) IsDrawCall() bool
```
IsDrawCall returns true if the atom is a draw call.

#### func (Flags) IsEndOfFrame

```go
func (f Flags) IsEndOfFrame() bool
```
IsEndOfFrame returns true if the atom represents the end of a frame.

#### type Group

```go
type Group struct {
	binary.Generate
	Name      string    // Name of this group.
	Range     Range     // The range of atoms this group (and sub-groups) represents.
	SubGroups GroupList // All sub-groups of this group.
}
```

Group represents a named, contiguous span of atoms with support for sparse
sub-groups. Groups are ideal for expressing nested hierarchies of atoms.

Groups have the concept of items. An item is either an immediate sub-group, or
an atom identifier that is within this group's span but outside of any
sub-group.

For example a Group spanning the atom identifier range [0 - 9] with two
sub-groups spanning [2 - 4] and [7 - 8] would have the following tree of items:

    Group
      │
      ├─── Item[0] ─── Atom[0]
      │
      ├─── Item[1] ─── Atom[1]
      │
      ├─── Item[2] ─── Sub-group 0
      │                   │
      │                   ├─── Item[0] ─── Atom[2]
      │                   │
      │                   ├─── Item[1] ─── Atom[3]
      │                   │
      │                   └─── Item[2] ─── Atom[4]
      │
      ├─── Item[3] ─── Atom[5]
      │
      ├─── Item[4] ─── Atom[6]
      │
      ├─── Item[5] ─── Sub-group 1
      │                   │
      │                   ├─── Item[0] ─── Atom[7]
      │                   │
      │                   └─── Item[1] ─── Atom[8]
      │
      └─── Item[6] ─── Atom[9]

#### func (*Group) Class

```go
func (*Group) Class() binary.Class
```

#### func (Group) Count

```go
func (g Group) Count() uint64
```
Count returns the number of immediate items this group contains.

#### func (Group) Index

```go
func (g Group) Index(index uint64) (baseAtomID ID, subgroup *Group)
```
Index returns the item with the specified index. If the item refers directly to
an atom identifier then the atom identifier is returned in baseAtomID and
subgroup is assigned nil. If the item is a sub-group then baseAtomID is returned
as the lowest atom identifier found in the sub-group and subgroup is assigned
the sub-group pointer.

#### func (Group) IndexOf

```go
func (g Group) IndexOf(atomID ID) uint64
```
IndexOf returns the item index that refers directly to, or contains the given
atom identifer.

#### func (*Group) Insert

```go
func (g *Group) Insert(atomID ID, count int)
```
Insert adjusts the spans of this group and all subgroups for an insertion of
count elements at atomID.

#### func (Group) String

```go
func (g Group) String() string
```
String returns a string representing the group's name, range and sub-groups.

#### type GroupList

```go
type GroupList []Group
```

GroupList is a list of Groups. Functions in this package expect the list to be
in ascending atom identifier order, and maintain that order on mutation.

#### func (*GroupList) Add

```go
func (l *GroupList) Add(start, end ID, name string)
```
Add inserts a new atom group into the list with the specified range and name. If
the new group does not overlap any existing groups in the list then it is
inserted into the list, keeping ascending atom-identifier order. If the new
group sits completely within an existing group then this new group will be added
to the existing group's sub-groups. If the new group completely wraps one or
more existing groups in the list then these existing groups are added as
sub-groups to the new group and then the new group is added to the list, keeping
ascending atom-identifier order. If the new group partially overlaps any
existing group then the function will panic.

#### func (GroupList) Copy

```go
func (l GroupList) Copy(to, from, count int)
```
Copy copies count groups within the list.

#### func (GroupList) GetSpan

```go
func (l GroupList) GetSpan(index int) interval.U64Span
```
GetSpan returns the atom identifier span for the group at index in the list.

#### func (*GroupList) IndexOf

```go
func (l *GroupList) IndexOf(atomID ID) int
```
IndexOf returns the index of the group that contains the atom identifier or -1
if not found.

#### func (GroupList) Length

```go
func (l GroupList) Length() int
```
Length returns the number of groups in the list.

#### func (*GroupList) Resize

```go
func (l *GroupList) Resize(length int)
```
Resize adjusts the length of the list.

#### func (GroupList) SetSpan

```go
func (l GroupList) SetSpan(index int, span interval.U64Span)
```
SetSpan sets the atom identifier span for the group at index in the list.

#### func (GroupList) String

```go
func (l GroupList) String() string
```
String returns a string representing all groups in the group list.

#### type ID

```go
type ID uint64
```

ID is the index of an atom in an atom stream.

#### type IDSet

```go
type IDSet map[ID]struct{}
```

IDSet is a set of IDs.

#### func (*IDSet) Add

```go
func (s *IDSet) Add(id ID)
```
Add adds id to the set. If the id was already in the set then the call does
nothing.

#### func (IDSet) Contains

```go
func (s IDSet) Contains(id ID) bool
```
Contains returns true if id is in the set, otherwise false.

#### func (*IDSet) Remove

```go
func (s *IDSet) Remove(id ID)
```
Remove removes id from the set. If the id was not in the set then the call does
nothing.

#### type List

```go
type List []Atom
```

List is a list of atoms.

#### func (*List) Add

```go
func (l *List) Add(a Atom)
```
Add adds a to the end of the atom list.

#### func (*List) AddAt

```go
func (l *List) AddAt(a Atom, id ID)
```
Add adds a to the list before the atom at id.

#### func (*List) Clone

```go
func (l *List) Clone() List
```
Clone makes and returns a shallow copy of the atom list.

#### func (*List) Decode

```go
func (l *List) Decode(d binary.Decoder) error
```
Encode decodes the atom list using the specified encoder.

#### func (*List) Encode

```go
func (l *List) Encode(e binary.Encoder) error
```
Encode encodes the atom list using the specified encoder.

#### func (*List) WriteTo

```go
func (l *List) WriteTo(w Writer)
```
WriteTo writes all atoms in the list to w, terminating with a single EOS atom.

#### type Observation

```go
type Observation struct {
	binary.Generate
	Range      memory.Range // The memory range that was observed.
	ResourceID binary.ID    // The resource identifier holding the memory that was observed.
}
```

Observation is an Atom describing a region of application space memory that was
observed at capture time.

#### func (*Observation) API

```go
func (a *Observation) API() gfxapi.API
```
Atom compliance

#### func (*Observation) Class

```go
func (*Observation) Class() binary.Class
```

#### func (*Observation) Flags

```go
func (a *Observation) Flags() Flags
```

#### func (*Observation) Mutate

```go
func (a *Observation) Mutate(s *gfxapi.State) error
```

#### func (*Observation) String

```go
func (a *Observation) String() string
```

#### func (*Observation) TypeID

```go
func (a *Observation) TypeID() TypeID
```

#### type Range

```go
type Range struct {
	binary.Generate
	Start ID // The first atom within the range.
	End   ID // One past the last atom within the range.
}
```

Range describes an interval of atoms in a stream.

#### func (*Range) Class

```go
func (*Range) Class() binary.Class
```

#### func (Range) Contains

```go
func (i Range) Contains(atomID ID) bool
```
Contains returns true if atomID is within the range, otherwise false.

#### func (Range) First

```go
func (i Range) First() ID
```
First returns the first atom identifier within the range.

#### func (Range) Last

```go
func (i Range) Last() ID
```
Last returns the last atom identifier within the range.

#### func (Range) Length

```go
func (i Range) Length() uint64
```
Length returns the number of atoms in the range.

#### func (Range) Range

```go
func (i Range) Range() (start, end ID)
```
Range returns the start and end of the range.

#### func (*Range) SetSpan

```go
func (i *Range) SetSpan(span interval.U64Span)
```
SetSpan sets the start and end range using a U64Span.

#### func (Range) Span

```go
func (i Range) Span() interval.U64Span
```
Span returns the start and end of the range as a U64Span.

#### func (Range) String

```go
func (i Range) String() string
```
String returns a string representing the range.

#### type RangeList

```go
type RangeList []Range
```

RangeList is a list of atom ranges.

#### func (RangeList) Copy

```go
func (l RangeList) Copy(to, from, count int)
```
Copy copies count ranges within the list.

#### func (RangeList) GetSpan

```go
func (l RangeList) GetSpan(index int) interval.U64Span
```
GetSpan returns the atom identifier span for the range at index in the list.

#### func (RangeList) Length

```go
func (l RangeList) Length() int
```
Length returns the number of ranges in the list.

#### func (*RangeList) Resize

```go
func (l *RangeList) Resize(length int)
```
Resize adjusts the length of the list.

#### func (RangeList) SetSpan

```go
func (l RangeList) SetSpan(index int, span interval.U64Span)
```
SetSpan sets the atom identifier span for the group at index in the list.

#### type Resource

```go
type Resource struct {
	binary.Generate
	ResourceID binary.ID // The resource identifier holding the memory that was observed.
	Data       []byte    // The resource data
}
```

Resource is an Atom that embeds a blob of memory into the stream. These atoms
are typically only used for .gfxtrace files as they are stripped from the stream
on import and placed into the database.

#### func (*Resource) API

```go
func (a *Resource) API() gfxapi.API
```
Atom compliance

#### func (*Resource) Class

```go
func (*Resource) Class() binary.Class
```

#### func (*Resource) Flags

```go
func (a *Resource) Flags() Flags
```

#### func (*Resource) Mutate

```go
func (a *Resource) Mutate(s *gfxapi.State) error
```

#### func (*Resource) String

```go
func (a *Resource) String() string
```

#### func (*Resource) TypeID

```go
func (a *Resource) TypeID() TypeID
```

#### type Transformer

```go
type Transformer interface {
	// Transform takes a given atom and identifier and Writes out a new atom and
	// identifier to the output atom Writer. Transform must not modify the atom in
	// any way.
	Transform(id ID, atom Atom, output Writer)
	// Flush is called at the end of an atom stream to cause Transformers that
	// cache atoms to send any they have stored into the output.
	Flush(output Writer)
}
```

Transformer is the interface that wraps the basic Transform method.

#### func  Transform

```go
func Transform(name string, f func(id ID, atom Atom, output Writer)) Transformer
```
Transform is a helper for building simple Transformers that are implemented by
function f. name is used to identify the transform when logging.

#### type Transforms

```go
type Transforms []Transformer
```

Transforms is a list of Transformer objects.

#### func (*Transforms) Add

```go
func (l *Transforms) Add(t ...Transformer)
```
Add is a convenience function for appending the list of Transformers t to the
end of the Transforms list.

#### func (Transforms) Transform

```go
func (l Transforms) Transform(atoms List, out Writer)
```
Transform sequentially transforms the atoms by each of the transformers in the
list, before writing the final output to the output atom Writer.

#### type TypeID

```go
type TypeID uint16
```

TypeID is an atom type identifier. Each implementation of the Atom interface
must have a unique type idenitifier. Any changes to the binary format of an atom
must result in a new type identifier to maintain binary compatability.

```go
const TypeIDEos TypeID = 0xffff
```
TypeIDEos is used as a special end of stream marker.

```go
const TypeIDObservation TypeID = 0xfffe
```

```go
const TypeIDResource TypeID = 0xfffd
```

#### type TypeInfo

```go
type TypeInfo struct {
	ID   TypeID      // The type identifier for the atom.
	New  func() Atom // The function for creating new instances of the atom type.
	Name string      // The name of the atom.
	Docs string      // The URL to the atom's documentation.
}
```

TypeInfo is the type information for a single Atom implementation.

#### type Writer

```go
type Writer interface {
	Write(id ID, atom Atom)
}
```

Writer is the interface that wraps the basic Write method.

Write writes or processes the given atom and identifier. Write must not modify
the atom in any way.
