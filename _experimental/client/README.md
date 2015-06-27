# client
--
    import "android.googlesource.com/platform/tools/gpu/_experimental/client"


## Usage

```go
const InvalidAtomID = ^atom.ID(0)
```

```go
var (
	LINE_NUMBER_COLOR        gxui.Color = gxui.ColorFromHex(0xFF1CAFFF)
	CODE_COLOR               gxui.Color = gxui.ColorFromHex(0xFFFFE8BB)
	COMMAND_COLOR            gxui.Color = gxui.ColorFromHex(0xFFFB9868)
	MEMORY_OBSERVATION_COLOR gxui.Color = gxui.ColorFromHex(0xFFA1CF8A)
	INACTIVE_COLOR           gxui.Color = gxui.ColorFromHex(0xFF505050)
	CONSTANT_COLOR           gxui.Color = gxui.ColorFromHex(0xFFDDB7FF)
	READ_MEMORY_COLOR        gxui.Color = gxui.Green
	WRITE_MEMORY_COLOR       gxui.Color = gxui.Red
	STALE_MEMORY_COLOR       gxui.Color = gxui.ColorFromHex(0xFF8A7753)
)
```

#### func  AddTextToolTip

```go
func AddTextToolTip(appCtx *ApplicationContext, target gxui.Control, text string)
```

#### func  CreateColorBufferPanel

```go
func CreateColorBufferPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateCommandsPanel

```go
func CreateCommandsPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateDepthBufferPanel

```go
func CreateDepthBufferPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateDocsPanel

```go
func CreateDocsPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateFramesPanel

```go
func CreateFramesPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateGotoCommandDialog

```go
func CreateGotoCommandDialog(appCtx *ApplicationContext) gxui.Window
```

#### func  CreateGxuiDebug

```go
func CreateGxuiDebug(appCtx *ApplicationContext, window gxui.Window, driver gxui.Driver) gxui.Control
```

#### func  CreateImageViewerPanel

```go
func CreateImageViewerPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateLabel

```go
func CreateLabel(t gxui.Theme, s string, c gxui.Color, active bool) gxui.Label
```

#### func  CreateLaunchAndroidDialog

```go
func CreateLaunchAndroidDialog(theme gxui.Theme, statusLogger log.Logger, capture func())
```

#### func  CreateLogPanel

```go
func CreateLogPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateMemoryPanel

```go
func CreateMemoryPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateMemoryPanelReinterpretButtons

```go
func CreateMemoryPanelReinterpretButtons(
	appCtx *ApplicationContext,
	memory_list gxui.List,
	rawAdapter *MemoryAdapter,
	imgAdapter *MemoryImageAdapter) gxui.Control
```

#### func  CreateMonospaceLabel

```go
func CreateMonospaceLabel(appCtx *ApplicationContext, s string, c gxui.Color, active bool) gxui.Label
```

#### func  CreateProfilerPanel

```go
func CreateProfilerPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateReportPanel

```go
func CreateReportPanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateStatePanel

```go
func CreateStatePanel(appCtx *ApplicationContext) gxui.Control
```

#### func  CreateTakeCaptureDialog

```go
func CreateTakeCaptureDialog(appCtx *ApplicationContext)
```

#### func  CreateWireframeButton

```go
func CreateWireframeButton(appCtx *ApplicationContext) gxui.Button
```

#### func  DoReplay

```go
func DoReplay(appCtx *ApplicationContext)
```

#### func  ImportCapture

```go
func ImportCapture(appCtx *ApplicationContext, path string, statusLogger log.Logger)
```

#### func  Run

```go
func Run(config Config)
```

#### type ASCII

```go
type ASCII uint8
```


#### func (ASCII) Name

```go
func (v ASCII) Name() string
```

#### func (ASCII) Read

```go
func (v ASCII) Read(buffer []byte) DataType
```

#### func (ASCII) SizeBytes

```go
func (v ASCII) SizeBytes() int
```

#### func (ASCII) String

```go
func (v ASCII) String() string
```

#### func (ASCII) Unknown

```go
func (v ASCII) Unknown() string
```

#### type ApplicationContext

```go
type ApplicationContext struct {
	Config
}
```


#### func  CreateApplicationContext

```go
func CreateApplicationContext(theme gxui.Theme, config Config) (*ApplicationContext, error)
```

#### func (*ApplicationContext) Atoms

