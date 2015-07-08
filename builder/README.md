# builder
--
    import "android.googlesource.com/platform/tools/gpu/builder"

Package builder implements builders for resources from requests typically stored
in the database, optionally depending on replay outputs.

## Usage

```go
var Namespace = registry.NewNamespace()
```

#### func  Captures

```go
func Captures(db database.Database, logger log.Logger) ([]service.CaptureId, error)
```
Captures returns all the captures stored by the database by identifier.

#### func  ImportCapture

```go
func ImportCapture(name string, atoms atom.List, d database.Database, l log.Logger) (service.CaptureId, error)
```
ImportCapture builds a new capture containing atoms, stores it into db and
returns the new capture identifier.

#### func  Resolve

```go
func Resolve(p path.Path, d database.Database, l log.Logger) (interface{}, error)
```
Resolve resolves and returns the object, value or memory at the path p.

#### func  ResolveAtoms

```go
func ResolveAtoms(p *path.Atoms, d database.Database, l log.Logger) ([]atom.Atom, error)
```
ResolveAtoms resolves and returns the atom list from the path p.

#### type BuildReport

```go
type BuildReport struct {
	binary.Generate
	Capture *path.Capture
}
```

BuildReport generates a service.Report for the given capture.

#### func (*BuildReport) BuildLazy

```go
func (r *BuildReport) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy writes to out the schema.Report resource resulting from the given
ResolveReport request.

#### func (*BuildReport) Class

```go
func (*BuildReport) Class() binary.Class
```

#### type Context

```go
type Context struct {
	ReplayManager *replay.Manager
}
```

Context is the type that should be passed to the database constructor's
buildContext parameter.

#### type ConvertImage

```go
type ConvertImage struct {
	binary.Generate
	Data       binary.ID
	Width      int
	Height     int
	FormatFrom image.Format
	FormatTo   image.Format
}
```

ConvertImage is a request to decode a compressed texture.

#### func (*ConvertImage) BuildLazy

```go
func (r *ConvertImage) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the byte array holding the converted image for the
ConvertImage request.

#### func (*ConvertImage) Class

```go
func (*ConvertImage) Class() binary.Class
```

#### type Get

```go
type Get struct {
	binary.Generate
	Path path.Path
}
```

Get resolves the object, value or memory at Path.

#### func (*Get) BuildLazy

```go
func (request *Get) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy resolves and returns the object, value or memory at the requested
path.

#### func (*Get) Class

```go
func (*Get) Class() binary.Class
```

#### type GetFramebufferColor

```go
type GetFramebufferColor struct {
	binary.Generate
	Device   service.DeviceId
	Capture  service.CaptureId
	API      service.ApiId
	After    atom.ID
	Settings service.RenderSettings
}
```

GetFramebufferColor records the parameters of a service.GetFramebufferColor RPC
request.

#### func (*GetFramebufferColor) BuildLazy

```go
func (r *GetFramebufferColor) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.ImageInfo resulting from the given
GetFramebufferColor request.

#### func (*GetFramebufferColor) Class

```go
func (*GetFramebufferColor) Class() binary.Class
```

#### type GetFramebufferDepth

```go
type GetFramebufferDepth struct {
	binary.Generate
	Device  service.DeviceId
	Capture service.CaptureId
	API     service.ApiId
	After   atom.ID
}
```

GetFramebufferDepth records the parameters of a service.GetFramebufferDepth RPC
request.

#### func (*GetFramebufferDepth) BuildLazy

```go
func (r *GetFramebufferDepth) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.ImageInfo resulting from the given
GetFramebufferDepth request.

#### func (*GetFramebufferDepth) Class

```go
func (*GetFramebufferDepth) Class() binary.Class
```

#### type GetHierarchy

```go
type GetHierarchy struct {
	binary.Generate
	Capture *path.Capture
}
```

GetHierarchy records the parameters of a service.GetHierarchy RPC request.

#### func (*GetHierarchy) BuildLazy

```go
func (r *GetHierarchy) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.Hierarchy resulting from the given GetHierarchy
request.

#### func (*GetHierarchy) Class

```go
func (*GetHierarchy) Class() binary.Class
```

