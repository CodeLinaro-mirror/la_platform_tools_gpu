# service
--
    import "android.googlesource.com/platform/tools/gpu/service"

Package service is the definition of the RPC GPU debugger service exposed by the
server.

It is not the actual implementation of the service functionality.

## Usage

```go
var ConstantValues schema.Constants
```

```go
var Namespace = registry.NewNamespace()
```

#### func  BindServer

```go
func BindServer(r io.Reader, w io.Writer, mtu int, l log.Logger, server RPC)
```

#### type ApiId

```go
type ApiId struct {
	binary.Generate
	ID binary.ID
}
```

Handle ApiId

#### func (*ApiId) Class

```go
func (*ApiId) Class() binary.Class
```

#### func (ApiId) Valid

```go
func (h ApiId) Valid() bool
```

#### type AtomRangeTimer

```go
type AtomRangeTimer struct {
	binary.Generate
	FromAtomId  uint64
	ToAtomId    uint64
	Nanoseconds uint64
}
```

Class AtomRangeTimer

#### func  CreateAtomRangeTimer

```go
func CreateAtomRangeTimer(
	FromAtomId uint64,
	ToAtomId uint64,
	Nanoseconds uint64,
) *AtomRangeTimer
```

#### func (*AtomRangeTimer) Class

```go
func (*AtomRangeTimer) Class() binary.Class
```

#### func (*AtomRangeTimer) GetFromAtomId

```go
func (c *AtomRangeTimer) GetFromAtomId() uint64
```

#### func (*AtomRangeTimer) GetNanoseconds

```go
func (c *AtomRangeTimer) GetNanoseconds() uint64
```

#### func (*AtomRangeTimer) GetToAtomId

```go
func (c *AtomRangeTimer) GetToAtomId() uint64
```

#### type AtomStream

```go
type AtomStream struct {
	binary.Generate
	Atoms []atom.Atom
}
```

Class AtomStream

#### func  CreateAtomStream

```go
func CreateAtomStream(
	Atoms []atom.Atom,
) *AtomStream
```

#### func  ResolveAtomStream

```go
func ResolveAtomStream(id AtomStreamId, d database.Database, l log.Logger) (res AtomStream, err error)
```
ResolveAtomStream loads and returns the AtomStream stored in the database d,
using id.

#### func (*AtomStream) Class

```go
func (*AtomStream) Class() binary.Class
```

#### func (*AtomStream) GetAtoms

```go
func (c *AtomStream) GetAtoms() []atom.Atom
```

#### type AtomStreamId

```go
type AtomStreamId struct {
	binary.Generate
	ID binary.ID
}
```

Handle AtomStreamId

#### func  StoreAtomStream

```go
func StoreAtomStream(v *AtomStream, d database.Database, l log.Logger) (AtomStreamId, error)
```
StoreAtomStream stores v into the database d, returning the AtomStreamId.

#### func (*AtomStreamId) Class

```go
func (*AtomStreamId) Class() binary.Class
```

#### func (AtomStreamId) Valid

```go
func (h AtomStreamId) Valid() bool
```

#### type AtomTimer

```go
type AtomTimer struct {
	binary.Generate
	AtomId      uint64
	Nanoseconds uint64
}
```

Class AtomTimer

#### func  CreateAtomTimer

```go
func CreateAtomTimer(
	AtomId uint64,
	Nanoseconds uint64,
) *AtomTimer
```

#### func (*AtomTimer) Class

```go
func (*AtomTimer) Class() binary.Class
```

#### func (*AtomTimer) GetAtomId

```go
func (c *AtomTimer) GetAtomId() uint64
```

#### func (*AtomTimer) GetNanoseconds

```go
func (c *AtomTimer) GetNanoseconds() uint64
```

#### type Binary

```go
type Binary struct {
	binary.Generate
	Data []uint8
}
```

Class Binary

#### func  CreateBinary

```go
func CreateBinary(
	Data []uint8,
) *Binary
```

#### func  ResolveBinary

```go
func ResolveBinary(id BinaryId, d database.Database, l log.Logger) (res Binary, err error)
```
ResolveBinary loads and returns the Binary stored in the database d, using id.

#### func (*Binary) Class

```go
func (*Binary) Class() binary.Class
```

#### func (*Binary) GetData

```go
func (c *Binary) GetData() []uint8
```

#### type BinaryId

```go
type BinaryId struct {
	binary.Generate
	ID binary.ID
}
```

Handle BinaryId

#### func  StoreBinary

