# adb
--
    import "android.googlesource.com/platform/tools/gpu/adb"

Package adb provides an interface to the Android Debug Bridge.

## Usage

```go
var ConstantValues schema.Constants
```

```go
var ErrADBNotFound = errors.New("ADB command not found on PATH")
```
ErrADBNotFound is returned when the ADB executable is not found.

```go
var ErrDeviceNotRooted = errors.New("Device is not rooted")
```
ErrDeviceNotRooted is returned by Device.Root when the device is running a
production build as is not 'rooted'.

```go
var ErrDeviceUnauthorized = errors.New("Device unauthorized")
```
ErrDeviceUnauthorized is returned by ADB commands when the device has not
authorized ADB debugging. Check the confirmation dialog on the device.

#### func  Devices

```go
func Devices() ([]*Device, error)
```
Devices returns the list of serial numbers of all the attached Android devices.

#### type Action

```go
type Action struct {
	Name     string // Example: android.intent.action.MAIN
	Package  *InstalledPackage
	Activity string // Example: .FooBarActivity
}
```

Action represents an Android action that can be sent as an intent.

#### func (*Action) String

```go
func (a *Action) String() string
```

#### type Cmd

```go
type Cmd struct {
	// Path is the path of the command to run on the device.
	//
	// If the string is empty, the command is treated as a ADB command for Device.
	Path string

	// Args holds the command line arguments to pass to the command.
	Args []string

	// The device this command should be run on. If nil, then any one of the
	// attached devices will execute the command.
	Device *Device

	// Stdout and Stderr specify the process's standard output and error.
	//
	// If either is nil, Run connects the corresponding file descriptor
	// to the null device (os.DevNull).
	//
	// If Stdout and Stderr are the same writer, at most one
	// goroutine at a time will call Write.
	Stdout io.Writer
	Stderr io.Writer
}
```

Cmd represents a command that can be run on an Android device.

#### func (*Cmd) Call

```go
func (c *Cmd) Call() (string, error)
```
Call starts the specified command and waits for it to complete, returning the
all stdout as a string. The returned error is nil if the command runs, has no
problems copying stdout and stderr, and exits with a zero exit status.

#### func (*Cmd) Run

```go
func (c *Cmd) Run() error
```
Run starts the specified command and waits for it to complete. The returned
error is nil if the command runs, has no problems copying stdout and stderr, and
exits with a zero exit status.

#### type Device

```go
type Device struct {
	Serial string
	State  DeviceState
}
```

Device represents an attached Android device.

#### func (*Device) Abi

```go
func (d *Device) Abi() string
```
String returns a string representing the device.

#### func (*Device) Command

```go
func (d *Device) Command(path string, args ...string) *Cmd
```
Command returns a new Cmd that will run the command with the specified name and
arguments on this device.

#### func (*Device) Forward

```go
func (d *Device) Forward(local, device Port) error
```
Forward will forward the specified device Port to the specified local Port.

#### func (*Device) InstalledPackages

```go
func (d *Device) InstalledPackages() (Packages, error)
```
InstalledPackages returns the sorted list of installed packages on the device.

#### func (*Device) Pull

```go
func (d *Device) Pull(remote, local string) error
```
Pulls the remote file to the local one.

#### func (*Device) Push

```go
func (d *Device) Push(local, remote string) error
```
Pushes the local file to the remote one.

#### func (*Device) Root

```go
func (d *Device) Root() error
```
Root restarts adb as root. If the device is running a production build then Root
will return ErrDeviceNotRooted.

#### func (*Device) SELinuxEnforcing

```go
func (d *Device) SELinuxEnforcing() (bool, error)
```
SELinuxEnforcing returns true if the device is currently in a SELinux enforcing
mode, or false if the device is currently in a SELinux permissive mode.

#### func (*Device) SetSELinuxEnforcing

```go
func (d *Device) SetSELinuxEnforcing(enforce bool) error
```
SetSELinuxEnforcing changes the SELinux-enforcing mode.

#### func (*Device) StartActivity

```go
func (d *Device) StartActivity(a Action) error
```
StartActivity launches the specified action.

#### func (*Device) String

```go
func (d *Device) String() string
```
String returns a string representing the device.

#### type DeviceState

```go
type DeviceState int
```

DeviceState represents the last queried state of an Android device.

```go
const (
	Offline DeviceState = iota
	Online
	Unauthorized
)
```
binary: DeviceState#Offline = offline binary: DeviceState#Online = device
binary: DeviceState#Unauthorized = unauthorized

#### func (*DeviceState) Parse

```go
func (v *DeviceState) Parse(s string) error
```

#### func (DeviceState) String

```go
func (v DeviceState) String() string
```

#### type InstalledPackage

```go
type InstalledPackage struct {
	Name    string    // Name of the package.
	Device  *Device   // The device this package is installed on.
	Actions []*Action // The actions this package supports.
}
```


#### func (*InstalledPackage) SetWrapProperties

```go
func (p *InstalledPackage) SetWrapProperties(props ...string) error
```
WrapProperties sets the list of wrap-properties for the given installed package.

#### func (*InstalledPackage) String

```go
func (p *InstalledPackage) String() string
```
String returns the package name.

#### func (*InstalledPackage) WrapProperties

```go
func (p *InstalledPackage) WrapProperties() ([]string, error)
```
WrapProperties returns the list of wrap-properties for the given installed
package.

#### type NamedAbstractSocket

```go
type NamedAbstractSocket string
```

NamedAbstractSocket represents an abstract UNIX domain socket name on either the
local machine or Android device. NamedAbstractSocket implements the Port
interface.

#### type Packages

```go
type Packages []*InstalledPackage
```


#### func (Packages) Len

```go
func (l Packages) Len() int
```

#### func (Packages) Less

```go
func (l Packages) Less(i, j int) bool
```

#### func (Packages) Swap

```go
func (l Packages) Swap(i, j int)
```

#### type Port

```go
type Port interface {
	// contains filtered or unexported methods
}
```

Port is the interface for sockets ports that can be forwarded from an Android
Device to the local machine.

#### type TCPPort

```go
type TCPPort int
```

TCPPort represents a TCP/IP port on either the local machine or Android device.
TCPPort implements the Port interface.
