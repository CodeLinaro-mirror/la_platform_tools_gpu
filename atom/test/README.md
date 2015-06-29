# test
--
    import "android.googlesource.com/platform/tools/gpu/atom/test"

Package test provides testing helpers for the atom package.

## Usage

```go
var (
	AtomAID = binary.ID{0x38, 0x1e, 0xef, 0x73, 0x50, 0xa0, 0x48, 0x6d, 0xa3, 0x1d, 0x8e, 0xb6, 0x5e, 0x47, 0xb7, 0xbf, 0x7b, 0xc8, 0x06, 0x33}
	AtomBID = binary.ID{0x32, 0x6a, 0x98, 0x0f, 0x59, 0xd2, 0x52, 0x34, 0x9c, 0xc2, 0x75, 0x25, 0x62, 0xb8, 0xb3, 0x0b, 0x48, 0x54, 0x3c, 0x85}
	AtomCID = binary.ID{0x02, 0x32, 0xdd, 0xd7, 0x4d, 0x7e, 0xbf, 0x43, 0x41, 0x47, 0xbb, 0xcb, 0xd8, 0xed, 0xd4, 0xd8, 0xb4, 0x42, 0xf9, 0xd5}
)
```

```go
var Namespace = registry.NewNamespace()
```

#### type AtomA

```go
type AtomA struct {
	binary.Generate `id:"AtomAID"`
	ID              atom.ID
	AtomFlags       atom.Flags
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
func (a *AtomA) Mutate(*gfxapi.State, database.Database, log.Logger) error
```

#### func (*AtomA) Observations

```go
func (a *AtomA) Observations() *atom.Observations
```

#### type AtomB

```go
type AtomB struct {
	binary.Generate `id:"AtomBID"`
	ID              atom.ID
	Bool            bool
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
func (a *AtomB) Mutate(*gfxapi.State, database.Database, log.Logger) error
```

#### func (*AtomB) Observations

```go
func (a *AtomB) Observations() *atom.Observations
```

#### type AtomC

```go
type AtomC struct {
	binary.Generate `id:"AtomCID"`
	String          string
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
func (a *AtomC) Mutate(*gfxapi.State, database.Database, log.Logger) error
```

#### func (*AtomC) Observations

```go
func (a *AtomC) Observations() *atom.Observations
```
