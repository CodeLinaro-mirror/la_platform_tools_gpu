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
	"sort"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

func api(ctx *context, out *semantic.API) {
	in := out.AST
	ctx.with(semantic.VoidType, func() {
		// Build and register the high level semantic objects
		out.Enums = make([]*semantic.Enum, len(in.Enums))
		for i, e := range in.Enums {
			out.Enums[i] = &semantic.Enum{AST: e, Name: e.Name.Value}
			ctx.addType(out.Enums[i])
		}
		out.Classes = make([]*semantic.Class, len(in.Classes))
		for i, c := range in.Classes {
			out.Classes[i] = &semantic.Class{AST: c, Name: c.Name.Value, Members: semantic.Members{}}
			ctx.addType(out.Classes[i])
		}
		out.Pseudonyms = make([]*semantic.Pseudonym, len(in.Pseudonyms))
		for i, p := range in.Pseudonyms {
			out.Pseudonyms[i] = &semantic.Pseudonym{AST: p, Name: p.Name.Value, Members: semantic.Members{}}
			ctx.addType(out.Pseudonyms[i])
		}
		out.Macros = make([]*semantic.Function, len(in.Macros))
		for i, m := range in.Macros {
			out.Macros[i] = &semantic.Function{AST: m, Name: m.Name.Value}
			out.Members[out.Macros[i].Name] = out.Macros[i]
		}
		out.Externs = make([]*semantic.Function, len(in.Externs))
		for i, e := range in.Externs {
			out.Externs[i] = &semantic.Function{AST: e, Name: e.Name.Value}
			out.Members[out.Externs[i].Name] = out.Externs[i]
		}
		for _, m := range in.Commands {
			if !m.Parameters[0].This {
				f := &semantic.Function{AST: m, Name: m.Name.Value}
				out.Functions = append(out.Functions, f)
				out.Members[f.Name] = f
			}
		}
		out.Globals = make([]*semantic.Global, len(in.Fields))
		for i, f := range in.Fields {
			out.Globals[i] = &semantic.Global{AST: f, Name: f.Name.Value}
			out.Members[out.Globals[i].Name] = out.Globals[i]
		}
		for name, member := range out.Members {
			ctx.add(name, member)
		}
		// Add all the alias remaps
		for _, a := range in.Aliases {
			ctx.addType(&alias{AST: a, Name: a.Name.Value})
		}
		// Now resolve all the references
		for _, e := range out.Enums {
			enum(ctx, e)
		}
		for _, g := range out.Globals {
			global(ctx, g)
		}
		for _, p := range out.Pseudonyms {
			pseudonym(ctx, p)
		}
		for _, m := range out.Macros {
			functionSignature(ctx, m)
			functionBody(ctx, nil, m)
		}
		for _, e := range out.Externs {
			functionSignature(ctx, e)
			functionBody(ctx, nil, e)
		}
		for _, c := range out.Classes {
			class(ctx, c)
		}
		for _, f := range out.Functions {
			functionSignature(ctx, f)
			functionBody(ctx, nil, f)
		}
		for _, m := range in.Commands {
			if m.Parameters[0].This {
				method(ctx, m)
			}
		}
	})
	sort.Sort(arraysByName(out.Arrays))
	sort.Sort(mapsByName(out.Maps))
}

func annotations(ctx *context, in ast.Annotations) semantic.Annotations {
	if len(in) == 0 {
		return nil
	}
	out := semantic.Annotations{}
	for _, a := range in {
		entry := &semantic.Annotation{AST: a, Name: a.Name.Value}
		for _, arg := range a.Arguments {
			entry.Arguments = append(entry.Arguments, expression(ctx, arg))
		}
		out = append(out, entry)
	}
	return out
}

func global(ctx *context, out *semantic.Global) {
	in := out.AST
	out.Annotations = annotations(ctx, in.Annotations)
	out.Type = type_(ctx, in.Type)
	if in.Default != nil {
		ctx.with(out.Type, func() {
			out.Default = expression(ctx, in.Default)
		})
		dt := out.Default.ExpressionType()
		if !assignable(out.Type, dt) {
			ctx.errorf(in, "cannot assign %s to %s", typename(dt), typename(out.Type))
		}
	}
}

// alias is used as a temporary type holder during type resolution,
// it is not present in the final tree returned.
type alias struct {
	AST  *ast.Alias
	Name string
	To   semantic.Type
}

func (t alias) Typename() string               { return t.Name }
func (t alias) Member(name string) interface{} { return nil }

// arraysByName is used to sort the array list by name for generated code stability
type arraysByName []*semantic.Array

func (a arraysByName) Len() int           { return len(a) }
func (a arraysByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a arraysByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// mapsByName is used to sort the map list by name for generated code stability
type mapsByName []*semantic.Map

func (a mapsByName) Len() int           { return len(a) }
func (a mapsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a mapsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
