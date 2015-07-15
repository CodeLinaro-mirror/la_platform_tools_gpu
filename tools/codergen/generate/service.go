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

package generate

import (
	"fmt"
	"sort"

	"android.googlesource.com/platform/tools/gpu/binary/schema"

	"golang.org/x/tools/go/types"
)

const (
	logLogger = "android.googlesource.com/platform/tools/gpu/log.Logger"
)

// Service is a description of an rpc service.
type Service struct {
	Name    string    // The name of the service.
	Prefix  string    // The file prefix for the service.
	Package string    // The package of the service.
	Methods []*Method // The methods of the service.
}

// Method describes a method of a service.
type Method struct {
	Name   string
	Call   Call
	Result Result
}

// Call wraps parameters to a service method.
type Call struct {
	*Struct
}

// Params returns the schema field list that represent the method parameters.
func (c Call) Params() schema.FieldList {
	return c.Struct.Class.Fields
}

// Result wraps the return values of a serice method.
type Result struct {
	*Struct
}

// Type returns the schema type of the return value if present, nil if not.
func (r Result) Type() schema.Type {
	if len(r.Struct.Class.Fields) < 1 {
		return nil
	}
	return r.Struct.Class.Fields[0].Type
}

// Params returns the schema field list that represent the method parameters.
func (r Result) List() schema.FieldList {
	return r.Struct.Class.Fields
}

func serviceStruct(m *Module, name string, tuple *types.Tuple, count int, b *types.Interface) *Struct {
	class := schema.Class{Name: name, Package: m.Source.Types.Name()}
	for i := 0; i < count; i++ {
		entry := tuple.At(i)
		class.Fields = append(class.Fields, schema.Field{
			Declared: entry.Name(),
			Type:     fromType(m.Source.Types, entry.Type(), "", m.Imports, b),
		})
	}
	s := &Struct{Class: class}
	// The generated structs will not be parsed by codergen, so they must be self registered.
	m.Structs = append(m.Structs, s)
	return s
}

func (m *Module) addService(n *types.TypeName, b *types.Interface) error {
	if n.Name() != m.Directive("service", "*.invalid.*") {
		return nil
	}
	t := n.Type().Underlying().(*types.Interface)
	s := &Service{
		Name:    n.Name(),
		Prefix:  fmt.Sprint(m.Directive("service.prefix", m.Source.Types.Name())),
		Package: m.Source.Types.Name(),
	}
	for i := 0; i < t.NumExplicitMethods(); i++ {
		decl := t.ExplicitMethod(i)
		sig := decl.Type().(*types.Signature)
		paramCount := sig.Params().Len()
		if paramCount <= 0 {
			return fmt.Errorf("RPC method %s.%s has %d parameters values, expected at least 1", s.Name, decl.Name(), paramCount)
		}
		lastParam := sig.Params().At(paramCount - 1).Type().String()
		if lastParam != logLogger {
			return fmt.Errorf("RPC method %s.%s param %d is %s, expected log.Logger", s.Name, decl.Name(), paramCount-1, lastParam)
		}
		resultCount := sig.Results().Len()
		if resultCount < 1 || resultCount > 2 {
			return fmt.Errorf("RPC method %s.%s has %d return values, expected either 1 or 2", s.Name, decl.Name(), resultCount)
		}
		lastResult := sig.Results().At(resultCount - 1).Type().String()
		if lastResult != "error" {
			return fmt.Errorf("RPC method %s.%s returns %s, expected error", s.Name, decl.Name(), lastResult)
		}
		method := &Method{Name: decl.Name()}
		method.Call.Struct = serviceStruct(m, "call"+decl.Name(), sig.Params(), paramCount-1, b)
		method.Result.Struct = serviceStruct(m, "result"+decl.Name(), sig.Results(), resultCount-1, b)
		if resultCount > 1 && method.Result.Struct.Class.Fields[0].Declared == "" {
			// for methods with an unnamed first return value, default the name to "value" to match legacy behaviour.
			method.Result.Struct.Class.Fields[0].Declared = "value"
		}
		method.Call.UpdateID()
		method.Result.UpdateID()
		s.Methods = append(s.Methods, method)
	}
	m.Services = append(m.Services, s)
	return nil
}

func (m *Module) finaliseServices() {
	sort.Sort(servicesByName(m.Services))
	for _, s := range m.Services {
		sort.Sort(methodsByName(s.Methods))
	}
}

type servicesByName []*Service

func (a servicesByName) Len() int           { return len(a) }
func (a servicesByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a servicesByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type methodsByName []*Method

func (a methodsByName) Len() int           { return len(a) }
func (a methodsByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a methodsByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
