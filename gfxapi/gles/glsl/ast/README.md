# ast
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/ast"

Package ast defines the abstract syntax tree of OpenGL ES Shading Language
programs.

It contains classes for individual types of expressions, statements and
declarations. The code manipulating (parsing, serializing, evaluating, ...) the
AST are located in other packages.

The type Ast represents the entire parsed program tree. It consists of a
sequence of declarations (types ***Decl), which can be variable, function,
interface declarations, etc.

Function bodies consist of a list of statements (types ***Stmt), for if, while,
for statements, etc. There are two special kinds of statements: DeclarationStmt
is a statement holding a declaration (described above) and ExpressionStmt is a
statement holding an expression.

The expressions of the program are represented by types ***Expr. Typical
examples are binary expressions, function call expressions, variable reference
expressions, etc.

The AST contains enough information to enable serialization to text while
preserving all whitespace and comments. The ***Cst members in all AST nodes are
used for storing this whitespace information. This way, whitespace gets
automatically preserved during AST modification. It is not necessary (but is
possible) to populate the ***Cst members when creating new nodes, as the
serialization is expected to succeed even without them.

The package also defines the type system of the GLES Shading Language programs
(***Type), their values (***Value) and several helper functions for manipulating
them.

## Usage

```go
var (
	BoAdd    = appendBinaryOperator("+")
	BoBand   = appendBinaryOperator("&")
	BoBor    = appendBinaryOperator("|")
	BoBxor   = appendBinaryOperator("^")
	BoComma  = appendBinaryOperator(",")
	BoDiv    = appendBinaryOperator("/")
	BoEq     = appendBinaryOperator("==")
	BoLand   = appendBinaryOperator("&&")
	BoLess   = appendBinaryOperator("<")
	BoLessEq = appendBinaryOperator("<=")
	BoLor    = appendBinaryOperator("||")
	BoLxor   = appendBinaryOperator("^^")
	BoMod    = appendBinaryOperator("%")
	BoMore   = appendBinaryOperator(">")
	BoMoreEq = appendBinaryOperator(">=")
	BoMul    = appendBinaryOperator("*")
	BoNotEq  = appendBinaryOperator("!=")
	BoShl    = appendBinaryOperator("<<")
	BoShr    = appendBinaryOperator(">>")
	BoSub    = appendBinaryOperator("-")

	BoAssign    = appendAssignOp("=")
	BoAddAssign = appendAssignOp("+=")
	BoAndAssign = appendAssignOp("&=")
	BoDivAssign = appendAssignOp("/=")
	BoModAssign = appendAssignOp("%=")
	BoMulAssign = appendAssignOp("*=")
	BoOrAssign  = appendAssignOp("|=")
	BoShlAssign = appendAssignOp("<<=")
	BoShrAssign = appendAssignOp(">>=")
	BoSubAssign = appendAssignOp("-=")
	BoXorAssign = appendAssignOp("^=")

	UoBnot    = appendPrefixOp("~")
	UoLnot    = appendPrefixOp("!")
	UoMinus   = appendPrefixOp("-")
	UoPlus    = appendPrefixOp("+")
	UoPredec  = appendPrefixOp("--")
	UoPreinc  = appendPrefixOp("++")
	UoPostdec = appendUnaryOperator("--")
	UoPostinc = appendUnaryOperator("++")
)
```
All the operators of the language, one variable per operator.

```go
var (
	TVoid                 = appendType("void")
	TFloat                = appendType("float")
	TInt                  = appendType("int")
	TUint                 = appendType("uint")
	TBool                 = appendType("bool")
	TMat2                 = appendType("mat2")
	TMat3                 = appendType("mat3")
	TMat4                 = appendType("mat4")
	TMat2x2               = appendType("mat2x2")
	TMat2x3               = appendType("mat2x3")
	TMat2x4               = appendType("mat2x4")
	TMat3x2               = appendType("mat3x2")
	TMat3x3               = appendType("mat3x3")
	TMat3x4               = appendType("mat3x4")
	TMat4x2               = appendType("mat4x2")
	TMat4x3               = appendType("mat4x3")
	TMat4x4               = appendType("mat4x4")
	TVec2                 = appendType("vec2")
	TVec3                 = appendType("vec3")
	TVec4                 = appendType("vec4")
	TIvec2                = appendType("ivec2")
	TIvec3                = appendType("ivec3")
	TIvec4                = appendType("ivec4")
	TBvec2                = appendType("bvec2")
	TBvec3                = appendType("bvec3")
	TBvec4                = appendType("bvec4")
	TUvec2                = appendType("uvec2")
	TUvec3                = appendType("uvec3")
	TUvec4                = appendType("uvec4")
	TSampler2D            = appendType("sampler2D")
	TSampler3D            = appendType("sampler3D")
	TSamplerCube          = appendType("samplerCube")
	TSampler2DShadow      = appendType("sampler2DShadow")
	TSamplerCubeShadow    = appendType("samplerCubeShadow")
	TSampler2DArray       = appendType("sampler2DArray")
	TSampler2DArrayShadow = appendType("sampler2DArrayShadow")
	TIsampler2D           = appendType("isampler2D")
	TIsampler3D           = appendType("isampler3D")
	TIsamplerCube         = appendType("isamplerCube")
	TIsampler2DArray      = appendType("isampler2DArray")
	TUsampler2D           = appendType("usampler2D")
	TUsampler3D           = appendType("usampler3D")
	TUsamplerCube         = appendType("usamplerCube")
	TUsampler2DArray      = appendType("usampler2DArray")
)
```
All types of the language.

