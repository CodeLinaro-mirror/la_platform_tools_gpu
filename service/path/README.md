# path
--
    import "android.googlesource.com/platform/tools/gpu/service/path"

Package path contains types that represent data references.

## Usage

```go
var Namespace = registry.NewNamespace()
```

#### func  Flatten

```go
func Flatten(p Path) []Path
```
Flatten returns the path p flattened into a list of path nodes, starting with
the root and ending with p.

#### type ArrayIndex

```go
type ArrayIndex struct {
	binary.Generate
	Array Path   // The path to the array.
	Index uint64 // The index of the element in the array.
}
```

ArrayIndex is a path that refers to a single element of an array.

#### func  FindArrayIndex

```go
func FindArrayIndex(p Path) *ArrayIndex
```
FindArrayIndex returns the first ArrayIndex found traversing the path p. If no
Atom was found, then nil is returned.

#### func (*ArrayIndex) ArrayIndex

```go
func (n *ArrayIndex) ArrayIndex(index uint64) *ArrayIndex
```
ArrayIndex returns the path to the i'th element on the array or slice
represented by this path. The represented value type must be of type array or
slice, otherwise the returned path is invalid.

#### func (*ArrayIndex) Base

```go
func (n *ArrayIndex) Base() Path
```
Base implements the Path interface, returning the path to the array.

#### func (*ArrayIndex) Class

```go
func (*ArrayIndex) Class() binary.Class
```

#### func (*ArrayIndex) Clone

```go
func (n *ArrayIndex) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*ArrayIndex) Field

```go
func (n *ArrayIndex) Field(name string) *Field
```
Field returns the path to the field value with the specified name on the struct
object represented by this path. The represented value type must be of type
struct, otherwise the returned path is invalid.

#### func (*ArrayIndex) MapIndex

```go
func (n *ArrayIndex) MapIndex(key interface{}) *MapIndex
```
MapIndex returns the path to the map element with key k on the map object
represented by this path. The represented value type must be of type map,
otherwise the returned path is invalid.

#### func (*ArrayIndex) Path

```go
func (n *ArrayIndex) Path() string
```
Path implements the Path interface.

#### func (*ArrayIndex) Slice

```go
func (n *ArrayIndex) Slice(start, end uint64) *Slice
```
Slice returns the path to the sliced subset of this array or slice represented
by this path. The represented value type must be of type array or slice,
otherwise the returned path is invalid.

#### func (*ArrayIndex) String

```go
func (n *ArrayIndex) String() string
```
String returns the string representation of the path.

#### func (*ArrayIndex) Validate

```go
func (n *ArrayIndex) Validate() error
```
Validate implements the Path interface.

#### type Atom

```go
type Atom struct {
	binary.Generate
	Atoms *Atoms // The path to the list of atoms.
	Index uint64 // The index of the atom in the array.
}
```

Atom is a path that refers to a single atom in an atom list.

#### func  FindAtom

```go
func FindAtom(p Path) *Atom
```
FindAtom returns the first Atom found traversing the path p. If no Atom was
found, then nil is returned.

#### func (*Atom) Base

```go
func (n *Atom) Base() Path
```
Base implements the Path interface, returning the path to the atom list.

#### func (*Atom) Class

```go
func (*Atom) Class() binary.Class
```

#### func (*Atom) Clone

```go
func (n *Atom) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*Atom) Field

```go
func (n *Atom) Field(name string) *Field
```
Field returns the path to the field value with the specified name on the atom
represented by this path.

#### func (*Atom) Path

```go
func (n *Atom) Path() string
```
Path implements the Path interface.

#### func (*Atom) StateAfter

```go
func (n *Atom) StateAfter() *State
```
StateAfter returns the path to the state immediately following this atom.

#### func (*Atom) String

```go
func (n *Atom) String() string
```
String returns the string representation of the path.

#### func (*Atom) Validate

```go
func (n *Atom) Validate() error
```
Validate implements the Path interface.

#### type Atoms

```go
type Atoms struct {
	binary.Generate
	Capture *Capture // The path to the capture containing the atoms.
}
```

Atoms is a path that refers to the full list of atoms in a capture.

#### func  FindAtoms

