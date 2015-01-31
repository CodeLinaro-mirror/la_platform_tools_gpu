package main

import (
	"flag"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	_ "android.googlesource.com/platform/tools/gpu/api/apic/reflow"
	_ "android.googlesource.com/platform/tools/gpu/api/apic/template"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		commands.Usage("Must supply a verb\n")
	}
	verb := args[0]
	matches := commands.Filter(verb)
	switch len(matches) {
	case 1:
		c := matches[0]
		commands.Log("Running %q\n", c.Name)
		c.Flags.Parse(args[1:])
		c.Run(c.Flags)
	case 0:
		commands.Usage("Verb '%s' is unknown\n", verb)
	default:
		commands.Usage("Verb '%s' is ambiguous\n", verb)
	}
}