```go
var BareTypes []BareType
```
All bare types of the language, in an array. Do not modify.

```go
var Operators []fmt.Stringer
```
All operators of the language, in an array, sorted according to string length
(longest first). Do not modify.

#### func  ConcatErrors

```go
func ConcatErrors(e1, e2 []error) []error
```
Concat concatenates two lists of errors.

#### func  GetVectorComponentInfo

```go
func GetVectorComponentInfo(r rune) (position uint8, kind VectorComponentKind)
```
GetVectorComponentInfo returns information about a rune, when used as a vector
swizzle character. It returns the VectorComponentKind this rune belongs to (or
ComponentKindNone if it is an invalid rune) and the vector position this rune
refers to. The position indexes are 0-based.

#### func  HasVectorExpansions

```go
func HasVectorExpansions(t BareType) bool
```
HasVectorExpansions returns true if the type t has corresponding vector
counterparts. This is currently true only for TBool, TFloat, TInt and TUint.

#### func  IsAssignmentOp

```go
func IsAssignmentOp(op BinaryOp) bool
```
Is op an assignment operator?

#### func  IsMatrixType

```go
func IsMatrixType(t BareType) bool
```
IsMatrixType returns true if t is a matrix type.

#### func  IsVectorType

```go
func IsVectorType(t BareType) bool
```
IsVectorType returns true if the given type is a vector type.

#### func  TypeDimensions

```go
func TypeDimensions(t BareType) (col uint8, row uint8)
```
TypeDimensions returns the number of columns and rows a given type contains.
Vectors are treated as column vectors (cols == 1). Scalars have 1 column and 1
row.

#### func  ValueEquals

```go
func ValueEquals(left, right Value) bool
```
ValueEquals compares two values for equality. The syntax of the GLES Shading
Language does not permit comparing function values and this function will always
return false for them.

#### func  VisitChildren

```go
func VisitChildren(n interface{}, v ChildVisitor)
```
VisitChildren is a helper function which calls the provided callback for each
child node of the argument.

#### type ArrayType

```go
type ArrayType struct {
	Base Type
	Size Expression

	LBracketCst *parse.Leaf
	RBracketCst *parse.Leaf

	ComputedSize UintValue // The size determined by semantic analysis.
}
```

ArrayType represents array types. Size is represented as an expression.
Unspecified size is represented by a nil expression. After semantic analysis,
the ComputedSize field will store the size of the array as an integer. This size
will be determined by evaluating the size expression, or from context. Precision
getters and setters use the precision of the Base type.

#### func (*ArrayType) Copy

```go
func (at *ArrayType) Copy() *ArrayType
```
Copy returns a (shallow) copy of the underlying object.

#### func (*ArrayType) Equal

```go
func (at *ArrayType) Equal(other Type) bool
```
Equal determines whether two objects represent the same type. Note that a
semantic analysis pass over the AST (to populate ComputedSize) is needed to
produce meaningful results.

#### type ArrayValue

```go
type ArrayValue []Value
```

ArrayValue represents values of array type. Arrays of size zero are illegal.

#### func (ArrayValue) Type

```go
func (v ArrayValue) Type() Type
```
Type returns the type of the value. The returned type will always be an
*ArrayType and it will have correctly populated Base, Size and ComputedSize
fields.

#### type Ast

```go
type Ast struct {
	Decls []interface{}
}
```

Ast represents an abstract syntax tree of a parsed program. It is just a holder
for a list of declarations.

#### type BareType

```go
type BareType struct {
}
```


#### func  GetFundamentalType

```go
func GetFundamentalType(t BareType) BareType
```
GetFundamentalType returns the fundamental type of a given builtin type. E.g.,
the fundamental type of TVec2 is TFloat, since vec2 is a vector of two flots.
For types which cannot be used in vectors or matrices (such as sampler types or
void), the function returns TVoid as their fundamental type.

