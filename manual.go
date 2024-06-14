package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"text/template"
)

// ManualPageTemplate is a basic man page[-ish] looking template
const ManualPageTemplate = `
{{define "manual"}}
NAME
  {{.BinName}} - compare byte and rune counts from stdin

SYNOPSIS
  $ {{.BinName}}
  $ {{.BinName}} [-h|help]

DESCRIPTION
  {{.BinName}} scans stdin and counts the number of bytes and runes to detect
  multi-byte characters. Be default, it will output the difference of
  bytes vs runes.

EXAMPLES
  $ {{.BinName}} -v < file.txt
  $ cat file.txt | {{.BinName}} -v

OPTIONS
{{.Options}}
VERSION
  version:  {{.BuildVersion}}
  compiled: {{.CompiledBy}}
  built:    {{.BuildTimestamp}}

{{end}}
`

// these vars are built at compile time, DO NOT ALTER
var (
	// Version adds build information
	BinName string
	// Version adds build information
	BuildVersion string
	// BuildTimestamp adds build information
	BuildTimestamp string
	// CompiledBy adds the make/model that was used to compile
	CompiledBy string
)

// Usage wraps a set of `Info` and creates a flag.Usage func
func RenderManualPage() func() {
	t := template.Must(template.New("manual").Parse(ManualPageTemplate))

	return func() {
		var def bytes.Buffer
		flag.CommandLine.SetOutput(&def)
		flag.PrintDefaults()

		t.Execute(os.Stdout, struct {
			Options        string
			BinName        string
			BuildVersion   string
			BuildTimestamp string
			CompiledBy     string
		}{
			Options:        def.String(),
			BinName:        BinName,
			BuildVersion:   BuildVersion,
			BuildTimestamp: BuildTimestamp,
			CompiledBy:     CompiledBy,
		})
	}
}

func RenderManualPageMulti(flags []*flag.FlagSet) func() {
	t := template.Must(template.New("manual").Parse(ManualPageTemplate))

	return func() {
		var def bytes.Buffer
		for _, f := range flags {
			fmt.Fprintf(&def, "\nSUBCOMMAND: %s\n", f.Name())
			f.SetOutput(&def)
			f.PrintDefaults()
		}

		t.Execute(os.Stdout, struct {
			Options        string
			BinName        string
			BuildVersion   string
			BuildTimestamp string
			CompiledBy     string
		}{
			Options:        def.String(),
			BinName:        BinName,
			BuildVersion:   BuildVersion,
			BuildTimestamp: BuildTimestamp,
			CompiledBy:     CompiledBy,
		})
	}
}

func GetVersionString() string {
	return fmt.Sprintf("%s version %s", BinName, BuildVersion)
}