```go
func FindAtoms(p Path) *Atoms
```
FindAtoms returns the first Atoms found traversing the path p. If no Atoms was
found, then nil is returned.

#### func (*Atoms) Base

```go
func (n *Atoms) Base() Path
```
Base implements the Path interface, returning the path to the atoms.

#### func (*Atoms) Class

```go
func (*Atoms) Class() binary.Class
```

#### func (*Atoms) Clone

```go
func (n *Atoms) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*Atoms) Index

```go
func (n *Atoms) Index(i uint64) *Atom
```
Index returns the path to the i'th atom in the atom list.

#### func (*Atoms) Path

```go
func (n *Atoms) Path() string
```
Path implements the Path interface.

#### func (*Atoms) Slice

```go
func (n *Atoms) Slice(start, end uint64) *Slice
```
Slice returns the path to the sliced subset of the atom list.

#### func (*Atoms) String

```go
func (n *Atoms) String() string
```
String returns the string representation of the path.

#### func (*Atoms) Validate

```go
func (n *Atoms) Validate() error
```
Validate implements the Path interface.

#### type Capture

```go
type Capture struct {
	binary.Generate
	ID binary.ID // The capture's unique identifier.
}
```

Capture is a path that refers to a capture.

#### func  FindCapture

```go
func FindCapture(p Path) *Capture
```
FindCapture returns the first Capture found traversing the path p. If no Capture
was found, then nil is returned.

#### func (*Capture) Atoms

```go
func (c *Capture) Atoms() *Atoms
```
Atoms returns the path to the full list of atoms in the capture.

#### func (*Capture) Base

```go
func (c *Capture) Base() Path
```
Base implements the Path interface, returning nil as this is a root.

#### func (*Capture) Class

```go
func (*Capture) Class() binary.Class
```

#### func (*Capture) Clone

```go
func (c *Capture) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*Capture) Hierarchy

```go
func (c *Capture) Hierarchy() *Hierarchy
```
Hierarchy returns the path to the capture's hierarchy.

#### func (*Capture) Path

```go
func (c *Capture) Path() string
```
Path implements the Path interface.

#### func (*Capture) Report

```go
func (c *Capture) Report() *Report
```
Report returns the path to the capture's report.

#### func (*Capture) String

```go
func (c *Capture) String() string
```
String returns the string representation of the path.

#### func (*Capture) Validate

```go
func (c *Capture) Validate() error
```
Validate implements the Path interface.

#### type Field

```go
type Field struct {
	binary.Generate
	Struct Path   // The path to the structure holding the field.
	Name   string // The name of the field.
}
```

Field is a path that refers to a single field of a struct object.

#### func  FindField

```go
func FindField(p Path) *Field
```
FindField returns the first Field found traversing the path p. If no Atom was
found, then nil is returned.

#### func (*Field) ArrayIndex

```go
func (n *Field) ArrayIndex(index uint64) *ArrayIndex
```
ArrayIndex returns the path to the i'th element on the array or slice
represented by this path. The represented value type must be of type array or
slice, otherwise the returned path is invalid.

#### func (*Field) Base

```go
func (n *Field) Base() Path
```
Base implements the Path interface, returning the path to the struct.

#### func (*Field) Class

```go
func (*Field) Class() binary.Class
```

#### func (*Field) Clone

```go
func (n *Field) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*Field) Field

```go
func (n *Field) Field(name string) *Field
```
Field returns the path to the field value with the specified name on the struct
object represented by this path. The represented value type must be of type
struct, otherwise the returned path is invalid.

#### func (*Field) MapIndex

```go
func (n *Field) MapIndex(key interface{}) *MapIndex
```
MapIndex returns the path to the map element with key k on the map object
represented by this path. The represented value type must be of type map,
otherwise the returned path is invalid.

#### func (*Field) Path

```go
func (n *Field) Path() string
```
Path implements the Path interface.

#### func (*Field) Slice

```go
func (n *Field) Slice(start, end uint64) *Slice
```
Slice returns the path to the sliced subset of this array or slice represented
by this path. The represented value type must be of type array or slice,
otherwise the returned path is invalid.

#### func (*Field) String

```go
func (n *Field) String() string
```
String returns the string representation of the path.

#### func (*Field) Validate

```go
func (n *Field) Validate() error
```
Validate implements the Path interface.

#### type Hierarchy

```go
type Hierarchy struct {
	binary.Generate
	Capture *Capture // The path to the capture containing the hierarchy.
}
```

Hierarchy is a path that refers to a capture's hierarchy.

#### func (*Hierarchy) Base

```go
func (n *Hierarchy) Base() Path
```
Base implements the Path interface, returning the path to the hierarchy.

#### func (*Hierarchy) Class

```go
func (*Hierarchy) Class() binary.Class
```

#### func (*Hierarchy) Clone

```go
func (n *Hierarchy) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*Hierarchy) Path

