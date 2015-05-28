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

func apiNames(ctx *context, in *ast.API) {
	// Build and register the high level semantic objects
	for _, e := range in.Enums {
		n := &semantic.Enum{AST: e, Name: e.Name.Value}
		ctx.api.Enums = append(ctx.api.Enums, n)
		ctx.api.Members[n.Name] = n
		ctx.addType(n)
	}
	for _, c := range in.Classes {
		n := &semantic.Class{AST: c, Name: c.Name.Value, Members: semantic.Members{}}
		ctx.api.Classes = append(ctx.api.Classes, n)
		ctx.api.Members[n.Name] = n
		ctx.addType(n)
	}
	for _, p := range in.Pseudonyms {
		n := &semantic.Pseudonym{AST: p, Name: p.Name.Value, Members: semantic.Members{}}
		ctx.api.Pseudonyms = append(ctx.api.Pseudonyms, n)
		ctx.api.Members[n.Name] = n
		ctx.addType(n)
	}
	for _, m := range in.Macros {
		stub := &macroStub{}
		ctx.macros = append(ctx.macros, stub)
		stub.function = &semantic.Function{AST: m, Name: m.Name.Value}
		ctx.add(stub.function.Name, stub)
	}
	for _, e := range in.Externs {
		n := &semantic.Function{AST: e, Name: e.Name.Value}
		ctx.api.Externs = append(ctx.api.Externs, n)
		ctx.api.Members[n.Name] = n
	}
	for _, m := range in.Commands {
		f := &semantic.Function{AST: m, Name: m.Name.Value}
		if !m.Parameters[0].This {
			ctx.api.Functions = append(ctx.api.Functions, f)
			ctx.api.Members[f.Name] = f
		} else {
			ctx.api.Methods = append(ctx.api.Methods, f)
		}
	}
	for _, f := range in.Fields {
		n := &semantic.Global{AST: f, Name: f.Name.Value}
		ctx.api.Globals = append(ctx.api.Globals, n)
		ctx.api.Members[n.Name] = n
	}
	// Add all the alias remaps
	for _, a := range in.Aliases {
		ctx.addType(&Alias{AST: a, Name: a.Name.Value})
	}
}

func resolve(ctx *context) {
	for name, member := range ctx.api.Members {
		ctx.add(name, member)
	}
	// Now resolve all the references
	for _, e := range ctx.api.Enums {
		enum(ctx, e)
	}
	for _, g := range ctx.api.Globals {
		global(ctx, g)
	}
	for _, p := range ctx.api.Pseudonyms {
		pseudonym(ctx, p)
	}
	for _, m := range ctx.macros {
		functionSignature(ctx, m.function)
		m.scope = ctx.scope
	}
	for _, e := range ctx.api.Externs {
		functionSignature(ctx, e)
		functionBody(ctx, nil, e)
	}
	for _, c := range ctx.api.Classes {
		class(ctx, c)
	}
	for _, f := range ctx.api.Functions {
		functionSignature(ctx, f)
		functionBody(ctx, nil, f)
	}
	for _, m := range ctx.api.Methods {
		method(ctx, m)
	}
	sort.Sort(slicesByName(ctx.api.Slices))
	sort.Sort(mapsByName(ctx.api.Maps))
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
		ctx.mappings[a] = entry
	}
	return out
}

func global(ctx *context, out *semantic.Global) {
	in := out.AST
	out.Annotations = annotations(ctx, in.Annotations)
	out.Type = type_(ctx, in.Type)
	if isVoid(out.Type) {
		ctx.errorf(in, "void typed global variable %s", out.Name)
	}
	if in.Default != nil {
		ctx.with(out.Type, func() {
			out.Default = expression(ctx, in.Default)
		})
		dt := out.Default.ExpressionType()
		if !assignable(out.Type, dt) {
			ctx.errorf(in, "cannot assign %s to %s", typename(dt), typename(out.Type))
		}
	}
	ctx.mappings[in] = out
}

// Alias is used as a temporary type holder during type resolution.
// It is not present in the final semantic tree returned, but may be present
// in the AST -> semantic map.
type Alias struct {
	AST  *ast.Alias
	Name string
	To   semantic.Type
}

func (t Alias) Typename() string                 { return t.Name }
func (t Alias) Member(name string) semantic.Node { return nil }

// slicesByName is used to sort the slice list by name for generated code stability
type slicesByName []*semantic.Slice

func (a slicesByName) Len() int           { return len(a) }
func (a slicesByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slicesByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

// mapsByName is used to sort the map list by name for generated code stability
type mapsByName []*semantic.Map

func (a mapsByName) Len() int           { return len(a) }
func (a mapsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a mapsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
