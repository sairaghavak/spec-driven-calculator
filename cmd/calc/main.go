// Command calc is a command-line calculator for one operation with two numbers.
// See specs/001-basic-arithmetic/spec.md.
package main

import "os"

func main() {
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
