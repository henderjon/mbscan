package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	verbose bool
	quiet   bool
	path    string
)

func init() {
	bi := getBuildInfo()
	flag.BoolVar(&verbose, "v", false, "verbose; shows multi-byte characters found in the input stream and the byte and rune counts")
	flag.StringVar(&path, "path", "", "an opaque string that is used to identify the source of the input stream")
	flag.BoolVar(&quiet, "s", false, "silent; no output only exit codes")
	flag.Usage = Usage(Info{
		Bin:            bi.getBinName(),
		Version:        bi.getBuildVersion(),
		CompiledBy:     bi.getCompiledBy(),
		BuildTimestamp: bi.getBuildTimestamp(),
	})
	flag.Parse()
}

func main() {
	bc, rc := scan(os.Stdin, func(w io.Writer) tokenFunc {
		return func(runePos int, byt []byte) {
			if verbose {
				fmt.Fprintf(w, "rune %d `%s` is %v\n", runePos, string(byt), byt)
			}
		}
	}(stdout))

	if verbose {
		fmt.Fprintln(stdout, " Bytes: ", bc)
		fmt.Fprintln(stdout, " Runes: ", rc)
	}

	if !quiet {
		fmt.Fprintf(stdout, "%d %s\n", bc-rc, path)
	}

	if bc != rc {
		os.Exit(1)
	}

	os.Exit(0)
}

type tokenFunc func(runePos int, byt []byte)

func scan(r io.Reader, tf tokenFunc) (int, int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	bc, rc := 0, 0
	for scanner.Scan() {
		if len(scanner.Bytes()) > 1 {
			tf(rc, scanner.Bytes())
		}
		bc += len(scanner.Bytes())
		rc++
	}
	return bc, rc
}
