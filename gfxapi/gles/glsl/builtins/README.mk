# builtins
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles/glsl/builtins"

Package builtins contains the definitions of all OpenGL ES Shading Language
builtin functions.

It is not expected to be used directly by the user, it just provides the other
packages with the list of builtin functions, their signatures and
implementations.

## Usage

```go
var BuiltinFunctions = []ast.BuiltinFunction{

	ast.BuiltinFunction{
		RetType: &ast.BuiltinType{Precision: ast.HighP, Type: ast.TUint},
		SymName: "packSnorm2x16",
		Params:  makeParams([]ast.Type{&ast.BuiltinType{Type: ast.TVec2}}),
		Impl: func(v []ast.Value) ast.Value {
			vec := v[0].(ast.VectorValue)
			x, y := toFloat(vec.Members[0]), toFloat(vec.Members[1])
			x = round(math.Min(math.Max(x, -1.0), 1.0) * factorSPack)
			y = round(math.Min(math.Max(y, -1.0), 1.0) * factorSPack)
			return ast.UintValue(int16(x)) + ast.UintValue(int16(y))<<16
		},
	},
	ast.BuiltinFunction{
		RetType: &ast.BuiltinType{Precision: ast.HighP, Type: ast.TUint},
		SymName: "packUnorm2x16",
		Params:  makeParams([]ast.Type{&ast.BuiltinType{Type: ast.TVec2}}),
		Impl: func(v []ast.Value) ast.Value {
			vec := v[0].(ast.VectorValue)
			x, y := toFloat(vec.Members[0]), toFloat(vec.Members[1])
			x = round(math.Min(math.Max(x, 0.0), 1.0) * factorUPack)
			y = round(math.Min(math.Max(y, 0.0), 1.0) * factorUPack)
			return ast.UintValue(uint16(x)) + ast.UintValue(uint16(y))<<16
		},
	},
	ast.BuiltinFunction{
		RetType: &ast.BuiltinType{Precision: ast.HighP, Type: ast.TVec2},
		SymName: "unpackSnorm2x16",
		Params:  makeParams([]ast.Type{&ast.BuiltinType{Precision: ast.HighP, Type: ast.TUint}}),
		Impl: func(v []ast.Value) ast.Value {
			pack := toUint(v[0])
			vec := []int16{int16(pack), int16(pack >> 16)}
			return makeFloatVector(2, func(i uint8) ast.Value {
				return ast.FloatValue(math.Min(
					math.Max(float64(vec[i])/factorSPack, -1.0), 1.0))
			})
		},
	},
	ast.BuiltinFunction{
		RetType: &ast.BuiltinType{Precision: ast.HighP, Type: ast.TVec2},
		SymName: "unpackUnorm2x16",
		Params:  makeParams([]ast.Type{&ast.BuiltinType{Precision: ast.HighP, Type: ast.TUint}}),
		Impl: func(v []ast.Value) ast.Value {
			pack := toUint(v[0])
			vec := []uint16{uint16(pack), uint16(pack >> 16)}
			return makeFloatVector(2, func(i uint8) ast.Value {
				return ast.FloatValue(math.Min(
					math.Max(float64(vec[i])/factorUPack, -1.0), 1.0))
			})
		},
	},
	ast.BuiltinFunction{
		RetType: &ast.BuiltinType{Precision: ast.HighP, Type: ast.TUint},
		SymName: "packHalf2x16",
		Params:  makeParams([]ast.Type{&ast.BuiltinType{Precision: ast.MediumP, Type: ast.TVec2}}),
		Impl: func(v []ast.Value) ast.Value {
			vec := v[0].(ast.VectorValue)
			x := binary.NewFloat16(float32(toFloat(vec.Members[0])))
			y := binary.NewFloat16(float32(toFloat(vec.Members[1])))
			return ast.UintValue(x) + ast.UintValue(y)<<16
		},
	},
	ast.BuiltinFunction{
		RetType: &ast.BuiltinType{Precision: ast.MediumP, Type: ast.TVec2},
		SymName: "unpackHalf2x16",
		Params:  makeParams([]ast.Type{&ast.BuiltinType{Precision: ast.HighP, Type: ast.TUint}}),
		Impl: func(v []ast.Value) ast.Value {
			pack := toUint(v[0])
			vec := []binary.Float16{binary.Float16(pack), binary.Float16(pack >> 16)}
			return makeFloatVector(2,
				func(i uint8) ast.Value { return ast.FloatValue(vec[i].Float32()) })
		},
	},

	ast.BuiltinFunction{
		RetType: &ast.BuiltinType{Type: ast.TVec3},
		SymName: "cross",
		Params: makeParams([]ast.Type{&ast.BuiltinType{Type: ast.TVec3},
			&ast.BuiltinType{Type: ast.TVec3}}),

		Impl: func(v []ast.Value) ast.Value {
			x, y := v[0].(ast.VectorValue).Members, v[1].(ast.VectorValue).Members

			return ast.NewVectorValue(ast.TVec3, func(i uint8) ast.Value {
				j, k := (i+1)%3, (i+2)%3
				return x[j].(ast.FloatValue)*y[k].(ast.FloatValue) -
					y[j].(ast.FloatValue)*x[k].(ast.FloatValue)
			})
		},
	},
}
```
BuiltinFunctions contains the list of builtin functions defined by the package,
in no particular order.

PS: The list showed by godoc is not complete. More functions are added during
package init(). For a definitive list, see The OpenGL ES Shading Language
specification, section 8.
