package template

import (
	"bytes"
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

type Functions struct {
	templates *template.Template
	funcs     template.FuncMap
	globals   globalMap
}

func newFunctions() *Functions {
	f := &Functions{
		templates: template.New("FunctionHolder"),
		funcs:     template.FuncMap{},
		globals:   globalMap{},
	}
	v := reflect.ValueOf(f)
	t := v.Type()
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		r, _ := utf8.DecodeRuneInString(m.Name)
		if unicode.IsUpper(r) {
			c := v.MethodByName(m.Name)
			f.funcs[m.Name] = c.Interface()
		}
	}
	initNodeTypes(f)
	initGlobals(f)
	if *tracer != "" {
		pattern := regexp.MustCompile(*tracer)
		for n, c := range f.funcs {
			if pattern.MatchString(n) {
				f.funcs[n] = trace(n, c)
			}
		}
	}
	f.templates.Funcs(f.funcs)
	return f
}

func trace(name string, f interface{}) func(values ...interface{}) (interface{}, error) {
	depth := ""
	return func(values ...interface{}) (interface{}, error) {
		fmt.Print(depth, name)
		depth += " |  "
		args := make([]reflect.Value, len(values))
		for i, v := range values {
			if v == nil {
				args[i] = reflect.ValueOf(&v).Elem()
			} else {
				args[i] = reflect.ValueOf(v)
			}
			switch v := v.(type) {
			case string:
				fmt.Printf(" %q", v)
			case semantic.Type:
				fmt.Printf(" [%s]", v.Typename())
			default:
				fmt.Printf(" <%T>", v)
			}
		}
		fmt.Println()
		defer func() {
			depth = depth[:len(depth)-4]
		}()
		result := reflect.ValueOf(f).Call(args)
		value := result[0].Interface()
		var err error
		if len(result) > 1 {
			err, _ = result[1].Interface().(error)
		}
		return value, err
	}
}

// Error raises an error terminating execution of the template.
//   {{Error "Foo returned error: %s" $err}}
func (f *Functions) Error(s string, args ...interface{}) (string, error) {
	return "", fmt.Errorf(s, args...)
}

// Log prints s and optional format arguments to stdout. Example:
//   {{Log "%s %s" "Hello" "world}}
func (f *Functions) Log(s string, args ...interface{}) string {
	fmt.Printf(s+"\n", args...)
	return ""
}

// File calls the macro templateName with the arguments values and writes the
// output to the file fileName, overwriting the file if it already exists. The
// file name is relative to the master output file and adopts the same file
// extension.
// See Macro for more information about argument passing.
// Example:
//  {{File "methods" "EmitMethods" $class}}
func (f *Functions) File(fileName, templateName string, arguments ...interface{}) (string, error) {
	result, err := f.Macro(templateName, arguments...)
	filePath := filepath.Join(filepath.Dir(*outputFilename), fileName+filepath.Ext(*outputFilename))
	commands.MaybeError(filePath, err)
	return "", write(f.templates.Name(), filePath, result)
}

func (*Functions) buildArgs(base map[string]interface{}, values ...interface{}) (map[string]interface{}, error) {
	data := make(map[string]interface{}, len(values)/2)
	for k, v := range base {
		data[k] = v
	}
	if len(values)%2 != 0 {
		return nil, errors.New("bad argument count to macro, must be in pairs")
	}
	for i := 0; i < len(values)-1; i += 2 {
		switch k := values[i].(type) {
		case string:
			data[k] = values[i+1]
		default:
			return nil, errors.New("invoke keys must be strings")
		}
	}
	return data, nil
}

// Macro invokes the template macro with the specified name and returns the
// template output as a string. If no arguments are passed then $ will be nil
// for the called macro. If a single argument is passed then $ will be the value
// of that argument. If more than one argument is passed then arguments is used
// as name-value pairs, where name is a field on $. For example:
//  {{define "SingleParameterMacro"}}
//      $ is: {{$}}
//  {{end}}
//
//  {{define "MultipleParameterMacro"}}
//      $.ArgA is: {{$.ArgA}}, $.ArgB is: {{$.ArgB}}
//  {{end}}
//
//  {{Macro "SingleParameterMacro"}}
//  {{/* Returns "$ is: nil" */}}
//
//  {{Macro "SingleParameterMacro" 42}}
//  {{/* Returns "$ is: 42" */}}
//
//  {{Macro "MultipleParameterMacro" "ArgA" 4 "ArgB" 2}}
//  {{/* Returns "$.ArgA is: 4, $.ArgB is: 2" */}}
func (f *Functions) Macro(name string, arguments ...interface{}) (string, error) {
	var arg interface{}
	switch len(arguments) {
	case 0:
		arg = nil
	case 1:
		arg = arguments[0]
	default:
		inherit, ok := arguments[0].(map[string]interface{})
		if ok {
			arguments = arguments[1:]
		}
		data, err := f.buildArgs(inherit, arguments...)
		if err != nil {
			return "", err
		}
		arg = data
	}
	var buf bytes.Buffer
	err := f.templates.ExecuteTemplate(&buf, name, arg)
	return strings.TrimSpace(buf.String()), err
}
