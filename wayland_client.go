// Warning. This file was generated by wayland-generator
// and should not be changed manually.
package wayland

import (
	"os"
)

const (
	DisplayErrorInvalidObject = 0
	DisplayErrorInvalidMethod = 1
	DisplayErrorNoMemory      = 2
)

const (
	ShmErrorInvalidFormat = 0
	ShmErrorInvalidStride = 1
	ShmErrorInvalidFd     = 2
)

const (
	ShmFormatArgb8888    = 0
	ShmFormatXrgb8888    = 1
	ShmFormatC8          = 0x20203843
	ShmFormatRgb332      = 0x38424752
	ShmFormatBgr233      = 0x38524742
	ShmFormatXrgb4444    = 0x32315258
	ShmFormatXbgr4444    = 0x32314258
	ShmFormatRgbx4444    = 0x32315852
	ShmFormatBgrx4444    = 0x32315842
	ShmFormatArgb4444    = 0x32315241
	ShmFormatAbgr4444    = 0x32314241
	ShmFormatRgba4444    = 0x32314152
	ShmFormatBgra4444    = 0x32314142
	ShmFormatXrgb1555    = 0x35315258
	ShmFormatXbgr1555    = 0x35314258
	ShmFormatRgbx5551    = 0x35315852
	ShmFormatBgrx5551    = 0x35315842
	ShmFormatArgb1555    = 0x35315241
	ShmFormatAbgr1555    = 0x35314241
	ShmFormatRgba5551    = 0x35314152
	ShmFormatBgra5551    = 0x35314142
	ShmFormatRgb565      = 0x36314752
	ShmFormatBgr565      = 0x36314742
	ShmFormatRgb888      = 0x34324752
	ShmFormatBgr888      = 0x34324742
	ShmFormatXbgr8888    = 0x34324258
	ShmFormatRgbx8888    = 0x34325852
	ShmFormatBgrx8888    = 0x34325842
	ShmFormatAbgr8888    = 0x34324241
	ShmFormatRgba8888    = 0x34324152
	ShmFormatBgra8888    = 0x34324142
	ShmFormatXrgb2101010 = 0x30335258
	ShmFormatXbgr2101010 = 0x30334258
	ShmFormatRgbx1010102 = 0x30335852
	ShmFormatBgrx1010102 = 0x30335842
	ShmFormatArgb2101010 = 0x30335241
	ShmFormatAbgr2101010 = 0x30334241
	ShmFormatRgba1010102 = 0x30334152
	ShmFormatBgra1010102 = 0x30334142
	ShmFormatYuyv        = 0x56595559
	ShmFormatYvyu        = 0x55595659
	ShmFormatUyvy        = 0x59565955
	ShmFormatVyuy        = 0x59555956
	ShmFormatAyuv        = 0x56555941
	ShmFormatNv12        = 0x3231564e
	ShmFormatNv21        = 0x3132564e
	ShmFormatNv16        = 0x3631564e
	ShmFormatNv61        = 0x3136564e
	ShmFormatYuv410      = 0x39565559
	ShmFormatYvu410      = 0x39555659
	ShmFormatYuv411      = 0x31315559
	ShmFormatYvu411      = 0x31315659
	ShmFormatYuv420      = 0x32315559
	ShmFormatYvu420      = 0x32315659
	ShmFormatYuv422      = 0x36315559
	ShmFormatYvu422      = 0x36315659
	ShmFormatYuv444      = 0x34325559
	ShmFormatYvu444      = 0x34325659
)

const (
	DataDeviceErrorRole = 0
)

const (
	ShellErrorRole = 0
)

const (
	ShellSurfaceResizeNone        = 0
	ShellSurfaceResizeTop         = 1
	ShellSurfaceResizeBottom      = 2
	ShellSurfaceResizeLeft        = 4
	ShellSurfaceResizeTopLeft     = 5
	ShellSurfaceResizeBottomLeft  = 6
	ShellSurfaceResizeRight       = 8
	ShellSurfaceResizeTopRight    = 9
	ShellSurfaceResizeBottomRight = 10
)

const (
	ShellSurfaceTransientInactive = 0x1
)

const (
	ShellSurfaceFullscreenMethodDefault = 0
	ShellSurfaceFullscreenMethodScale   = 1
	ShellSurfaceFullscreenMethodDriver  = 2
	ShellSurfaceFullscreenMethodFill    = 3
)

const (
	SurfaceErrorInvalidScale     = 0
	SurfaceErrorInvalidTransform = 1
)

const (
	SeatCapabilityPointer  = 1
	SeatCapabilityKeyboard = 2
	SeatCapabilityTouch    = 4
)

const (
	PointerErrorRole = 0
)

const (
	PointerButtonStateReleased = 0
	PointerButtonStatePressed  = 1
)