```go
func StoreBinary(v *Binary, d database.Database, l log.Logger) (BinaryId, error)
```
StoreBinary stores v into the database d, returning the BinaryId.

#### func (*BinaryId) Class

```go
func (*BinaryId) Class() binary.Class
```

#### func (BinaryId) Valid

```go
func (h BinaryId) Valid() bool
```

#### type Capture

```go
type Capture struct {
	binary.Generate
	Name   string
	Atoms  AtomStreamId
	Report ReportId
	Apis   []ApiId
}
```

Class Capture

#### func  CreateCapture

```go
func CreateCapture(
	Name string,
	Atoms AtomStreamId,
	Report ReportId,
	Apis []ApiId,
) *Capture
```

#### func  ResolveCapture

```go
func ResolveCapture(id CaptureId, d database.Database, l log.Logger) (res Capture, err error)
```
ResolveCapture loads and returns the Capture stored in the database d, using id.

#### func (*Capture) Class

```go
func (*Capture) Class() binary.Class
```

#### func (*Capture) GetApis

```go
func (c *Capture) GetApis() []ApiId
```

#### func (*Capture) GetAtoms

```go
func (c *Capture) GetAtoms() AtomStreamId
```

#### func (*Capture) GetName

```go
func (c *Capture) GetName() string
```

#### func (*Capture) GetReport

```go
func (c *Capture) GetReport() ReportId
```

#### type CaptureId

```go
type CaptureId struct {
	binary.Generate
	ID binary.ID
}
```

Handle CaptureId

#### func  StoreCapture

```go
func StoreCapture(v *Capture, d database.Database, l log.Logger) (CaptureId, error)
```
StoreCapture stores v into the database d, returning the CaptureId.

#### func (*CaptureId) Class

```go
func (*CaptureId) Class() binary.Class
```

#### func (CaptureId) Path

```go
func (c CaptureId) Path() *path.Capture
```
Path returns a path.Capture representing the capture with this identifier.

#### func (CaptureId) Valid

```go
func (h CaptureId) Valid() bool
```

#### type Client

```go
type Client interface {
	// Client exposes all the RPC interface methods.
	RPC
	// Multiplexer returns the multiplexer used for communication to the server.
	Multiplexer() *multiplexer.Multiplexer
	// Namespace returns the custom namespace used for decoding responses from the
	// server, or nil if no custom namespace has been specified.
	Namespace() *registry.Namespace
}
```

Client is the client interface for RPC calls.

#### func  NewClient

```go
func NewClient(m *multiplexer.Multiplexer, n *registry.Namespace) Client
```
NewClient creates a new rpc client object that uses the multiplexer m for
communication the namespace n for decoding objects. If n is nil then the global
namespace is used.

#### type Device

```go
type Device struct {
	binary.Generate
	Name             string
	Model            string
	OS               string
	PointerSize      uint8
	PointerAlignment uint8
	MaxMemorySize    uint64
	Extensions       string
	Renderer         string
	Vendor           string
	Version          string
}
```

Class Device

#### func  CreateDevice

```go
func CreateDevice(
	Name string,
	Model string,
	OS string,
	PointerSize uint8,
	PointerAlignment uint8,
	MaxMemorySize uint64,
	Extensions string,
	Renderer string,
	Vendor string,
	Version string,
) *Device
```

#### func  ResolveDevice

```go
func ResolveDevice(id DeviceId, d database.Database, l log.Logger) (res Device, err error)
```
ResolveDevice loads and returns the Device stored in the database d, using id.

#### func (Device) Architecture

```go
func (d Device) Architecture() device.Architecture
```
Architecture return's the device's architecture.

#### func (*Device) Class

```go
func (*Device) Class() binary.Class
```

#### func (*Device) GetExtensions

```go
func (c *Device) GetExtensions() string
```

#### func (*Device) GetMaxMemorySize

```go
func (c *Device) GetMaxMemorySize() uint64
```

#### func (*Device) GetModel

```go
func (c *Device) GetModel() string
```

#### func (*Device) GetName

```go
func (c *Device) GetName() string
```

#### func (*Device) GetOS

```go
func (c *Device) GetOS() string
```

#### func (*Device) GetPointerAlignment

```go
func (c *Device) GetPointerAlignment() uint8
```

#### func (*Device) GetPointerSize

```go
func (c *Device) GetPointerSize() uint8
```

#### func (*Device) GetRenderer

```go
func (c *Device) GetRenderer() string
```

#### func (*Device) GetVendor

