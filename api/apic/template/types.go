package template

import (
	"errors"
	"fmt"
	"unicode"

	"reflect"

	"strings"

	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

var (
	// NodeTypeList exposes all the types valid as Is*name* tests
	NodeTypeList = []interface{}{
		// Primitive types
		"",    // string
		false, // bool
		// primitive value node types
		semantic.BoolValue(true),
		semantic.Int32Value(0),
		semantic.Uint32Value(0),
		semantic.Int64Value(0),
		semantic.Uint64Value(0),
		semantic.StringValue(""),
		semantic.Float32Value(1.0),
		semantic.Float64Value(1.0),
		// semantic node types
		semantic.API{},
		semantic.ArrayIndex{},
		semantic.Array{},
		semantic.Assert{},
		semantic.Assign{},
		semantic.BinaryOp{},
		semantic.Branch{},
		semantic.Builtin{},
		semantic.Call{},
		semantic.Cast{},
		semantic.Choice{},
		semantic.ClassInitializer{},
		semantic.Class{},
		semantic.DeclareLocal{},
		semantic.EnumEntry{},
		semantic.Enum{},
		semantic.Field{},
		semantic.Function{},
		semantic.Global{},
		semantic.Iteration{},
		semantic.Local{},
		semantic.MapIndex{},
		semantic.Map{},
		semantic.Member{},
		semantic.Observed{},
		semantic.Parameter{},
		semantic.Pointer{},
		semantic.Pseudonym{},
		semantic.Return{},
		semantic.Select{},
		semantic.StaticArray{},
		semantic.Switch{},
		semantic.Unknown{},
		// node interface types
		(*semantic.Annotated)(nil),
		(*semantic.Expression)(nil),
		(*semantic.Type)(nil),
	}

	nodeTypes = map[string]reflect.Type{}
)

func init() {
	for _, n := range NodeTypeList {
		nt := baseType(n)
		name := nt.Name()
		nodeTypes[name] = nt
	}
}

func initNodeTypes(f *Functions) {
	for _, b := range semantic.BuiltinTypes {
		b := b
		name := "Is" + strings.Title(b.Name)
		if b == semantic.PointerType {
			name = name + "Type"
		}
		f.funcs[name] = func(t semantic.Type) bool {
			return t == b
		}
	}
	for name, t := range nodeTypes {
		for _, r := range name {
			if unicode.IsUpper(r) {
				f.funcs["Is"+name] = isTypeTest(t)
			}
			break
		}
	}
}

// Returns the resolved semantic type of an expression node.
func (*Functions) TypeOf(v interface{}) (semantic.Type, error) {
	if v == nil {
		return semantic.VoidType, nil
	}
	switch e := v.(type) {
	case *semantic.Field:
		return e.Type, nil
	case semantic.Expression:
		return e.ExpressionType(), nil
	default:
		return nil, fmt.Errorf("Type \"%T\" is not an expression", v)
	}
}

// Returns true if v is one of the primitive numeric types.
func (*Functions) IsNumericValue(v interface{}) bool {
	switch v.(type) {
	case semantic.Int8Value,
		semantic.Uint8Value,
		semantic.Int16Value,
		semantic.Uint16Value,
		semantic.Int32Value,
		semantic.Uint32Value,
		semantic.Int64Value,
		semantic.Uint64Value,
		semantic.Float32Value,
		semantic.Float64Value:
		return true
	default:
		return false
	}
}

// Returns the base name of the type of v
func baseType(v interface{}) reflect.Type {
	ty := reflect.TypeOf(v)
	for ty != nil && ty.Kind() == reflect.Ptr {
		return ty.Elem()
	}
	return ty
}

func isTypeTest(t reflect.Type) func(v interface{}) bool {
	return func(v interface{}) bool {
		ty := baseType(v)
		if ty == nil {
			return false
		}
		return ty.AssignableTo(t)
	}
}

// Asserts that the type of v is in the list of expected types
func (*Functions) AssertType(v interface{}, expected ...string) (string, error) {
	got := baseType(v)
	if got == nil {
		return "", fmt.Errorf("Calling AssertType with nil value")
	}
	matched := 0
	for _, e := range expected {
		et, found := nodeTypes[e]
		if !found {
			return "", fmt.Errorf("%s is not a valid type", e)
		}
		if got.AssignableTo(et) {
			matched++
		}
	}
	if matched > 0 {
		return "", nil
	}

	msg := fmt.Sprintf("Type assertion. Got: %s, Expected: ", got)
	if c := len(expected); c > 1 {
		msg += strings.Join(expected[:c-1], ", ")
		msg += " or " + expected[c-1]
	} else {
		msg += expected[0]
	}
	return "", errors.New(msg)
}