const (
	PointerAxisVerticalScroll   = 0
	PointerAxisHorizontalScroll = 1
)

const (
	KeyboardKeymapFormatNoKeymap = 0
	KeyboardKeymapFormatXkbV1    = 1
)

const (
	KeyboardKeyStateReleased = 0
	KeyboardKeyStatePressed  = 1
)

const (
	OutputSubpixelUnknown       = 0
	OutputSubpixelNone          = 1
	OutputSubpixelHorizontalRgb = 2
	OutputSubpixelHorizontalBgr = 3
	OutputSubpixelVerticalRgb   = 4
	OutputSubpixelVerticalBgr   = 5
)

const (
	OutputTransformNormal     = 0
	OutputTransform90         = 1
	OutputTransform180        = 2
	OutputTransform270        = 3
	OutputTransformFlipped    = 4
	OutputTransformFlipped90  = 5
	OutputTransformFlipped180 = 6
	OutputTransformFlipped270 = 7
)

const (
	OutputModeCurrent   = 0x1
	OutputModePreferred = 0x2
)

const (
	SubcompositorErrorBadSurface = 0
)

const (
	SubsurfaceErrorBadSurface = 0
)

type DisplayErrorEvent struct {
	ObjectId Proxy
	Code     uint32
	Message  string
}

type Display struct {
	BaseProxy
	ErrorChan    chan<- DisplayErrorEvent
	DeleteIdChan chan<- uint32
	context      *connection
}

func (p *Display) Sync() (*Callback, error) {
	ret := Callback{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *Display) GetRegistry() (*Registry, error) {
	ret := Registry{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 1)
}

type RegistryGlobalEvent struct {
	Name    uint32
	Ifc     string
	Version uint32
}

type Registry struct {
	BaseProxy
	GlobalChan       chan<- RegistryGlobalEvent
	GlobalRemoveChan chan<- uint32
}

func (p *Registry) Bind(name uint32, id Proxy) error {
	return Proxy(p).GetDisplay().SendRequest(p, 0, name, id)
}

type Callback struct {
	BaseProxy
	DoneChan chan<- uint32
}

type Compositor struct {
	BaseProxy
}

func (p *Compositor) CreateSurface() (*Surface, error) {
	ret := Surface{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *Compositor) CreateRegion() (*Region, error) {
	ret := Region{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 1)
}

type ShmPool struct {
	BaseProxy
}

func (p *ShmPool) CreateBuffer(offset int, width int, height int, stride int, format uint32) (*Buffer, error) {
	ret := Buffer{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 0, offset, width, height, stride, format)
}

func (p *ShmPool) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 1)
}

func (p *ShmPool) Resize(size int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 2, size)
}

type Shm struct {
	BaseProxy
	FormatChan chan<- uint32
}

func (p *Shm) CreatePool(fd *os.File, size int) (*ShmPool, error) {
	ret := ShmPool{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 0, fd, size)
}

type BufferReleaseEvent struct {
}

type Buffer struct {
	BaseProxy
	ReleaseChan chan<- BufferReleaseEvent
}

func (p *Buffer) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 0)
}

type DataOffer struct {
	BaseProxy
	OfferChan chan<- string
}

func (p *DataOffer) Accept(serial uint32, mimeType string) error {
	return Proxy(p).GetDisplay().SendRequest(p, 0, serial, mimeType)
}

func (p *DataOffer) Receive(mimeType string, fd *os.File) error {
	return Proxy(p).GetDisplay().SendRequest(p, 1, mimeType, fd)
}

func (p *DataOffer) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 2)
}

type DataSourceSendEvent struct {
	MimeType string
	Fd       *os.File
}

type DataSourceCancelledEvent struct {
}

type DataSource struct {
	BaseProxy
	TargetChan    chan<- string
	SendChan      chan<- DataSourceSendEvent
	CancelledChan chan<- DataSourceCancelledEvent
}

func (p *DataSource) Offer(mimeType string) error {
	return Proxy(p).GetDisplay().SendRequest(p, 0, mimeType)
}

func (p *DataSource) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 1)
}

type DataDeviceEnterEvent struct {
	Serial  uint32
	Surface Surface
	X       float32
	Y       float32
	Id      DataOffer
}

type DataDeviceLeaveEvent struct {
}

type DataDeviceMotionEvent struct {
	Time uint32
	X    float32
	Y    float32
}

type DataDeviceDropEvent struct {
}

type DataDevice struct {
	BaseProxy
	DataOfferChan chan<- DataOffer
	EnterChan     chan<- DataDeviceEnterEvent
	LeaveChan     chan<- DataDeviceLeaveEvent
	MotionChan    chan<- DataDeviceMotionEvent
	DropChan      chan<- DataDeviceDropEvent
	SelectionChan chan<- DataOffer
}