```go
func (n *Hierarchy) Path() string
```
Path implements the Path interface.

#### func (*Hierarchy) String

```go
func (n *Hierarchy) String() string
```
String returns the string representation of the path.

#### func (*Hierarchy) Validate

```go
func (n *Hierarchy) Validate() error
```
Validate implements the Path interface.

#### type MapIndex

```go
type MapIndex struct {
	binary.Generate
	Map Path        // The path to the map containing the value.
	Key interface{} // The key to the value in the map.
}
```

MapIndex is a path that refers to a single value in a map.

#### func (*MapIndex) ArrayIndex

```go
func (n *MapIndex) ArrayIndex(index uint64) *ArrayIndex
```
ArrayIndex returns the path to the i'th element on the array or slice
represented by this path. The represented value type must be of type array or
slice, otherwise the returned path is invalid.

#### func (*MapIndex) Base

```go
func (n *MapIndex) Base() Path
```
Base implements the Path interface, returning the path to the map.

#### func (*MapIndex) Class

```go
func (*MapIndex) Class() binary.Class
```

#### func (*MapIndex) Clone

```go
func (n *MapIndex) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*MapIndex) Field

```go
func (n *MapIndex) Field(name string) *Field
```
Field returns the path to the field value with the specified name on the struct
object represented by this path. The represented value type must be of type
struct, otherwise the returned path is invalid.

#### func (*MapIndex) MapIndex

```go
func (n *MapIndex) MapIndex(key interface{}) *MapIndex
```
MapIndex returns the path to the map element with key k on the map object
represented by this path. The represented value type must be of type map,
otherwise the returned path is invalid.

#### func (*MapIndex) Path

```go
func (n *MapIndex) Path() string
```
Path implements the Path interface.

#### func (*MapIndex) Slice

```go
func (n *MapIndex) Slice(start, end uint64) *Slice
```
Slice returns the path to the sliced subset of this array or slice represented
by this path. The represented value type must be of type array or slice,
otherwise the returned path is invalid.

#### func (*MapIndex) String

```go
func (n *MapIndex) String() string
```
String returns the string representation of the path.

#### func (*MapIndex) Validate

```go
func (n *MapIndex) Validate() error
```
Validate implements the Path interface.

#### type Path

```go
type Path interface {
	binary.Object

	// Path returns the string representation of the path.
	// The returned string must be consistent for equal paths.
	Path() string

	// Base returns the path that this path derives from.
	// If this path is a root, then Base returns nil.
	Base() Path

	// Clone returns a deep-copy of the path.
	Clone() Path

	// Validate checks the path for correctness, returning an error if any
	// issues are found.
	Validate() error
}
```

Path is the interface for types that represent a reference to a capture, atom
list, single atom, memory, state or sub-object. A path can be passed between
client and server using RPCs in order to describe some data in a capture.

#### type Report

```go
type Report struct {
	binary.Generate
	Capture *Capture // The path to the capture containing the report.
}
```

Report is a path that refers to a capture's report.

#### func (*Report) Base

```go
func (n *Report) Base() Path
```
Base implements the Path interface, returning the path to the report.

#### func (*Report) Class

```go
func (*Report) Class() binary.Class
```

#### func (*Report) Clone

```go
func (n *Report) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*Report) Path

```go
func (n *Report) Path() string
```
Path implements the Path interface.

#### func (*Report) String

```go
func (n *Report) String() string
```
String returns the string representation of the path.

#### func (*Report) Validate