```go
func (c *ApplicationContext) Atoms() []Atom
```

#### func (*ApplicationContext) Capture

```go
func (c *ApplicationContext) Capture() service.Capture
```

#### func (*ApplicationContext) CaptureID

```go
func (c *ApplicationContext) CaptureID() service.CaptureId
```

#### func (*ApplicationContext) ColorBuffer

```go
func (c *ApplicationContext) ColorBuffer() gxui.Texture
```

#### func (*ApplicationContext) DecodeAtoms

```go
func (c *ApplicationContext) DecodeAtoms(stream service.AtomStream, s service.Schema) ([]Atom, error)
```
DecodeAtoms decodes all atoms from the AtomStream stream.

#### func (*ApplicationContext) DepthBuffer

```go
func (c *ApplicationContext) DepthBuffer() gxui.Texture
```

#### func (*ApplicationContext) DropDownOverlay

```go
func (c *ApplicationContext) DropDownOverlay() gxui.BubbleOverlay
```

#### func (*ApplicationContext) Hierarchy

```go
func (c *ApplicationContext) Hierarchy() atom.Group
```

#### func (*ApplicationContext) LoadCapture

```go
func (c *ApplicationContext) LoadCapture(captureID service.CaptureId, resetSelected bool)
```

#### func (*ApplicationContext) LoadHierarchy

```go
func (c *ApplicationContext) LoadHierarchy()
```

#### func (*ApplicationContext) LoadReport

```go
func (c *ApplicationContext) LoadReport()
```

#### func (*ApplicationContext) LoadState

```go
func (c *ApplicationContext) LoadState()
```

#### func (*ApplicationContext) Logger

```go
func (c *ApplicationContext) Logger() *log.Splitter
```

#### func (*ApplicationContext) OnAddressSelected