func (p *DataDevice) StartDrag(source DataSource, origin Surface, icon Surface, serial uint32) error {
	return Proxy(p).GetDisplay().SendRequest(p, 0, source, origin, icon, serial)
}

func (p *DataDevice) SetSelection(source DataSource, serial uint32) error {
	return Proxy(p).GetDisplay().SendRequest(p, 1, source, serial)
}

func (p *DataDevice) Release() error {
	return Proxy(p).GetDisplay().SendRequest(p, 2)
}

type DataDeviceManager struct {
	BaseProxy
}

func (p *DataDeviceManager) CreateDataSource() (*DataSource, error) {
	ret := DataSource{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *DataDeviceManager) GetDataDevice(seat Seat) (*DataDevice, error) {
	ret := DataDevice{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 1, seat)
}

type Shell struct {
	BaseProxy
}

func (p *Shell) GetShellSurface(surface Surface) (*ShellSurface, error) {
	ret := ShellSurface{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 0, surface)
}

type ShellSurfaceConfigureEvent struct {
	Edges  uint32
	Width  int
	Height int
}

type ShellSurfacePopupDoneEvent struct {
}

type ShellSurface struct {
	BaseProxy
	PingChan      chan<- uint32
	ConfigureChan chan<- ShellSurfaceConfigureEvent
	PopupDoneChan chan<- ShellSurfacePopupDoneEvent
}

func (p *ShellSurface) Pong(serial uint32) error {
	return Proxy(p).GetDisplay().SendRequest(p, 0, serial)
}

func (p *ShellSurface) Move(seat Seat, serial uint32) error {
	return Proxy(p).GetDisplay().SendRequest(p, 1, seat, serial)
}

func (p *ShellSurface) Resize(seat Seat, serial uint32, edges uint32) error {
	return Proxy(p).GetDisplay().SendRequest(p, 2, seat, serial, edges)
}

func (p *ShellSurface) SetToplevel() error {
	return Proxy(p).GetDisplay().SendRequest(p, 3)
}

func (p *ShellSurface) SetTransient(parent Surface, x int, y int, flags uint32) error {
	return Proxy(p).GetDisplay().SendRequest(p, 4, parent, x, y, flags)
}

func (p *ShellSurface) SetFullscreen(method uint32, framerate uint32, output Output) error {
	return Proxy(p).GetDisplay().SendRequest(p, 5, method, framerate, output)
}

func (p *ShellSurface) SetPopup(seat Seat, serial uint32, parent Surface, x int, y int, flags uint32) error {
	return Proxy(p).GetDisplay().SendRequest(p, 6, seat, serial, parent, x, y, flags)
}

func (p *ShellSurface) SetMaximized(output Output) error {
	return Proxy(p).GetDisplay().SendRequest(p, 7, output)
}

func (p *ShellSurface) SetTitle(title string) error {
	return Proxy(p).GetDisplay().SendRequest(p, 8, title)
}

func (p *ShellSurface) SetClass(class string) error {
	return Proxy(p).GetDisplay().SendRequest(p, 9, class)
}

type Surface struct {
	BaseProxy
	EnterChan chan<- Output
	LeaveChan chan<- Output
}

func (p *Surface) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *Surface) Attach(buffer Buffer, x int, y int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 1, buffer, x, y)
}

func (p *Surface) Damage(x int, y int, width int, height int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 2, x, y, width, height)
}

func (p *Surface) Frame() (*Callback, error) {
	ret := Callback{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 3)
}

func (p *Surface) SetOpaqueRegion(region Region) error {
	return Proxy(p).GetDisplay().SendRequest(p, 4, region)
}

func (p *Surface) SetInputRegion(region Region) error {
	return Proxy(p).GetDisplay().SendRequest(p, 5, region)
}

func (p *Surface) Commit() error {
	return Proxy(p).GetDisplay().SendRequest(p, 6)
}

func (p *Surface) SetBufferTransform(transform int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 7, transform)
}

func (p *Surface) SetBufferScale(scale int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 8, scale)
}

type Seat struct {
	BaseProxy
	CapabilitiesChan chan<- uint32
	NameChan         chan<- string
}

func (p *Seat) GetPointer() (*Pointer, error) {
	ret := Pointer{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *Seat) GetKeyboard() (*Keyboard, error) {
	ret := Keyboard{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 1)
}

func (p *Seat) GetTouch() (*Touch, error) {
	ret := Touch{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 2)
}

type PointerEnterEvent struct {
	Serial   uint32
	Surface  Surface
	SurfaceX float32
	SurfaceY float32
}

type PointerLeaveEvent struct {
	Serial  uint32
	Surface Surface
}

type PointerMotionEvent struct {
	Time     uint32
	SurfaceX float32
	SurfaceY float32
}

type PointerButtonEvent struct {
	Serial uint32
	Time   uint32
	Button uint32
	State  uint32
}

type PointerAxisEvent struct {
	Time  uint32
	Axis  uint32
	Value float32
}

type Pointer struct {
	BaseProxy
	EnterChan  chan<- PointerEnterEvent
	LeaveChan  chan<- PointerLeaveEvent
	MotionChan chan<- PointerMotionEvent
	ButtonChan chan<- PointerButtonEvent
	AxisChan   chan<- PointerAxisEvent
}

func (p *Pointer) SetCursor(serial uint32, surface Surface, hotspotX int, hotspotY int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 0, serial, surface, hotspotX, hotspotY)
}

