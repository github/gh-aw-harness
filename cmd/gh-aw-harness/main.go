package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/github/gh-aw-harness/pkg/harness"
)

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s %s\n", harness.Name(), harness.Version)
		return
	}

	fmt.Fprintln(os.Stdout, harness.Name())
}
