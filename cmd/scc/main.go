// Package main is the entry point for the `scc` (spiffe-compliance-checker) CLI.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "scc: subcommands not yet wired up")
	os.Exit(2)
}
