# cc
--
    import "android.googlesource.com/platform/tools/gpu/cc"


## Usage

```go
var (
	ExternalRoot = build.RepoRoot.Path.Join("external")
	GPURoot      = build.RepoRoot.Path.Join("tools", "gpu", "src", "android.googlesource.com", "platform", "tools", "gpu")
	BinRoot      = build.RepoRoot.Path.Join("tools", "gpu", "bin")
	CCRoot       = GPURoot.Join("cc")
	GapicRoot    = CCRoot.Join("gapic")
	GapiiRoot    = CCRoot.Join("gapii")
	GapirRoot    = CCRoot.Join("gapir")
	ReplaydRoot  = CCRoot.Join("replayd")
	GmockRoot    = ExternalRoot.Join("gmock")
	GtestRoot    = ExternalRoot.Join("gtest")
)
```

#### func  Graph

```go
func Graph(targetNames []string)
```
Generate "maker" graph for building the cpp code.

#### type Target

```go
type Target struct {
	SourceFiles []string
	Gtest       cpp.Config
	Gmock       cpp.Config
	Gapic       cpp.Config
	GapicTests  cpp.Config
	Gapii       cpp.Config
	Gapir       cpp.Config
	GapirTests  cpp.Config
	Spy         cpp.Config
	Replayd     cpp.Config
}
```


#### func (Target) Build

```go
func (t Target) Build(env build.Environment)
```

#### func (Target) Extend

```go
func (t Target) Extend(n Target) Target
```
