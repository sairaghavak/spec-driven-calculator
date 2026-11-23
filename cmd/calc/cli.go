package main

import "io"

// run is the whole program behind main, with its inputs and outputs passed in so tests can
// drive it. With arguments it runs one operation (one-shot mode); without, an interactive
// session. It returns the exit code.
//
// Empty for the red commit (T003): the tests can run and fail on a wrong answer.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	return 1
}
