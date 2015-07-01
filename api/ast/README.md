# ast
--
    import "android.googlesource.com/platform/tools/gpu/api/ast"

Package ast holds the set of types used in the abstract syntax tree
representation of the api language.

## Usage

```go
const (
	// Keyword strings represent places in the syntax where a word has special
	// meaning.
	KeywordAPI       = "api"
	KeywordAlias     = "alias"
	KeywordBitfield  = "bitfield"
	KeywordCase      = "case"
	KeywordClass     = "class"
	KeywordCmd       = "cmd"
	KeywordConst     = "const"
	KeywordElse      = "else"
	KeywordEnum      = "enum"
	KeywordExtern    = "extern"
	KeywordFalse     = "false"
	KeywordFor       = "for"
	KeywordIf        = "if"
	KeywordImport    = "import"
	KeywordIn        = "in"
	KeywordMacro     = "macro"
	KeywordNull      = "null"
	KeywordReturn    = "return"
	KeywordPseudonym = "type"
	KeywordSwitch    = "switch"
	KeywordThis      = "this"
	KeywordTrue      = "true"
	KeywordWhen      = "when"
)
```

```go
const (
	// The set of operators understood by the parser.
	OpUnknown       = "?"
	OpBlockStart    = "{"
	OpBlockEnd      = "}"
	OpIndexStart    = "["
	OpIndexEnd      = "]"
	OpSlice         = ":"
	OpListStart     = "("
	OpListSeparator = ","
	OpListEnd       = ")"
	OpAssign        = "="
	OpAssignPlus    = "+="
	OpAssignMinus   = "-="
	OpDeclare       = ":="
	OpMember        = "."
	OpExtends       = ":"
	OpAnnotation    = "@"
	OpInitialise    = ":"
	OpPointer       = "*"
	OpEQ            = "=="
	OpGT            = ">"
	OpLT            = "<"
	OpGE            = ">="
	OpLE            = "<="
	OpNE            = "!="
	OpOr            = "||"
	OpAnd           = "&&"
	OpPlus          = "+"
	OpMinus         = "-"
	OpMultiply      = "*"
	OpDivide        = "/"
	OpRange         = ".."
	OpNot           = "!"
	OpIn            = "in"
	OpGeneric       = "!"
)
```

```go
const (
	// Runes with special meaning to the parser.
	Quote = '"'
)
```

```go
var (
	Operators       = []string{}            // all valid operator strings, sorted in descending length order
	UnaryOperators  = map[string]struct{}{} // the map of valid unary operators
	BinaryOperators = map[string]struct{}{} // the map of valid boolean operators
)
```

#### type API

```go
type API struct {
	CST        *parse.Branch // underlying parse structure for this node
	Imports    []*Import     // api files imported with the "import" keyword
	Macros     []*Function   // functions declared with the "macro" keyword
	Externs    []*Function   // functions declared with the "extern" keyword
	Commands   []*Function   // functions declared with the "cmd" keyword
	Pseudonyms []*Pseudonym  // strong type aliases declared with the "type" keyword
	Aliases    []*Alias      // weak type aliases declared with the "alias" keyword
	Enums      []*Enum       // enumerated types, declared with the "enum" keyword
	Classes    []*Class      // class types, declared with the "class" keyword
	Fields     []*Field      // variables declared at the global scope
}
```

API is the root of the AST tree, and constitutes one entire parsed file. It
holds the set of top level AST nodes, grouped by type.

#### func (API) Fragment

```go
func (t API) Fragment() parse.Fragment
```

#### type Alias

```go
type Alias struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the alias
	Name        *Identifier   // the name of the alias
	To          Node          // the type it is an alias for
}
```

Alias represents a weak type alias, with structure «"alias" type name». An alias
does not declare a new type, just a reusable name for a common type.

#### func (Alias) Fragment

```go
func (t Alias) Fragment() parse.Fragment
```

#### type Annotation

```go
type Annotation struct {
	CST       *parse.Branch // underlying parse structure for this node
	Name      *Identifier   // the name part (between the @ and the brackets)
	Arguments []Node        // the list of arguments (the bit in brackets)
}
```

Annotation is the AST node that represents «@name(arguments) constructs»

#### func (Annotation) Fragment

```go
func (t Annotation) Fragment() parse.Fragment
```

#### type Annotations

```go
type Annotations []*Annotation
```

