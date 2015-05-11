# service
--
    import "android.googlesource.com/platform/tools/gpu/service"

Package service is the definition of the RPC GPU debugger service exposed by the
server.

It is not the actual implementation of the service functionality.

## Usage

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

#### type ApiIdArray

```go
type ApiIdArray []ApiId
```

Array ApiIdArray

#### func (ApiIdArray) Format

```go
func (a ApiIdArray) Format(f fmt.State, c rune)
```

#### type ApiSchema

```go
type ApiSchema struct {
	binary.Generate
	Api   ApiId
	State StructInfo
}
```

Class ApiSchema

#### func  CreateApiSchema

```go
func CreateApiSchema(
	Api ApiId,
	State StructInfo,
) *ApiSchema
```

#### func (*ApiSchema) Class

```go
func (*ApiSchema) Class() binary.Class
```

#### func (*ApiSchema) GetApi

```go
func (c *ApiSchema) GetApi() ApiId
```

#### func (*ApiSchema) GetState

```go
func (c *ApiSchema) GetState() StructInfo
```

#### type ApiSchemaArray

```go
type ApiSchemaArray []ApiSchema
```

Array ApiSchemaArray

#### func (ApiSchemaArray) Format

```go
func (a ApiSchemaArray) Format(f fmt.State, c rune)
```

#### type ArrayInfo

```go
type ArrayInfo struct {
	binary.Generate
	Name        string
	Kind        TypeKind
	ElementType TypeInfo
}
```

Class ArrayInfo

#### func  CreateArrayInfo

```go
func CreateArrayInfo(
	Name string,
	Kind TypeKind,
	ElementType TypeInfo,
) *ArrayInfo
```

#### func (*ArrayInfo) Class

```go
func (*ArrayInfo) Class() binary.Class
```

#### func (*ArrayInfo) GetElementType

```go
func (c *ArrayInfo) GetElementType() TypeInfo
```

#### func (*ArrayInfo) GetKind

```go
func (c *ArrayInfo) GetKind() TypeKind
```

#### func (*ArrayInfo) GetName

```go
func (c *ArrayInfo) GetName() string
```

#### type AtomGroup

```go
type AtomGroup struct {
	binary.Generate
	Name      string
	Range     AtomRange
	SubGroups AtomGroupArray
}
```

Class AtomGroup

#### func  CreateAtomGroup

```go
func CreateAtomGroup(
	Name string,
	Range AtomRange,
	SubGroups AtomGroupArray,
) *AtomGroup
```

#### func (*AtomGroup) Class

```go
func (*AtomGroup) Class() binary.Class
```

#### func (*AtomGroup) GetName

```go
func (c *AtomGroup) GetName() string
```

#### func (*AtomGroup) GetRange

```go
func (c *AtomGroup) GetRange() AtomRange
```

#### func (*AtomGroup) GetSubGroups

```go
func (c *AtomGroup) GetSubGroups() AtomGroupArray
```

#### func (*AtomGroup) Pack

```go
func (g *AtomGroup) Pack(o atom.Group)
```
Pack packs the atom Group o into the RPC-friendly AtomGroup structure.

#### func (AtomGroup) Unpack

```go
func (g AtomGroup) Unpack(o *atom.Group)
```
Unpack unpacks the RPC-friendly AtomGroup structure into the atom Group o.

#### type AtomGroupArray

```go
type AtomGroupArray []AtomGroup
```

Array AtomGroupArray

#### func (AtomGroupArray) Format

```go
func (a AtomGroupArray) Format(f fmt.State, c rune)
```

#### type AtomInfo

```go
type AtomInfo struct {
	binary.Generate
	Api              ApiId
	Type             uint16
	Name             string
	Parameters       ParameterInfoArray
	IsCommand        bool
	IsDrawCall       bool
	IsEndOfFrame     bool
	DocumentationUrl string
}
```

Class AtomInfo

#### func  CreateAtomInfo