func (p *Pointer) Release() error {
	return Proxy(p).GetDisplay().SendRequest(p, 1)
}

type KeyboardKeymapEvent struct {
	Format uint32
	Fd     *os.File
	Size   uint32
}

type KeyboardEnterEvent struct {
	Serial  uint32
	Surface Surface
	Keys    []int
}

type KeyboardLeaveEvent struct {
	Serial  uint32
	Surface Surface
}

type KeyboardKeyEvent struct {
	Serial uint32
	Time   uint32
	Key    uint32
	State  uint32
}

type KeyboardModifiersEvent struct {
	Serial        uint32
	ModsDepressed uint32
	ModsLatched   uint32
	ModsLocked    uint32
	Group         uint32
}

type KeyboardRepeatInfoEvent struct {
	Rate  int
	Delay int
}

type Keyboard struct {
	BaseProxy
	KeymapChan     chan<- KeyboardKeymapEvent
	EnterChan      chan<- KeyboardEnterEvent
	LeaveChan      chan<- KeyboardLeaveEvent
	KeyChan        chan<- KeyboardKeyEvent
	ModifiersChan  chan<- KeyboardModifiersEvent
	RepeatInfoChan chan<- KeyboardRepeatInfoEvent
}

func (p *Keyboard) Release() error {
	return Proxy(p).GetDisplay().SendRequest(p, 0)
}

type TouchDownEvent struct {
	Serial  uint32
	Time    uint32
	Surface Surface
	Id      int
	X       float32
	Y       float32
}

type TouchUpEvent struct {
	Serial uint32
	Time   uint32
	Id     int
}

type TouchMotionEvent struct {
	Time uint32
	Id   int
	X    float32
	Y    float32
}

type TouchFrameEvent struct {
}

type TouchCancelEvent struct {
}

type Touch struct {
	BaseProxy
	DownChan   chan<- TouchDownEvent
	UpChan     chan<- TouchUpEvent
	MotionChan chan<- TouchMotionEvent
	FrameChan  chan<- TouchFrameEvent
	CancelChan chan<- TouchCancelEvent
}

func (p *Touch) Release() error {
	return Proxy(p).GetDisplay().SendRequest(p, 0)
}

type OutputGeometryEvent struct {
	X              int
	Y              int
	PhysicalWidth  int
	PhysicalHeight int
	Subpixel       int
	Make           string
	Model          string
	Transform      int
}

type OutputModeEvent struct {
	Flags   uint32
	Width   int
	Height  int
	Refresh int
}

type OutputDoneEvent struct {
}

type Output struct {
	BaseProxy
	GeometryChan chan<- OutputGeometryEvent
	ModeChan     chan<- OutputModeEvent
	DoneChan     chan<- OutputDoneEvent
	ScaleChan    chan<- int
}

type Region struct {
	BaseProxy
}

func (p *Region) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *Region) Add(x int, y int, width int, height int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 1, x, y, width, height)
}

func (p *Region) Subtract(x int, y int, width int, height int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 2, x, y, width, height)
}

type Subcompositor struct {
	BaseProxy
}

func (p *Subcompositor) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *Subcompositor) GetSubsurface(surface Surface, parent Surface) (*Subsurface, error) {
	ret := Subsurface{}
	Proxy(p).GetDisplay().Register(&ret)
	return &ret, Proxy(p).GetDisplay().SendRequest(p, 1, surface, parent)
}

type Subsurface struct {
	BaseProxy
}

func (p *Subsurface) Destroy() error {
	return Proxy(p).GetDisplay().SendRequest(p, 0)
}

func (p *Subsurface) SetPosition(x int, y int) error {
	return Proxy(p).GetDisplay().SendRequest(p, 1, x, y)
}

func (p *Subsurface) PlaceAbove(sibling Surface) error {
	return Proxy(p).GetDisplay().SendRequest(p, 2, sibling)
}

func (p *Subsurface) PlaceBelow(sibling Surface) error {
	return Proxy(p).GetDisplay().SendRequest(p, 3, sibling)
}

func (p *Subsurface) SetSync() error {
	return Proxy(p).GetDisplay().SendRequest(p, 4)
}

func (p *Subsurface) SetDesync() error {
	return Proxy(p).GetDisplay().SendRequest(p, 5)
}