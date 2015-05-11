# template
--
    import "android.googlesource.com/platform/tools/gpu/api/apic/template"


## Usage

```go
var (
	// NodeTypeList exposes all the types valid as Is*name* tests
	NodeTypeList = []interface{}{

		"",
		false,

		semantic.BoolValue(true),
		semantic.Int32Value(0),
		semantic.Uint32Value(0),
		semantic.Int64Value(0),
		semantic.Uint64Value(0),
		semantic.StringValue(""),
		semantic.Float32Value(1.0),
		semantic.Float64Value(1.0),

		semantic.API{},
		semantic.ArrayIndex{},
		semantic.Array{},
		semantic.Assert{},
		semantic.Assign{},
		semantic.BinaryOp{},
		semantic.BitTest{},
		semantic.Branch{},
		semantic.Buffer{},
		semantic.Builtin{},
		semantic.Call{},
		semantic.Cast{},
		semantic.Choice{},
		semantic.ClassInitializer{},
		semantic.Class{},
		semantic.Copy{},
		semantic.DeclareLocal{},
		semantic.EnumEntry{},
		semantic.Enum{},
		semantic.Field{},
		semantic.Function{},
		semantic.Global{},
		semantic.Iteration{},
		semantic.Length{},
		semantic.Local{},
		semantic.MapAssign{},
		semantic.MapContains{},
		semantic.MapIndex{},
		semantic.Map{},
		semantic.Member{},
		semantic.New{},
		semantic.Null{},
		semantic.Observed{},
		semantic.Parameter{},
		semantic.Pointer{},
		semantic.Pseudonym{},
		semantic.Return{},
		semantic.Select{},
		semantic.StaticArray{},
		semantic.Switch{},
		semantic.UnaryOp{},
		semantic.Unknown{},

		(*semantic.Annotated)(nil),
		(*semantic.Expression)(nil),
		(*semantic.Type)(nil),
	}
)
```

#### type Functions

```go
type Functions struct {
}
```


#### func  NewFunctions

```go
func NewFunctions(apiFile string, api *semantic.API, loader func(filename string) ([]byte, error), funcs template.FuncMap) *Functions
```
NewFunctions builds a new template management object that can be used to run
templates over an API file. The apiFile name is used in error messages, and
should be the name of the file the api was loaded from. loader can be used to
intercept file system access from within the templates, specifically used when
including other templates. The functions in funcs are made available to the
templates, and can override the functions from this package if needed.

#### func (*Functions) AllCommands

```go
func (f *Functions) AllCommands(api interface{}) ([]interface{}, error)
```
AllCommands returns a list of all cmd entries for a given API, regardless of
whether they are free functions, class methods or pseudonym methods.

#### func (*Functions) AssertType

```go
func (*Functions) AssertType(v interface{}, expected ...string) (string, error)
```
Asserts that the type of v is in the list of expected types

#### func (Functions) Contains

```go
func (Functions) Contains(substr string, v ...interface{}) bool
```
Contains returns true if any string segment contains substr.

#### func (*Functions) Copyright

```go
func (f *Functions) Copyright(name string, tool string) (string, error)
```
Copyright emits the copyright header specified by name with the «Tool» set to
tool.

#### func (*Functions) Error

```go
func (f *Functions) Error(s string, args ...interface{}) (string, error)
```
Error raises an error terminating execution of the template.

    {{Error "Foo returned error: %s" $err}}

#### func (Functions) FilterOut

```go
func (Functions) FilterOut(v, from stringList) stringList
```
FilterOut returns from with all occurances of v removed.

#### func (*Functions) ForEach

```go
func (f *Functions) ForEach(arr interface{}, m string) (stringList, error)
```
ForEach returns a string list containing the strings emitted by calling the
macro m for each sequential item in the array arr. 0 length strings will be
ommitted from the returned list.

#### func (*Functions) Format

```go
func (f *Functions) Format(command stringList, value string) string
```
Format reflows the string using an external command.

#### func (*Functions) GetAnnotation

```go
func (*Functions) GetAnnotation(ty interface{}, name string) *semantic.Annotation
```
GetAnnotation finds and returns the annotation on ty with the specified name. If
the annotation cannot be found, or ty does not support annotations then
GetAnnotation returns nil.

#### func (*Functions) GetArrayParamCount

```go
func (*Functions) GetArrayParamCount(param *semantic.Parameter) interface{}
```
GetArrayParamCount returns the inferred array size for param as a semantic
expression. If the array size cannot be inferred, then GetArrayParamCount
returns nil.

#### func (*Functions) Global

```go
func (f *Functions) Global(name string, values ...interface{}) (interface{}, error)
```
Gets or sets a template global variable Example:

    {{Global "CatSays" "Meow"}}
    The cat says: {{Global "CatSays"}}

#### func (*Functions) GoFmt

```go
func (f *Functions) GoFmt(value string) string
```
GoFmt reflows the string as if it were go code using the standard go fmt
library.

#### func (*Functions) HasMore

```go
func (*Functions) HasMore(i int, l interface{}) bool
```
HasMore returns true if the i'th indexed item in l is not the last.

#### func (*Functions) Include

```go
func (f *Functions) Include(templates ...string) error
```
Include loads each of the templates and executes their main bodies. The
filenames are relative to the template doing the include.

#### func (*Functions) IndexOf

```go
func (*Functions) IndexOf(array interface{}, value interface{}) int
```
IndexOf returns the index of value in the list array. If value is not found in
array then IndexOf returns -1.