```go
func (n *Report) Validate() error
```
Validate implements the Path interface.

#### type Slice

```go
type Slice struct {
	binary.Generate
	Array Path // The path to the array.
	Start uint64
	End   uint64
}
```

Slice is a path that refers to a subset of the elements in an array.

#### func  FindAtomSlice

```go
func FindAtomSlice(p Path) (*Slice, *Atoms)
```
FindAtomSlice returns the first slice of Atoms found traversing the path p. If
no Atoms was found, then nil is returned.

#### func  FindSlice

```go
func FindSlice(p Path) *Slice
```
FindSlice returns the first Slice found traversing the path p. If no Slice was
found, then nil is returned.

#### func (*Slice) Base

```go
func (n *Slice) Base() Path
```
Base implements the Path interface, returning the path to the array.

#### func (*Slice) Class

```go
func (*Slice) Class() binary.Class
```

#### func (*Slice) Clone

```go
func (n *Slice) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*Slice) Index

```go
func (n *Slice) Index(i uint64) Path
```
Index returns the path to the i'th element in the slice.

#### func (*Slice) Path

```go
func (n *Slice) Path() string
```
Path implements the Path interface.

#### func (*Slice) String

```go
func (n *Slice) String() string
```
String returns the string representation of the path.

#### func (*Slice) Validate

```go
func (n *Slice) Validate() error
```
Validate implements the Path interface.

#### type State

```go
type State struct {
	binary.Generate
	After *Atom // The path to the atom the state immediately follows.
}
```

State is a path that refers to the driver state immediately after an atom.

#### func (*State) ArrayIndex

```go
func (n *State) ArrayIndex(index uint64) *ArrayIndex
```
ArrayIndex returns the path to the i'th element on the array or slice
represented by this path. The represented value type must be of type array or
slice, otherwise the returned path is invalid.

#### func (*State) Base

```go
func (n *State) Base() Path
```
Base implements the Path interface, returning the path to the atom the state is
after.

#### func (*State) Class

```go
func (*State) Class() binary.Class
```

#### func (*State) Clone

```go
func (n *State) Clone() Path
```
Clone implements the Path interface, returning a deep-copy of this path.

#### func (*State) Field

```go
func (n *State) Field(name string) *Field
```
Field returns the path to the field value with the specified name on the struct
object represented by this path. The represented value type must be of type
struct, otherwise the returned path is invalid.

#### func (*State) MapIndex

```go
func (n *State) MapIndex(key interface{}) *MapIndex
```
MapIndex returns the path to the map element with key k on the map object
represented by this path. The represented value type must be of type map,
otherwise the returned path is invalid.

#### func (*State) Path

```go
func (n *State) Path() string
```
Path implements the Path interface.

#### func (*State) Slice

```go
func (n *State) Slice(start, end uint64) *Slice
```
Slice returns the path to the sliced subset of this array or slice represented
by this path. The represented value type must be of type array or slice,
otherwise the returned path is invalid.

#### func (*State) String

```go
func (n *State) String() string
```
String returns the string representation of the path.

#### func (*State) Validate

```go
func (n *State) Validate() error
```
Validate implements the Path interface.

#### type Value

```go
type Value interface {
	// Value extends the Path interface.
	Path

	// Field returns the path to the field value with the specified name on the
	// struct object represented by this path.
	// The represented value type must be of type struct, otherwise the returned
	// path is invalid.
	Field(name string) *Field

	// Slice returns the path to the sliced subset of this array or slice
	// represented by this path.
	// The represented value type must be of type array or slice, otherwise the
	// returned path is invalid.
	Slice(start, end uint64) *Slice

	// ArrayIndex returns the path to the i'th element on the array or slice
	// represented by this path.
	// The represented value type must be of type array or slice, otherwise the
	// returned path is invalid.
	ArrayIndex(i uint64) *ArrayIndex

	// MapIndex returns the path to the map element with key k on the map object
	// represented by this path.
	// The represented value type must be of type map, otherwise the returned path
	// is invalid.
	MapIndex(k interface{}) *MapIndex
}
```

Value is the expanded Path interface for types that represent a reference to a
value type. The value referenced by this path may be a struct, array, slice, map
or POD type.