```go
func CreateAtomInfo(
	Api ApiId,
	Type uint16,
	Name string,
	Parameters ParameterInfoArray,
	IsCommand bool,
	IsDrawCall bool,
	IsEndOfFrame bool,
	DocumentationUrl string,
) *AtomInfo
```

#### func (*AtomInfo) Class

```go
func (*AtomInfo) Class() binary.Class
```

#### func (*AtomInfo) GetApi

```go
func (c *AtomInfo) GetApi() ApiId
```

#### func (*AtomInfo) GetDocumentationUrl

```go
func (c *AtomInfo) GetDocumentationUrl() string
```

#### func (*AtomInfo) GetIsCommand

```go
func (c *AtomInfo) GetIsCommand() bool
```

#### func (*AtomInfo) GetIsDrawCall

```go
func (c *AtomInfo) GetIsDrawCall() bool
```

#### func (*AtomInfo) GetIsEndOfFrame

```go
func (c *AtomInfo) GetIsEndOfFrame() bool
```

#### func (*AtomInfo) GetName

```go
func (c *AtomInfo) GetName() string
```

#### func (*AtomInfo) GetParameters

```go
func (c *AtomInfo) GetParameters() ParameterInfoArray
```

#### func (*AtomInfo) GetType

```go
func (c *AtomInfo) GetType() uint16
```

#### type AtomInfoArray

```go
type AtomInfoArray []AtomInfo
```

Array AtomInfoArray

#### func (AtomInfoArray) Format

```go
func (a AtomInfoArray) Format(f fmt.State, c rune)
```

#### type AtomRange

```go
type AtomRange struct {
	binary.Generate
	First uint64
	Count uint64
}
```

Class AtomRange

#### func  CreateAtomRange

```go
func CreateAtomRange(
	First uint64,
	Count uint64,
) *AtomRange
```

#### func (*AtomRange) Class

```go
func (*AtomRange) Class() binary.Class
```

#### func (*AtomRange) GetCount

```go
func (c *AtomRange) GetCount() uint64
```

#### func (*AtomRange) GetFirst

```go
func (c *AtomRange) GetFirst() uint64
```

#### func (*AtomRange) Pack

```go
func (r *AtomRange) Pack(o atom.Range)
```
Pack packs the atom Range o into the RPC-friendly AtomRange structure.

#### func (AtomRange) Unpack

```go
func (r AtomRange) Unpack(o *atom.Range)
```
Unpack unpacks the RPC-friendly AtomRange structure into the atom Range o.

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

#### type AtomRangeTimerArray

```go
type AtomRangeTimerArray []AtomRangeTimer
```

Array AtomRangeTimerArray

#### func (AtomRangeTimerArray) Format

```go
func (a AtomRangeTimerArray) Format(f fmt.State, c rune)
```

#### type AtomStream

```go
type AtomStream struct {
	binary.Generate
	Data U8Array
}
```

Class AtomStream

#### func  CreateAtomStream

```go
func CreateAtomStream(
	Data U8Array,
) *AtomStream
```

#### func  NewAtomStream

```go
func NewAtomStream(list atom.List) (AtomStream, error)
```
NewAtomStream creates a fully-encoded AtomStream from the atom list.

#### func (*AtomStream) Class

```go
func (*AtomStream) Class() binary.Class
```

#### func (*AtomStream) GetData

```go
func (c *AtomStream) GetData() U8Array
```

#### func (AtomStream) List

```go
func (s AtomStream) List() (atom.List, error)
```
List decodes and returns the AtomStream decoded to an atom list.

#### type AtomStreamId

```go
type AtomStreamId struct {
	binary.Generate
	ID binary.ID
}
```

Handle AtomStreamId

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

#### type AtomTimerArray

```go
type AtomTimerArray []AtomTimer
```

Array AtomTimerArray

#### func (AtomTimerArray) Format

```go
func (a AtomTimerArray) Format(f fmt.State, c rune)
```

#### type Binary

