# test
--
    import "android.googlesource.com/platform/tools/gpu/atom/test"

Package test provides testing helpers for the atom package.

## Usage

```go
const AtomIDA = atom.TypeID(1)
```

```go
const AtomIDB = atom.TypeID(2)
```

```go
const AtomIDC = atom.TypeID(3)
```

#### type AtomA

```go
type AtomA struct {
	binary.Generate
	ID        atom.ID
	AtomFlags atom.Flags
}
```


#### func (*AtomA) API

```go
func (a *AtomA) API() gfxapi.API
```

#### func (*AtomA) Class

```go
func (*AtomA) Class() binary.Class
```

#### func (*AtomA) Flags

```go
func (a *AtomA) Flags() atom.Flags
```

#### func (*AtomA) Mutate

```go
func (a *AtomA) Mutate(*gfxapi.State) error
```

#### func (*AtomA) TypeID

```go
func (a *AtomA) TypeID() atom.TypeID
```

#### type AtomB

```go
type AtomB struct {
	binary.Generate
	ID   atom.ID
	Bool bool
}
```


#### func (*AtomB) API

```go
func (a *AtomB) API() gfxapi.API
```

#### func (*AtomB) Class

```go
func (*AtomB) Class() binary.Class
```

#### func (*AtomB) Flags

```go
func (a *AtomB) Flags() atom.Flags
```

#### func (*AtomB) Mutate

```go
func (a *AtomB) Mutate(*gfxapi.State) error
```

#### func (*AtomB) TypeID

```go
func (a *AtomB) TypeID() atom.TypeID
```

#### type AtomC

```go
type AtomC struct {
	binary.Generate
	String string
}
```


#### func (*AtomC) API

```go
func (a *AtomC) API() gfxapi.API
```

#### func (*AtomC) Class

```go
func (*AtomC) Class() binary.Class
```

#### func (*AtomC) Flags

```go
func (a *AtomC) Flags() atom.Flags
```

#### func (*AtomC) Mutate

```go
func (a *AtomC) Mutate(*gfxapi.State) error
```

#### func (*AtomC) TypeID

```go
func (a *AtomC) TypeID() atom.TypeID
```
