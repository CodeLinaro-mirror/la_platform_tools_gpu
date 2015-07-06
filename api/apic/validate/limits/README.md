# limits
--
    import "android.googlesource.com/platform/tools/gpu/api/apic/validate/limits"

Package limits is used to calculate the possible values for a variable.

## Usage

```go
const (
	False = boolLimit(iota) // A boolean limit that is false.
	True                    // A boolean limit that is true.
	Maybe                   // A boolean limit that can be either true of false.
)
```

```go
var ConstantValues schema.Constants
```

#### type Limits

```go
type Limits interface {
	// Binary returns the unary operator op performed with the limits.
	Unary(op string) Limits
	// Binary returns the binary operator op performed with the limits and rhs.
	Binary(op string, rhs Limits) Limits
}
```

Limits represent the possible values for a given type.

#### func  Uint

```go
func Uint(v uint64) Limits
```

#### func  Unbound

```go
func Unbound(ty semantic.Type) Limits
```
Unbound returns the unbound limits for the given semantic type.
