# copyright
--
    import "android.googlesource.com/platform/tools/gpu/tools/copyright"


## Usage

```go
var (
	External  = []*regexp.Regexp{}
	Generated = []*regexp.Regexp{}
	Normal    = []*regexp.Regexp{}
)
```

#### func  Build

```go
func Build(name string, i Info) string
```

#### func  MatchExternal

```go
func MatchExternal(file []byte) int
```

#### func  MatchGenerated

```go
func MatchGenerated(file []byte) int
```

#### func  MatchNormal

```go
func MatchNormal(file []byte) int
```

#### func  Regexp

```go
func Regexp(name string, i Info, trim bool) *regexp.Regexp
```

#### type Info

```go
type Info struct {
	Year string
	Tool string
}
```


#### type Language

```go
type Language struct {
	Name       string
	Extensions []string
	License    string
	Emit       string
	Current    []*regexp.Regexp
	Old        []*regexp.Regexp
}
```


#### func  FindExtension

```go
func FindExtension(ext string) *Language
```

#### func  FindLanguage

```go
func FindLanguage(name string) *Language
```

#### func (*Language) MatchCurrent

```go
func (l *Language) MatchCurrent(file []byte) int
```

#### func (*Language) MatchOld

```go
func (l *Language) MatchOld(file []byte) int
```
