package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"text/template"
)

// Tmpl is a basic man page[-ish] looking template
const Tmpl = `
{{define "manual"}}
NAME
  {{.Bin}} - compare byte and rune counts from stdin

SYNOPSIS
  $ {{.Bin}}
  $ {{.Bin}} [-h|help]

DESCRIPTION
  {{.Bin}} scans stdin and counts the number of bytes and runes to detect
  multi-byte characters. Be default, it will output the difference of
  bytes vs runes.

EXAMPLES
  $ {{.Bin}} -v < file.txt
  $ cat file.txt | {{.Bin}} -v

OPTIONS
{{.Options}}
VERSION
  version:  {{.Version}}
  compiled: {{.CompiledBy}}
  built:    {{.BuildTimestamp}}

{{end}}
`

// Info represents the information used in the default Tmpl string
type Info struct {
	Tmpl           string
	Bin            string
	Version        string
	CompiledBy     string
	BuildTimestamp string
	Options        string
	QuickHelp      string
}

// Usage wraps a set of `Info` and creates a flag.Usage func
func Usage(info Info) func() {
	if len(info.Tmpl) == 0 {
		info.Tmpl = Tmpl
	}

	t := template.Must(template.New("manual").Parse(info.Tmpl))

	return func() {
		var def bytes.Buffer
		flag.CommandLine.SetOutput(&def)
		flag.PrintDefaults()

		info.Options = def.String()
		t.Execute(os.Stdout, info)
	}
}

func MultiUsage(flags []*flag.FlagSet, info Info) func() {
	if len(info.Tmpl) == 0 {
		info.Tmpl = Tmpl
	}

	t := template.Must(template.New("manual").Parse(info.Tmpl))

	return func() {
		var def bytes.Buffer
		for _, f := range flags {
			fmt.Fprintf(&def, "\nSUBCOMMAND: %s\n", f.Name())
			f.SetOutput(&def)
			f.PrintDefaults()
		}
		info.Options = def.String()
		t.Execute(os.Stdout, info)
	}
}