```go
type Binary struct {
	binary.Generate
	Data U8Array
}
```

Class Binary

#### func  CreateBinary

```go
func CreateBinary(
	Data U8Array,
) *Binary
```

#### func (*Binary) Class

```go
func (*Binary) Class() binary.Class
```

#### func (*Binary) GetData

```go
func (c *Binary) GetData() U8Array
```

#### type BinaryId

```go
type BinaryId struct {
	binary.Generate
	ID binary.ID
}
```

Handle BinaryId

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
	Apis   ApiIdArray
	Schema SchemaId
}
```

Class Capture

#### func  CreateCapture

```go
func CreateCapture(
	Name string,
	Atoms AtomStreamId,
	Apis ApiIdArray,
	Schema SchemaId,
) *Capture
```

#### func (*Capture) Class

```go
func (*Capture) Class() binary.Class
```

#### func (*Capture) GetApis

```go
func (c *Capture) GetApis() ApiIdArray
```

#### func (*Capture) GetAtoms

```go
func (c *Capture) GetAtoms() AtomStreamId
```

#### func (*Capture) GetName

```go
func (c *Capture) GetName() string
```

#### func (*Capture) GetSchema

```go
func (c *Capture) GetSchema() SchemaId
```

#### type CaptureId

```go
type CaptureId struct {
	binary.Generate
	ID binary.ID
}
```

Handle CaptureId

#### func (*CaptureId) Class

```go
func (*CaptureId) Class() binary.Class
```

#### func (CaptureId) Valid

```go
func (h CaptureId) Valid() bool
```

#### type CaptureIdArray

```go
type CaptureIdArray []CaptureId
```

Array CaptureIdArray

#### func (CaptureIdArray) Format

```go
func (a CaptureIdArray) Format(f fmt.State, c rune)
```

#### type ClassInfo

```go
type ClassInfo struct {
	binary.Generate
	Name    string
	Kind    TypeKind
	Fields  FieldInfoArray
	Extends ClassInfoArray
}
```

Class ClassInfo

#### func  CreateClassInfo

```go
func CreateClassInfo(
	Name string,
	Kind TypeKind,
	Fields FieldInfoArray,
	Extends ClassInfoArray,
) *ClassInfo
```

#### func (*ClassInfo) Class

```go
func (*ClassInfo) Class() binary.Class
```

#### func (*ClassInfo) GetExtends

```go
func (c *ClassInfo) GetExtends() ClassInfoArray
```

#### func (*ClassInfo) GetFields

```go
func (c *ClassInfo) GetFields() FieldInfoArray
```

#### func (*ClassInfo) GetKind

```go
func (c *ClassInfo) GetKind() TypeKind
```

#### func (*ClassInfo) GetName

```go
func (c *ClassInfo) GetName() string
```

#### type ClassInfoArray

```go
type ClassInfoArray []*ClassInfo
```

Array ClassInfoRefArray

#### func (ClassInfoArray) Format

```go
func (a ClassInfoArray) Format(f fmt.State, c rune)
```

#### type Device

```go
type Device struct {
	binary.Generate
	Name                   string
	Model                  string
	OS                     string
	PointerSize            uint8
	PointerAlignment       uint8
	MaxMemorySize          uint64
	RequiresShaderPatching bool
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
	RequiresShaderPatching bool,
) *Device
```

#### func (*Device) Class

```go
func (*Device) Class() binary.Class
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

#### func (*Device) GetRequiresShaderPatching

```go
func (c *Device) GetRequiresShaderPatching() bool
```

#### type DeviceId

```go
type DeviceId struct {
	binary.Generate
	ID binary.ID
}
```

Handle DeviceId

#### func (*DeviceId) Class

```go
func (*DeviceId) Class() binary.Class
```

#### func (DeviceId) Valid

```go
func (h DeviceId) Valid() bool
```

#### type DeviceIdArray

```go
type DeviceIdArray []DeviceId
```

Array DeviceIdArray

