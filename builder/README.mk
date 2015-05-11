# builder
--
    import "android.googlesource.com/platform/tools/gpu/builder"

Package builder implements builders for resources from requests typically stored
in the database, optionally depending on replay outputs.

## Usage

#### func  Captures

```go
func Captures(db database.Database, logger log.Logger) (service.CaptureIdArray, error)
```
Captures returns all the captures stored by the database.

#### func  ImportCapture

```go
func ImportCapture(name string, atoms atom.List, db database.Database, logger log.Logger) (service.CaptureId, error)
```
ImportCapture builds a new capture containing atoms, stores it into db and
returns the new capture identifier.

#### func  New

```go
func New() *builder
```
New creates a database.builder which can hold a replayManager, potentially
required to build request outputs.

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

#### func (*GetFramebufferDepth) Class

```go
func (*GetFramebufferDepth) Class() binary.Class
```

#### type GetHierarchy

```go
type GetHierarchy struct {
	binary.Generate
	Capture service.CaptureId
}
```

GetHierarchy records the parameters of a service.GetHierarchy RPC request.

#### func (*GetHierarchy) Class

```go
func (*GetHierarchy) Class() binary.Class
```

#### type GetMemoryInfo

```go
type GetMemoryInfo struct {
	binary.Generate
	Capture service.CaptureId
	After   atom.ID
	Range   memory.Range
}
```

GetMemoryInfo records the parameters of a service.GetMemoryInfo RPC request.

#### func (*GetMemoryInfo) Class

```go
func (*GetMemoryInfo) Class() binary.Class
```

#### type GetState

```go
type GetState struct {
	binary.Generate
	Capture service.CaptureId
	After   atom.ID
}
```

GetState records the parameters of a service.GetState RPC request.

#### func (*GetState) Class

```go
func (*GetState) Class() binary.Class
```

#### type GetTimingInfo

```go
type GetTimingInfo struct {
	binary.Generate
	Device     service.DeviceId
	Capture    service.CaptureId
	TimingMask service.TimingMask
}
```

GetTimingInfo records the parameters of a service.GetTimingInfo RPC request.

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

#### func (*RenderFramebufferDepth) Class

```go
func (*RenderFramebufferDepth) Class() binary.Class
```

#### type ReplaceAtom

```go
type ReplaceAtom struct {
	binary.Generate
	Capture service.CaptureId
	Atom    atom.ID
	Type    atom.TypeID
	Data    service.Binary
}
```

ReplaceAtom records the parameters of a service.ReplaceAtom RPC request.

#### func (*ReplaceAtom) Class

```go
func (*ReplaceAtom) Class() binary.Class
```
