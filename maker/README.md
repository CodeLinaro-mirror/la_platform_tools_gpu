# maker
--
    import "android.googlesource.com/platform/tools/gpu/maker"

Package maker provides a system for building and running a dependancy graph of
tasks.

## Usage

```go
const (
	// Default is the name of the entity that is built if none are supplied on the
	// command line.
	Default = "default"
)
```

```go
const GoPkgResources = "go_packages"
```

```go
const HostExecutableExtension = ""
```

```go
const HostExecutableExtension = ""
```

```go
const HostExecutableExtension = ".exe"
```

```go
const HostOS = "linux"
```

```go
const HostOS = "osx"
```

```go
const HostOS = "windows"
```

```go
var (
	// Config holds the current configuration of the maker system.
	Config struct {
		// Verbose enables increased logging output.
		Verbose int
		// DisableParallel turns of all parallel build support.
		DisableParallel bool
		// StopOnError makes the system quit faster once an error has been found.
		StopOnError bool
	}
	// Paths holds the set of path roots for the build.
	Paths struct {
		// The root path of the build
		Root string
		// The dependancy cache directory
		Deps string
		// The application binary directory.
		Bin string
	}
	//GoPath is the GOPATH environment setting
	GoPath []string
)
```

```go
var (
	// EnvVars holds the environment overrides used when spawning external commands.
	EnvVars = map[string][]string{}
)
```

```go
var (
	// Errors is the main error list, it holds all errors that the system has
	// encountered.
	Errors errors
)
```

#### func  AddEntity

```go
func AddEntity(e Entity)
```
AddEntity adds a new entity into the system. The entity must not have the same
name as an already exixting entity.

#### func  CommonPath

```go
func CommonPath(path ...string) string
```
CommonPath returns the '/' delimited file-path formed from the joined path
segments, correcting non-forward slash delimiters found in any path segment.

#### func  CopyFile

```go
func CopyFile(dst, src Entity)
```
CopyFile builds a new Step that copies the file from src to dst, and returns the
dst entity. The step will depend on the src, and the existance of the directory
dst is inside.

#### func  Dir

```go
func Dir(path ...string) *file
```
Dir returns an Entity that represents a directory. The entities name will be the
absolute path of the directory. The entity map will be checked for a matching
directory entry and if one is not found, a new one will be added and returned.
It also adds the rules to create the directory if needed.

#### func  DirOf

```go
func DirOf(path interface{}) *file
```
DirOf attempts to make a Dir for the parent of the specified File.

#### func  EntityHook

```go
func EntityHook(f func(e Entity))
```
EntityHook registers a function that is invoked when a new entity is added to
the system.

#### func  ExecAt

```go
func ExecAt(wd string, verbose int, path string, args ...string) error
```
ExecAt executes "path" with the specified arguments with the working directory
set to "wd".

#### func  File

```go
func File(path ...string) *file
```
File returns an Entity that represents a file. The entities name will be the
absolute path of the file. The entity map will be checked for a matching file
entry and if one is not found, a new one will be added and returned.

#### func  FilesOf

```go
func FilesOf(path string, filter func(os.FileInfo) bool) []*file
```
FilesOf reads the list of files in path, filters them with the supplied filter
and returns the set of file entities that matched.

#### func  FindEntities

```go
func FindEntities(pattern string) []Entity
```
FindEntities tries to find all entities who's names match the supplied string.
The string is treated as a case insensitive regular expression.

#### func  GoSrcPath

```go
func GoSrcPath(path string) string
```
GoSrcPath returns the full path to a file or directory inside the GoPath.

#### func  IsFile

```go
func IsFile(e Entity) bool
```
IsFile returns true if the supplied entity is of file type.

#### func  IsVirtual

```go
func IsVirtual(e Entity) bool
```
IsVirtual returns true if the supplied entity is a virtual type.

#### func  Newest

```go
func Newest(t1, t2 time.Time) time.Time
```
Newest returns the newest of two timestamps. Timestamps with a zero value do not
count as newer.

#### func  OSPath

```go
func OSPath(path ...string) (string, error)
```
OSPath returns the OS delimited file-path formed from the joined path segments.

#### func  Oldest

```go
func Oldest(t1, t2 time.Time) time.Time
```
Oldest returns the oldest of two timestamps. Timestamps with a zero value do not
count as older.

#### func  Path

```go
func Path(path ...string) string
```
Path returns the '/' delimited file-path formed from the joined path segments,
without correcting non-forward slash delimiters found in any path segment.

#### func  PathSplit

```go
func PathSplit(path ...string) (dir, base string)
```

#### func  Register

```go
func Register(f func())
```
Register a new graph building function with the maker system. The function will
be invoked during Run to add entities and steps to the build graph.

#### func  Run

```go
func Run()
```
Run should be invoked once from main. It parses the command line, builds the
graph, and then performs the required action.

#### func  StepHook

```go
func StepHook(f func(s *Step))
```
StepHook adds a function that is invoked each time a step is added or modified.

#### func  Virtual

