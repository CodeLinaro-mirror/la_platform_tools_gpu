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

#### func  Add

```go
func Add(p Owner, c Owned)
```
Add connects an Owned to its Owner.

#### func  Visit

```go
func Visit(node Node, visitor func(Node))
```
Visit invokes visitor for all the children of the supplied node.

#### type API

```go
type API struct {
	Named
	Enums        []*Enum        // the set of enums
	Classes      []*Class       // the set of classes
	Pseudonyms   []*Pseudonym   // the set of pseudo types
	Externs      []*Function    // the external function references
	Functions    []*Function    // the global functions
	Methods      []*Function    // the method functions
	Globals      []*Global      // the global variables
	StaticArrays []*StaticArray // the fixed size array types used
	Maps         []*Map         // the map types used
	Pointers     []*Pointer     // the pointer types used
	Slices       []*Slice       // the pointer types used
	References   []*Reference   // the reference types used
	Signatures   []*Signature   // the function signature types used
	Imported     *Symbols       // the symbols imported into this api
}
```

API is the root of the ASG, and holds a fully resolved api.

#### func (*API) Member

```go
func (m *API) Member(name string) Owned
```

#### func (*API) VisitMembers

```go
func (m *API) VisitMembers(visitor func(Owned))
```

#### type Alias

```go
type Alias struct {
	AST *ast.Alias
	Named
	To Type
}
```

Alias is used as a temporary type holder during type resolution. It is not
present in the final semantic tree returned, but may be present in the AST ->
semantic map.

#### func (Alias) Member

```go
func (Alias) Member(string) Owned
```

#### func (*Alias) Owner

```go
func (o *Alias) Owner() Owner
```

#### func (Alias) VisitMembers

```go
func (Alias) VisitMembers(func(Owned))
```

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
	Named                     // the name of the annotation
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

#### type ArrayAssign

```go
type ArrayAssign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	To       *ArrayIndex // the array index to assign to
	Operator string      // the assignment operator being applied
	Value    Expression  // the value to set in the array
}
```

ArrayAssign represents assigning to a static-array index expression.

#### type ArrayIndex

```go
type ArrayIndex struct {
	AST   *ast.Index   // the underlying syntax node this was built from
	Type  *StaticArray // the array type
	Array Expression   // the expression that returns the array to be indexed
	Index Expression   // the index to use on the array
}
```

ArrayIndex represents using the indexing operator on a static-array type.

#### func (*ArrayIndex) ExpressionType

```go
func (i *ArrayIndex) ExpressionType() Type
```
ExpressionType implements Expression. It returns the element type of the array.

#### type ArrayInitializer

```go
type ArrayInitializer struct {
	AST    *ast.Call    // the underlying syntax node this was built from
	Array  Type         // the array type to initialize (may be aliased)
	Values []Expression // the list of element values
}
```

ArrayInitializer represents an expression that creates a new StaticArray
instance using a value list, of the form T(v0, v1, v2)

#### func (*ArrayInitializer) ExpressionType

```go
func (c *ArrayInitializer) ExpressionType() Type
```
ExpressionType implements Expression returning the class type being initialized.

#### type Assert

```go
type Assert struct {
	AST       *ast.Call  // the underlying syntax node this was built from
	Condition Expression // the condition is being asserted must be true
}
```

Assert represents a runtime assertion. Assertions are also used to infer
required behavior from the expressions.

#### func (*Assert) ExpressionType

```go
func (a *Assert) ExpressionType() Type
```
ExpressionType implements Expression

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

#### type Builtin

```go
type Builtin struct {
	Named // the primitive type name
}
```

Builtin represents one of the primitive types.

#### func (Builtin) Member

```go
func (Builtin) Member(string) Owned
```

#### func (*Builtin) Owner

```go
func (o *Builtin) Owner() Owner
```

#### func (Builtin) VisitMembers

```go
func (Builtin) VisitMembers(func(Owned))
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
	AST    *ast.Call  // the underlying syntax node this was built from
	Object Expression // the expression to cast the result of
	Type   Type       // the type to cast to
}
```

Cast represents a type reinterpret expresssion.

#### func (*Cast) ExpressionType