#### func (DeviceIdArray) Format

```go
func (a DeviceIdArray) Format(f fmt.State, c rune)
```

#### type EnumEntry

```go
type EnumEntry struct {
	binary.Generate
	Name  string
	Value uint32
}
```

Class EnumEntry

#### func  CreateEnumEntry

```go
func CreateEnumEntry(
	Name string,
	Value uint32,
) *EnumEntry
```

#### func (*EnumEntry) Class

```go
func (*EnumEntry) Class() binary.Class
```

#### func (*EnumEntry) GetName

```go
func (c *EnumEntry) GetName() string
```

#### func (*EnumEntry) GetValue

```go
func (c *EnumEntry) GetValue() uint32
```

#### type EnumEntryArray

```go
type EnumEntryArray []EnumEntry
```

Array EnumEntryArray

#### func (EnumEntryArray) Format

```go
func (a EnumEntryArray) Format(f fmt.State, c rune)
```

#### type EnumInfo

```go
type EnumInfo struct {
	binary.Generate
	Name    string
	Kind    TypeKind
	Entries EnumEntryArray
	Extends EnumInfoArray
}
```

Class EnumInfo

#### func  CreateEnumInfo

```go
func CreateEnumInfo(
	Name string,
	Kind TypeKind,
	Entries EnumEntryArray,
	Extends EnumInfoArray,
) *EnumInfo
```

#### func (*EnumInfo) Class

```go
func (*EnumInfo) Class() binary.Class
```

#### func (*EnumInfo) GetEntries

```go
func (c *EnumInfo) GetEntries() EnumEntryArray
```

#### func (*EnumInfo) GetExtends

```go
func (c *EnumInfo) GetExtends() EnumInfoArray
```

#### func (*EnumInfo) GetKind

```go
func (c *EnumInfo) GetKind() TypeKind
```

#### func (*EnumInfo) GetName

```go
func (c *EnumInfo) GetName() string
```

#### type EnumInfoArray

```go
type EnumInfoArray []*EnumInfo
```

Array EnumInfoRefArray

#### func (EnumInfoArray) Format

```go
func (a EnumInfoArray) Format(f fmt.State, c rune)
```

#### type FieldInfo

```go
type FieldInfo struct {
	binary.Generate
	Name string
	Type TypeInfo
}
```

Class FieldInfo

#### func  CreateFieldInfo

```go
func CreateFieldInfo(
	Name string,
	Type TypeInfo,
) *FieldInfo
```

#### func (*FieldInfo) Class

```go
func (*FieldInfo) Class() binary.Class
```

#### func (*FieldInfo) GetName

```go
func (c *FieldInfo) GetName() string
```

#### func (*FieldInfo) GetType

```go
func (c *FieldInfo) GetType() TypeInfo
```

#### type FieldInfoArray

```go
type FieldInfoArray []*FieldInfo
```

Array FieldInfoRefArray

#### func (FieldInfoArray) Format

```go
func (a FieldInfoArray) Format(f fmt.State, c rune)
```

#### type Hierarchy

```go
type Hierarchy struct {
	binary.Generate
	Root AtomGroup
}
```

Class Hierarchy

#### func  CreateHierarchy

```go
func CreateHierarchy(
	Root AtomGroup,
) *Hierarchy
```

#### func (*Hierarchy) Class

```go
func (*Hierarchy) Class() binary.Class
```

#### func (*Hierarchy) GetRoot

```go
func (c *Hierarchy) GetRoot() AtomGroup
```

#### type HierarchyId

```go
type HierarchyId struct {
	binary.Generate
	ID binary.ID
}
```

Handle HierarchyId

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

#### func (ImageFormat) String

