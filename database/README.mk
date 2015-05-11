# database
--
    import "android.googlesource.com/platform/tools/gpu/database"

Package database implements the persistence layer for the gpu debugger tools.

## Usage

#### type Database

```go
type Database interface {
	StoreLink(to, id binary.ID, logger log.Logger) error
	StoreRequest(request binary.Object, logger log.Logger) (id binary.ID, err error)
	Store(binary.Object, log.Logger) (binary.ID, error)
	Load(binary.ID, log.Logger, binary.Object) error
	Contains(binary.ID, log.Logger) bool
	Close()
}
```

Database is the interface to a resource store.

#### func  Create

```go
func Create(path string, maxDataCacheSize, metaDataCompactionSize, maxDerivedCacheSize int, builder builder) Database
```
Create builds a new database.

#### func  CreateTransientDatabase

```go
func CreateTransientDatabase(inner Database) Database
```