Annotations represents the set of Annotation objects that apply to another AST
node.

#### type Assign

```go
type Assign struct {
	CST      *parse.Branch // underlying parse structure for this node
	LHS      Node          // the location to store the value into
	Operator string        // the assignment operator being applied
	RHS      Node          // the value to store
}
```

Assign represents a «location {,+,-}= value» statement that assigns a value to
an existing mutable location.

#### func (Assign) Fragment

```go
func (t Assign) Fragment() parse.Fragment
```

#### type BinaryOp

```go
type BinaryOp struct {
	CST      *parse.Branch // underlying parse structure for this node
	LHS      Node          // the expression on the left of the operator
	Operator string        // the operator being applied
	RHS      Node          // the expression on the right of the operator
}
```

BinaryOp represents any binary operation applied to two expressions.

#### func (BinaryOp) Fragment

```go
func (t BinaryOp) Fragment() parse.Fragment
```

#### type Block

```go
type Block struct {
	CST        *parse.Branch // underlying parse structure for this node
	Statements []Node        // The set of statements that make up the block
}
```

Block represents a linear sequence of statements, most often the contents of a
{} pair.

#### func (Block) Fragment

```go
func (t Block) Fragment() parse.Fragment
```

#### type Bool

```go
type Bool struct {
	CST   *parse.Leaf // underlying parse leaf for this node
	Value bool        // The value of the boolean
}
```

Bool is used for the "true" and "false" keywords.

#### func (Bool) Fragment

```go
func (t Bool) Fragment() parse.Fragment
```

#### type Branch

```go
type Branch struct {
	CST       *parse.Branch // underlying parse structure for this node
	Condition Node          // the condition to use to select which block is active
	True      *Block        // the block to use if condition is true
	False     *Block        // the block to use if condition is false
}
```

Branch represents an «"if" condition { trueblock } "else" { falseblock }»
structure.

#### func (Branch) Fragment

```go
func (t Branch) Fragment() parse.Fragment
```

#### type Call

```go
type Call struct {
	CST       *parse.Branch // underlying parse structure for this node
	Target    Node          // the function to invoke
	Arguments []Node        // the arguments to the function
}
```

Call is an expression that invokes a function with a set of arguments. It has
the structure «target(arguments)» where target must be a function and arguments
is a comma separated list of expressions.

#### func (Call) Fragment

```go
func (t Call) Fragment() parse.Fragment
```

#### type Case

```go
type Case struct {
	CST        *parse.Branch // underlying parse structure for this node.
	Conditions []Node        // the set of conditions that would select this case
	Block      *Block        // the block to run if this case is selected
}
```

Case represents a «"case" conditions: block» structure within a switch
statement. The conditions are a comma separated list of expressions the switch
statement value will be compared against.

#### func (Case) Fragment

```go
func (t Case) Fragment() parse.Fragment
```

#### type Class

```go
type Class struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the class
	Name        *Identifier   // the name of the class
	Fields      []*Field      // the fields of the class
}
```

Class represents a class type declaration of the form «"class" name { fields }»

#### func (Class) Fragment

```go
func (t Class) Fragment() parse.Fragment
```

#### type DeclareLocal

```go
type DeclareLocal struct {
	CST  *parse.Branch // underlying parse structure for this node
	Name *Identifier   // the name to give the new local
	RHS  Node          // the value to store in that local
}
```

DeclareLocal represents a «name := value» statement that declares a new
immutable local variable with the specified value and inferred type.

#### func (DeclareLocal) Fragment

```go
func (t DeclareLocal) Fragment() parse.Fragment
```

#### type Enum

```go
type Enum struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the enum
	Name        *Identifier   // the name of the enum
	IsBitfield  bool          // whether this enum represents a bitfield form
	Entries     []*EnumEntry  // the set of valid entries for this enum
	Extends     []*Identifier // deprecated list of enums this extends
}
```

Enum represents an enumerated type declaration, of the form «"enum" name {
entries }» where entries is a comma separated list of «name = value»

#### func (Enum) Fragment

```go
func (t Enum) Fragment() parse.Fragment
```

#### type EnumEntry

```go
type EnumEntry struct {
	CST   *parse.Branch // underlying parse structure for this node
	Owner *Enum         // the enum this entry is a part of
	Name  *Identifier   // the name this entry is given
	Value *Number       // the value of this entry
}
```

