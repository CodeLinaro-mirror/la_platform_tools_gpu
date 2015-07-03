# check
--
    import "android.googlesource.com/platform/tools/gpu/check"

Package check contains test helper functions that verify expected state.

## Usage

#### func  SlicesEqual

```go
func SlicesEqual(t *testing.T, got interface{}, expected interface{}) (equal bool)
```
SlicesEqual checks the array or slice of got matches expected. If the arrays or
slices are equal then true is returned. If any differences are found then these
are logged to t, the test fails and false is returned.
