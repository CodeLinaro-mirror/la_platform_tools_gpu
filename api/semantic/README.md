# semantic
--
    import "android.googlesource.com/platform/tools/gpu/api/semantic"

Package semantic holds the set of types used in the abstract semantic graph
representation of the api language.

## Usage

```go
var (

	// Special types
	VoidType   = builtin("void")
	AnyType    = builtin("any")
	StringType = builtin("string")
	// Unsized primitives
	BoolType = builtin("bool")
	CharType = builtin("char")
	IntType  = builtin("int")
	UintType = builtin("uint")
	// Fixed size integer forms
	Int8Type   = builtin("s8")
	Uint8Type  = builtin("u8")
	Int16Type  = builtin("s16")
	Uint16Type = builtin("u16")
	Int32Type  = builtin("s32")
	Uint32Type = builtin("u32")
	Int64Type  = builtin("s64")
	Uint64Type = builtin("u64")
	// Floating point forms
	Float32Type = builtin("f32")
	Float64Type = builtin("f64")
)
```

```go
var BuiltinTypes []*Builtin
```

#### type API

```go
type API struct {
	AST          *ast.API       // the underlying syntax node this was built from
	Enums        []*Enum        // the set of enums
	Classes      []*Class       // the set of classes
	Pseudonyms   []*Pseudonym   // the set of pseudo types
	Externs      []*Function    // the external function references
	Functions    []*Function    // the global functions
	Globals      []*Global      // the global variables
	Arrays       []*Array       // the array types used
	StaticArrays []*StaticArray // the fixed size array types used
	Maps         []*Map         // the map types used
	Pointers     []*Pointer     // the pointer types used
	Buffers      []*Buffer      // the buffer types used
	Signatures   []*Signature   // the function signature types used
	Members                     // a map of name to member for top level symbols
}
```

API is the root of the ASG, and holds a fully resolved api.

#### type Annotated

```go
type Annotated interface {
	// GetAnnotation returns the annotation with the matching name, if present.
	GetAnnotation(name string) *Annotation
}
```

Annotated is the common interface to objects that can carry annotations.

#### type Annotation

```go
type Annotation struct {
	AST       *ast.Annotation // the underlying syntax node this was built from
	Name      string          // the name of the annotation
	Arguments []Expression    // the arguments to the annotation
}
```


#### type Annotations

```go
type Annotations []*Annotation
```

Annotations is an array of Annotation objects that implements the Annotated
interface. It is used as an anonymous field on objects that carry annotations.

#### func (Annotations) GetAnnotation

```go
func (a Annotations) GetAnnotation(name string) *Annotation
```
GetAnnotation implements the Annotated interface for the Annotations type.

#### type Array

```go
type Array struct {
	Name      string // the full name of the array type
	ValueType Type   // the value type stored in the array
}
```

Array represents an array type declaration.

#### func (Array) Member

```go
func (t Array) Member(name string) Node
```

#### func (Array) Typename

```go
func (t Array) Typename() string
```

#### type ArrayIndex

```go
type ArrayIndex struct {
	AST       *ast.Index // the underlying syntax node this was built from
	ValueType Type       // the value type of the array being indexed
	Array     Expression // the expression that returns the array to be indexed
	Index     Expression // the index to use on the array
}
```

ArrayIndex represents using the indexing operator on an array type.

#### func (*ArrayIndex) ExpressionType

```go
func (i *ArrayIndex) ExpressionType() Type
```
ExpressionType implements Expression returning the value type of the array.

#### type Assert

```go
type Assert struct {
	AST       *ast.Assert // the underlying syntax node this was built from
	Condition Expression  // the condition is being asserted must be true
}
```

Assert represents a runtime assertion. Assertions are also used to infer
required behavior from the expressions.

#### type Assign

```go
type Assign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	LHS      Expression  // the expression that gives the location to store into
	Operator string      // the assignment operator being applied
	RHS      Expression  // the value to store
}
```

Assign is the only "mutating" construct. It assigns the value from the rhs into
the slot described by the lhs, as defined by the operator.

#### type BinaryOp

```go
type BinaryOp struct {
	AST      *ast.BinaryOp // the underlying syntax node this was built from
	Type     Type          // the resolved type of this binary expression
	LHS      Expression    // the expression that appears on the left of the operator
	Operator string        // the operator being applied
	RHS      Expression    // the expression that appears on the right of the operator
}
```

BinaryOp represents any operator applied to two arguments. The resolved type of
the expression depends on the types of the two arguments and which operator it
represents.

