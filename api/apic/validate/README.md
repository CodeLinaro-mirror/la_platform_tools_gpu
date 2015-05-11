# validate
--
    import "android.googlesource.com/platform/tools/gpu/api/apic/validate"

Package validate registers and implements the "validate" apic command.

The validate command analyses the specified API for correctness, reporting
errors if any problems are found.

## Usage

#### func  Validate

```go
func Validate(apiName string, api *semantic.API) []error
```
Validate performs a number of checks on the api file for correctness. If any
problems are found then they are returned as errors.
