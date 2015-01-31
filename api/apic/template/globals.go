package template

import "strings"

var (
	globalList stringSetFlag
)

type globalMap map[string]interface{}

type stringSetFlag []string

func (f *stringSetFlag) String() string    { return strings.Join(f.Strings(), ":") }
func (f *stringSetFlag) Strings() []string { return ([]string)(*f) }

func (f *stringSetFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func initGlobals(f *Functions) {
	for _, g := range globalList.Strings() {
		v := strings.SplitN(g, "=", 2)
		f.globals[v[0]] = v[1]
	}
}

// Gets or sets a template global variable
// Example:
//  {{Global "CatSays" "Meow"}}
//  The cat says: {{Global "CatSays"}}
func (f *Functions) Global(name string, values ...interface{}) (interface{}, error) {
	switch len(values) {
	case 0:
		value, _ := f.globals[name]
		return value, nil
	case 1:
		f.globals[name] = values[0]
		return "", nil
	default:
		f.globals[name] = values
		return "", nil
	}
}