#### func (*BinaryOp) ExpressionType

```go
func (b *BinaryOp) ExpressionType() Type
```
ExpressionType implements Expression

#### type BitTest

```go
type BitTest struct {
	AST      *ast.BinaryOp // the underlying syntax node this was built from
	Bitfield Expression    // the bitfield being tested
	Bits     Expression    // the bits to test for
}
```

BitTest is the "in" operator applied to a bitfield.

#### func (BitTest) ExpressionType

```go
func (BitTest) ExpressionType() Type
```
ExpressionType implements Expression

#### type Block

```go
type Block struct {
	AST        *ast.Block // the underlying syntax node this was built from
	Statements []Node     // the set of statements this block represents
}
```

Block represents a collection of statements, used as the body of other nodes.

#### type BoolValue

```go
type BoolValue bool
```

BoolValue is a bool that implements Expression so it can be in the semantic
graph

#### func (BoolValue) ExpressionType

```go
func (v BoolValue) ExpressionType() Type
```
ExpressionType implements Expression with a type of BoolType

#### type Branch

```go
type Branch struct {
	AST       *ast.Branch // the underlying syntax node this was built from
	Condition Expression  // the condition to select on
	True      *Block      // use if Condition is true
	False     *Block      // used if Condition is false
}
```

Branch represents the basic conditional execution statement. If Condition is
true we use the True block, otherwise the False block.

#### type Buffer

```go
type Buffer struct {
	Name      string // the full type name
	To        Type   // the type this is a pointer to
	Array     bool   // points to multiple elements, rather than one
	FakeArray *Array // TODO:Remove - Hold a fake array for now as a schema compatibility measure
}
```

Buffer represents a state pointer type declaration.

#### func (Buffer) Member

```go
func (t Buffer) Member(name string) Node
```

#### func (Buffer) Typename

```go
func (t Buffer) Typename() string
```

#### type Builtin

```go
type Builtin struct {
	Name string // the primitive type name
}
```

Builtin represents one of the primitive types.

#### func (Builtin) Member

```go
func (t Builtin) Member(name string) Node
```

#### func (Builtin) Typename

```go
func (t Builtin) Typename() string
```

#### type Call

```go
type Call struct {
	AST       *ast.Call    // the underlying syntax node this was built from
	Target    *Callable    // the function expression this invokes
	Arguments []Expression // the arguments to pass to the function
	Type      Type         // the return type of the call
}
```

Call represents a function call. It binds an Callable to the arguments it will
be passed.

#### func (*Call) ExpressionType

```go
func (c *Call) ExpressionType() Type
```
ExpressionType implements Expression returning the underlying function return
type.

#### type Callable

```go
type Callable struct {
	Object   Expression // the object to use as the this parameter for a method
	Function *Function  // the function this expression represents
}
```

Callable wraps a Function declaration into a first class function value
expression, optionally binding to an object if its a method.

#### func (*Callable) ExpressionType

```go
func (c *Callable) ExpressionType() Type
```
ExpressionType implements Expression returning the function type signature.

#### type Case

```go
type Case struct {
	AST        *ast.Case    // the underlying syntax node this was built from
	Conditions []Expression // the set of expressions to match the switch value against
	Block      *Block       // the block to use if a condition matches
}
```

Case represents a possible choice in a switch.

#### type Cast

```go
type Cast struct {
	AST    *ast.Cast  // the underlying syntax node this was built from
	Object Expression // the actual expression being wrapped
	Type   Type       // the type to coerce the expression to
}
```

Cast represents a type coercion expression. It reports it's type as the one
specified, rather than the expression it wraps.

#### func (*Cast) ExpressionType

```go
func (c *Cast) ExpressionType() Type
```
ExpressionType implements Expression returning the type being cast to.

#### type Choice

```go
type Choice struct {
	AST        *ast.Case    // the underlying syntax node this was built from
	Conditions []Expression // the set of expressions to match the select value against
	Expression Expression   // the expression to use if a condition matches
}
```

Choice represents a possible choice in a select

#### type Class

```go
type Class struct {
	AST         *ast.Class  // the underlying syntax node this was built from
	Annotations             // the annotations applied to this class
	Name        string      // the name of the class
	Docs        []string    // the documentation for the class
	Extends     []*Class    // the classes this extends
	ExtendedBy  []*Class    // the classes that declared they extended this class
	Fields      []*Field    // the set of fields the class declares
	Methods     []*Function // the set of functions associated with the class
	Members                 // the name->member map, to implement Type
}
```

