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
	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func functionSignature(ctx *context, out *semantic.Function) {
	in := out.AST
	out.FullParameters = make([]*semantic.Parameter, len(in.Parameters))
	for i, p := range in.Parameters {
		if i > 0 && p.This {
			ctx.errorf(p, "this only allowed on arg 0")
		}
		out.FullParameters[i] = parameter(ctx, out, p)
	}
	if in.Parameters[0].This {
		out.This = out.FullParameters[0]
	}
	out.Return = out.FullParameters[len(out.FullParameters)-1]
	if out.Return.Type == semantic.VoidType {
		out.FullParameters = out.FullParameters[0 : len(out.FullParameters)-1]
	} else {
		out.Return.Name = "result"
	}
	for _, p := range out.FullParameters {
		if p.Output {
			out.Outputs = append(out.Outputs, p)
		}
	}
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
}