#### type GetMemoryInfo

```go
type GetMemoryInfo struct {
	binary.Generate
	After *path.Atom
	Range memory.Range
}
```

GetMemoryInfo records the parameters of a service.GetMemoryInfo RPC request.

#### func (*GetMemoryInfo) BuildLazy

```go
func (r *GetMemoryInfo) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.MemoryInfo resulting from the given GetMemoryInfo
request.

#### func (*GetMemoryInfo) Class

```go
func (*GetMemoryInfo) Class() binary.Class
```

#### type GetState

```go
type GetState struct {
	binary.Generate
	Capture service.CaptureId
	API     service.ApiId
	After   atom.ID
}
```

GetState records the parameters of a service.GetState RPC request.

#### func (*GetState) BuildLazy

```go
func (r *GetState) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.Binary resulting from the given GetState request.

#### func (*GetState) Class

```go
func (*GetState) Class() binary.Class
```

#### type GetTimingInfo

```go
type GetTimingInfo struct {
	binary.Generate
	Device  service.DeviceId
	Capture service.CaptureId
	Flags   service.TimingFlags
}
```

GetTimingInfo records the parameters of a service.GetTimingInfo RPC request.

#### func (*GetTimingInfo) BuildLazy

```go
func (r *GetTimingInfo) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.TimingInfo resulting from the given GetTimingInfo
request.

#### func (*GetTimingInfo) Class

```go
func (*GetTimingInfo) Class() binary.Class
```

#### type PrerenderFramebuffers

```go
type PrerenderFramebuffers struct {
	binary.Generate
	Device  service.DeviceId
	Capture service.CaptureId
	API     service.ApiId
	AtomIDs []uint64
	Width   uint32
	Height  uint32
}
```

PrerenderFramebuffers records the parameters of a service.PrerenderFramebuffers
RPC request.

#### func (*PrerenderFramebuffers) BuildLazy

```go
func (r *PrerenderFramebuffers) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy renders and caches all the framebuffer color buffers in the
GetFramebufferDepth request, returning an empty *service.Binary.

#### func (*PrerenderFramebuffers) Class

```go
func (*PrerenderFramebuffers) Class() binary.Class
```

#### type RenderFramebufferColor

```go
type RenderFramebufferColor struct {
	binary.Generate
	Device    service.DeviceId
	Capture   service.CaptureId
	API       service.ApiId
	After     atom.ID
	Width     uint32
	Height    uint32
	Wireframe bool
}
```

RenderFramebufferColor records the parameters of an internal
RenderFramebufferColor request.

#### func (*RenderFramebufferColor) BuildLazy

```go
func (r *RenderFramebufferColor) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.Binary data for the given RenderFramebufferColor
request.

#### func (*RenderFramebufferColor) Class

```go
func (*RenderFramebufferColor) Class() binary.Class
```

#### type RenderFramebufferDepth

```go
type RenderFramebufferDepth struct {
	binary.Generate
	Device            service.DeviceId
	Capture           service.CaptureId
	API               service.ApiId
	After             atom.ID
	FramebufferWidth  uint32
	FramebufferHeight uint32
}
```

RenderFramebufferDepth records the parameters of an internal
RenderFramebufferDepth request.

#### func (*RenderFramebufferDepth) BuildLazy

```go
func (r *RenderFramebufferDepth) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy returns the *service.Binary data for the given RenderFramebufferDepth
request.

#### func (*RenderFramebufferDepth) Class

```go
func (*RenderFramebufferDepth) Class() binary.Class
```

#### type Set

```go
type Set struct {
	binary.Generate
	Path  path.Path
	Value interface{}
}
```

Set creates a copy of the capture referenced by Path, but with the object, value
or memory at p replaced with v. The path returned is identical to Path, but with
the base changed to refer to the new capture.

#### func (*Set) BuildLazy

```go
func (request *Set) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error)
```
BuildLazy creates a copy of the capture referenced by the request's path, but
with the object, value or memory at p replaced with v. The path returned is
identical to p, but with the base changed to refer to the new capture.

#### func (*Set) Class

```go
func (*Set) Class() binary.Class
```