#### func (*Functions) IsNil

```go
func (f *Functions) IsNil(v interface{}) bool
```
IsNil returns true if v is nil.

#### func (*Functions) IsNumericValue

```go
func (*Functions) IsNumericValue(v interface{}) bool
```
Returns true if v is one of the primitive numeric types.

#### func (*Functions) Join

```go
func (*Functions) Join(slices ...interface{}) []interface{}
```
Join returns the concatenation of all the passed array or slices into a single,
sequential list.

#### func (Functions) JoinWith

```go
func (Functions) JoinWith(sep string, v ...interface{}) string
```
Join returns the concatenation of all the string segments with the specified
separator.

#### func (*Functions) Log

```go
func (f *Functions) Log(s string, args ...interface{}) string
```
Log prints s and optional format arguments to stdout. Example:

    {{Log "%s %s" "Hello" "world}}

#### func (Functions) Lower

```go
func (Functions) Lower(v ...interface{}) stringList
```
Lower lower-cases all letters of each string segment.

#### func (*Functions) Macro

```go
func (f *Functions) Macro(name string, arguments ...interface{}) (string, error)
```
Macro invokes the template macro with the specified name and returns the
template output as a string. If no arguments are passed then $ will be nil for
the called macro. If a single argument is passed then $ will be the value of
that argument. If more than one argument is passed then arguments is used as
name-value pairs, where name is a field on $. For example:

    {{define "SingleParameterMacro"}}
        $ is: {{$}}
    {{end}}

    {{define "MultipleParameterMacro"}}
        $.ArgA is: {{$.ArgA}}, $.ArgB is: {{$.ArgB}}
    {{end}}

    {{Macro "SingleParameterMacro"}}
    {{/* Returns "$ is: nil" */}}

    {{Macro "SingleParameterMacro" 42}}
    {{/* Returns "$ is: 42" */}}

    {{Macro "MultipleParameterMacro" "ArgA" 4 "ArgB" 2}}
    {{/* Returns "$.ArgA is: 4, $.ArgB is: 2" */}}

#### func (*Functions) Reflow

```go
func (f *Functions) Reflow(indentSize int, value string) string
```
Reflow does the primitive reflow, but no language specific handling.

#### func (Functions) Replace

```go
func (Functions) Replace(old string, new string, v ...interface{}) stringList
```
Replace any occurance of old with new in the string segments.

#### func (*Functions) Reverse

```go
func (*Functions) Reverse(in interface{}) interface{}
```
Reverse returns a new list with all the elements of in reversed.

#### func (*Functions) SortBy

```go
func (f *Functions) SortBy(arr interface{}, keyMacro string) (interface{}, error)
```
SortBy returns a new list containing the sorted elements of arr. The list is
sorted by the string values returned by calling the template function keyMacro
with each element in the list.

#### func (Functions) SplitOn

```go
func (Functions) SplitOn(sep string, v ...interface{}) stringList
```
Split slices each string segement into all substrings separated by sep. The
returned stringList will not contain any occurances of sep.

#### func (Functions) SplitPascalCase

```go
func (Functions) SplitPascalCase(v ...interface{}) stringList
```
SplitPascalCase slices each string segment at each transition from an letter
rune to a upper-case letter rune.

#### func (Functions) SplitUpperCase

```go
func (Functions) SplitUpperCase(v ...interface{}) stringList
```
SplitUpperCase slices each string segment before and after each upper-case rune.

#### func (Functions) Strings

```go
func (Functions) Strings(v ...interface{}) stringList
```
Strings returns the arguments as a string list.

#### func (*Functions) Tail

```go
func (*Functions) Tail(start int, array interface{}) interface{}
```
Tail returns a slice of the list from start to len(array).

#### func (Functions) Title

```go
func (Functions) Title(v ...interface{}) stringList
```
Title capitalizes each letter of each string segment.

#### func (Functions) TrimLeft

```go
func (Functions) TrimLeft(cutset string, v ...interface{}) stringList
```

#### func (Functions) TrimRight

```go
func (Functions) TrimRight(cutset string, v ...interface{}) stringList
```

#### func (*Functions) TypeOf

```go
func (*Functions) TypeOf(v interface{}) (semantic.Type, error)
```
Returns the resolved semantic type of an expression node.

#### func (*Functions) Underlying

```go
func (f *Functions) Underlying(ty semantic.Type) semantic.Type
```
Underlying returns the underlying type for ty by recursively traversing the
pseudonym chain until reaching and returning the first non-pseudoym type. If ty
is not a pseudonym then it is simply returned.

#### func (*Functions) Unpack

```go
func (*Functions) Unpack(v interface{}) interface{}
```
Unpack throws away type information to work around template system limitations
When you have a value of an interface type that carries methods, it fails to
introspect the concrete type for it's members, so the template can't see them.
The result of Upack no longer has a type, so the concrete type members become
visible.

#### func (Functions) Untitle

```go
func (Functions) Untitle(v ...interface{}) stringList
```
Untitle lower-cases each letter of each string segment.

#### func (Functions) Upper

```go
func (Functions) Upper(v ...interface{}) stringList
```
Upper upper-cases all letters of each string segment.

#### func (*Functions) Write

```go
func (f *Functions) Write(fileName string, value string) (string, error)
```
Write takes a string and writes it into the specified file. The filename is
relative to the output directory.