#### func  GetMatrixType

```go
func GetMatrixType(col, row uint8) BareType
```
Returns a matrix type with the given number of columns and rows. In case the
number of rows or columns is 1, returns a vector type. In case both are 1,
returns TFloat.

#### func  GetPrecisionType

```go
func GetPrecisionType(t BareType) (prect BareType, hasPrecision bool)
```
GetPrecisionType returns the type from which the provided type inherits is
precision. E.g., vec2 inherits its type from float as they have the same basic
type. If the type cannot be qualified with a precision the second result is
false.

#### func  GetVectorType

```go
func GetVectorType(fund BareType, size uint8) BareType
```
GetVectorType returns the builtin vector type with the given fundamental type
(TFloat, TInt, TUint and TBool) and the given size. E.g., GetVectorType(TBool,
4) == TBvec4. In case size is 1 it returns the corresponding scalar type. It is
illegal to call this function with any other type except the four types
mentioned.

#### func (BareType) Canonicalize

```go
func (t BareType) Canonicalize() BareType
```
Canonicalize returns a canonical representation of the type. Specifically, it
returns TMat2x2 instead of TMat2, etc.

#### func (BareType) String

```go
func (t BareType) String() string
```

#### type BinaryExpr

```go
type BinaryExpr struct {
	Left  Expression
	Op    BinaryOp
	Right Expression

	OpCst *parse.Leaf

	ExprType Type // Found by semantic analysis.
}
```

E_Left op E_Right

#### func (*BinaryExpr) Constant

```go
func (e *BinaryExpr) Constant() bool
```

#### func (*BinaryExpr) LValue

```go
func (e *BinaryExpr) LValue() bool
```

#### func (*BinaryExpr) Type

```go
func (e *BinaryExpr) Type() Type
```

#### type BinaryOp

```go
type BinaryOp struct {
}
```

A type for binary operators of the language.

#### func (BinaryOp) String

```go
func (op BinaryOp) String() string
```

#### type BoolValue

```go
type BoolValue bool
```

BoolValue represents a boolean value.

#### func (BoolValue) Type

```go
func (v BoolValue) Type() Type
```

#### type BreakStmt

