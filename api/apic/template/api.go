package template

import (
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

// Unpack throws away type information to work around template system limitations
// When you have a value of an interface type that carries methods, it fails to
// introspect the concrete type for it's members, so the template can't see them.
// The result of Upack no longer has a type, so the concrete type members become
// visible.
func (*Functions) Unpack(v interface{}) interface{} { return v }

// GetAnnotation finds and returns the annotation on ty with the specified name.
// If the annotation cannot be found, or ty does not support annotations then
// GetAnnotation returns nil.
func (*Functions) GetAnnotation(ty interface{}, name string) *semantic.Annotation {
	a, ok := ty.(semantic.Annotated)
	if !ok {
		return nil
	}
	return a.GetAnnotation(name)
}

// GetArrayParamCount returns the inferred array size for param as a semantic
// expression. If the array size cannot be inferred, then GetArrayParamCount
// returns nil.
func (*Functions) GetArrayParamCount(param *semantic.Parameter) interface{} {
	f := param.Function
	for _, s := range f.Block.Statements {
		assert, ok := s.(*semantic.Assert)
		if !ok {
			break
		}
		binary, ok := assert.Condition.(*semantic.BinaryOp)
		if !ok {
			continue
		}
		if (binary.Operator != ast.OpGE) && (binary.Operator != ast.OpEQ) {
			continue
		}
		length, ok := binary.LHS.(*semantic.Length)
		if !ok {
			continue
		}
		p, ok := length.Object.(*semantic.Parameter)
		if !ok {
			continue
		}
		if p != param {
			continue
		}
		return binary.RHS
	}
	return nil
}