```go
func (i ImageFormat) String() string
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

#### func (*ImageInfoId) Class

```go
func (*ImageInfoId) Class() binary.Class
```

#### func (ImageInfoId) Valid

```go
func (h ImageInfoId) Valid() bool
```

#### type MapInfo

```go
type MapInfo struct {
	binary.Generate
	Name      string
	Kind      TypeKind
	KeyType   TypeInfo
	ValueType TypeInfo
}
```

Class MapInfo

#### func  CreateMapInfo

```go
func CreateMapInfo(
	Name string,
	Kind TypeKind,
	KeyType TypeInfo,
	ValueType TypeInfo,
) *MapInfo
```

#### func (*MapInfo) Class

```go
func (*MapInfo) Class() binary.Class
```

#### func (*MapInfo) GetKeyType

```go
func (c *MapInfo) GetKeyType() TypeInfo
```

#### func (*MapInfo) GetKind

```go
func (c *MapInfo) GetKind() TypeKind
```

#### func (*MapInfo) GetName

```go
func (c *MapInfo) GetName() string
```

#### func (*MapInfo) GetValueType

```go
func (c *MapInfo) GetValueType() TypeInfo
```

#### type MemoryInfo

```go
type MemoryInfo struct {
	binary.Generate
	Data    U8Array
	Stale   MemoryRangeArray
	Current MemoryRangeArray
	Unknown MemoryRangeArray
}
```

Class MemoryInfo

#### func  CreateMemoryInfo

```go
func CreateMemoryInfo(
	Data U8Array,
	Stale MemoryRangeArray,
	Current MemoryRangeArray,
	Unknown MemoryRangeArray,
) *MemoryInfo
```

#### func (*MemoryInfo) Class

```go
func (*MemoryInfo) Class() binary.Class
```

#### func (*MemoryInfo) GetCurrent

```go
func (c *MemoryInfo) GetCurrent() MemoryRangeArray
```

#### func (*MemoryInfo) GetData

```go
func (c *MemoryInfo) GetData() U8Array
```

#### func (*MemoryInfo) GetStale

```go
func (c *MemoryInfo) GetStale() MemoryRangeArray
```

#### func (*MemoryInfo) GetUnknown

```go
func (c *MemoryInfo) GetUnknown() MemoryRangeArray
```

#### type MemoryInfoId

```go
type MemoryInfoId struct {
	binary.Generate
	ID binary.ID
}
```

Handle MemoryInfoId

#### func (*MemoryInfoId) Class

```go
func (*MemoryInfoId) Class() binary.Class
```

#### func (MemoryInfoId) Valid

```go
func (h MemoryInfoId) Valid() bool
```

#### type MemoryRange

```go
type MemoryRange struct {
	binary.Generate
	Base uint64
	Size uint64
}
```

Class MemoryRange

#### func  CreateMemoryRange

```go
func CreateMemoryRange(
	Base uint64,
	Size uint64,
) *MemoryRange
```

#### func (*MemoryRange) Class

```go
func (*MemoryRange) Class() binary.Class
```

#### func (*MemoryRange) GetBase

```go
func (c *MemoryRange) GetBase() uint64
```

#### func (*MemoryRange) GetSize

```go
func (c *MemoryRange) GetSize() uint64
```

#### func (*MemoryRange) Pack

```go
func (r *MemoryRange) Pack(o memory.Range)
```
Pack packs the memory Range o into the RPC-friendly MemoryRange structure.

#### func (MemoryRange) Unpack

```go
func (r MemoryRange) Unpack(o *memory.Range)
```
Unpack unpacks the RPC-friendly MemoryRange structure into the memory Range o.

#### type MemoryRangeArray

```go
type MemoryRangeArray []MemoryRange
```

Array MemoryRangeArray

#### func (MemoryRangeArray) Format

```go
func (a MemoryRangeArray) Format(f fmt.State, c rune)
```

#### func (*MemoryRangeArray) Pack

```go
func (l *MemoryRangeArray) Pack(o memory.RangeList)
```
Pack packs the memory RangeList o into the RPC-friendly MemoryRangeArray
structure.

#### func (MemoryRangeArray) Unpack

```go
func (l MemoryRangeArray) Unpack(o *memory.RangeList)
```
Unpack unpacks the RPC-friendly MemoryRangeArray structure into the memory
RangeList o.

#### type ParameterInfo

```go
type ParameterInfo struct {
	binary.Generate
	Name string
	Type TypeInfo
	Out  bool
}
```

Class ParameterInfo

#### func  CreateParameterInfo

```go
func CreateParameterInfo(
	Name string,
	Type TypeInfo,
	Out bool,
) *ParameterInfo
```

#### func (*ParameterInfo) Class

```go
func (*ParameterInfo) Class() binary.Class
```

#### func (*ParameterInfo) GetName

```go
func (c *ParameterInfo) GetName() string
```

#### func (*ParameterInfo) GetOut

```go
func (c *ParameterInfo) GetOut() bool
```

#### func (*ParameterInfo) GetType

```go
func (c *ParameterInfo) GetType() TypeInfo
```

#### type ParameterInfoArray

```go
type ParameterInfoArray []ParameterInfo
```

Array ParameterInfoArray

#### func (ParameterInfoArray) Format

```go
func (a ParameterInfoArray) Format(f fmt.State, c rune)
```

#### type RPC

```go
type RPC interface {
	Import(l log.Logger, name string, Data U8Array) (CaptureId, error)
	GetCaptures(l log.Logger) (CaptureIdArray, error)
	GetDevices(l log.Logger) (DeviceIdArray, error)
	GetState(l log.Logger, capture CaptureId, after uint64) (BinaryId, error)
	GetHierarchy(l log.Logger, capture CaptureId) (HierarchyId, error)
	GetMemoryInfo(l log.Logger, capture CaptureId, after uint64, rng MemoryRange) (MemoryInfoId, error)
	GetFramebufferColor(l log.Logger, device DeviceId, capture CaptureId, api ApiId, after uint64, settings RenderSettings) (ImageInfoId, error)
	GetFramebufferDepth(l log.Logger, device DeviceId, capture CaptureId, api ApiId, after uint64) (ImageInfoId, error)
	GetTimingInfo(l log.Logger, device DeviceId, capture CaptureId, mask TimingMask) (TimingInfoId, error)
	PrerenderFramebuffers(l log.Logger, device DeviceId, capture CaptureId, api ApiId, width uint32, height uint32, atomIds U64Array) (BinaryId, error)
	ReplaceAtom(l log.Logger, capture CaptureId, atomId uint64, atomType uint16, data Binary) (CaptureId, error)
	ResolveAtomStream(l log.Logger, id AtomStreamId) (AtomStream, error)
	ResolveBinary(l log.Logger, id BinaryId) (Binary, error)
	ResolveCapture(l log.Logger, id CaptureId) (Capture, error)
	ResolveDevice(l log.Logger, id DeviceId) (Device, error)
	ResolveHierarchy(l log.Logger, id HierarchyId) (Hierarchy, error)
	ResolveImageInfo(l log.Logger, id ImageInfoId) (ImageInfo, error)
	ResolveMemoryInfo(l log.Logger, id MemoryInfoId) (MemoryInfo, error)
	ResolveSchema(l log.Logger, id SchemaId) (Schema, error)
	ResolveTimingInfo(l log.Logger, id TimingInfoId) (TimingInfo, error)
}
```


#### func  CreateClient

```go
func CreateClient(r io.Reader, w io.Writer, mtu int) RPC
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

