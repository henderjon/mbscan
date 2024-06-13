package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	verbose bool
	quiet   bool
	out     LogWriter
)

func init() {
	bi := getBuildInfo()
	flag.BoolVar(&verbose, "v", false, "verbose; shows multi-byte characters found in the input stream and the byte and rune counts")
	flag.BoolVar(&quiet, "s", false, "silent; no output only exit codes")
	flag.Usage = Usage(Info{
		Bin:            bi.getBinName(),
		Version:        bi.getBuildVersion(),
		CompiledBy:     bi.getCompiledBy(),
		BuildTimestamp: bi.getBuildTimestamp(),
	})
	flag.Parse()

	out = stdout
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	bc, rc := 0, 0
	for scanner.Scan() {
		if verbose && len(scanner.Bytes()) > 1 {
			fmt.Fprintf(out, "rune %d `%s` is %v\n", rc, scanner.Text(), scanner.Bytes())
		}
		bc += len(scanner.Bytes())
		rc++
	}

	if verbose {
		fmt.Fprintln(out, " Bytes: ", bc)
		fmt.Fprintln(out, " Runes: ", rc)
	}

	if !quiet {
		fmt.Fprintln(os.Stdout, bc-rc)
	}

	if bc != rc {
		os.Exit(1)
	}

	os.Exit(0)
}
