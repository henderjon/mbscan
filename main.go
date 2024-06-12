package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
	verbose bool
)

func init() {
	bi := getBuildInfo()
	flag.BoolVar(&verbose, "v", false, "verbose; show any multi-byte characters found in the input stream")
	flag.Usage = Usage(Info{
		Bin:            bi.getBinName(),
		Version:        bi.getBuildVersion(),
		CompiledBy:     bi.getCompiledBy(),
		BuildTimestamp: bi.getBuildTimestamp(),
	})
	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	bc, rc := 0, 0
	for scanner.Scan() {
		if verbose && len(scanner.Bytes()) > 1 {
			fmt.Fprintf(os.Stderr, "rune %d `%s` is %v\n", rc, scanner.Text(), scanner.Bytes())
		}
		bc += len(scanner.Bytes())
		rc++
	}

	// fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, " Bytes: ", bc)
	fmt.Fprintln(os.Stderr, " Runes: ", rc)

	if bc != rc {
		os.Exit(1)
	}

	os.Exit(0)
}