Class represents an api class construct.

#### func (Class) Typename

```go
func (t Class) Typename() string
```
Implements Type to return the class name as the type name

#### type ClassInitializer

```go
type ClassInitializer struct {
	AST    *ast.ClassInitializer // the underlying syntax node this was built from
	Class  *Class                // the class to initialize
	Fields []*FieldInitializer   // the set of field assignments
}
```

ClassInitializer represents an expression that can assign values to multiple
fields of a class.

#### func (*ClassInitializer) ExpressionType

```go
func (c *ClassInitializer) ExpressionType() Type
```
ExpressionType implements Expression returning the class type being initialized.

#### type Copy

```go
type Copy struct {
	AST *ast.Assign // the underlying syntax node this was built from
	Dst *Slice      // the slice to copy to
	Src *Slice      // the slice to copy from
}
```

Copy is the special form of assign that copies data between slices. One of LHS
or RHS may be missing, and if both are present the upper bound on one may be
inferred from the other.

#### type DeclareLocal

```go
type DeclareLocal struct {
	AST   *ast.DeclareLocal // the underlying syntax node this was built from
	Local *Local            // the local variable that was declared by this statement
}
```

DeclareLocal represents a local variable declaration statement. Variables cannot
be modified after declaration.

#### type Enum

```go
type Enum struct {
	AST         *ast.Enum    // the underlying syntax node this was built from
	Annotations              // the annotations applied to this enum
	Name        string       // the type name of the enum
	Docs        []string     // the documentation for the enum
	IsBitfield  bool         // whether this enum is actually a bitfield
	Extends     []*Enum      // the enums this enum extends
	Entries     []*EnumEntry // the entries of this enum
	AllEntries  []*EnumEntry // the flattened list of all entries including inherited ones
}
```

Enum represents the api enum construct.

#### func (Enum) Member

```go
func (t Enum) Member(name string) Node
```
Implements Type returning the matching enum entry if there is one.

#### func (Enum) Typename

```go
func (t Enum) Typename() string
```
Implements Type to return the enum name as the type name

#### type EnumEntry

```go
type EnumEntry struct {
	AST   *ast.EnumEntry // the underlying syntax node this was built from
	Enum  *Enum          // the enum this entry belongs to
	Name  string         // the name of this entry
	Docs  []string       // the documentation for the enum entry
	Value uint32         // the value this entry represents
}
```

EnumEntry represents a single entry in an Enum.

#### func (*EnumEntry) ExpressionType

```go
func (e *EnumEntry) ExpressionType() Type
```
ExpressionType implements Expression returning the enum type.

#### type Expression

```go
type Expression interface {
	Node
	ExpressionType() Type // returns the expression value type.
}
```

Expression represents anything that can act as an expression in the api
language, it must be able to correctly report the type of value it would return
if executed.

#### type Field

```go
type Field struct {
	AST         *ast.Field // the underlying syntax node this was built from
	Annotations            // the annotations applied to this field
	Class       *Class     // the class this field belongs to
	Type        Type       // the type the field stores
	Name        string     // the name of the field
	Docs        []string   // the documentation for the field
	Default     Expression // the default value of the field
}
```

Field represents a field entry in a class.

#### func (*Field) ExpressionType

```go
func (f *Field) ExpressionType() Type
```
Implements Expression to return the type stored in the field.

#### type FieldInitializer

```go
type FieldInitializer struct {
	AST   *ast.FieldInitializer // the underlying syntax node this was built from
	Field *Field                // the field to assign to
	Value Expression            // the value to assign
}
```

FieldInitializer

#### type Float32Value

```go
type Float32Value float32
```

Float32Value is a float32 that implements Expression so it can be in the
semantic graph

#### func (Float32Value) ExpressionType

