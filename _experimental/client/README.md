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
func CreateGotoCommandDialog(appCtx *ApplicationContext, atoms *path.Atoms) gxui.Window
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
func CreateWireframeButton(appCtx *ApplicationContext, changed func(bool)) gxui.Button
```

#### func  ImportCapture

```go
func ImportCapture(appCtx *ApplicationContext, path string, statusLogger log.Logger)
```

#### func  NewColorTexture

```go
func NewColorTexture(driver gxui.Driver, width, height int, rgba []byte) gxui.Texture
```
NewColorTexture returns a gxui.Texture from the rgba-8888 data.

#### func  NewDepthTexture

```go
func NewDepthTexture(driver gxui.Driver, width, height int, depths []byte) gxui.Texture
```
NewDepthTexture returns a gxui.Texture from the depth data.

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

#### func (*ApplicationContext) Change

```go
func (c *ApplicationContext) Change(p path.Path, v interface{}) error
```
Change modifies the value at p to v, and selects the new path. The call is
blocking.

#### func (*ApplicationContext) Run

```go
func (c *ApplicationContext) Run(f func()) bool
```
Run enqueues f to be called on the UI go-routine. Run can return before f is
called.

#### func (*ApplicationContext) RunSync

```go
func (c *ApplicationContext) RunSync(f func()) bool
```
RunSync calls f on the UI go-routine, blocking until f has returned.

#### type Atom

```go
type Atom struct {
}
```


#### func (*Atom) API

```go
func (a *Atom) API() gfxapi.ID
```

#### func (*Atom) Class

```go
func (a *Atom) Class() binary.Class
```

#### func (*Atom) Field

```go
func (a *Atom) Field(index int) (schema.Field, interface{})
```

#### func (*Atom) FieldCount

```go
func (a *Atom) FieldCount() int
```

#### func (*Atom) Flags

```go
func (a *Atom) Flags() atom.Flags
```

#### func (*Atom) Mutate

```go
func (*Atom) Mutate(*gfxapi.State, database.Database, log.Logger) error
```

#### func (*Atom) Observations

```go
func (a *Atom) Observations() *atom.Observations
```

#### func (*Atom) SetField

```go
func (a *Atom) SetField(index int, value interface{})
```

#### type AtomClass

```go
type AtomClass struct {
}
```


#### func  NewAtomClass

```go
func NewAtomClass(base *schema.Class, meta *atom.Metadata) *AtomClass
```

#### func (*AtomClass) Decode

```go
func (c *AtomClass) Decode(d binary.Decoder) (binary.Object, error)
```

#### func (*AtomClass) DecodeTo

```go
func (c *AtomClass) DecodeTo(d binary.Decoder, object binary.Object) error
```

#### func (*AtomClass) Encode

```go
func (c *AtomClass) Encode(e binary.Encoder, object binary.Object) error
```

#### func (*AtomClass) ID

```go
func (c *AtomClass) ID() binary.ID
```

#### func (*AtomClass) New

```go
func (c *AtomClass) New() binary.Object
```

#### func (*AtomClass) Schema

```go
func (c *AtomClass) Schema() *schema.Class
```

#### func (*AtomClass) Skip

```go
func (c *AtomClass) Skip(d binary.Decoder) error
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
func (a CommandAdapter) Item(p path.Path) gxui.AdapterItem
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

#### func (CommandAdapter) Path

```go
func (a CommandAdapter) Path(item gxui.AdapterItem) path.Path
```

#### func (CommandAdapter) Size

```go
func (a CommandAdapter) Size(theme gxui.Theme) math.Size
```
gxui.TreeAdapter compliance

#### func (*CommandAdapter) UpdateAtoms

```go
func (a *CommandAdapter) UpdateAtoms(capture *path.Capture, atoms []atom.Atom, root atom.Group)
```

#### func (*CommandAdapter) UpdateDevice

```go
func (a *CommandAdapter) UpdateDevice(device *path.Device)
```

#### func (*CommandAdapter) UpdateTimings

```go
func (a *CommandAdapter) UpdateTimings(timings service.TimingInfo)
```

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


#### type Events

```go
type Events struct {
}
```


#### func (*Events) Init

```go
func (e *Events) Init()
```

#### func (*Events) OnSelect

```go
func (e *Events) OnSelect(f func(p path.Path)) gxui.EventSubscription
```

#### func (*Events) Select

```go
func (e *Events) Select(p path.Path)
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
func (a *FilmStripAdapter) Create(theme gxui.Theme, index int) gxui.Control
```

#### func (*FilmStripAdapter) ItemAt