```go
func (c *Device) GetVendor() string
```

#### func (*Device) GetVersion

```go
func (c *Device) GetVersion() string
```

#### func (Device) HasExtension

```go
func (d Device) HasExtension(extension string) bool
```

#### type DeviceId

```go
type DeviceId struct {
	binary.Generate
	ID binary.ID
}
```

Handle DeviceId

#### func  StoreDevice

```go
func StoreDevice(v *Device, d database.Database, l log.Logger) (DeviceId, error)
```
StoreDevice stores v into the database d, returning the DeviceId.

#### func (*DeviceId) Class

```go
func (*DeviceId) Class() binary.Class
```

#### func (DeviceId) Valid

```go
func (h DeviceId) Valid() bool
```

#### type Hierarchy

```go
type Hierarchy struct {
	binary.Generate
	Root atom.Group
}
```

Class Hierarchy

#### func  CreateHierarchy

```go
func CreateHierarchy(
	Root atom.Group,
) *Hierarchy
```

#### func  ResolveHierarchy

```go
func ResolveHierarchy(id HierarchyId, d database.Database, l log.Logger) (res Hierarchy, err error)
```
ResolveHierarchy loads and returns the Hierarchy stored in the database d, using
id.

#### func (*Hierarchy) Class

```go
func (*Hierarchy) Class() binary.Class
```

#### func (*Hierarchy) GetRoot

```go
func (c *Hierarchy) GetRoot() atom.Group
```

#### type HierarchyId

```go
type HierarchyId struct {
	binary.Generate
	ID binary.ID
}
```

Handle HierarchyId

#### func  StoreHierarchy

```go
func StoreHierarchy(v *Hierarchy, d database.Database, l log.Logger) (HierarchyId, error)
```
StoreHierarchy stores v into the database d, returning the HierarchyId.

#### func (*HierarchyId) Class

```go
func (*HierarchyId) Class() binary.Class
```

#### func (HierarchyId) Valid

```go
func (h HierarchyId) Valid() bool
```

#### type ImageFormat

```go
type ImageFormat int
```

Enum ImageFormat

```go
const (
	ImageFormatRGBA8   ImageFormat = 0
	ImageFormatFloat32 ImageFormat = 1
)
```

#### func (ImageFormat) IsFloat32

```go
func (i ImageFormat) IsFloat32() bool
```

#### func (ImageFormat) IsRGBA8

```go
func (i ImageFormat) IsRGBA8() bool
```

#### func (*ImageFormat) Parse

```go
func (v *ImageFormat) Parse(s string) error
```

#### func (ImageFormat) String

```go
func (v ImageFormat) String() string
```

#### type ImageInfo

```go
type ImageInfo struct {
	binary.Generate
	Format ImageFormat
	Width  uint32
	Height uint32
	Data   BinaryId
}
```

Class ImageInfo

#### func  CreateImageInfo

```go
func CreateImageInfo(
	Format ImageFormat,
	Width uint32,
	Height uint32,
	Data BinaryId,
) *ImageInfo
```

#### func  ResolveImageInfo

```go
func ResolveImageInfo(id ImageInfoId, d database.Database, l log.Logger) (res ImageInfo, err error)
```
ResolveImageInfo loads and returns the ImageInfo stored in the database d, using
id.

#### func (*ImageInfo) Class

```go
func (*ImageInfo) Class() binary.Class
```

#### func (*ImageInfo) GetData

```go
func (c *ImageInfo) GetData() BinaryId
```

#### func (*ImageInfo) GetFormat

```go
func (c *ImageInfo) GetFormat() ImageFormat
```

#### func (*ImageInfo) GetHeight

```go
func (c *ImageInfo) GetHeight() uint32
```

#### func (*ImageInfo) GetWidth

```go
func (c *ImageInfo) GetWidth() uint32
```

#### type ImageInfoId

```go
type ImageInfoId struct {
	binary.Generate
	ID binary.ID
}
```

Handle ImageInfoId

#### func  StoreImageInfo

```go
func StoreImageInfo(v *ImageInfo, d database.Database, l log.Logger) (ImageInfoId, error)
```
StoreImageInfo stores v into the database d, returning the ImageInfoId.

#### func (*ImageInfoId) Class

```go
func (*ImageInfoId) Class() binary.Class
```

#### func (ImageInfoId) Valid

```go
func (h ImageInfoId) Valid() bool
```

#### type MemoryInfo