```go
func (c *Cast) ExpressionType() Type
```
ExpressionType implements Expression

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
	Named                   // implement Child
	Docs        []string    // the documentation for the class
	Fields      []*Field    // the set of fields the class declares
	Methods     []*Function // the set of functions associated with the class
}
```

Class represents an api class construct.

#### func (*Class) Member

```go
func (m *Class) Member(name string) Owned
```

#### func (*Class) Owner

```go
func (o *Class) Owner() Owner
```

#### func (*Class) VisitMembers

```go
func (m *Class) VisitMembers(visitor func(Owned))
```

#### type ClassInitializer

```go
type ClassInitializer struct {
	AST    *ast.Call           // the underlying syntax node this was built from
	Class  *Class              // the class to initialize
	Fields []*FieldInitializer // the set of field assignments
}
```

ClassInitializer represents an expression that can assign values to multiple
fields of a class.

#### func (*ClassInitializer) ExpressionType

```go
func (c *ClassInitializer) ExpressionType() Type
```
ExpressionType implements Expression returning the class type being initialized.

#### type Clone

```go
type Clone struct {
	AST   *ast.Call // the underlying syntax node this was built from
	Slice Expression
	Type  *Slice
}
```

Clone represents a call to make.

#### func (*Clone) ExpressionType

```go
func (m *Clone) ExpressionType() Type
```
ExpressionType implements Expression

#### type Copy

```go
type Copy struct {
	AST *ast.Call // the underlying syntax node this was built from
	Src Expression
	Dst Expression
}
```

Copy represents a call to make.

#### func (*Copy) ExpressionType

```go
func (*Copy) ExpressionType() Type
```
ExpressionType implements Expression

#### type Create

```go
type Create struct {
	AST         *ast.Call // the underlying syntax node this was built from
	Type        *Reference
	Initializer *ClassInitializer
}
```

Create represents a call to new on a class type.

#### func (*Create) ExpressionType

```go
func (n *Create) ExpressionType() Type
```
ExpressionType implements Expression

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
	Named                    // the type name of the enum
	Docs        []string     // the documentation for the enum
	IsBitfield  bool         // whether this enum is actually a bitfield
	Extends     []*Enum      // the enums this enum extends
	Entries     []*EnumEntry // the entries of this enum
}
```

Enum represents the api enum construct.

#### func (*Enum) Member

```go
func (m *Enum) Member(name string) Owned
```

#### func (*Enum) Owner

```go
func (o *Enum) Owner() Owner
```

#### func (*Enum) VisitMembers

```go
func (m *Enum) VisitMembers(visitor func(Owned))
```

#### type EnumEntry