```go
func (a *FilmStripAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*FilmStripAdapter) ItemIndex

```go
func (a *FilmStripAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*FilmStripAdapter) Size

```go
func (a *FilmStripAdapter) Size(theme gxui.Theme) math.Size
```

#### func (*FilmStripAdapter) UpdateDevice

```go
func (a *FilmStripAdapter) UpdateDevice(device *path.Device)
```

#### func (*FilmStripAdapter) UpdateFrames

```go
func (a *FilmStripAdapter) UpdateFrames(capture *path.Capture, frames []uint64)
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
func (a *MemoryAdapter) AddressAtIndex(index int) uint64
```

#### func (*MemoryAdapter) Count

```go
func (a *MemoryAdapter) Count() int
```

#### func (*MemoryAdapter) Create

```go
func (a *MemoryAdapter) Create(theme gxui.Theme, index int) gxui.Control
```

#### func (*MemoryAdapter) IndexOfAddress

```go
func (a *MemoryAdapter) IndexOfAddress(addr uint64) int
```

#### func (*MemoryAdapter) ItemAt

```go
func (a *MemoryAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*MemoryAdapter) ItemIndex

```go
func (a *MemoryAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*MemoryAdapter) SetDataType

```go
func (a *MemoryAdapter) SetDataType(dataType DataType)
```

#### func (*MemoryAdapter) Size

```go
func (a *MemoryAdapter) Size(theme gxui.Theme) math.Size
```

#### func (*MemoryAdapter) Update

```go
func (a *MemoryAdapter) Update(after *path.Atom, baseAddress uint64)
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
func (a *MemoryImageAdapter) AddressAtIndex(index int) uint64
```

#### func (*MemoryImageAdapter) Count

```go
func (a *MemoryImageAdapter) Count() int
```

#### func (*MemoryImageAdapter) Create

```go
func (a *MemoryImageAdapter) Create(theme gxui.Theme, index int) gxui.Control
```

#### func (*MemoryImageAdapter) IndexOfAddress

```go
func (a *MemoryImageAdapter) IndexOfAddress(addr uint64) int
```

#### func (*MemoryImageAdapter) ItemAt

```go
func (a *MemoryImageAdapter) ItemAt(index int) gxui.AdapterItem
```

#### func (*MemoryImageAdapter) ItemIndex

```go
func (a *MemoryImageAdapter) ItemIndex(item gxui.AdapterItem) int
```

#### func (*MemoryImageAdapter) SetPixelType

```go
func (a *MemoryImageAdapter) SetPixelType(pixelType PixelType)
```

#### func (*MemoryImageAdapter) Size

```go
func (a *MemoryImageAdapter) Size(theme gxui.Theme) math.Size
```

#### func (*MemoryImageAdapter) Update

```go
func (a *MemoryImageAdapter) Update(after *path.Atom, baseAddress uint64)
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

#### func (*ReportAdapter) Update

```go
func (a *ReportAdapter) Update(report service.Report)
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
	StateAdapterNode
	gxui.AdapterBase
}
```


#### func  NewStateAdapter

```go
func NewStateAdapter(appCtx *ApplicationContext) *StateAdapter
```

#### func (*StateAdapter) Size

```go
func (r *StateAdapter) Size(theme gxui.Theme) math.Size
```

#### func (*StateAdapter) Update

```go
func (a *StateAdapter) Update(value interface{}, path *path.State)
```

#### type StateAdapterNode

```go
type StateAdapterNode struct {
}
```


#### func (*StateAdapterNode) Count

```go
func (n *StateAdapterNode) Count() int
```

#### func (*StateAdapterNode) Create

```go
func (n *StateAdapterNode) Create(t gxui.Theme, index int) gxui.Control
```

#### func (*StateAdapterNode) ItemAt

```go
func (n *StateAdapterNode) ItemAt(index int) gxui.AdapterItem
```

#### func (*StateAdapterNode) ItemIndex

```go
func (n *StateAdapterNode) ItemIndex(item gxui.AdapterItem) int
```

#### func (*StateAdapterNode) NodeAt

```go
func (n *StateAdapterNode) NodeAt(index int) gxui.TreeNode
```

#### type StateAdapterNodeList

```go
type StateAdapterNodeList []*StateAdapterNode
```


#### func (StateAdapterNodeList) Len

```go
func (l StateAdapterNodeList) Len() int
```

#### func (StateAdapterNodeList) Less

```go
func (l StateAdapterNodeList) Less(a, b int) bool
```

#### func (StateAdapterNodeList) Swap

```go
func (l StateAdapterNodeList) Swap(a, b int)
```

#### type TimingData

```go
type TimingData struct {
}
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