```go
type MemoryInfo struct {
	binary.Generate
	Data     []uint8
	Reads    memory.RangeList
	Writes   memory.RangeList
	Observed memory.RangeList
}
```

Class MemoryInfo

#### func  CreateMemoryInfo

```go
func CreateMemoryInfo(
	Data []uint8,
	Reads memory.RangeList,
	Writes memory.RangeList,
	Observed memory.RangeList,
) *MemoryInfo
```

#### func  ResolveMemoryInfo

```go
func ResolveMemoryInfo(id MemoryInfoId, d database.Database, l log.Logger) (res MemoryInfo, err error)
```
ResolveMemoryInfo loads and returns the MemoryInfo stored in the database d,
using id.

#### func (*MemoryInfo) Class

```go
func (*MemoryInfo) Class() binary.Class
```

#### func (*MemoryInfo) GetData

```go
func (c *MemoryInfo) GetData() []uint8
```

#### func (*MemoryInfo) GetObserved

```go
func (c *MemoryInfo) GetObserved() memory.RangeList
```

#### func (*MemoryInfo) GetReads

```go
func (c *MemoryInfo) GetReads() memory.RangeList
```

#### func (*MemoryInfo) GetWrites

```go
func (c *MemoryInfo) GetWrites() memory.RangeList
```

#### type MemoryInfoId

```go
type MemoryInfoId struct {
	binary.Generate
	ID binary.ID
}
```

Handle MemoryInfoId

#### func  StoreMemoryInfo

```go
func StoreMemoryInfo(v *MemoryInfo, d database.Database, l log.Logger) (MemoryInfoId, error)
```
StoreMemoryInfo stores v into the database d, returning the MemoryInfoId.

#### func (*MemoryInfoId) Class

```go
func (*MemoryInfoId) Class() binary.Class
```

#### func (MemoryInfoId) Valid

```go
func (h MemoryInfoId) Valid() bool
```

#### type RPC

```go
type RPC interface {
	GetSchema(l log.Logger) (Schema, error)
	Import(name string, Data []uint8, l log.Logger) (CaptureId, error)
	GetCaptures(l log.Logger) ([]CaptureId, error)
	GetDevices(l log.Logger) ([]DeviceId, error)
	GetHierarchy(capture CaptureId, l log.Logger) (HierarchyId, error)
	GetMemoryInfo(capture CaptureId, after uint64, rng memory.Range, l log.Logger) (MemoryInfoId, error)
	GetFramebufferColor(device DeviceId, capture CaptureId, api ApiId, after uint64, settings RenderSettings, l log.Logger) (ImageInfoId, error)
	GetFramebufferDepth(device DeviceId, capture CaptureId, api ApiId, after uint64, l log.Logger) (ImageInfoId, error)
	GetTimingInfo(device DeviceId, capture CaptureId, mask TimingMask, l log.Logger) (TimingInfoId, error)
	PrerenderFramebuffers(device DeviceId, capture CaptureId, api ApiId, width uint32, height uint32, atomIds []uint64, l log.Logger) (BinaryId, error)
	ReplaceAtom(capture CaptureId, atomId uint64, atom atom.Atom, l log.Logger) (CaptureId, error)
	Get(p path.Path, l log.Logger) (interface{}, error)
	ResolveAtomStream(id AtomStreamId, l log.Logger) (AtomStream, error)
	ResolveBinary(id BinaryId, l log.Logger) (Binary, error)
	ResolveCapture(id CaptureId, l log.Logger) (Capture, error)
	ResolveDevice(id DeviceId, l log.Logger) (Device, error)
	ResolveHierarchy(id HierarchyId, l log.Logger) (Hierarchy, error)
	ResolveImageInfo(id ImageInfoId, l log.Logger) (ImageInfo, error)
	ResolveMemoryInfo(id MemoryInfoId, l log.Logger) (MemoryInfo, error)
	ResolveReport(id ReportId, l log.Logger) (Report, error)
	ResolveTimingInfo(id TimingInfoId, l log.Logger) (TimingInfo, error)
}
```


#### type RenderSettings

```go
type RenderSettings struct {
	binary.Generate
	MaxWidth  uint32
	MaxHeight uint32
	Wireframe bool
}
```

Class RenderSettings

#### func  CreateRenderSettings

```go
func CreateRenderSettings(
	MaxWidth uint32,
	MaxHeight uint32,
	Wireframe bool,
) *RenderSettings
```

#### func (*RenderSettings) Class

```go
func (*RenderSettings) Class() binary.Class
```

#### func (*RenderSettings) GetMaxHeight

