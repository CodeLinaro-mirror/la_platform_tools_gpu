# charts
--
    import "android.googlesource.com/platform/tools/gpu/_experimental/client/charts"


## Usage

```go
const (
	ScrollX = InputKind(iota)
	ScrollY
	DragX
	DragY
)
```

```go
const (
	Nop = ActionKind(iota)
	PanX
	PanY
	ZoomX
	ZoomY
	ZoomXY
)
```

```go
var DefaultGridlines = GridLines{
	Brush:     gxui.CreateBrush(gxui.Gray40),
	TextColor: gxui.Gray70,
	Font:      nil,
	Format:    func(v int) string { return fmt.Sprintf("%d", v) },
	Density:   30,
	Multiples: []int{2, 5},
}
```

#### type ActionKind

```go
type ActionKind int
```


#### func  DefaultInputHandler

```go
func DefaultInputHandler(input InputKind, o gxui.Orientation, ev gxui.MouseEvent) ActionKind
```

#### type BarChart

```go
type BarChart struct {
	base.Container
	ZoomWindow
	parts.BackgroundBorderPainter
}
```


#### func  NewBarChart

```go
func NewBarChart(theme gxui.Theme) *BarChart
```

#### func (*BarChart) Click

```go
func (c *BarChart) Click(e gxui.MouseEvent) bool
```

#### func (*BarChart) DesiredSize

```go
func (c *BarChart) DesiredSize(min, max math.Size) math.Size
```

#### func (*BarChart) DoubleClick

```go
func (c *BarChart) DoubleClick(e gxui.MouseEvent) bool
```

#### func (*BarChart) LayoutChildren

```go
func (c *BarChart) LayoutChildren()
```

#### func (*BarChart) MouseMove

```go
func (c *BarChart) MouseMove(e gxui.MouseEvent)
```

#### func (*BarChart) OnBarClicked

```go
func (c *BarChart) OnBarClicked(f func(idx int, ev gxui.MouseEvent)) gxui.EventSubscription
```

#### func (*BarChart) OnBarDoubleClicked

```go
func (c *BarChart) OnBarDoubleClicked(f func(idx int, ev gxui.MouseEvent)) gxui.EventSubscription
```

#### func (*BarChart) Orientation

```go
func (c *BarChart) Orientation() gxui.Orientation
```

#### func (*BarChart) Paint

```go
func (c *BarChart) Paint(canvas gxui.Canvas)
```

#### func (*BarChart) SetBarPen

```go
func (c *BarChart) SetBarPen(pen gxui.Pen)
```

#### func (*BarChart) SetData

```go
func (c *BarChart) SetData(data BarChartData)
```

#### func (*BarChart) SetGridlines

```go
func (c *BarChart) SetGridlines(gridlines GridLines)
```

#### func (*BarChart) SetHighlightedBarPen

```go
func (c *BarChart) SetHighlightedBarPen(pen gxui.Pen)
```

#### func (*BarChart) SetOrientation

```go
func (c *BarChart) SetOrientation(o gxui.Orientation)
```

#### func (*BarChart) SetSize

```go
func (c *BarChart) SetSize(s math.Size)
```

#### type BarChartData

```go
type BarChartData interface {
	Count() int
	Values(bar int) []int
	Limits() (int, int)
	BarBrush(bar int, stack int, highlighted bool) gxui.Brush
	LabelBackgroundBrush(bar int, stack int) gxui.Brush
	LabelTextColor(bar int, stack int) gxui.Color
}
```


#### type GridLines

```go
type GridLines struct {
	Brush     gxui.Brush
	TextColor gxui.Color
	Font      gxui.Font
	Format    func(int) string
	Density   int // DIPs per line
	Multiples []int
}
```


#### func (GridLines) PaintHorizontal

```go
func (g GridLines) PaintHorizontal(theme gxui.Theme, canvas gxui.Canvas, zoomWindow math.Rect, viewBounds math.Rect)
```

#### func (GridLines) PaintVertical

```go
func (g GridLines) PaintVertical(theme gxui.Theme, canvas gxui.Canvas, zoomWindow math.Rect, viewBounds math.Rect)
```

#### type InputHandler

```go
type InputHandler func(InputKind, gxui.Orientation, gxui.MouseEvent) ActionKind
```


#### type InputKind

```go
type InputKind int
```


#### type ZoomWindow

```go
type ZoomWindow struct {
	InputHandler InputHandler
}
```


#### func (*ZoomWindow) Init

```go
func (z *ZoomWindow) Init(outer ZoomWindowOuter, theme gxui.Theme)
```

#### func (*ZoomWindow) MouseDown

```go
func (z *ZoomWindow) MouseDown(e gxui.MouseEvent)
```

#### func (*ZoomWindow) MouseMove

```go
func (z *ZoomWindow) MouseMove(e gxui.MouseEvent)
```

#### func (*ZoomWindow) MouseScroll

```go
func (z *ZoomWindow) MouseScroll(e gxui.MouseEvent) bool
```

#### func (*ZoomWindow) MouseUp

```go
func (z *ZoomWindow) MouseUp(e gxui.MouseEvent)
```

#### func (*ZoomWindow) SetMinZoomWindowSize

```go
func (z *ZoomWindow) SetMinZoomWindowSize(size math.Size)
```

#### func (*ZoomWindow) SetZoomBounds

```go
func (z *ZoomWindow) SetZoomBounds(zoomBounds math.Rect)
```

#### func (*ZoomWindow) SetZoomWindow

```go
func (z *ZoomWindow) SetZoomWindow(zoomWindow math.Rect)
```

#### func (*ZoomWindow) ViewBounds

```go
func (z *ZoomWindow) ViewBounds() math.Rect
```

#### type ZoomWindowOuter

```go
type ZoomWindowOuter interface {
	base.ContainerOuter
	Orientation() gxui.Orientation
}
```