```go
type BreakStmt struct {
	BreakCst     *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

break;

#### type BuiltinFunction

```go
type BuiltinFunction struct {
	RetType Type
	SymName string
	Params  []*FuncParameterSym
	Impl    func(v []Value) Value
}
```

BuiltinFunction is the definition of a GLSL builtin function. It implements the
Function interface. The definition of the function is provided as an opaque go
function.

#### func (*BuiltinFunction) Constant

```go
func (f *BuiltinFunction) Constant() bool
```

#### func (*BuiltinFunction) LValue

```go
func (f *BuiltinFunction) LValue() bool
```

#### func (*BuiltinFunction) Name

```go
func (f *BuiltinFunction) Name() string
```

#### func (*BuiltinFunction) Parameters

```go
func (f *BuiltinFunction) Parameters() []*FuncParameterSym
```

#### func (*BuiltinFunction) ReturnType

```go
func (f *BuiltinFunction) ReturnType() Type
```

#### func (*BuiltinFunction) Type

```go
func (f *BuiltinFunction) Type() Type
```

#### type BuiltinType

```go
type BuiltinType struct {
	Precision Precision
	Type      BareType

	PrecisionCst *parse.Leaf
	TypeCst      *parse.Leaf
}
```

BuiltinType represents a builtin type of the language: int, float, vec2, etc. It
also stores the precision of the type.

#### func (*BuiltinType) Equal

```go
func (bt *BuiltinType) Equal(other Type) bool
```
Two builtin types are equal if their Type members are equal. Precision is
ignored.

#### type CallExpr

```go
type CallExpr struct {
	Callee Expression
	Args   []Expression

	LParenCst   *parse.Leaf
	CommaCst    []*parse.Leaf
	RParenCst   *parse.Leaf
	VoidPresent bool
	VoidCst     *parse.Leaf

	// Found by semantic analysis.
	ExprType     Type
	ExprConstant bool
}
```

Callee(Arg_1, ... Arg_n).

If the function is called with no arguments the VoidPresent field is set to
indicate the usage of the f(void) call syntax.

#### func (*CallExpr) Constant

```go
func (e *CallExpr) Constant() bool
```

#### func (*CallExpr) LValue

```go
func (e *CallExpr) LValue() bool
```

#### func (*CallExpr) Type

```go
func (e *CallExpr) Type() Type
```

#### type CaseStmt

```go
type CaseStmt struct {
	Expr Expression

	CaseCst  *parse.Leaf
	ColonCst *parse.Leaf
}
```

case expr:

#### type ChildVisitor

```go
type ChildVisitor func(interface{})
```

ChildVisitor is a callback function used by VisitChildren to visit nodes.

#### type CompoundStmt

```go
type CompoundStmt struct {
	Stmts []interface{}

	LBraceCst *parse.Leaf
	RBraceCst *parse.Leaf
}
```

{ stmts... }

#### type ConditionalExpr

```go
type ConditionalExpr struct {
	Cond      Expression
	TrueExpr  Expression
	FalseExpr Expression

	QuestionCst *parse.Leaf
	ColonCst    *parse.Leaf
}
```

E_c ? E_t : E_f

#### func (*ConditionalExpr) Constant

```go
func (e *ConditionalExpr) Constant() bool
```

#### func (*ConditionalExpr) LValue

```go
func (e *ConditionalExpr) LValue() bool
```

#### func (*ConditionalExpr) Type

```go
func (e *ConditionalExpr) Type() Type
```

#### type ConstantExpr

```go
type ConstantExpr struct {
	Value Value

	ValCst *parse.Leaf
}
```

ConstantExpr holds integer, floating point and boolean constants.

#### func (*ConstantExpr) Constant

```go
func (e *ConstantExpr) Constant() bool
```

#### func (*ConstantExpr) LValue

```go
func (e *ConstantExpr) LValue() bool
```

#### func (*ConstantExpr) Type

```go
func (e *ConstantExpr) Type() Type
```

#### type ContinueStmt

```go
type ContinueStmt struct {
	ContinueCst  *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

continue;

#### type DeclarationStmt

```go
type DeclarationStmt struct {
	Decl interface{}
}
```

DeclarationStmt is a statement holding a declaration.

#### type DefaultStmt

```go
type DefaultStmt struct {
	DefaultCst *parse.Leaf
	ColonCst   *parse.Leaf
}
```

default:

#### type DiscardStmt

```go
type DiscardStmt struct {
	DiscardCst   *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

discard;

#### type DoStmt

```go
type DoStmt struct {
	Stmt interface{}
	Expr Expression

	DoCst        *parse.Leaf
	LParenCst    *parse.Leaf
	RParenCst    *parse.Leaf
	WhileCst     *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

do stmt while(expr);

#### type DotExpr

```go
type DotExpr struct {
	Expr  Expression
	Field string

	DotCst   *parse.Leaf
	FieldCst *parse.Leaf

	// Found by semantic analysis.
	ExprType     Type
	ExprConstant bool
	ExprLValue   bool
}
```

E . field

This expression represents struct member selection (struct.member), vector
swizzle (vec4.xyz) and array length (array.length) expressions.

#### func (*DotExpr) Constant

```go
func (e *DotExpr) Constant() bool
```

#### func (*DotExpr) LValue

```go
func (e *DotExpr) LValue() bool
```

#### func (*DotExpr) Type

```go
func (e *DotExpr) Type() Type
```

#### type EmptyStmt

```go
type EmptyStmt struct {
	SemicolonCst *parse.Leaf
}
```

;

#### type ErrorCollector

```go
type ErrorCollector []error
```

ErrorCollector is an object which collects parsing errors.

#### func (*ErrorCollector) Error

```go
func (e *ErrorCollector) Error(err ...error)
```
Error adds the arguments to the collection of errors.

#### func (*ErrorCollector) Errorf

```go
func (e *ErrorCollector) Errorf(message string, args ...interface{})
```
Errorf adds an error defined by a format string to the collector.

#### func (*ErrorCollector) GetErrors

```go
func (e *ErrorCollector) GetErrors() []error
```
GetErrors returns the collected errors.

#### type Expression

```go
type Expression interface {
	LValue() bool
	Type() Type
	Constant() bool
}
```

The interface for all expressions in the program. It defines methods retrieving
the expression type, its const-ness, and lvalued-ness. These are to be called
only after a semantic analysis pass over the program.

#### type ExpressionCond

```go
type ExpressionCond struct {
	Expr Expression
}
```

ExpressionCond is a condition holding an expression. Conditions are used in
ForStmt and WhileStmt.

#### type ExpressionStmt

```go
type ExpressionStmt struct {
	Expr Expression

	SemicolonCst *parse.Leaf
}
```

ExpressionStmt is a statement holding an expression.

#### type FloatValue

```go
type FloatValue float64
```

FloatValue represents a floating point value, as a float64.

#### func (FloatValue) String

```go
func (v FloatValue) String() string
```

#### func (FloatValue) Type

```go
func (v FloatValue) Type() Type
```

#### type ForStmt

```go
type ForStmt struct {
	Init interface{} // statement
	Cond interface{} // condition
	Loop Expression
	Body interface{} // statement

	ForCst        *parse.Leaf
	LParenCst     *parse.Leaf
	Semicolon2Cst *parse.Leaf // semicolon1 is a part of Init
	RParenCst     *parse.Leaf
}
```

for(init; cond; loop) body

The second member Cond is expected to hold an expression (ExpressionCond) or a
declaration of a single variable (VarDeclCond).

#### type FuncParameterSym

```go
type FuncParameterSym struct {
	Const     bool
	Direction ParamDirection
	SymType   Type
	SymName   string

	ConstCst      *parse.Leaf
	DirectionCst  *parse.Leaf
	NameCst       *parse.Leaf
	ArrayAfterVar bool // Whether the variable was declared as int foo[].
}
```

FuncParameter is a formal function parameter in a function declaration.

#### func (*FuncParameterSym) Constant

```go
func (v *FuncParameterSym) Constant() bool
```

#### func (*FuncParameterSym) LValue

```go
func (v *FuncParameterSym) LValue() bool
```

#### func (*FuncParameterSym) Name

```go
func (v *FuncParameterSym) Name() string
```

#### func (*FuncParameterSym) Type

```go
func (v *FuncParameterSym) Type() Type
```

#### type Function

```go
type Function interface {
	ValueSymbol
	ReturnType() Type
	Parameters() []*FuncParameterSym
}
```

Function is an interface for function declarations. In addition to ValueSymbol
methods, it provides methods to retrieve function return type and function
parameters.

#### type FunctionDecl

```go
type FunctionDecl struct {
	RetType Type
	SymName string
	Params  []*FuncParameterSym
	Stmts   *CompoundStmt

	NameCst      *parse.Leaf
	LParenCst    *parse.Leaf
	CommaCst     []*parse.Leaf
	RParenCst    *parse.Leaf
	SemicolonCst *parse.Leaf
	VoidPresent  bool
	VoidCst      *parse.Leaf
}
```

FunctionDecl is a function declaration in a GLSL program. If field Stsms is not
nil, then it represents a function definition. FunctionDecl implements
ValueSymbol and Function.

If the function is declared with no arguments the VoidPresent field is set to
indicate the usage of the f(void) call syntax.

#### func (*FunctionDecl) Constant

```go
func (f *FunctionDecl) Constant() bool
```

#### func (*FunctionDecl) LValue

```go
func (f *FunctionDecl) LValue() bool
```

#### func (*FunctionDecl) Name

```go
func (f *FunctionDecl) Name() string
```

#### func (*FunctionDecl) Parameters

```go
func (f *FunctionDecl) Parameters() []*FuncParameterSym
```

#### func (*FunctionDecl) ReturnType

```go
func (f *FunctionDecl) ReturnType() Type
```

#### func (*FunctionDecl) Type

```go
func (f *FunctionDecl) Type() Type
```

#### type FunctionType

```go
type FunctionType struct {
	ArgTypes []Type
	RetType  Type
}
```

FunctionType represents function types. The GLES Shading Language does not
specify the presence of function types, but we add it to be able to handle
expressions uniformly. A function type can appear in the AST in the following
manner:

- VarRefExpr to a function declaration,

- "array.length" expression,

- the "type" part of the "type(expression)" type constructor syntax.

In all valid programs a expression of a function type can be present only as a
child of a call expression.

#### func (*FunctionType) Equal

```go
func (t *FunctionType) Equal(other Type) bool
```
Two function types are equal if they have the same return and argument types.

#### type FunctionValue

```go
type FunctionValue struct {
	Func    func(v []Value) Value
	ValType *FunctionType
}
```


#### func (FunctionValue) Type

```go
func (v FunctionValue) Type() Type
```

#### type IfStmt

```go
type IfStmt struct {
	IfExpr   Expression
	ThenStmt interface{}
	ElseStmt interface{}

	IfCst     *parse.Leaf
	LParenCst *parse.Leaf
	RParenCst *parse.Leaf
	ElseCst   *parse.Leaf
}
```

if(ifExpr) thenStmt else elseStmt

if ElseStms is nil, the statement has no else clause.

#### type IndexExpr

```go
type IndexExpr struct {
	Base  Expression
	Index Expression

	LBracketCst *parse.Leaf
	RBracketCst *parse.Leaf

	ExprType Type // Found by semantic analysis.
}
```

Base[Index]

#### func (*IndexExpr) Constant

```go
func (e *IndexExpr) Constant() bool
```

#### func (*IndexExpr) LValue

```go
func (e *IndexExpr) LValue() bool
```

#### func (*IndexExpr) Type

```go
func (e *IndexExpr) Type() Type
```

#### type IntValue

```go
type IntValue int32
```

IntValue represents an integer value, as a int32.

#### func (IntValue) String

```go
func (v IntValue) String() string
```

#### func (IntValue) Type

```go
func (v IntValue) Type() Type
```

#### type InterpolationQualifier

```go
type InterpolationQualifier uint8
```


```go
const (
	IntNone InterpolationQualifier = iota
	IntSmooth
	IntFlat
)
```

#### func (InterpolationQualifier) String

```go
func (i InterpolationQualifier) String() string
```

#### type InvariantDecl

```go
type InvariantDecl struct {
	Vars []*VarRefExpr

	InvariantCst *parse.Leaf
	CommaCst     []*parse.Leaf
	SemicolonCst *parse.Leaf
}
```

invariant foo, bar, baz;

#### type Language

```go
type Language uint8
```


```go
const (
	LangVertexShader   Language = iota // Vertex Shader language
	LangFragmentShader                 // Fragment Shader language
	LangPreprocessor                   // Language of the #if preprocessor expressions
)
```
Constants identifying various sub-languages.

#### type LayoutDecl

```go
type LayoutDecl struct {
	Layout *LayoutQualifier

	UniformCst   *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

layout(...) uniform;

#### type LayoutQualifier

```go
type LayoutQualifier struct {
	Ids []LayoutQualifierID

	LayoutCst *parse.Leaf
	LParenCst *parse.Leaf
	CommaCst  []*parse.Leaf
	RParenCst *parse.Leaf
}
```

layout(...)

#### type LayoutQualifierID

```go
type LayoutQualifierID struct {
	Name  string
	Value Value

	NameCst  *parse.Leaf
	EqualCst *parse.Leaf
	ValueCst *parse.Leaf
}
```

A single name = value pair in the layout qualifier. If the qualifier is present
just as a bare name (without the value) then Value is nil.

#### type MatrixValue

```go
type MatrixValue struct {
	Members [][]Value
	ValType BuiltinType
}
```

MatrixValue represents matrix. Its elements are stored column-major.

#### func  NewMatrixValue

```go
func NewMatrixValue(t BareType, member func(col, row uint8) Value) MatrixValue
```
NewMatrixValue constructs a new matrix value. The number of colums and rows is
determined using the type argument. The member callback function should return a
value for each matrix member, given its column and row. The function is called
in column-major order. The type of the created value will always be in a
canonical (TMat2x2 instead of TMat2) form.

#### func  NewMatrixValueCR

```go
func NewMatrixValueCR(col, row uint8, member func(col, row uint8) Value) MatrixValue
```
NewMatrixValueCR constructs a new matrix value, given the number of its columns
and rows. It is possible (and occasionally useful) to create matrix values with
invalid sizes (number of columns of rows equal to 1 or greater than 4), but then
the ValType field will not be populated correctly. The function is called in
column-major order. The type of the created value will always be in a canonical
(TMat2x2 instead of TMat2) form.

#### func (MatrixValue) Type

```go
func (v MatrixValue) Type() Type
```

#### type MultiVarDecl

```go
type MultiVarDecl struct {
	Vars  []*VariableSym
	Quals *TypeQualifiers
	Type  Type

	FunnyComma    bool // for declarations like: `int ,a;`
	FunnyCommaCst *parse.Leaf
	CommaCst      []*parse.Leaf
	SemicolonCst  *parse.Leaf
}
```

MultiVarDecl is a declaration of a (possibly more than one) variable. If no
qualifiers are present, the Quals field can be null.

#### type ParamDirection

```go
type ParamDirection uint8
```

The direction of a parameter in a function declaration.

```go
const (
	DirNone ParamDirection = iota
	DirIn
	DirOut
	DirInout
)
```

#### func (ParamDirection) String

```go
func (d ParamDirection) String() string
```

#### type ParenExpr

```go
type ParenExpr struct {
	Expr Expression

	LParenCst *parse.Leaf
	RParenCst *parse.Leaf
}
```

( E )

#### func (*ParenExpr) Constant

```go
func (e *ParenExpr) Constant() bool
```

#### func (*ParenExpr) LValue

```go
func (e *ParenExpr) LValue() bool
```

#### func (*ParenExpr) Type

```go
func (e *ParenExpr) Type() Type
```

#### type Precision

```go
type Precision uint8
```

Precision of a type

```go
const (
	NoneP Precision = iota
	LowP
	MediumP
	HighP
)
```

#### func (Precision) String

```go
func (p Precision) String() string
```

#### type PrecisionDecl

```go
type PrecisionDecl struct {
	Type *BuiltinType

	PrecisionCst *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

precision lowp int;

#### type ReturnStmt

```go
type ReturnStmt struct {
	Expr Expression

	ReturnCst    *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

return expr;

#### type StorageQualifier

```go
type StorageQualifier uint8
```


```go
const (
	StorNone StorageQualifier = iota
	StorConst
	StorUniform
	StorIn
	StorOut
	StorCentroidIn
	StorCentroidOut
	StorAttribute
	StorVarying
)
```

#### func (StorageQualifier) String

```go
func (s StorageQualifier) String() string
```

#### type StructSym

```go
type StructSym struct {
	SymName string
	Vars    []*MultiVarDecl

	StructCst *parse.Leaf
	NameCst   *parse.Leaf
	LBraceCst *parse.Leaf
	RBraceCst *parse.Leaf
}
```

StructSym is the struct declaration. It has a name and a list of variable
declarations.

#### func (*StructSym) Name

```go
func (s *StructSym) Name() string
```

#### type StructType

```go
type StructType struct {
	Sym       *StructSym
	StructDef bool // if true, the structure was defined at this point in the AST.

	NameCst *parse.Leaf // Only used if StructDef is false
}
```

StructType represents structure types as a pointer to the corresponding
structure symbol.

#### func (*StructType) Equal

```go
func (s *StructType) Equal(other Type) bool
```

#### type StructValue

```go
type StructValue struct {
	Members map[string]Value
	ValType *StructType
}
```

StructValue represents values of structure types.

#### func  NewStructValue

```go
func NewStructValue(t *StructType, member func(name string) Value) StructValue
```
NewStructValue construct a new struct value, given its type and a callback
function. The function should return the value for a member, given its name.

#### func (StructValue) Type

```go
func (v StructValue) Type() Type
```

#### type SwitchStmt

```go
type SwitchStmt struct {
	Expr  Expression
	Stmts *CompoundStmt

	SwitchCst *parse.Leaf
	LParenCst *parse.Leaf
	RParenCst *parse.Leaf
}
```

switch(expr) stmts

#### type Symbol

```go
type Symbol interface {
	Name() string
}
```

Symbol is an interface for all declared symbols (anything that has a name and is
scoped) in the language. It does not always coincide with the "declaration"
concept. E.g.: 'int a,b;' is a single declaration, but declares two symbols; the
declaration 'struct A { ... } a;' declares one symbol for the struct and another
for the variable, etc.

#### type Type

```go
type Type interface {
	// Whether two objects represent the same type, according to the language semantics.
	Equal(other Type) bool
}
```

Type represents a language type along with its precision.

#### type TypeConversionExpr

```go
type TypeConversionExpr struct {
	RetType Type

	ArgTypes []Type // Found by semantic analysis.
}
```

TypeConversionExpr represents a type conversion expression. It is only used as
the "callee" child of CallExpr. The RetType member is the type specified in the
program (the type we are converting to). The ArgTypes member is computed by
semantic analysis from the arguments. The full type of the expression (what
Type() returns) is then a FunctionType{RetType, ArgTypes}.

#### func (*TypeConversionExpr) Constant

```go
func (e *TypeConversionExpr) Constant() bool
```

#### func (*TypeConversionExpr) LValue

```go
func (e *TypeConversionExpr) LValue() bool
```

#### func (*TypeConversionExpr) Type

```go
func (e *TypeConversionExpr) Type() Type
```

#### type TypeQualifiers

```go
type TypeQualifiers struct {
	Invariant     bool
	Interpolation InterpolationQualifier
	Layout        *LayoutQualifier
	Storage       StorageQualifier

	InvariantCst       *parse.Leaf
	InterpolationCst   *parse.Leaf
	CentroidStorageCst *parse.Leaf
	StorageCst         *parse.Leaf
}
```

Invariant, interpolation, layout and storage qualifiers of declared variables.

#### type UintValue

```go
type UintValue uint32
```

UintValue represents an unsigned integer value, as a uint32.

#### func (UintValue) String

```go
func (v UintValue) String() string
```

#### func (UintValue) Type

```go
func (v UintValue) Type() Type
```

#### type UnaryExpr

```go
type UnaryExpr struct {
	Op   UnaryOp
	Expr Expression

	OpCst *parse.Leaf
}
```

UnaryExpr represents both prefix and postfix unary operators.

#### func (*UnaryExpr) Constant

```go
func (e *UnaryExpr) Constant() bool
```

#### func (*UnaryExpr) LValue

```go
func (e *UnaryExpr) LValue() bool
```

#### func (*UnaryExpr) Type

```go
func (e *UnaryExpr) Type() Type
```

#### type UnaryOp

```go
type UnaryOp BinaryOp
```

A type for unary operators of the language.

#### func  GetPrefixOp

```go
func GetPrefixOp(opname string) (op UnaryOp, present bool)
```
GetPrefixOp returns the prefix unary operator associated with the given string.
For "++" it returns UoPreinc instead of UoPostinc. If an operator is not found,
the second result is false.

#### func (UnaryOp) String

```go
func (o UnaryOp) String() string
```

#### type UniformBlock

```go
type UniformBlock struct {
	Layout  *LayoutQualifier
	SymName string
	Vars    []*MultiVarDecl

	UniformCst *parse.Leaf
	NameCst    *parse.Leaf
	LBraceCst  *parse.Leaf
	RBraceCst  *parse.Leaf
}
```

UniformBlock represents the "block" part of the uniform interface declarations.

    layout(...) uniform BlockName { int Var; ... } InstanceName;
    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                   UniformBlock
    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                           UniformDecl

Implements Symbol with Name = BlockName.

#### func (*UniformBlock) Name

```go
func (u *UniformBlock) Name() string
```

#### type UniformDecl

```go
type UniformDecl struct {
	Block   *UniformBlock
	SymName string
	Size    Expression

	NameCst      *parse.Leaf
	LBracketCst  *parse.Leaf
	RBracketCst  *parse.Leaf
	SemicolonCst *parse.Leaf
}
```

UniformDecl represents a whole uniform interface declaration.

    layout(...) uniform BlockName { int Var; ... } InstanceName;
    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                   UniformBlock
    ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
                           UniformDecl

Implements Symbol with Name = InstanceName.

#### func (*UniformDecl) Name

```go
func (u *UniformDecl) Name() string
```

#### type Value

```go
type Value interface {
	Type() Type
}
```

A langauge value. Each value has a type.

#### type ValueSymbol

```go
type ValueSymbol interface {
	Symbol
	Type() Type
	LValue() bool
	Constant() bool
}
```

ValueSymbol represents a Symbol (a named object in the language), which can be
used in a value context in the language (basically, referenced in an expression
through VarRefExpr). After a semantic analysis pass the three functions provide
information about the nature of the symbol.

#### type VarDeclCond

```go
type VarDeclCond struct {
	Sym *VariableSym
}
```

VarDeclCond is a condition holding a declaration of a single variable.
Conditions are used in ForStmt and WhileStmt.

#### type VarRefExpr

```go
type VarRefExpr struct {
	Sym ValueSymbol

	SymCst *parse.Leaf
}
```

VarRefExpr represents a reference to a declared symbol used in an expression.

#### func (*VarRefExpr) Constant

```go
func (e *VarRefExpr) Constant() bool
```

#### func (*VarRefExpr) LValue

```go
func (e *VarRefExpr) LValue() bool
```

#### func (*VarRefExpr) Type

```go
func (e *VarRefExpr) Type() Type
```

#### type VariableSym

```go
type VariableSym struct {
	SymType Type
	SymName string
	Init    Expression

	Quals *TypeQualifiers

	NameCst  *parse.Leaf
	EqualCst *parse.Leaf

	SymLValue bool // Found by semantic analysis.
}
```

VariableSym is a declaration of a single variable, possibly with an
initialization expression. It is normally a part of a MultiVarDecl (for normal
declaration) or a VarDeclCond (for declarations which are a part of while and
for statements).

#### func (*VariableSym) Constant

```go
func (v *VariableSym) Constant() bool
```

#### func (*VariableSym) LValue

```go
func (v *VariableSym) LValue() bool
```

#### func (*VariableSym) Name

```go
func (v *VariableSym) Name() string
```

#### func (*VariableSym) Type

```go
func (v *VariableSym) Type() Type
```

#### type VectorComponentKind

```go
type VectorComponentKind uint8
```

VectorComponentKind is a type representing various ways of addressing vector
components.

```go
const (
	ComponentKindNone VectorComponentKind = iota
	ComponentKindXYZW
	ComponentKindRGBA
	ComponentKindSTPQ
)
```
Constants representing ways of adressing vector components.

#### type VectorValue

```go
type VectorValue struct {
	Members []Value
	ValType BuiltinType
}
```

VectorValue represents vector values.

#### func  NewVectorValue

```go
func NewVectorValue(t BareType, member func(i uint8) Value) VectorValue
```
NewVectorValue constructs a new vector value, given its type and a member
callback function. The number of elements is determined from the type. The
member function should return the member value, given its id. It is called in
the natural order.

#### func (VectorValue) Type

```go
func (v VectorValue) Type() Type
```

#### type WhileStmt

```go
type WhileStmt struct {
	Cond interface{}
	Stmt interface{}

	WhileCst  *parse.Leaf
	LParenCst *parse.Leaf
	RParenCst *parse.Leaf
}
```

while(cond) stmt