```go
func (c *RenderSettings) GetMaxHeight() uint32
```

#### func (*RenderSettings) GetMaxWidth

```go
func (c *RenderSettings) GetMaxWidth() uint32
```

#### func (*RenderSettings) GetWireframe

```go
func (c *RenderSettings) GetWireframe() bool
```

#### type Report

```go
type Report struct {
	binary.Generate
	Items []ReportItem
}
```

Class Report

#### func  CreateReport

```go
func CreateReport(
	Items []ReportItem,
) *Report
```

#### func  ResolveReport

```go
func ResolveReport(id ReportId, d database.Database, l log.Logger) (res Report, err error)
```
ResolveReport loads and returns the Report stored in the database d, using id.

#### func (*Report) Class

```go
func (*Report) Class() binary.Class
```

#### func (*Report) GetItems

```go
func (c *Report) GetItems() []ReportItem
```

#### type ReportId

```go
type ReportId struct {
	binary.Generate
	ID binary.ID
}
```

Handle ReportId

#### func  StoreReport

```go
func StoreReport(v *Report, d database.Database, l log.Logger) (ReportId, error)
```
StoreReport stores v into the database d, returning the ReportId.

#### func (*ReportId) Class

```go
func (*ReportId) Class() binary.Class
```

#### func (ReportId) Valid

```go
func (h ReportId) Valid() bool
```

#### type ReportItem

```go
type ReportItem struct {
	binary.Generate
	Severity Severity
	Message  string
	Atom     uint64
}
```

Class ReportItem

#### func  CreateReportItem

```go
func CreateReportItem(
	Severity Severity,
	Message string,
	Atom uint64,
) *ReportItem
```

#### func (*ReportItem) Class

```go
func (*ReportItem) Class() binary.Class
```

#### func (*ReportItem) GetAtom

```go
func (c *ReportItem) GetAtom() uint64
```

#### func (*ReportItem) GetMessage

```go
func (c *ReportItem) GetMessage() string
```

#### func (*ReportItem) GetSeverity

```go
func (c *ReportItem) GetSeverity() Severity
```

#### type Resolver

```go
type Resolver struct {
	Database database.Database
}
```


#### func (Resolver) ResolveAtomStream

```go
func (r Resolver) ResolveAtomStream(id AtomStreamId, l log.Logger) (AtomStream, error)
```
ResolveAtomStream loads and returns the AtomStream stored in the resolver's
database, using id.

#### func (Resolver) ResolveBinary

```go
func (r Resolver) ResolveBinary(id BinaryId, l log.Logger) (Binary, error)
```
ResolveBinary loads and returns the Binary stored in the resolver's database,
using id.

#### func (Resolver) ResolveCapture

```go
func (r Resolver) ResolveCapture(id CaptureId, l log.Logger) (Capture, error)
```
ResolveCapture loads and returns the Capture stored in the resolver's database,
using id.

#### func (Resolver) ResolveDevice

```go
func (r Resolver) ResolveDevice(id DeviceId, l log.Logger) (Device, error)
```
ResolveDevice loads and returns the Device stored in the resolver's database,
using id.

#### func (Resolver) ResolveHierarchy

```go
func (r Resolver) ResolveHierarchy(id HierarchyId, l log.Logger) (Hierarchy, error)
```
ResolveHierarchy loads and returns the Hierarchy stored in the resolver's
database, using id.

#### func (Resolver) ResolveImageInfo

```go
func (r Resolver) ResolveImageInfo(id ImageInfoId, l log.Logger) (ImageInfo, error)
```
ResolveImageInfo loads and returns the ImageInfo stored in the resolver's
database, using id.

#### func (Resolver) ResolveMemoryInfo

```go
func (r Resolver) ResolveMemoryInfo(id MemoryInfoId, l log.Logger) (MemoryInfo, error)
```
ResolveMemoryInfo loads and returns the MemoryInfo stored in the resolver's
database, using id.

#### func (Resolver) ResolveReport

```go
func (r Resolver) ResolveReport(id ReportId, l log.Logger) (Report, error)
```
ResolveReport loads and returns the Report stored in the resolver's database,
using id.

#### func (Resolver) ResolveTimingInfo

```go
func (r Resolver) ResolveTimingInfo(id TimingInfoId, l log.Logger) (TimingInfo, error)
```
ResolveTimingInfo loads and returns the TimingInfo stored in the resolver's
database, using id.

#### type Schema