```go
func (c *ApplicationContext) OnAddressSelected(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnAtomSelected

```go
func (c *ApplicationContext) OnAtomSelected(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnAtomsUpdated

```go
func (c *ApplicationContext) OnAtomsUpdated(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnColorBufferUpdate

```go
func (c *ApplicationContext) OnColorBufferUpdate(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnDepthBufferUpdate

```go
func (c *ApplicationContext) OnDepthBufferUpdate(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnDeviceSelected

```go
func (c *ApplicationContext) OnDeviceSelected(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnHierarchyUpdated

```go
func (c *ApplicationContext) OnHierarchyUpdated(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnObjectSelected

```go
func (c *ApplicationContext) OnObjectSelected(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnReportUpdated

```go
func (c *ApplicationContext) OnReportUpdated(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnRequestReplay

```go
func (c *ApplicationContext) OnRequestReplay(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnStateUpdated

```go
func (c *ApplicationContext) OnStateUpdated(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnTimingInfoUpdated

```go
func (c *ApplicationContext) OnTimingInfoUpdated(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) OnWireframeChanged

```go
func (c *ApplicationContext) OnWireframeChanged(f func()) gxui.EventSubscription
```

#### func (*ApplicationContext) ReplaceAtom

```go
func (c *ApplicationContext) ReplaceAtom(a Atom, id atom.ID)
```

#### func (*ApplicationContext) RequestMemory

```go
func (c *ApplicationContext) RequestMemory(after atom.ID, base memory.Pointer, size uint64, callback MemoryCallback) chan<- struct{}
```

#### func (*ApplicationContext) RequestReplay

```go
func (c *ApplicationContext) RequestReplay()
```

#### func (*ApplicationContext) RequestThumbnail

```go
func (c *ApplicationContext) RequestThumbnail(after atom.ID, maxWidth, maxHeight uint32, callback ImageCallback) chan<- struct{}
```

#### func (*ApplicationContext) Rpc

```go
func (c *ApplicationContext) Rpc() service.RPC
```

#### func (*ApplicationContext) Run

```go
func (c *ApplicationContext) Run(f func())
```

#### func (*ApplicationContext) Schema

```go
func (c *ApplicationContext) Schema() service.Schema
```

#### func (*ApplicationContext) SelectAddress

```go
func (c *ApplicationContext) SelectAddress(address memory.Pointer)
```

#### func (*ApplicationContext) SelectAtom

```go
func (c *ApplicationContext) SelectAtom(id atom.ID)
```

#### func (*ApplicationContext) SelectDevice

```go
func (c *ApplicationContext) SelectDevice(device service.DeviceId)
```

#### func (*ApplicationContext) SelectObject

```go
func (c *ApplicationContext) SelectObject(object interface{})
```

#### func (*ApplicationContext) SelectedAddress

```go
func (c *ApplicationContext) SelectedAddress() memory.Pointer
```

#### func (*ApplicationContext) SelectedAtomID

```go
func (c *ApplicationContext) SelectedAtomID() atom.ID
```
func (c *ApplicationContext) State() schema.Struct { return c.state }

#### func (*ApplicationContext) SelectedDevice

```go
func (c *ApplicationContext) SelectedDevice() service.DeviceId
```

#### func (*ApplicationContext) SelectedObject

```go
func (c *ApplicationContext) SelectedObject() interface{}
```

#### func (*ApplicationContext) SetWireframe

```go
func (c *ApplicationContext) SetWireframe(value bool)
```

#### func (*ApplicationContext) Theme

```go
func (c *ApplicationContext) Theme() gxui.Theme
```

#### func (*ApplicationContext) ToolTipController

```go
func (c *ApplicationContext) ToolTipController() *gxui.ToolTipController
```

#### func (*ApplicationContext) ToolTipOverlay

```go
func (c *ApplicationContext) ToolTipOverlay() gxui.BubbleOverlay
```

#### func (*ApplicationContext) UpdateSchema

```go
func (c *ApplicationContext) UpdateSchema()
```

#### func (*ApplicationContext) Wireframe

```go
func (c *ApplicationContext) Wireframe() bool
```

#### type Atom

```go
type Atom struct {
}
```


#### func (*Atom) Api

```go
func (a *Atom) Api() service.ApiId
```

#### func (*Atom) DisplayName

```go
func (a *Atom) DisplayName() string
```

#### func (*Atom) DocumentationUrl

```go
func (a *Atom) DocumentationUrl() string
```

#### func (*Atom) Field

```go
func (a *Atom) Field(index int) (schema.Field, interface{})
```

#### func (*Atom) FieldCount

```go
func (a *Atom) FieldCount() int
```

#### func (*Atom) IsCommand

```go
func (a *Atom) IsCommand() bool
```

#### func (*Atom) IsDrawCall

```go
func (a *Atom) IsDrawCall() bool
```

#### func (*Atom) IsEndOfFrame

```go
func (a *Atom) IsEndOfFrame() bool
```

#### func (*Atom) Observations

```go
func (a *Atom) Observations() *atom.Observations
```

#### func (*Atom) SetField

```go
func (a *Atom) SetField(index int, value interface{})
```

#### type CommandAdapter

```go
type CommandAdapter struct {
	gxui.AdapterBase
}
```


#### func  CreateCommandAdapter

```go
func CreateCommandAdapter(appCtx *ApplicationContext) *CommandAdapter
```

#### func (CommandAdapter) AtomRange

```go
func (a CommandAdapter) AtomRange(item gxui.AdapterItem) atom.Range
```

#### func (CommandAdapter) Count

```go
func (n CommandAdapter) Count() int
```

#### func (CommandAdapter) Create

```go
func (n CommandAdapter) Create(theme gxui.Theme, index int) gxui.Control
```

#### func (CommandAdapter) Item

```go
func (a CommandAdapter) Item(id atom.ID) gxui.AdapterItem
```

#### func (CommandAdapter) ItemAt

```go
func (n CommandAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (CommandAdapter) ItemIndex

```go
func (n CommandAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (CommandAdapter) NodeAt

```go
func (n CommandAdapter) NodeAt(index int) gxui.TreeNode
```

#### func (*CommandAdapter) SetRoot

```go
func (a *CommandAdapter) SetRoot(root atom.Group)
```

#### func (CommandAdapter) Size

```go
func (a CommandAdapter) Size(theme gxui.Theme) math.Size
```
gxui.TreeAdapter compliance

#### type Config

```go
type Config struct {
	DataPath       string
	Gapis          string
	GXUIDebug      bool
	InitialCapture string
	ReplayDevice   string
}
```


#### type DataType

```go
type DataType interface {
	Name() string
	SizeBytes() int
	String() string
	Unknown() string
	Read(buffer []byte) DataType
}
```


#### type EnableDebugger

```go
type EnableDebugger interface {
	EnableDebug(bool)
}
```


#### type F32

```go
type F32 float32
```


#### func (F32) Name

```go
func (v F32) Name() string
```

#### func (F32) Read

```go
func (v F32) Read(buffer []byte) DataType
```

#### func (F32) SizeBytes

```go
func (v F32) SizeBytes() int
```

#### func (F32) String

```go
func (v F32) String() string
```

#### func (F32) Unknown

```go
func (v F32) Unknown() string
```

#### type F64

```go
type F64 float64
```


#### func (F64) Name

```go
func (v F64) Name() string
```

#### func (F64) Read

```go
func (v F64) Read(buffer []byte) DataType
```

#### func (F64) SizeBytes

```go
func (v F64) SizeBytes() int
```

#### func (F64) String

```go
func (v F64) String() string
```

#### func (F64) Unknown

```go
func (v F64) Unknown() string
```

#### type FilmStripAdapter

```go
type FilmStripAdapter struct {
	gxui.AdapterBase
}
```


#### func  CreateFilmStripAdapter

```go
func CreateFilmStripAdapter(appCtx *ApplicationContext) *FilmStripAdapter
```

#### func (*FilmStripAdapter) Count

```go
func (a *FilmStripAdapter) Count() int
```

#### func (*FilmStripAdapter) Create

```go
func (a *FilmStripAdapter) Create(t gxui.Theme, index int) gxui.Control
```

#### func (*FilmStripAdapter) ItemAt

```go
func (a *FilmStripAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*FilmStripAdapter) ItemIndex

```go
func (a *FilmStripAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*FilmStripAdapter) SetAtoms

```go
func (a *FilmStripAdapter) SetAtoms(atoms []Atom)
```

#### func (*FilmStripAdapter) Size

```go
func (a *FilmStripAdapter) Size(theme gxui.Theme) math.Size
```

#### type ImageCallback

```go
type ImageCallback func(gxui.Texture)
```


#### type InvokationHandler

```go
type InvokationHandler struct {
}
```


#### type LogAdapter

```go
type LogAdapter struct {
	gxui.AdapterBase
}
```


#### func  CreateLogAdapter

```go
func CreateLogAdapter(maxLength int, callOnUI func(func()) bool) *LogAdapter
```

#### func (*LogAdapter) Clear

```go
func (a *LogAdapter) Clear()
```

#### func (*LogAdapter) Count

```go
func (a *LogAdapter) Count() int
```

#### func (*LogAdapter) Create

```go
func (a *LogAdapter) Create(theme gxui.Theme, index int) gxui.Control
```

#### func (*LogAdapter) Entry

```go
func (a *LogAdapter) Entry(index int) log.Entry
```

#### func (*LogAdapter) ItemAt

```go
func (a *LogAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*LogAdapter) ItemIndex

```go
func (a *LogAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*LogAdapter) Logger

```go
func (a *LogAdapter) Logger() log.Logger
```

#### func (*LogAdapter) Size

```go
func (a *LogAdapter) Size(theme gxui.Theme) math.Size
```

#### type MemoryAdapter

```go
type MemoryAdapter struct {
	gxui.AdapterBase
}
```


#### func  CreateMemoryAdapter

```go
func CreateMemoryAdapter(appCtx *ApplicationContext) *MemoryAdapter
```

#### func (*MemoryAdapter) AddressAtIndex

```go
func (a *MemoryAdapter) AddressAtIndex(index int) memory.Pointer
```

#### func (*MemoryAdapter) Count

```go
func (a *MemoryAdapter) Count() int
```

#### func (*MemoryAdapter) Create

```go
func (a *MemoryAdapter) Create(t gxui.Theme, index int) gxui.Control
```

#### func (*MemoryAdapter) IndexOfAddress

```go
func (a *MemoryAdapter) IndexOfAddress(addr memory.Pointer) int
```

#### func (*MemoryAdapter) ItemAt

```go
func (a *MemoryAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*MemoryAdapter) ItemIndex

```go
func (a *MemoryAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*MemoryAdapter) SetData

```go
func (a *MemoryAdapter) SetData(atomID atom.ID, baseAddress memory.Pointer)
```

#### func (*MemoryAdapter) SetDataType

```go
func (a *MemoryAdapter) SetDataType(dataType DataType)
```

#### func (*MemoryAdapter) Size

```go
func (a *MemoryAdapter) Size(theme gxui.Theme) math.Size
```

#### type MemoryCallback

```go
type MemoryCallback func(service.MemoryInfo)
```


#### type MemoryImageAdapter

```go
type MemoryImageAdapter struct {
	gxui.AdapterBase
}
```

//////////////////////////////////////////////////////////////////////////////
MemoryImageAdapter
//////////////////////////////////////////////////////////////////////////////

#### func  CreateMemoryImageAdapter

```go
func CreateMemoryImageAdapter(appCtx *ApplicationContext) *MemoryImageAdapter
```

#### func (*MemoryImageAdapter) AddressAtIndex

```go
func (a *MemoryImageAdapter) AddressAtIndex(index int) memory.Pointer
```

#### func (*MemoryImageAdapter) Count

```go
func (a *MemoryImageAdapter) Count() int
```

#### func (*MemoryImageAdapter) Create

```go
func (a *MemoryImageAdapter) Create(t gxui.Theme, index int) gxui.Control
```

#### func (*MemoryImageAdapter) IndexOfAddress

```go
func (a *MemoryImageAdapter) IndexOfAddress(addr memory.Pointer) int
```

#### func (*MemoryImageAdapter) ItemAt

```go
func (a *MemoryImageAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*MemoryImageAdapter) ItemIndex

```go
func (a *MemoryImageAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*MemoryImageAdapter) SetData

```go
func (a *MemoryImageAdapter) SetData(atomID atom.ID, baseAddress memory.Pointer)
```

#### func (*MemoryImageAdapter) SetPixelType

```go
func (a *MemoryImageAdapter) SetPixelType(pixelType PixelType)
```

#### func (*MemoryImageAdapter) Size

```go
func (a *MemoryImageAdapter) Size(theme gxui.Theme) math.Size
```

#### type MemoryRequestCallback

```go
type MemoryRequestCallback func(atomId atom.ID, address memory.Pointer, bytes []byte)
```


#### type MemoryRequester

```go
type MemoryRequester interface {
	RequestMemory(atomId atom.ID, address memory.Pointer, size uint64, callback MemoryRequestCallback)
}
```


#### type PixelType

```go
type PixelType interface {
	Name() string
	SizeBytes() int
	Read(buffer []byte) gxui.Color
}
```


#### type RGB888

```go
type RGB888 int
```


#### func (RGB888) Name

```go
func (v RGB888) Name() string
```

#### func (RGB888) Read

```go
func (v RGB888) Read(data []byte) gxui.Color
```

#### func (RGB888) SizeBytes

```go
func (v RGB888) SizeBytes() int
```

#### type RGBA8888

```go
type RGBA8888 int
```


#### func (RGBA8888) Name

```go
func (v RGBA8888) Name() string
```

#### func (RGBA8888) Read

```go
func (v RGBA8888) Read(data []byte) gxui.Color
```

#### func (RGBA8888) SizeBytes

```go
func (v RGBA8888) SizeBytes() int
```

#### type ReportAdapter

```go
type ReportAdapter struct {
	gxui.AdapterBase
}
```


#### func (*ReportAdapter) Count

```go
func (a *ReportAdapter) Count() int
```

#### func (*ReportAdapter) Create

```go
func (a *ReportAdapter) Create(t gxui.Theme, index int) gxui.Control
```

#### func (*ReportAdapter) ItemAt

```go
func (a *ReportAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*ReportAdapter) ItemIndex

```go
func (a *ReportAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*ReportAdapter) Size

```go
func (a *ReportAdapter) Size(theme gxui.Theme) math.Size
```

#### type S16

```go
type S16 int16
```


#### func (S16) Name

```go
func (v S16) Name() string
```

#### func (S16) Read

```go
func (v S16) Read(buffer []byte) DataType
```

#### func (S16) SizeBytes

```go
func (v S16) SizeBytes() int
```

#### func (S16) String

```go
func (v S16) String() string
```

#### func (S16) Unknown

```go
func (v S16) Unknown() string
```

#### type S32

```go
type S32 int32
```


#### func (S32) Name

```go
func (v S32) Name() string
```

#### func (S32) Read

```go
func (v S32) Read(buffer []byte) DataType
```

#### func (S32) SizeBytes

```go
func (v S32) SizeBytes() int
```

#### func (S32) String

```go
func (v S32) String() string
```

#### func (S32) Unknown

```go
func (v S32) Unknown() string
```

#### type S64

```go
type S64 int64
```


#### func (S64) Name

```go
func (v S64) Name() string
```

#### func (S64) Read

```go
func (v S64) Read(buffer []byte) DataType
```

#### func (S64) SizeBytes

```go
func (v S64) SizeBytes() int
```

#### func (S64) String

```go
func (v S64) String() string
```

#### func (S64) Unknown

```go
func (v S64) Unknown() string
```

#### type S8

```go
type S8 int8
```


#### func (S8) Name

```go
func (v S8) Name() string
```

#### func (S8) Read

```go
func (v S8) Read(buffer []byte) DataType
```

#### func (S8) SizeBytes

```go
func (v S8) SizeBytes() int
```

#### func (S8) String

```go
func (v S8) String() string
```

#### func (S8) Unknown

```go
func (v S8) Unknown() string
```

#### type StateAdapter

```go
type StateAdapter struct {
	StateAdapterItem
	gxui.AdapterBase
}
```


#### func  CreateStateAdapter

```go
func CreateStateAdapter(appCtx *ApplicationContext, state schema.Struct) *StateAdapter
```

#### func (*StateAdapter) Size

```go
func (r *StateAdapter) Size(theme gxui.Theme) math.Size
```

#### type StateAdapterItem

```go
type StateAdapterItem struct {
}
```


#### func (*StateAdapterItem) Count

```go
func (i *StateAdapterItem) Count() int
```

#### func (*StateAdapterItem) Create

```go
func (i *StateAdapterItem) Create(t gxui.Theme, index int) gxui.Control
```

#### func (*StateAdapterItem) Init

```go
func (i *StateAdapterItem) Init(appCtx *ApplicationContext, key string, value interface{})
```

#### func (*StateAdapterItem) ItemAt

```go
func (i *StateAdapterItem) ItemAt(index int) gxui.AdapterItem
```

#### func (*StateAdapterItem) ItemIndex

```go
func (i *StateAdapterItem) ItemIndex(item gxui.AdapterItem) int
```

#### func (*StateAdapterItem) NodeAt

```go
func (i *StateAdapterItem) NodeAt(index int) gxui.TreeNode
```

#### type TimingData

```go
type TimingData struct {
}
```


#### func  NewTimingData

```go
func NewTimingData(t service.TimingInfo) TimingData
```

#### func (TimingData) BarBrush

```go
func (d TimingData) BarBrush(bar int, stack int, highlighted bool) gxui.Brush
```

#### func (TimingData) Count

```go
func (d TimingData) Count() int
```

#### func (TimingData) LabelBackgroundBrush

```go
func (d TimingData) LabelBackgroundBrush(bar int, stack int) gxui.Brush
```

#### func (TimingData) LabelTextColor

```go
func (d TimingData) LabelTextColor(bar int, stack int) gxui.Color
```

#### func (TimingData) Limits

```go
func (d TimingData) Limits() (int, int)
```

#### func (TimingData) Values

```go
func (d TimingData) Values(i int) []int
```

#### type U16

```go
type U16 uint16
```


#### func (U16) Name

```go
func (v U16) Name() string
```

#### func (U16) Read

```go
func (v U16) Read(buffer []byte) DataType
```

#### func (U16) SizeBytes

```go
func (v U16) SizeBytes() int
```

#### func (U16) String

```go
func (v U16) String() string
```

#### func (U16) Unknown

```go
func (v U16) Unknown() string
```

#### type U32

```go
type U32 uint32
```


#### func (U32) Name

```go
func (v U32) Name() string
```

#### func (U32) Read

```go
func (v U32) Read(buffer []byte) DataType
```

#### func (U32) SizeBytes

```go
func (v U32) SizeBytes() int
```

#### func (U32) String

```go
func (v U32) String() string
```

#### func (U32) Unknown

```go
func (v U32) Unknown() string
```

#### type U64

```go
type U64 uint64
```


#### func (U64) Name

```go
func (v U64) Name() string
```

#### func (U64) Read

```go
func (v U64) Read(buffer []byte) DataType
```

#### func (U64) SizeBytes

```go
func (v U64) SizeBytes() int
```

#### func (U64) String

```go
func (v U64) String() string
```

#### func (U64) Unknown

```go
func (v U64) Unknown() string
```

#### type U8

```go
type U8 uint8
```


#### func (U8) Name

```go
func (v U8) Name() string
```

#### func (U8) Read

```go
func (v U8) Read(buffer []byte) DataType
```

#### func (U8) SizeBytes

```go
func (v U8) SizeBytes() int
```

#### func (U8) String

```go
func (v U8) String() string
```

#### func (U8) Unknown

```go
func (v U8) Unknown() string
```
