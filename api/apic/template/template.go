package template

import (
	"bytes"
	"flag"
	"io/ioutil"
	"text/template"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
)

var (
	command = &commands.Command{
		Name:      "template",
		ShortHelp: "Passes the ast to a template for code generation",
	}
	mainTemplate   = command.Flags.String("m", "Main", "The main template name to execute")
	outputFilename = command.Flags.String("o", "out", "The output file path")
	tracer         = command.Flags.String("t", "", "The template function trace expression")
	formatEnable   = command.Flags.Bool("f", true, "Reformat output")
	indentSize     = command.Flags.Int("i", 2, "Indentation size")
)

func init() {
	command.Flags.Var(&globalList, "G", "A global value setting for the template")
	command.Run = doTemplate
	commands.Register(command)
}

func write(templateName, outputPath, data string) error {
	result := reformat(outputPath, data)
	commands.Log("Writing output to %q\n", outputPath)
	return ioutil.WriteFile(outputPath, []byte(result), 0666)
}

func doTemplate(flags flag.FlagSet) {
	args := flags.Args()
	if len(args) < 1 {
		commands.Usage("Missing api file\n")
	}
	apiName := args[0]
	templateNames := args[1:]
	commands.Log("Reading api file %q\n", apiName)
	info, err := ioutil.ReadFile(apiName)
	commands.MaybeError(apiName, err)
	commands.Log("Compiling api file %q\n", apiName)
	parsed, errs := parser.Parse(string(info[:]))
	commands.CheckErrors(apiName, errs)
	compiled, errs := resolver.Resolve(parsed)
	commands.CheckErrors(apiName, errs)

	f := newFunctions()
	for _, templateName := range templateNames {
		commands.Log("Reading template %q\n", templateName)
		tmplData, err := ioutil.ReadFile(templateName)
		commands.MaybeError(templateName, err)
		template.Must(f.templates.New(templateName).Parse(string(tmplData)))
	}
	commands.Log("Executing template %q\n", *mainTemplate)
	var buf bytes.Buffer
	err = f.templates.ExecuteTemplate(&buf, *mainTemplate, compiled)
	commands.MaybeError("", err)
	err = write(templateNames[0], *outputFilename, buf.String())
	commands.MaybeError(*outputFilename, err)
}