```go
func (v Float32Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Float32Type

#### type Float64Value

```go
type Float64Value float64
```

Float64Value is a float64 that implements Expression so it can be in the
semantic graph

#### func (Float64Value) ExpressionType

```go
func (v Float64Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Float64Type

#### type Function

```go
type Function struct {
	AST            *ast.Function // the underlying syntax node this was built from
	Annotations                  // the annotations applied to the function
	Name           string        // the name of the function
	Docs           []string      // the documentation for the function
	Owner          Type          // the owner of the function
	Return         *Parameter    // the return parameter
	This           *Parameter    // the this parameter, missing for non method functions
	FullParameters []*Parameter  // all the parameters, including This at the start if valid, and Return at the end if not void
	Outputs        []*Parameter  // only the output parameters
	Block          *Block        // the body of the function, missing for externs
	Signature      *Signature    // the type signature of the function
}
```

Function represents function like objects in the semantic graph.

#### func (*Function) CallParameters

```go
func (f *Function) CallParameters() []*Parameter
```
CallParameters returns the full set of parameters with the return value filtered
out.

#### type Global

```go
type Global struct {
	AST         *ast.Field // the underlying syntax node this was built from
	Annotations            // the annotations applied to this global
	Type        Type       // the type the global stores
	Name        string     // the name of the global
	Default     Expression // the initial value of the global
}
```

Global represents a global variable.

#### func (*Global) ExpressionType

```go
func (g *Global) ExpressionType() Type
```
Implements Expression to return the type stored in the global.

#### type Ignore

```go
type Ignore struct {
	AST ast.Node // the underlying syntax node this was built from
}
```

Ignore represents an _ expression.

#### func (Ignore) ExpressionType

```go
func (i Ignore) ExpressionType() Type
```
ExpressionType implements Expression.

#### type Int16Value

```go
type Int16Value int16
```

Int16Value is an int16 that implements Expression so it can be in the semantic
graph

#### func (Int16Value) ExpressionType

```go
func (v Int16Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Int16Type

#### type Int32Value

```go
type Int32Value int32
```

Int32Value is an int32 that implements Expression so it can be in the semantic
graph

#### func (Int32Value) ExpressionType

```go
func (v Int32Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Int32Type

#### type Int64Value

```go
type Int64Value int64
```

Int64Value is an int64 that implements Expression so it can be in the semantic
graph

#### func (Int64Value) ExpressionType

```go
func (v Int64Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Int64Type

#### type Int8Value

```go
type Int8Value int8
```

Int8Value is an int8 that implements Expression so it can be in the semantic
graph

#### func (Int8Value) ExpressionType

```go
func (v Int8Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Int8Type

#### type Iteration

```go
type Iteration struct {
	AST      *ast.Iteration // the underlying syntax node this was built from
	Iterator *Local         // the iteration control variable
	Iterable Expression     // the expression to iterate over
	Block    *Block         // the block to run for each entry from Iterable
}
```

Iteration is the basic looping construct. It will set Iterator to each value
from Iterable in turn, and run Block for each one.

#### type Length

```go
type Length struct {
	AST    *ast.Length // the underlying syntax node this was built from
	Object Expression  // the object go get the length of
	Type   Type        // the resolved type of the length operation
}
```

Represents a length of object expression. Object must be of either Array, Map or
string type. The length expression is allowed to be of any numeric type

#### func (Length) ExpressionType

```go
func (l Length) ExpressionType() Type
```
ExpressionType implements Expression

#### type Local

```go
type Local struct {
	Declaration *DeclareLocal // the statement that created the local
	Type        Type          // the type of the storage
	Name        string        // the identifier that will resolve to this local
	Value       Expression    // the expression the local was assigned on creation
}
```

Local represents an immutable local storage slot, created by a DeclareLocal, and
referred to by identifiers that resolve to that slot.

#### func (*Local) ExpressionType

```go
func (l *Local) ExpressionType() Type
```
ExpressionType implements Expression

#### type Map

```go
type Map struct {
	Name      string // the full type name
	KeyType   Type   // the type used as an indexing key
	ValueType Type   // the type stored in the map
	Members          // holds the map built-in methods
}
```

Map represents an api map type declaration.

#### func (Map) Typename

```go
func (t Map) Typename() string
```

#### type MapAssign

```go
type MapAssign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	To       *MapIndex   // the map index to assign to
	Operator string      // the assignment operator being applied
	Value    Expression  // the value to set in the map
}
```

MapAssign represents assigning to a map index expression.

#### type MapContains

```go
type MapContains struct {
	AST *ast.BinaryOp // the underlying syntax node this was built from
	Map Expression    // the map being tested
	Key Expression    // the key to test for
}
```

MapContains is the "in" operator applied to a map.

#### func (MapContains) ExpressionType

```go
func (MapContains) ExpressionType() Type
```
ExpressionType implements Expression

#### type MapIndex

```go
type MapIndex struct {
	AST       *ast.Index // the underlying syntax node this was built from
	ValueType Type       // the value type of the array being indexed
	Map       Expression // the expression that returns the map to be indexed
	Index     Expression // the index to use on the map
}
```

MapIndex represents using the indexing operator on a map type.

#### func (*MapIndex) ExpressionType

```go
func (i *MapIndex) ExpressionType() Type
```
ExpressionType implements Expression returning the value type of the map.

#### type Member

```go
type Member struct {
	AST    *ast.Member // the underlying syntax node this was built from
	Object Expression  // the object to look up a field of
	Field  *Field      // the field to look up
}
```

Member is an expression that looks up a field by name from an object.

#### func (*Member) ExpressionType

```go
func (m *Member) ExpressionType() Type
```
ExpressionType implements Expression returning the type of the field.

#### type Members

```go
type Members map[string]Node
```

Members wraps a map and implements part of the Type interface. It is used as a
mixin helper.

#### func (Members) Member

```go
func (t Members) Member(name string) Node
```
Member returns the entry in the map that matches name, or nil if none does.

#### type New

```go
type New struct {
	AST         *ast.New          // the underlying syntax node this was built from
	Initializer *ClassInitializer // The initialization for the new instance
	Type        Type              // The pointer type returned from the new
}
```

New represents an expression that allocates a new class instance.

#### func (*New) ExpressionType

```go
func (n *New) ExpressionType() Type
```
ExpressionType implements Expression returning a pointer to the class type being
initialized.

#### type Node

```go
type Node interface{}
```

Node represents any semantic-tree node type.

#### type Null

```go
type Null struct {
	AST  *ast.Null // the underlying syntax node this was built from
	Type Type      // the resolved type of this null
}
```

Null represents a default value.

#### func (Null) ExpressionType

```go
func (n Null) ExpressionType() Type
```
ExpressionType implements Expression with the inferred type of the null.

#### type Observed

```go
type Observed struct {
	Parameter *Parameter // the output parameter to infer from
}
```

Observed represents the final observed value of an output parameter. It is never
produced directly from the ast, but is inserted when inferring the value of an
unknown from observed outputs.

#### func (*Observed) ExpressionType

```go
func (e *Observed) ExpressionType() Type
```
ExpressionType implements Expression for observed parameter lookup.

#### type Parameter

```go
type Parameter struct {
	AST         *ast.Parameter // the underlying syntax node this was built from
	Annotations                // the annotations applied to the parameter
	Function    *Function      // the function this parameter belongs to
	Name        string         // the name of the parameter
	Docs        []string       // the documentation for the parameter
	Input       bool           // true if the parameter is an input
	Output      bool           // true if the parameter is an output
	Type        Type           // the type of the parameter
}
```

Parameter represents a single parameter declaration for a function.

#### func (*Parameter) ExpressionType

```go
func (p *Parameter) ExpressionType() Type
```
ExpressionType implements Expression for parameter lookup.

#### func (*Parameter) IsThis

```go
func (p *Parameter) IsThis() bool
```
IsThis returns true if this parameter is the This parameter of it's function.

#### type Pointer

```go
type Pointer struct {
	Name  string // the full type name
	To    Type   // the type this is a pointer to
	Array bool   // points to multiple elements, rather than one
}
```

Pointer represents an api pointer type declaration.

#### func (Pointer) Member

```go
func (t Pointer) Member(name string) Node
```

#### func (Pointer) Typename

```go
func (t Pointer) Typename() string
```

#### type Pseudonym

```go
type Pseudonym struct {
	AST         *ast.Pseudonym // the underlying syntax node this was built from
	Annotations                // the annotations applied to this pseudonym
	Name        string         // the type name
	Docs        []string       // the documentation for the pseudonym
	To          Type           // the underlying type
	Methods     []*Function    // the methods added directly to the pseudonym
	Members                    // the direct members
}
```

Pseudonym represents the type construct. It acts as a type in it's own right
that can carry methods, but is defined in terms of another type.

#### func (Pseudonym) Member

```go
func (t Pseudonym) Member(name string) Node
```
Implements Type returning the direct member if it has it, otherwise delegating
the lookup to the underlying type.

#### func (Pseudonym) Typename

```go
func (t Pseudonym) Typename() string
```
Implements Type to return the type name

#### type Return

```go
type Return struct {
	AST      *ast.Return // the underlying syntax node this was built from
	Function *Function   // the function this statement returns from
	Value    Expression  // the value to be returned
}
```

Return represents return statement for a function or macro.

#### type Select

```go
type Select struct {
	AST     *ast.Switch // the underlying syntax node this was built from
	Type    Type        // the return type of the select if valid
	Value   Expression  // the value to match the cases against
	Choices []*Choice   // The set of possible choices to match
}
```

Select is the expression form of a switch.

#### func (*Select) ExpressionType

```go
func (s *Select) ExpressionType() Type
```
ExpressionType implements Expression with the unified type of the choices

#### type Signature

```go
type Signature struct {
	Name      string // the full type name
	Return    Type   // the return type of the callable
	Arguments []Type // the required callable arguments
}
```

Signature represents a callable type signature

#### func (Signature) Member

```go
func (Signature) Member(name string) Node
```

#### func (Signature) Typename

```go
func (t Signature) Typename() string
```

#### type Slice

```go
type Slice struct {
	AST   *ast.Index // the underlying syntax node this was built from
	Array Expression // the expression that returns the array to be indexed
	Lower Expression // the inclusive lower bound to slice at
	Upper Expression // the non-inclusive upper bound to slice at
}
```

Slice represents using the slicing operator on an array type.

#### func (*Slice) ExpressionType

```go
func (i *Slice) ExpressionType() Type
```
ExpressionType implements Expression. It returns VoidType as slices are only
valid in Copy assignments.

#### type StaticArray

```go
type StaticArray struct {
	Name      string // the full type name
	ValueType Type   // the storage type of the elements
	Size      uint32 // the dimension of the array
}
```

StaticArray represents a multi-dimensional fixed size array type.

#### func (StaticArray) Member

```go
func (t StaticArray) Member(name string) Node
```

#### func (StaticArray) Typename

```go
func (t StaticArray) Typename() string
```

#### type StringValue

```go
type StringValue string
```

StringValue is a string that implements Expression so it can be in the semantic
graph

#### func (StringValue) ExpressionType

```go
func (v StringValue) ExpressionType() Type
```
ExpressionType implements Expression with a type of StringType

#### type Switch

```go
type Switch struct {
	AST   *ast.Switch // the underlying syntax node this was built from
	Value Expression  // the value to match the cases against
	Cases []*Case     // the set of case statements to choose from
}
```

Switch represents a resolved ast.Switch statement.

#### type Type

```go
type Type interface {
	Node
	Typename() string        // returns the full name of the type, must be unique
	Member(Name string) Node // looks up a member by name from a type
}
```

Type is the interface to any object that can act as a type to the api langauge.

#### type Uint16Value

```go
type Uint16Value uint16
```

Uint16Value is a uint16 that implements Expression so it can be in the semantic
graph

#### func (Uint16Value) ExpressionType

```go
func (v Uint16Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Uint16Type

#### type Uint32Value

```go
type Uint32Value uint32
```

Uint32Value is a uint32 that implements Expression so it can be in the semantic
graph

#### func (Uint32Value) ExpressionType

```go
func (v Uint32Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Uint32Type

#### type Uint64Value

```go
type Uint64Value uint64
```

Uint64Value is a uint64 that implements Expression so it can be in the semantic
graph

#### func (Uint64Value) ExpressionType

```go
func (v Uint64Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Uint64Type

#### type Uint8Value

```go
type Uint8Value uint8
```

Uint8Value is a uint8 that implements Expression so it can be in the semantic
graph

#### func (Uint8Value) ExpressionType

```go
func (v Uint8Value) ExpressionType() Type
```
ExpressionType implements Expression with a type of Uint8Type

#### type UnaryOp

```go
type UnaryOp struct {
	AST        *ast.UnaryOp // the underlying syntax node this was built from
	Type       Type         // the resolved type of the operation
	Operator   string       // the operator being applied
	Expression Expression   // the expression to apply the operator to
}
```

UnaryOp represents an operator applied to a single expression. It's type depends
on the operator and the type of the expression it is being applied to.

#### func (*UnaryOp) ExpressionType

```go
func (b *UnaryOp) ExpressionType() Type
```
ExpressionType implements Expression

#### type Unknown

```go
type Unknown struct {
	AST      *ast.Unknown // the underlying syntax node this was built from
	Inferred Expression   // the inferred expression to derive the unknown from the outputs
}
```

Unknown represents a value that cannot be predicted. These values are
non-deterministic with regard to the API specification and may vary between
implementations of the API.

#### func (Unknown) ExpressionType

```go
func (u Unknown) ExpressionType() Type
```
ExpressionType implements Expression with the inferred type of the unknown. If
the unknown could not be inferred, it will be of type "any" so allow expressions
using it to resolve anyway.