EnumEntry represents a single value in an enumerated type.

#### func (EnumEntry) Fragment

```go
func (t EnumEntry) Fragment() parse.Fragment
```

#### type Field

```go
type Field struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the field
	Type        Node          // the type the field holds
	Name        *Identifier   // the name of the field
	Default     Node          // the default value expression for the field
}
```

Field represents a field of a class or api, with the structure «type name =
expression»

#### func (Field) Fragment

```go
func (t Field) Fragment() parse.Fragment
```

#### type Function

```go
type Function struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to this function
	Name        *Identifier   // the name of the function
	Parameters  []*Parameter  // the parameters the function takes
	Block       *Block        // the body of the function if present
}
```

Function represents the declaration of a callable entity, any of "cmd", "extern"
or "macro". Its structure is «return_type name(parameters) body» where
parameters is a comma separated list and body is an optional block.

#### func (Function) Fragment

```go
func (t Function) Fragment() parse.Fragment
```

#### type Generic

```go
type Generic struct {
	CST       *parse.Branch // underlying parse structure for this node
	Name      *Identifier   // the generic identifier.
	Arguments []Node        // the type arguments to the generic.
}
```

Generic represents a identifier modified by type arguments. It looks like:
«identifier ! ( arg | <arg {, arg} )>»

#### func (Generic) Fragment

```go
func (t Generic) Fragment() parse.Fragment
```

#### type Group

```go
type Group struct {
	CST        *parse.Branch // underlying parse structure for this node
	Expression Node          // the expression within the parentheses
}
```

Group represents the «(expression)» construct, a single parenthesized
expression.

#### func (Group) Fragment

```go
func (t Group) Fragment() parse.Fragment
```

#### type Identifier

```go
type Identifier struct {
	CST   *parse.Leaf // underlying parse leaf for this node
	Value string      // the identifier
}
```

Identifier holds a parsed identifier in the parse tree.

#### func (Identifier) Fragment

```go
func (t Identifier) Fragment() parse.Fragment
```

#### type Import

```go
type Import struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the import
	Name        *Identifier   // the name to import an api file as
	Path        *String       // the relative path to the api file
}
```

Import is the AST node that represents «import name "path"» constructs

#### func (Import) Fragment

```go
func (t Import) Fragment() parse.Fragment
```

#### type Imported

```go
type Imported struct {
	CST  *parse.Branch // underlying parse structure for this node
	From *Identifier   // the import this name is from
	Name *Identifier   // the name being imported
}
```

Imported represents an imported type name.

#### func (Imported) Fragment

```go
func (t Imported) Fragment() parse.Fragment
```

#### type Index

```go
type Index struct {
	CST    *parse.Branch // underlying parse structure for this node
	Object Node          // the object to index
	Index  Node          // the index to lookup
}
```

Index represents any expression of the form «object[index]» Used for arrays,
maps and bitfields.

#### func (Index) Fragment

```go
func (t Index) Fragment() parse.Fragment
```

#### type IndexedType

```go
type IndexedType struct {
	CST       *parse.Branch // underlying parse structure for this node
	ValueType Node          // The element type exposed by the indexed type
	Index     Node          // the index of the type
}
```

IndexedType represents a type declaration with an indexing suffix, which looks
like «type[index]»

#### func (IndexedType) Fragment

```go
func (t IndexedType) Fragment() parse.Fragment
```

#### type Invalid

```go
type Invalid struct{}
```

Invalid is used when an error was encountered in the parsing, but we want to
keep going. If there are no errors, this will never be in the tree.

#### func (Invalid) Fragment

```go
func (t Invalid) Fragment() parse.Fragment
```

#### type Iteration

```go
type Iteration struct {
	CST      *parse.Branch // underlying parse structure for this node
	Variable *Identifier   // the variable to use for the iteration value
	Iterable Node          // the expression that produces the iterable to loop over
	Block    *Block        // the block to run once per item in the iterable
}
```

Iteration represents a «"for" variable "in" iterable { block }» structure.

#### func (Iteration) Fragment

```go
func (t Iteration) Fragment() parse.Fragment
```

#### type Member

```go
type Member struct {
	CST    *parse.Branch // underlying parse structure for this node
	Object Node          // the object to get a member of
	Name   *Identifier   // the name of the member to get
}
```