```go
type Schema struct {
	binary.Generate
	Classes   []*schema.Class
	Constants []schema.ConstantSet
}
```

Class Schema

#### func  CreateSchema

```go
func CreateSchema(
	Classes []*schema.Class,
	Constants []schema.ConstantSet,
) *Schema
```

#### func (*Schema) Class

```go
func (*Schema) Class() binary.Class
```

#### func (*Schema) GetClasses

```go
func (c *Schema) GetClasses() []*schema.Class
```

#### func (*Schema) GetConstants

```go
func (c *Schema) GetConstants() []schema.ConstantSet
```

#### type Severity

```go
type Severity int
```

Enum Severity

```go
const (
	SeverityEmergency     Severity = 0
	SeverityAlert         Severity = 1
	SeverityCritical      Severity = 2
	SeverityError         Severity = 3
	SeverityWarning       Severity = 4
	SeverityNotice        Severity = 5
	SeverityInformational Severity = 6
	SeverityDebug         Severity = 7
)
```

#### func (Severity) IsAlert

```go
func (i Severity) IsAlert() bool
```

#### func (Severity) IsCritical

```go
func (i Severity) IsCritical() bool
```

#### func (Severity) IsDebug

```go
func (i Severity) IsDebug() bool
```

#### func (Severity) IsEmergency

```go
func (i Severity) IsEmergency() bool
```

#### func (Severity) IsError

```go
func (i Severity) IsError() bool
```

#### func (Severity) IsInformational

```go
func (i Severity) IsInformational() bool
```

#### func (Severity) IsNotice

```go
func (i Severity) IsNotice() bool
```

#### func (Severity) IsWarning

```go
func (i Severity) IsWarning() bool
```

#### func (*Severity) Parse

```go
func (v *Severity) Parse(s string) error
```

#### func (Severity) String

```go
func (v Severity) String() string
```

#### type TimingInfo

```go
type TimingInfo struct {
	binary.Generate
	PerCommand  []AtomTimer
	PerDrawCall []AtomRangeTimer
	PerFrame    []AtomRangeTimer
}
```

Class TimingInfo

#### func  CreateTimingInfo

```go
func CreateTimingInfo(
	PerCommand []AtomTimer,
	PerDrawCall []AtomRangeTimer,
	PerFrame []AtomRangeTimer,
) *TimingInfo
```

#### func  ResolveTimingInfo

```go
func ResolveTimingInfo(id TimingInfoId, d database.Database, l log.Logger) (res TimingInfo, err error)
```
ResolveTimingInfo loads and returns the TimingInfo stored in the database d,
using id.

#### func (*TimingInfo) Class

```go
func (*TimingInfo) Class() binary.Class
```

#### func (*TimingInfo) GetPerCommand

```go
func (c *TimingInfo) GetPerCommand() []AtomTimer
```

#### func (*TimingInfo) GetPerDrawCall

```go
func (c *TimingInfo) GetPerDrawCall() []AtomRangeTimer
```

#### func (*TimingInfo) GetPerFrame

```go
func (c *TimingInfo) GetPerFrame() []AtomRangeTimer
```

#### type TimingInfoId

```go
type TimingInfoId struct {
	binary.Generate
	ID binary.ID
}
```

Handle TimingInfoId

#### func  StoreTimingInfo

```go
func StoreTimingInfo(v *TimingInfo, d database.Database, l log.Logger) (TimingInfoId, error)
```
StoreTimingInfo stores v into the database d, returning the TimingInfoId.

#### func (*TimingInfoId) Class

```go
func (*TimingInfoId) Class() binary.Class
```

#### func (TimingInfoId) Valid

```go
func (h TimingInfoId) Valid() bool
```

#### type TimingMask

```go
type TimingMask int
```

Enum TimingMask

```go
const (
	TimingMaskTimingPerCommand  TimingMask = 1
	TimingMaskTimingPerDrawCall TimingMask = 2
	TimingMaskTimingPerFrame    TimingMask = 4
)
```

#### func (TimingMask) IsTimingPerCommand

```go
func (i TimingMask) IsTimingPerCommand() bool
```

#### func (TimingMask) IsTimingPerDrawCall

```go
func (i TimingMask) IsTimingPerDrawCall() bool
```

#### func (TimingMask) IsTimingPerFrame

```go
func (i TimingMask) IsTimingPerFrame() bool
```

#### func (*TimingMask) Parse

```go
func (v *TimingMask) Parse(s string) error
```

#### func (TimingMask) String

```go
func (v TimingMask) String() string
```
