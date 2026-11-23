package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// The exact six-line welcome from FR-012, and what follows every answer in a session.
const welcome = "Welcome to the interactive calculator session.\n" +
	"A spec-driven development demo, built with GitHub Spec Kit.\n" +
	"Operations: +, -, *, /\n" +
	"Use exactly 2 numbers and 1 operator, for example 2+3\n" +
	"===========================================================\n" +
	"Please enter your input. Example: 10/5\n"

const afterAnswer = "-----------------------------------------------------------\n" +
	"Please enter your input.\n"

// errorsBlock formats problems as the spec's Errors: block (FR-015).
func errorsBlock(problems ...string) string {
	var b strings.Builder
	b.WriteString("Errors:\n")
	for _, p := range problems {
		b.WriteString("- " + p + "\n")
	}
	return b.String()
}

// T016: one-shot mode prints the bare result or the Errors: block, and sets the exit code
// (FR-011; research R2, R3, R4).
func TestOneShot(t *testing.T) {
	tests := []struct {
		fr         string
		name       string
		args       []string
		wantStdout string
		wantStderr string
		wantCode   int
	}{
		{fr: "FR-011", name: "three pieces", args: []string{"2", "+", "3"}, wantStdout: "5\n"},
		{fr: "FR-011", name: "one piece", args: []string{"2+3"}, wantStdout: "5\n"},
		{fr: "FR-011", name: "quoted multiplication", args: []string{"6 * 7"}, wantStdout: "42\n"},
		{fr: "FR-011", name: "multiplication without spaces", args: []string{"6*7"}, wantStdout: "42\n"},
		{fr: "FR-011", name: "division by zero", args: []string{"8/0"},
			wantStderr: errorsBlock(msgDivisionByZero), wantCode: 1},
		{fr: "FR-015", name: "two problems", args: []string{"1000 * 5"},
			wantStderr: errorsBlock(msgTooManyDigits("1000"), msgMultLimit("1000")), wantCode: 1},
		{fr: "FR-005", name: "empty argument", args: []string{""},
			wantStderr: errorsBlock(msgNothingEntered), wantCode: 1},
		{fr: "FR-013", name: "quit is not a command", args: []string{"quit"},
			wantStderr: errorsBlock(msgMissingNumber), wantCode: 1},
	}
	for _, tt := range tests {
		t.Run(tt.fr+"/"+tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(tt.args, strings.NewReader(""), &stdout, &stderr)
			if stdout.String() != tt.wantStdout {
				t.Errorf("stdout = %q; want %q", stdout.String(), tt.wantStdout)
			}
			if stderr.String() != tt.wantStderr {
				t.Errorf("stderr = %q; want %q", stderr.String(), tt.wantStderr)
			}
			if code != tt.wantCode {
				t.Errorf("exit code = %d; want %d", code, tt.wantCode)
			}
		})
	}
}

// T017: the session prints the welcome once, then for each line an answer, the separator,
// and the prompt; it ends quietly with exit code 0 (FR-012, FR-013).
func TestInteractive(t *testing.T) {
	tests := []struct {
		fr         string
		name       string
		stdin      string
		wantStdout string
	}{
		{fr: "FR-012", name: "result", stdin: "2+3\nquit\n",
			wantStdout: welcome + "Result: 5\n" + afterAnswer},
		{fr: "FR-012", name: "errors", stdin: "abc + 3\nexit\n",
			wantStdout: welcome + errorsBlock(msgNotNumber("abc")) + afterAnswer},
		{fr: "FR-012", name: "multiplication with spaces", stdin: "6 * 7\n",
			wantStdout: welcome + "Result: 42\n" + afterAnswer},
		{fr: "FR-012", name: "several lines", stdin: "2+3\n8/0\n",
			wantStdout: welcome + "Result: 5\n" + afterAnswer + errorsBlock(msgDivisionByZero) + afterAnswer},
		{fr: "FR-013", name: "blank and space-only lines", stdin: "\n   \n2+3\n",
			wantStdout: welcome + "Result: 5\n" + afterAnswer},
		{fr: "FR-013", name: "quit with spaces", stdin: "  quit  \n2+3\n",
			wantStdout: welcome},
		{fr: "FR-013", name: "end of input", stdin: "",
			wantStdout: welcome},
	}
	for _, tt := range tests {
		t.Run(tt.fr+"/"+tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := run(nil, strings.NewReader(tt.stdin), &stdout, &stderr)
			if stdout.String() != tt.wantStdout {
				t.Errorf("stdout = %q; want %q", stdout.String(), tt.wantStdout)
			}
			if stderr.String() != "" {
				t.Errorf("stderr = %q; want nothing", stderr.String())
			}
			if code != 0 {
				t.Errorf("exit code = %d; want 0", code)
			}
		})
	}
}

// failingReader fails on the first read, like a broken input stream.
type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("input stream broke") }

// T018: if reading the input fails, the error goes to the error output and the exit code is 1
// (research R4).
func TestInteractiveReadFailure(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := run(nil, failingReader{}, &stdout, &stderr)
	if !strings.Contains(stderr.String(), "input stream broke") {
		t.Errorf("stderr = %q; want it to contain the read error", stderr.String())
	}
	if code != 1 {
		t.Errorf("exit code = %d; want 1", code)
	}
}