```go
func Virtual(name string) *virtual
```
Virtual returns an Entity that represents a virtual object in the build graph.
If name already represents a virtual entity, it will be returned, otherwise a
new one will be built. This is used as the output for commands that have no file
outputs.

#### type Entity

```go
type Entity interface {
	// Name returns the unique name of this entity.
	// If Name is the empty string, the entity does not appear in the main entity
	// list and cannot be looked up.
	Name() string
	// Timestamp returns the last modified time of the entity, or the zero time if
	// not available.
	Timestamp() time.Time
	// NeedsUpdate returns true if the entity requires it's generating step to run.
	// The supplied timestamp is the newest timestamp of the entities this one
	// depends on, and can be used to decide if an update is neccesary.
	NeedsUpdate(t time.Time) bool
	// Updated is called when a step that modifies this entity completes.
	Updated()
}
```

Entity represents an object in the build graph.

#### func  EntityOf

```go
func EntityOf(v interface{}) Entity
```
EntityOf tries to find an entity for the supplied value. If the value is an
entity, it will be returned directly. If the value is a string it will use
FindPathEntity to look it up. It is an error for v to not match an existing
entity.

#### func  FindEntity

```go
func FindEntity(name string) Entity
```
FindEntity tries to look up an entity by name, and returns nil if the entity
cannot be found.

#### func  FindPathEntity

```go
func FindPathEntity(name string) Entity
```
FindPathEntity tries to look up an entity by name. If an exact match cannot be
found, then the name is treated as a path, and the absolute path is looked up
instead. If that also does not match, it returns nil.

#### func  FindTool

```go
func FindTool(name string) Entity
```
FindTool finds an executable on the host search path, and returns a File entity
for the tool if found.

#### func  GoInstall

```go
func GoInstall(module string) Entity
```
GoInstall builds a new Step that runs "go install" on the supplied module. It
will return the resulting binary entity. The step will depend on the go tool,
and will be set to always run if depended on.

#### func  GoTest

```go
func GoTest(module string) Entity
```
GoTest creates a new Step that runs "go test" on the supplied module. It returns
a virtual entity that represents the test output.

#### type Step

```go
type Step struct {
}
```

Step represents an action that creates it's outputs from it's inputs. It is the
linking entity in the build graph. Each entity can be built by only one step,
but a step may build many entites, and may depend on many entities. It is an
error to have a cycle in the build graph.

#### func  Command

```go
func Command(binary Entity, args ...string) *Step
```
Command builds and returns a new Step that runs the specified external binary
with the supplied arguments. The newly created Step will be made to depend on
the binary.

#### func  Creator

```go
func Creator(of interface{}) *Step
```
Creator looks up the step that builds an entity. The entity will be looked up
using EntityOf.

#### func  GoCommand

```go
func GoCommand(args ...string) *Step
```
GoCommand runs "go" with the specified arguments.

#### func  GoRun

```go
func GoRun(gofile string, args ...string) *Step
```
GoRun returns a Step that runs "go run" with the supplied go file and arguments.

#### func  List

```go
func List(name string) *Step
```
List returns the Step for a virtual entity. If the virtual entity does not
exist, it will be created, and if it does not have a Creator, then a new Step
with no action will be created. This is used to build lists of entities that
want to be updated together.

#### func  NewStep

```go
func NewStep(a func(*Step) error) *Step
```
NewStep creates and returns a new step that runs the supplied action.

#### func (*Step) Access

```go
func (s *Step) Access(v ...string) *Step
```
Access adds the supplied shared resource names to the list of resources accessed
by this step.

#### func (*Step) AlwaysRun

```go
func (s *Step) AlwaysRun() *Step
```
AlwaysRun marks a step as always needing to run, rather than running only when
it's outputs need updating.

#### func (*Step) Creates

```go
func (s *Step) Creates(out ...Entity) *Step
```
Creates adds new outputs to the Step. It will complain if any of the outputs
already have a Creator.

#### func (*Step) DependsOn

```go
func (s *Step) DependsOn(in ...interface{}) *Step
```
DependsOn adds a new input dependancy to a Step. If the Step already depends on
an input, it will not be added again.

#### func (*Step) DependsStruct

```go
func (s *Step) DependsStruct(v interface{})
```
DependsStruct adds all the fields of the supplied struct as input dependacies of
the step. It is an error if the struct has any fields that are not public fields
of types that implement Entity.

#### func (*Step) Disable

```go
func (s *Step) Disable() *Step
```
Disable marks a step as disabled, so it will not run.

#### func (*Step) HasInput

```go
func (s *Step) HasInput(e Entity) bool
```
HasInput returns true if the step already has the supplied entity in it's
inputs.

#### func (*Step) HasOutput

```go
func (s *Step) HasOutput(e Entity) bool
```
HasOutput returns true if the step already has the supplied entity in it's
outputs.

#### func (*Step) String

```go
func (s *Step) String() string
```
String returns the name of the first output if present, for logging.

#### func (*Step) UseDepsFile

```go
func (s *Step) UseDepsFile(deps Entity)
```
UseDepsFile reads in a deps file and uses it to populuate the inputs and outputs
of the Step.
