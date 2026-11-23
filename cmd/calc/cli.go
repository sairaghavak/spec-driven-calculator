package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// The exact session texts from FR-012.
const (
	welcomeText = "Welcome to the interactive calculator session.\n" +
		"A spec-driven development demo, built with GitHub Spec Kit.\n" +
		"Operations: +, -, *, /\n" +
		"Use exactly 2 numbers and 1 operator, for example 2+3\n" +
		"===========================================================\n" +
		"Please enter your input. Example: 10/5\n"
	separatorLine = "-----------------------------------------------------------\n"
	promptLine    = "Please enter your input.\n"
)

// run is the whole program behind main, with its inputs and outputs passed in so tests can
// drive it. With arguments it runs one operation (one-shot mode); without, an interactive
// session. It returns the exit code.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	if len(args) > 0 {
		return runOnce(strings.Join(args, " "), stdout, stderr)
	}
	return runSession(stdin, stdout, stderr)
}

// runOnce evaluates one operation (T019, FR-011; research R2 to R4): the bare result to
// stdout and exit code 0, or the Errors: block to stderr and exit code 1.
func runOnce(input string, stdout, stderr io.Writer) int {
	result, err := Evaluate(input)
	if err != nil {
		fmt.Fprint(stderr, errorsText(err))
		return 1
	}
	fmt.Fprintln(stdout, result)
	return 0
}

// runSession reads one operation per line until quit, exit, or the end of input (T020,
// FR-012, FR-013). Everything goes to stdout; a read failure goes to stderr with exit code 1
// (FR-016).
func runSession(stdin io.Reader, stdout, stderr io.Writer) int {
	fmt.Fprint(stdout, welcomeText)
	lines := bufio.NewScanner(stdin)
	for lines.Scan() {
		line := strings.TrimSpace(lines.Text())
		switch line {
		case "":
			continue
		case "quit", "exit":
			return 0
		}
		if result, err := Evaluate(line); err != nil {
			fmt.Fprint(stdout, errorsText(err))
		} else {
			fmt.Fprintf(stdout, "Result: %s\n", result)
		}
		fmt.Fprint(stdout, separatorLine+promptLine)
	}
	if err := lines.Err(); err != nil {
		fmt.Fprintf(stderr, "error reading input: %v\n", err)
		return 1
	}
	return 0
}

// errorsText formats an error as the Errors: block, one "- " line per problem (FR-015).
func errorsText(err error) string {
	problems := []string{err.Error()}
	var inputErr *InputError
	if errors.As(err, &inputErr) {
		problems = inputErr.Problems
	}
	var b strings.Builder
	b.WriteString("Errors:\n")
	for _, p := range problems {
		b.WriteString("- " + p + "\n")
	}
	return b.String()
}
