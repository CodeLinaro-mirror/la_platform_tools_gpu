# schema
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/schema"


## Usage

```go
var Any = service.CreateSimpleInfo("any", service.TypeKindAny)
```

```go
var Bool = service.CreateSimpleInfo("bool", service.TypeKindBool)
```

```go
var Double = service.CreateSimpleInfo("double", service.TypeKindF64)
```

```go
var Float = service.CreateSimpleInfo("float", service.TypeKindF32)
```

```go
var ID = service.CreateSimpleInfo("id", service.TypeKindID)
```

```go
var Int = service.CreateSimpleInfo("int", service.TypeKindS64)
```

```go
var Memory = service.CreateSimpleInfo("memory", service.TypeKindMemory)
```

```go
var Pointer = service.CreateSimpleInfo("pointer", service.TypeKindPointer)
```

```go
var S16 = service.CreateSimpleInfo("s16", service.TypeKindS16)
```

```go
var S32 = service.CreateSimpleInfo("s32", service.TypeKindS32)
```

```go
var S64 = service.CreateSimpleInfo("s64", service.TypeKindS64)
```

```go
var S8 = service.CreateSimpleInfo("s8", service.TypeKindS8)
```

```go
var String = service.CreateSimpleInfo("string", service.TypeKindString)
```

```go
var U16 = service.CreateSimpleInfo("u16", service.TypeKindU16)
```

```go
var U32 = service.CreateSimpleInfo("u32", service.TypeKindU32)
```

```go
var U64 = service.CreateSimpleInfo("u64", service.TypeKindU64)
```

```go
var U8 = service.CreateSimpleInfo("u8", service.TypeKindU8)
```

```go
var Uint = service.CreateSimpleInfo("uint", service.TypeKindU64)
```

#### func  RegisterAtom

```go
func RegisterAtom(a service.AtomInfo)
```
RegisterAtom registers the atom info a with the schema.

#### func  Schema

```go
func Schema() service.Schema
```
Schema returns the schema of all registered atoms and APIs.