```go
type EnumEntry struct {
	AST   *ast.EnumEntry // the underlying syntax node this was built from
	Named                // the name of this entry
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

#### func (*EnumEntry) Owner

```go
func (o *EnumEntry) Owner() Owner
```

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

#### type Fence

```go
type Fence struct {
	Statement Node
}
```

Fence is a marker to indicate the point between all statements to be executed
before (pre-fence) the call to the API function and all statements to be
executed after (post-fence) the call to the API function.

The Statement member is the first statement that is classified as post-fence,
but may be nil if the fence is being added at the end of a function that has no
post operations.

Note that some statements are classified as both pre-fence and post-fence, and
require logic to be executed either side of the API function call.

#### type Field

```go
type Field struct {
	AST         *ast.Field // the underlying syntax node this was built from
	Annotations            // the annotations applied to this field
	Type        Type       // the type the field stores
	Named                  // the name of the field
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

#### func (*Field) Owner

```go
func (o *Field) Owner() Owner
```

#### type FieldInitializer

```go
type FieldInitializer struct {
	AST   ast.Node   // the underlying syntax node this was built from
	Field *Field     // the field to assign to
	Value Expression // the value to assign
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
	Named                        // the name of the function
	Docs           []string      // the documentation for the function
	Return         *Parameter    // the return parameter
	This           *Parameter    // the this parameter, missing for non method functions
	FullParameters []*Parameter  // all the parameters, including This at the start if valid, and Return at the end if not void
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

#### func (*Function) Owner

```go
func (o *Function) Owner() Owner
```

#### type Global

```go
type Global struct {
	AST         *ast.Field // the underlying syntax node this was built from
	Annotations            // the annotations applied to this global
	Type        Type       // the type the global stores
	Named                  // the name of the global
	Default     Expression // the initial value of the global
}
```

Global represents a global variable.

#### func (*Global) ExpressionType

```go
func (g *Global) ExpressionType() Type
```
Implements Expression to return the type stored in the global.

#### func (*Global) Owner

```go
func (o *Global) Owner() Owner
```

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

#### type Import

```go
type Import struct {
	Named      // the full type name
	API   *API // the API being imported
}
```

Import wraps an API with it's imported name.

#### func (Import) Member

```go
func (i Import) Member(name string) Owned
```
Implement the Owner interface delegating member lookup to the imported API

#### func (*Import) Owner

```go
func (o *Import) Owner() Owner
```

#### func (Import) VisitMembers

```go
func (Import) VisitMembers(func(Owned))
```

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
	AST    *ast.Call  // the underlying syntax node this was built from
	Object Expression // the object go get the length of
	Type   Type       // the resolved type of the length operation
}
```

Length represents a length of object expression. Object must be of either
pointer, slice, map or string type. The length expression is allowed to be of
any numeric type

#### func (*Length) ExpressionType

```go
func (l *Length) ExpressionType() Type
```
ExpressionType implements Expression

#### type Local

```go
type Local struct {
	Declaration *DeclareLocal // the statement that created the local
	Type        Type          // the type of the storage
	Named                     // the identifier that will resolve to this local
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

#### type Make

```go
type Make struct {
	AST  *ast.Call // the underlying syntax node this was built from
	Type *Slice
	Size Expression
}
```

Make represents a call to make.

#### func (*Make) ExpressionType

```go
func (m *Make) ExpressionType() Type
```
ExpressionType implements Expression

#### type Map

```go
type Map struct {
	Named          // the full type name
	KeyType   Type // the type used as an indexing key
	ValueType Type // the type stored in the map
}
```

Map represents an api map type declaration, of the form map!(KeyType, ValueType)

#### func (*Map) Member

```go
func (m *Map) Member(name string) Owned
```

#### func (*Map) Owner

```go
func (o *Map) Owner() Owner
```

#### func (*Map) VisitMembers

```go
func (m *Map) VisitMembers(visitor func(Owned))
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
	AST   *ast.Index // the underlying syntax node this was built from
	Type  *Map       // the value type of the map being indexed
	Map   Expression // the expression that returns the map to be indexed
	Index Expression // the index to use on the map
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

#### type Named

```go
type Named string
```

Named is mixed in to implement the Name method of NamedNode.

#### func (Named) Name

```go
func (n Named) Name() string
```

#### type NamedNode

```go
type NamedNode interface {
	Node
	Name() string // Returns the partial name of the object.
}
```

NamedNode represents any semantic-tree node that carries a name.

#### type New

```go
type New struct {
	AST  *ast.Call // the underlying syntax node this was built from
	Type *Reference
}
```

New represents a call to new.

#### func (*New) ExpressionType

```go
func (n *New) ExpressionType() Type
```
ExpressionType implements Expression

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

#### type Owned

```go
type Owned interface {
	NamedNode
	Owner() Owner // Returns the owner of this node.
	// contains filtered or unexported methods
}
```

Owned is the interface to an object with a unique name and an owner.

#### type Owner

```go
type Owner interface {
	NamedNode
	Member(string) Owned      // looks up a member by name from an owner
	VisitMembers(func(Owned)) // invokes the supplied function once for each member
	// contains filtered or unexported methods
}
```

Owner is the interface for an object that has named members.

#### type Parameter

```go
type Parameter struct {
	AST         *ast.Parameter // the underlying syntax node this was built from
	Annotations                // the annotations applied to the parameter
	Function    *Function      // the function this parameter belongs to
	Named                      // the name of the parameter
	Docs        []string       // the documentation for the parameter
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
	Named        // the full type name
	To    Type   // the type this is a pointer to
	Const bool   // wether the pointer was declared with the const attribute
	Slice *Slice // The complementary slice type for this pointer.
}
```

Pointer represents an api pointer type declaration, of the form To*

#### func (*Pointer) Member

```go
func (t *Pointer) Member(name string) Owned
```

#### func (*Pointer) Owner

```go
func (o *Pointer) Owner() Owner
```

#### func (*Pointer) VisitMembers

```go
func (t *Pointer) VisitMembers(visitor func(Owned))
```

#### type PointerRange

```go
type PointerRange struct {
	AST     *ast.Index // the underlying syntax node this was built from
	Type    *Slice     // the slice type returned.
	Pointer Expression // the expression that returns the pointer to be indexed
	Range   *BinaryOp  // the range to use on the slice
}
```

PointerRange represents using the indexing operator on a pointer type with a
range expression.

#### func (*PointerRange) ExpressionType

```go
func (i *PointerRange) ExpressionType() Type
```
ExpressionType implements Expression. It returns the same slice type being
sliced.

#### type Pseudonym

```go
type Pseudonym struct {
	AST         *ast.Pseudonym // the underlying syntax node this was built from
	Annotations                // the annotations applied to this pseudonym
	Named                      // the type name
	Docs        []string       // the documentation for the pseudonym
	To          Type           // the underlying type
	Methods     []*Function    // the methods added directly to the pseudonym
}
```

Pseudonym represents the type construct. It acts as a type in it's own right
that can carry methods, but is defined in terms of another type.

#### func (*Pseudonym) Member

```go
func (t *Pseudonym) Member(name string) Owned
```
Implements Type returning the direct member if it has it, otherwise delegating
the lookup to the underlying type.

#### func (*Pseudonym) Owner

```go
func (o *Pseudonym) Owner() Owner
```

#### func (*Pseudonym) VisitMembers

```go
func (t *Pseudonym) VisitMembers(visitor func(Owned))
```

#### type Read

```go
type Read struct {
	AST   *ast.Call // the underlying syntax node this was built from
	Slice Expression
}
```

Read represents a call to make.

#### func (*Read) ExpressionType

```go
func (*Read) ExpressionType() Type
```
ExpressionType implements Expression

#### type Reference

```go
type Reference struct {
	Named      // the full type name
	To    Type // the type this is a reference to
}
```

Reference represents an api reference type declaration, of the form ref!To

#### func (*Reference) Member

```go
func (t *Reference) Member(name string) Owned
```

#### func (*Reference) Owner

```go
func (o *Reference) Owner() Owner
```

#### func (*Reference) VisitMembers

```go
func (t *Reference) VisitMembers(visitor func(Owned))
```

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
	Named            // the full type name
	Return    Type   // the return type of the callable
	Arguments []Type // the required callable arguments
}
```

Signature represents a callable type signature

#### func (Signature) Member

```go
func (Signature) Member(string) Owned
```

#### func (*Signature) Owner

```go
func (o *Signature) Owner() Owner
```

#### func (Signature) VisitMembers

```go
func (Signature) VisitMembers(func(Owned))
```

#### type Slice

```go
type Slice struct {
	Named            // the full type name
	To      Type     // The type this is a slice of
	Pointer *Pointer // The complementary pointer type for this slice.
}
```

Slice represents an api slice type declaration, of the form To[]

#### func (Slice) Member

```go
func (Slice) Member(string) Owned
```

#### func (*Slice) Owner

```go
func (o *Slice) Owner() Owner
```

#### func (Slice) VisitMembers

```go
func (Slice) VisitMembers(func(Owned))
```

#### type SliceAssign

```go
type SliceAssign struct {
	AST      *ast.Assign // the underlying syntax node this was built from
	To       *SliceIndex // the slice index to assign to
	Operator string      // the assignment operator being applied
	Value    Expression  // the value to set in the slice
}
```

SliceAssign represents assigning to a slice index expression.

#### type SliceIndex

```go
type SliceIndex struct {
	AST   *ast.Index // the underlying syntax node this was built from
	Type  *Slice     // the slice type
	Slice Expression // the expression that returns the slice to be indexed
	Index Expression // the index to use on the slice
}
```

SliceIndex represents using the indexing operator on a slice type.

#### func (*SliceIndex) ExpressionType

```go
func (i *SliceIndex) ExpressionType() Type
```
ExpressionType implements Expression. It returns the value type of the slice.

#### type SliceRange

```go
type SliceRange struct {
	AST   *ast.Index // the underlying syntax node this was built from
	Type  *Slice     // the slice type
	Slice Expression // the expression that returns the slice to be indexed
	Range *BinaryOp  // the range to use on the slice
}
```

SliceRange represents using the indexing operator on a slice type with a range
expression.

#### func (*SliceRange) ExpressionType

```go
func (i *SliceRange) ExpressionType() Type
```
ExpressionType implements Expression. It returns the same slice type being
sliced.

#### type StaticArray

```go
type StaticArray struct {
	Named            // the full type name
	ValueType Type   // the storage type of the elements
	Size      uint32 // the dimension of the array
}
```

StaticArray represents a multi-dimensional fixed size array type, of the form
T[8]

#### func (StaticArray) Member

```go
func (StaticArray) Member(string) Owned
```

#### func (*StaticArray) Owner

```go
func (o *StaticArray) Owner() Owner
```

#### func (StaticArray) VisitMembers

```go
func (StaticArray) VisitMembers(func(Owned))
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

#### type Symbols

```go
type Symbols struct {
}
```

Symbols is an object with named members and no other functionality.

#### func (*Symbols) Add

```go
func (s *Symbols) Add(name string, entry Node)
```
Add inserts a node into the symbol space with the specified name.

#### func (*Symbols) AddNamed

```go
func (s *Symbols) AddNamed(entry NamedNode)
```
Add inserts a named node into the symbol space.

#### func (*Symbols) Find

```go
func (s *Symbols) Find(name string) (Node, error)
```

#### func (*Symbols) FindAll

```go
func (s *Symbols) FindAll(name string) []Node
```

#### func (*Symbols) Visit

```go
func (s *Symbols) Visit(visitor func(string, Node))
```

#### type Type

```go
type Type interface {
	Owner
	// contains filtered or unexported methods
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

#### type Write

```go
type Write struct {
	AST   *ast.Call // the underlying syntax node this was built from
	Slice Expression
}
```

Write represents a call to make.

#### func (*Write) ExpressionType

```go
func (*Write) ExpressionType() Type
```
ExpressionType implements Expression