#### type Schema

```go
type Schema struct {
	binary.Generate
	Atoms AtomInfoArray
	Apis  ApiSchemaArray
}
```

Class Schema

#### func  CreateSchema

```go
func CreateSchema(
	Atoms AtomInfoArray,
	Apis ApiSchemaArray,
) *Schema
```

#### func (*Schema) Class

```go
func (*Schema) Class() binary.Class
```

#### func (*Schema) GetApis

```go
func (c *Schema) GetApis() ApiSchemaArray
```

#### func (*Schema) GetAtoms

```go
func (c *Schema) GetAtoms() AtomInfoArray
```

#### type SchemaId

```go
type SchemaId struct {
	binary.Generate
	ID binary.ID
}
```

Handle SchemaId

#### func (*SchemaId) Class

```go
func (*SchemaId) Class() binary.Class
```

#### func (SchemaId) Valid

```go
func (h SchemaId) Valid() bool
```

#### type SimpleInfo

```go
type SimpleInfo struct {
	binary.Generate
	Name string
	Kind TypeKind
}
```

Class SimpleInfo

#### func  CreateSimpleInfo

```go
func CreateSimpleInfo(
	Name string,
	Kind TypeKind,
) *SimpleInfo
```

#### func (*SimpleInfo) Class

```go
func (*SimpleInfo) Class() binary.Class
```

