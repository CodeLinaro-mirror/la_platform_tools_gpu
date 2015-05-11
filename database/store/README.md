# store
--
    import "android.googlesource.com/platform/tools/gpu/database/store"

Package store implements the storage layers of the database system.

## Usage

#### func  CopyResource

```go
func CopyResource(out interface{}, value interface{})
```
CopyResource assigns the value object to the variable out points to

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

#### type Store

```go
type Store interface {
	// Stores the resource r with the key id. d holds r as binary encoded.
	Store(id binary.ID, r binary.Object, d []byte, l log.Logger) error
	// Load puts the resource with key id into the out parameter
	Load(id binary.ID, l log.Logger, out binary.Object) (size int, err error)
	// Contains returns true if this store contains the resource identified by id.
	Contains(id binary.ID) bool
	// Close shuts down the store, it is an error to call any other method after this one.
	Close()
}
```

Store is the interface to a database storage layer.

#### func  CreateCache

```go
func CreateCache(maxSize int, inner Store) Store
```
CreateCache builds a store that uses a memory bounded lru cache over the top of
another store.

#### func  CreateCompactingArchive

```go
func CreateCompactingArchive(path string) Store
```
CreateCompactingArchive returns a Store that attempts to reclaim space by
dropping old resources that can be rebuilt on demand, and compacting the
remaining resources to reclaim space.

#### func  CreateSmallArchive

```go
func CreateSmallArchive(path string, compactionSize int) Store
```
CreateSmallArchive creates a Store optimized for small objects. When the
resources are small enough the index entry of a regular archive uses more space
than the data itself. The small archive stores the data unsorted in a single
file. New records with a new value are appended to the data file. All current
records are stored in RAM (this uses less RAM than storing all the index entries
for the equivalent archive in RAM).

#### func  CreateStoreIfNew

```go
func CreateStoreIfNew(inner Store) Store
```
CreateStoreIfNew wraps a store with a helper that discards attempts to store a
resource with a duplicate id to one already in the store.

#### func  CreateUnboundedArchive

```go
func CreateUnboundedArchive(path string) Store
```
CreateUnboundedArchive returns a store that appends new resources, and never
attempts to reclaim obsolete ones.
