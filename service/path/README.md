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
	Array Value  // The path to the array.
	Index uint64 // The index of the element in the array.
}
```

ArrayIndex is a path that refers to a single element of an array.

#### func (*ArrayIndex) ArrayIndex

```go
func (n *ArrayIndex) ArrayIndex(index uint64) Value
```
ArrayIndex implements the Value interface.

#### func (*ArrayIndex) Base

```go
func (n *ArrayIndex) Base() Path
```
Base implements the Path interface, returning the path to the array.

#### func (*ArrayIndex) Class

```go
func (*ArrayIndex) Class() binary.Class
```

#### func (*ArrayIndex) Field

```go
func (n *ArrayIndex) Field(name string) Value
```
Field implements the Value interface.

#### func (*ArrayIndex) MapIndex

```go
func (n *ArrayIndex) MapIndex(key interface{}) Value
```
MapIndex implements the Value interface.

#### func (*ArrayIndex) Path

```go
func (n *ArrayIndex) Path() string
```
Path implements the Path interface.

#### func (*ArrayIndex) String

```go
func (n *ArrayIndex) String() string
```
String returns the string representation of the path.

#### type Atom

```go
type Atom struct {
	binary.Generate
	Atoms *Atoms // The path to the list of atoms.
	Index uint64 // The index of the atom in the array.
}
```

Atom is a path that refers to a single atom in an atom list.

#### func (*Atom) ArrayIndex

```go
func (n *Atom) ArrayIndex(index uint64) Value
```
ArrayIndex implements the Value interface.

#### func (*Atom) Base

```go
func (n *Atom) Base() Path
```
Base implements the Path interface, returning the path to the atom list.

#### func (*Atom) Class

```go
func (*Atom) Class() binary.Class
```

#### func (*Atom) Field

```go
func (n *Atom) Field(name string) Value
```
Field implements the Value interface.

#### func (*Atom) MapIndex

```go
func (n *Atom) MapIndex(key interface{}) Value
```
MapIndex implements the Value interface.

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

#### type Atoms

```go
type Atoms struct {
	binary.Generate
	Capture *Capture // The path to the capture containing the atoms.
}
```

Atoms is a path that refers to the full list of atoms in a capture.

#### func (*Atoms) Base

```go
func (n *Atoms) Base() Path
```
Base implements the Path interface, returning the path to the capture.

#### func (*Atoms) Class

```go
func (*Atoms) Class() binary.Class
```

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

#### func (*Atoms) String

```go
func (n *Atoms) String() string
```
String returns the string representation of the path.

#### type Capture

```go
type Capture struct {
	binary.Generate
	ID binary.ID // The capture's unique identifier.
}
```

Capture is a path that refers to a capture.

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

#### func (*Capture) Path

```go
func (c *Capture) Path() string
```
Path implements the Path interface.

#### func (*Capture) String

```go
func (c *Capture) String() string
```
String returns the string representation of the path.

#### type Field

```go
type Field struct {
	binary.Generate
	Struct Value  // The path to the structure holding the field.
	Name   string // The name of the field.
}
```

Field is a path that refers to a single field of a struct object.

#### func (*Field) ArrayIndex

```go
func (n *Field) ArrayIndex(index uint64) Value
```
ArrayIndex implements the Value interface.

#### func (*Field) Base

```go
func (n *Field) Base() Path
```
Base implements the Path interface, returning the path to the struct.

#### func (*Field) Class

```go
func (*Field) Class() binary.Class
```

#### func (*Field) Field

```go
func (n *Field) Field(name string) Value
```
Field implements the Value interface.

#### func (*Field) MapIndex

```go
func (n *Field) MapIndex(key interface{}) Value
```
MapIndex implements the Value interface.

#### func (*Field) Path

```go
func (n *Field) Path() string
```
Path implements the Path interface.

#### func (*Field) String

```go
func (n *Field) String() string
```
String returns the string representation of the path.

#### type MapIndex

```go
type MapIndex struct {
	binary.Generate
	Map Value       // The path to the map containing the value.
	Key interface{} // The key to the value in the map.
}
```

MapIndex is a path that refers to a single value in a map.

#### func (*MapIndex) ArrayIndex

```go
func (n *MapIndex) ArrayIndex(index uint64) Value
```
ArrayIndex implements the Value interface.

#### func (*MapIndex) Base

```go
func (n *MapIndex) Base() Path
```
Base implements the Path interface, returning the path to the map.

#### func (*MapIndex) Class

```go
func (*MapIndex) Class() binary.Class
```

#### func (*MapIndex) Field

```go
func (n *MapIndex) Field(name string) Value
```
Field implements the Value interface.

#### func (*MapIndex) MapIndex

```go
func (n *MapIndex) MapIndex(key interface{}) Value
```
MapIndex implements the Value interface.

#### func (*MapIndex) Path

```go
func (n *MapIndex) Path() string
```
Path implements the Path interface.

#### func (*MapIndex) String

```go
func (n *MapIndex) String() string
```
String returns the string representation of the path.

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
}
```

Path is the interface for types that represent a reference to a capture, atom
list, single atom, memory, state or sub-object. A path can be passed between
client and server using RPCs in order to describe some data in a capture.

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
func (n *State) ArrayIndex(index uint64) Value
```
ArrayIndex implements the Value interface.

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

#### func (*State) Field

```go
func (n *State) Field(name string) Value
```
Field implements the Value interface.

#### func (*State) MapIndex

```go
func (n *State) MapIndex(key interface{}) Value
```
MapIndex implements the Value interface.

#### func (*State) Path

```go
func (n *State) Path() string
```
Path implements the Path interface.

#### func (*State) String

```go
func (n *State) String() string
```
String returns the string representation of the path.

#### type Value

```go
type Value interface {
	// Value extends the Path interface.
	Path

	// Field returns the path to the field value with the specified name on the
	// struct object represented by this path.
	// The represented value type must be of type struct, otherwise the returned
	// path is invalid.
	Field(name string) Value

	// ArrayIndex returns the path to the i'th array element on the array object
	// represented by this path.
	// The represented value type must be of type array or slice, otherwise the
	// returned path is invalid.
	ArrayIndex(i uint64) Value

	// MapIndex returns the path to the map element with key k on the map object
	// represented by this path.
	// The represented value type must be of type map, otherwise the returned path
	// is invalid.
	MapIndex(k interface{}) Value
}
```

Value is the expanded Path interface for types that represent a reference to a
value type. The value referenced by this path may be a struct, array, slice, map
or POD type.