#### func (*SimpleInfo) GetKind

```go
func (c *SimpleInfo) GetKind() TypeKind
```

#### func (*SimpleInfo) GetName

```go
func (c *SimpleInfo) GetName() string
```

#### type StaticArrayInfo

```go
type StaticArrayInfo struct {
	binary.Generate
	Name        string
	Kind        TypeKind
	ElementType TypeInfo
	Size        uint32
}
```

Class StaticArrayInfo

#### func  CreateStaticArrayInfo

```go
func CreateStaticArrayInfo(
	Name string,
	Kind TypeKind,
	ElementType TypeInfo,
	Size uint32,
) *StaticArrayInfo
```

#### func (*StaticArrayInfo) Class

```go
func (*StaticArrayInfo) Class() binary.Class
```

#### func (*StaticArrayInfo) GetElementType

```go
func (c *StaticArrayInfo) GetElementType() TypeInfo
```

#### func (*StaticArrayInfo) GetKind

```go
func (c *StaticArrayInfo) GetKind() TypeKind
```

#### func (*StaticArrayInfo) GetName

```go
func (c *StaticArrayInfo) GetName() string
```

#### func (*StaticArrayInfo) GetSize

```go
func (c *StaticArrayInfo) GetSize() uint32
```

#### type StructInfo

```go
type StructInfo struct {
	binary.Generate
	Name   string
	Kind   TypeKind
	Fields FieldInfoArray
}
```

Class StructInfo

#### func  CreateStructInfo

```go
func CreateStructInfo(
	Name string,
	Kind TypeKind,
	Fields FieldInfoArray,
) *StructInfo
```

#### func (*StructInfo) Class

```go
func (*StructInfo) Class() binary.Class
```

#### func (*StructInfo) GetFields

```go
func (c *StructInfo) GetFields() FieldInfoArray
```

#### func (*StructInfo) GetKind

```go
func (c *StructInfo) GetKind() TypeKind
```

#### func (*StructInfo) GetName

```go
func (c *StructInfo) GetName() string
```

#### type TimingInfo

```go
type TimingInfo struct {
	binary.Generate
	PerCommand  AtomTimerArray
	PerDrawCall AtomRangeTimerArray
	PerFrame    AtomRangeTimerArray
}
```

Class TimingInfo

#### func  CreateTimingInfo

```go
func CreateTimingInfo(
	PerCommand AtomTimerArray,
	PerDrawCall AtomRangeTimerArray,
	PerFrame AtomRangeTimerArray,
) *TimingInfo
```

#### func (*TimingInfo) Class

```go
func (*TimingInfo) Class() binary.Class
```

#### func (*TimingInfo) GetPerCommand

```go
func (c *TimingInfo) GetPerCommand() AtomTimerArray
```

#### func (*TimingInfo) GetPerDrawCall

```go
func (c *TimingInfo) GetPerDrawCall() AtomRangeTimerArray
```

#### func (*TimingInfo) GetPerFrame

```go
func (c *TimingInfo) GetPerFrame() AtomRangeTimerArray
```

#### type TimingInfoId

