package reflow

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
)

var (
	command = &commands.Command{
		Name:      "reflow",
		ShortHelp: "Reflows an api file for smart formatting",
	}
)

func init() {
	command.Run = doReflow
	commands.Register(command)
}

func doReflow(flags flag.FlagSet) {
	args := flags.Args()
	if len(args) < 1 {
		commands.Usage("Missing api file\n")
	}
	for _, apiName := range args {
		info, err := ioutil.ReadFile(apiName)
		commands.MaybeError(apiName, err)
		commands.Log("Compiling api file %q\n", apiName)
		parsed, errs := parser.Parse(string(info[:]))
		commands.CheckErrors(apiName, errs)
		compiled, errs := resolver.Resolve(parsed)
		commands.CheckErrors(apiName, errs)
		commands.Log("Reflowing api %s\n", apiName)

		commands.Log("Writing output to %q\n", apiName)
		w, err := os.Create(apiName)
		commands.MaybeError(apiName, err)
		out := bufio.NewWriter(w)
		err = compiled.AST.CST.WriteTo(out)
		commands.MaybeError(apiName, err)
		err = out.Flush()
		commands.MaybeError(apiName, err)
		err = w.Close()
		commands.MaybeError(apiName, err)
	}
}