Member represents an expressions that access members of objects. Always of the
form «object.name» where object is an expression.

#### func (Member) Fragment

```go
func (t Member) Fragment() parse.Fragment
```

#### type NamedArg

```go
type NamedArg struct {
	CST   *parse.Branch // underlying parse structure for this node
	Name  *Identifier   // the name of the parameter this value is for
	Value Node          // the value to use for that parameter
}
```

NamedArg represents a «name = value» expressionas a function argument.

#### func (NamedArg) Fragment

```go
func (t NamedArg) Fragment() parse.Fragment
```

#### type Node

```go
type Node interface {
	// Fragment returns the CST parse Fragment for this AST node.
	Fragment() parse.Fragment
}
```

Node represents any AST-node type.

#### type Null

```go
type Null struct {
	CST *parse.Leaf // underlying parse leaf for this node
}
```

Null represents the null literal. This is the default value for the inferred
type, and must be used in a context where the type can be inferred.

#### func (Null) Fragment

```go
func (t Null) Fragment() parse.Fragment
```

#### type Number

```go
type Number struct {
	CST   *parse.Leaf // underlying parse leaf for this node
	Value string      // the string representation of the constant
}
```

Number represents a typeless numeric constant.

#### func (Number) Fragment

```go
func (t Number) Fragment() parse.Fragment
```

#### type Parameter

```go
type Parameter struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to this parameter
	This        bool          // true if the parameter is the this pointer of a method
	Type        Node          // the type of the parameter
	Name        *Identifier   // the name the parameter as exposed to the body
}
```

Parameter represents a single parameter in the set of parameters for a Function.
It has the structure «["in"|"out"|"inout"|"this"] type name»

#### func (Parameter) Fragment

```go
func (t Parameter) Fragment() parse.Fragment
```

#### type PointerType

```go
type PointerType struct {
	CST   *parse.Branch // underlying parse structure for this node
	To    Node          // the underlying type this pointer points to
	Const bool          // wether the pointer type has the const modifier applied
}
```

PointerType represents a pointer type declaration, of the form «type*»

#### func (PointerType) Fragment

```go
func (t PointerType) Fragment() parse.Fragment
```

#### type Pseudonym

```go
type Pseudonym struct {
	CST         *parse.Branch // underlying parse structure for this node
	Annotations Annotations   // the annotations applied to the type
	Name        *Identifier   // the name of the type
	To          Node          // the underlying type
}
```

Pseudonym declares a new type in terms of another type. Has the form «"type"
type name» Pseydonyms are proper types, but the underlying type can be
discovered.

#### func (Pseudonym) Fragment

```go
func (t Pseudonym) Fragment() parse.Fragment
```

#### type Return

```go
type Return struct {
	CST   *parse.Branch // underlying parse structure for this node.
	Value Node          // the value to return
}
```

Return represents the «"return" value» construct, that assigns the value to the
result slot of the function.

#### func (Return) Fragment

```go
func (t Return) Fragment() parse.Fragment
```

#### type String

```go
type String struct {
	CST   *parse.Leaf // underlying parse leaf for this node
	Value string      // The body of the string, not including the delimiters
}
```

String represents a quoted string constant.

#### func (String) Fragment

```go
func (t String) Fragment() parse.Fragment
```

#### type Switch

```go
type Switch struct {
	CST   *parse.Branch // underlying parse structure for this node
	Value Node          // the value to match against
	Cases []*Case       // the set of cases to match the value with
}
```

Switch represents a «"switch" value { cases }» structure. The first matching
case is selected. If a switch is used as an expression, the case blocks must all
be a single expression.

#### func (Switch) Fragment

```go
func (t Switch) Fragment() parse.Fragment
```

#### type UnaryOp

```go
type UnaryOp struct {
	CST        *parse.Branch // underlying parse structure for this node
	Operator   string        // the operator being applied
	Expression Node          // the expression the operator is being applied to
}
```

UnaryOp represents any unary operation applied to an expression.

#### func (UnaryOp) Fragment

```go
func (t UnaryOp) Fragment() parse.Fragment
```

#### type Unknown

```go
type Unknown struct {
	CST *parse.Leaf // underlying parse leaf for this node
}
```

Unknown represents the "?" construct. This is used in places where an expression
takes a value that is implementation defined.

#### func (Unknown) Fragment

```go
func (t Unknown) Fragment() parse.Fragment
```