```go
type TimingInfoId struct {
	binary.Generate
	ID binary.ID
}
```

Handle TimingInfoId

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

#### func (TimingMask) String

```go
func (i TimingMask) String() string
```

#### type TypeInfo

```go
type TypeInfo interface {
	binary.Object
	GetName() string
	GetKind() TypeKind
}
```

Interface TypeInfo

#### type TypeKind

```go
type TypeKind int
```

Enum TypeKind

```go
const (
	TypeKindBool        TypeKind = 0
	TypeKindS8          TypeKind = 1
	TypeKindU8          TypeKind = 2
	TypeKindS16         TypeKind = 3
	TypeKindU16         TypeKind = 4
	TypeKindS32         TypeKind = 5
	TypeKindU32         TypeKind = 6
	TypeKindF32         TypeKind = 7
	TypeKindS64         TypeKind = 8
	TypeKindU64         TypeKind = 9
	TypeKindF64         TypeKind = 10
	TypeKindString      TypeKind = 11
	TypeKindEnum        TypeKind = 12
	TypeKindStruct      TypeKind = 14
	TypeKindClass       TypeKind = 15
	TypeKindArray       TypeKind = 16
	TypeKindStaticArray TypeKind = 17
	TypeKindMap         TypeKind = 18
	TypeKindPointer     TypeKind = 19
	TypeKindMemory      TypeKind = 20
	TypeKindAny         TypeKind = 21
	TypeKindID          TypeKind = 22
)
```

#### func (TypeKind) IsAny

```go
func (i TypeKind) IsAny() bool
```

#### func (TypeKind) IsArray

```go
func (i TypeKind) IsArray() bool
```

#### func (TypeKind) IsBool

```go
func (i TypeKind) IsBool() bool
```

#### func (TypeKind) IsClass

```go
func (i TypeKind) IsClass() bool
```

#### func (TypeKind) IsEnum

```go
func (i TypeKind) IsEnum() bool
```

#### func (TypeKind) IsF32

```go
func (i TypeKind) IsF32() bool
```

#### func (TypeKind) IsF64

```go
func (i TypeKind) IsF64() bool
```

#### func (TypeKind) IsID

```go
func (i TypeKind) IsID() bool
```

#### func (TypeKind) IsMap

```go
func (i TypeKind) IsMap() bool
```

#### func (TypeKind) IsMemory

```go
func (i TypeKind) IsMemory() bool
```

#### func (TypeKind) IsPointer

```go
func (i TypeKind) IsPointer() bool
```

#### func (TypeKind) IsS16

```go
func (i TypeKind) IsS16() bool
```

#### func (TypeKind) IsS32

```go
func (i TypeKind) IsS32() bool
```

#### func (TypeKind) IsS64

```go
func (i TypeKind) IsS64() bool
```

#### func (TypeKind) IsS8

```go
func (i TypeKind) IsS8() bool
```

#### func (TypeKind) IsStaticArray

```go
func (i TypeKind) IsStaticArray() bool
```

#### func (TypeKind) IsString

```go
func (i TypeKind) IsString() bool
```

#### func (TypeKind) IsStruct

```go
func (i TypeKind) IsStruct() bool
```

#### func (TypeKind) IsU16

```go
func (i TypeKind) IsU16() bool
```

#### func (TypeKind) IsU32

```go
func (i TypeKind) IsU32() bool
```

#### func (TypeKind) IsU64

```go
func (i TypeKind) IsU64() bool
```

#### func (TypeKind) IsU8

```go
func (i TypeKind) IsU8() bool
```

#### func (TypeKind) String

```go
func (i TypeKind) String() string
```

#### type U64Array

```go
type U64Array []uint64
```

Array U64Array

#### func (U64Array) Format

```go
func (a U64Array) Format(f fmt.State, c rune)
```

#### type U8Array

```go
type U8Array []uint8
```

Array U8Array

#### func (U8Array) Format

```go
func (a U8Array) Format(f fmt.State, c rune)
```
