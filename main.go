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
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.Parse()
}

func main() {
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	bc, rc := 0, 0
	for scanner.Scan() {
		if verbose && len(scanner.Bytes()) > 1 {
			fmt.Printf("rune %d `%s` is %v\n", rc, scanner.Text(), scanner.Bytes())
		}
		bc += len(scanner.Bytes())
		rc++
	}
	fmt.Println("Bytes: ", bc)
	fmt.Println("Runes: ", rc)
}
