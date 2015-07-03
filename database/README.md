# database
--
    import "android.googlesource.com/platform/tools/gpu/database"

Package database implements the persistence layer for the gpu debugger tools.

## Usage

```go
var Namespace = registry.NewNamespace()
```

#### func  Build

```go
func Build(lazy Lazy, d Database, l log.Logger) (interface{}, error)
```
Build stores lazy into d, and then resolves and returns the lazy-built object.

#### func  Hash

```go
func Hash(v interface{}) (binary.ID, error)
```
Hash returns a unique binary.ID based on the contents of the object. Two objects
of identical content will return the same ID, and the probability of two objects
with different content generating the same ID will be ignorable. Objects with a
graph structure are allowed. Only members that would be encoded using a
binary.Encoder are considered.

#### func  LazyOutputID

```go
func LazyOutputID(in binary.ID) binary.ID
```
LazyOutputID returns the identifier of a LazyOutput object given the identifier
of the Lazy. The database will not contain the LazyObject with the returned
identifier until it is built.

#### func  ResolveBlob

```go
func ResolveBlob(id binary.ID, d Database, l log.Logger) ([]byte, error)
```
Resolve blob loads a Blob from the database, returning the byte slice.

#### func  Store

```go
func Store(v interface{}, d Database, l log.Logger) (binary.ID, error)
```
Store is a helper that stores v to the database with the id calculated by the
Hash function.

#### func  StoreBlob

```go
func StoreBlob(data []byte, d Database, l log.Logger) (binary.ID, error)
```
StoreBlob stores the byte slice data inside a Blob to the database d.

#### type Blob

```go
type Blob struct {
	binary.Generate
	Data []byte
}
```

Blob is an encodable wrapper for a byte array, used for storing raw data in
databases.

#### func (*Blob) Class

```go
func (*Blob) Class() binary.Class
```

#### type Database

```go
type Database interface {
	// Store adds a key-value pair to the database.
	// It is an error if the id is already mapped to an object.
	Store(binary.ID, interface{}, log.Logger) error
	// Resolve attempts to resolve the final value associated with an id.
	// It will traverse all Lazy objects, blocking until they are ready.
	Resolve(binary.ID, log.Logger) (interface{}, error)
	// Containts returns true if the database has an entry for the specified id.
	Contains(binary.ID, log.Logger) bool
}
```

Database is the interface to a resource store.

#### func  NewInMemory

```go
func NewInMemory(buildContext interface{}) Database
```
NewInMemory builds a new in memory database.

#### type Lazy

```go
type Lazy interface {
	binary.Object

	// BuildLazy constructs and returns the lazily-built object.
	// c is the build context that was passed to the database constructor.
	BuildLazy(c interface{}, d Database, l log.Logger) (interface{}, error)
}
```

Lazy is the interface for types that redirects database resolves to an object
lazily built using BuildLazy(). The BuildLazy() method will be called the first
time the object is resolved, and all subsequent resolves will return the same
pre-built object. Lazy is commonly implemented by objects that generate data
that is expensive to calculate but can be deterministically produced using the
information stored in the Lazy.
