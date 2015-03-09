// Copyright (C) 2014 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resolver

import (
	"bytes"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

type macroStub struct {
	scope    *scope
	function *semantic.Function
}

func (m *macroStub) ExpressionType() semantic.Type { return m.function.Signature }

func functionSignature(ctx *context, out *semantic.Function) {
	in := out.AST
	args := make([]semantic.Type, 0, len(in.Parameters)-1)
	out.Outputs = make([]*semantic.Parameter, 0, len(in.Parameters))
	out.FullParameters = make([]*semantic.Parameter, 0, len(in.Parameters))
	for i, inp := range in.Parameters {
		outp := parameter(ctx, out, inp)
		if inp.This {
			if i == 0 {
				out.This = outp
			} else {
				ctx.errorf(inp, "this only allowed on arg 0")
			}
		}
		if i < len(in.Parameters)-1 {
			if isVoid(outp.Type) {
				ctx.errorf(in, "void typed parameter %s on function %s", outp.Name, out.Name)
			}
			if !inp.This {
				args = append(args, outp.Type)
			}
			if inp.Output {
				out.Outputs = append(out.Outputs, outp)
			}
			out.FullParameters = append(out.FullParameters, outp)
		} else {
			out.Return = outp
			if !isVoid(outp.ExpressionType()) {
				out.Return.Name = "result"
				out.Outputs = append(out.Outputs, outp)
				out.FullParameters = append(out.FullParameters, outp)
			}
		}
	}
	out.Signature = getSignature(ctx, in, out.Return.Type, args)
	ctx.mappings[in] = out
}

func parameter(ctx *context, owner *semantic.Function, in *ast.Parameter) *semantic.Parameter {
	out := &semantic.Parameter{
		AST:      in,
		Input:    in.Input,
		Output:   in.Output,
		Function: owner,
	}
	if in.Name != nil {
		out.Name = in.Name.Value
	}
	out.Annotations = annotations(ctx, in.Annotations)
	out.Type = type_(ctx, in.Type)
	ctx.mappings[in] = out
	return out
}

func functionBody(ctx *context, owner semantic.Type, out *semantic.Function) {
	in := out.AST
	out.Owner = owner
	if in.Block != nil {
		if in.Block.Docs != nil {
			out.Docs = semantic.Docs(in.Block.Docs.URL)
		}
		ctx.with(semantic.VoidType, func() {
			for _, p := range out.FullParameters {
				if p != out.Return {
					ctx.add(p.Name, p)
				}
			}
			if out.This != nil {
				ctx.add(string(ast.KeywordThis), out.This)
			}
			out.Annotations = annotations(ctx, in.Annotations)
			out.Block = block(ctx, in.Block, out)
		})
	}
	ctx.mappings[in] = out
}

func method(ctx *context, in *ast.Function) {
	out := &semantic.Function{AST: in, Name: in.Name.Value}
	functionSignature(ctx, out)
	t := out.This.Type
	switch t := t.(type) {
	case *semantic.Pointer:
		if class, ok := t.To.(*semantic.Class); !ok {
			ctx.errorf(in, "expected this as a reference to a class, got %s[%T]", typename(t.To), t.To)
		} else {
			class.Methods = append(class.Methods, out)
			class.Members[out.Name] = out
			functionBody(ctx, class, out)
		}
	case *semantic.Pseudonym:
		t.Methods = append(t.Methods, out)
		t.Members[out.Name] = out
		functionBody(ctx, t, out)
	case *semantic.Class:
		t.Methods = append(t.Methods, out)
		t.Members[out.Name] = out
		functionBody(ctx, t, out)
	default:
		ctx.errorf(in, "invalid type for this , got %s[%T]", typename(t), t)
	}
	ctx.mappings[in] = out
}

func getSignature(ctx *context, at ast.Node, r semantic.Type, args []semantic.Type) *semantic.Signature {
	buffer := bytes.Buffer{}
	buffer.WriteString("fun_")
	buffer.WriteString(r.Typename())
	buffer.WriteString("_")
	for _, a := range args {
		buffer.WriteString("_")
		buffer.WriteString(a.Typename())
	}
	name := buffer.String()
	for _, s := range ctx.api.Signatures {
		if s.Name == name {
			if !equal(r, s.Return) {
				ctx.icef(at, "Signature %s found with non matching return type, got %s expected %s", name, typename(s.Return), typename(r))
			}
			if len(args) != len(s.Arguments) {
				ctx.icef(at, "Signature %s found with %d arguments, expected %s", name, len(s.Arguments), len(args))
			}
			for i, a := range args {
				if !equal(a, s.Arguments[i]) {
					ctx.icef(at, "Signature %s found with non matching arg at %d, got %s expected %s", name, i, typename(s.Arguments[i]), typename(a))
				}
			}
			return s
		}
	}
	out := &semantic.Signature{
		Name:      name,
		Return:    r,
		Arguments: args,
	}
	ctx.api.Signatures = append(ctx.api.Signatures, out)
	return out
}
