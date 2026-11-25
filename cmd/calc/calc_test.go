package main

import (
	"slices"
	"strings"
	"testing"
	"time"
)

// Exact message texts from the spec's "Error messages" table (FR-005 to FR-008, FR-014). The operator
// list includes %, as feature 002's table of changed texts says.
const expectedForm = "expected <number> <operator> <number>, where <operator> is one of +, -, *, /, %"

const (
	msgNothingEntered = "no operation was entered; " + expectedForm
	msgTooMany        = "only one operation with two numbers is accepted; " + expectedForm
	msgMissingNumber  = "a number is missing; " + expectedForm
	msgMissingOp      = "the operator is missing; " + expectedForm
	msgMisplacedOp    = "the operator must be between the two numbers; " + expectedForm
	msgDivisionByZero = "division by zero is not allowed"
)

func msgUnknownOp(text string) string {
	return `"` + text + `" is not a supported operator; ` + expectedForm
}

func msgNotNumber(text string) string {
	return `"` + text + `" is not a number; numbers are written like 7, -3, or 2.5`
}

func msgLeadingZeros(text string) string {
	return `"` + text + `" has leading zeros, which are not allowed; numbers are written like 7, -3, or 2.5`
}

func msgTooManyDigits(text string) string {
	return `"` + text + `" has too many digits; numbers may have at most 3 digits before and 2 digits after the decimal point (largest 999.99)`
}

func msgMultLimit(text string) string {
	return `"` + text + `" is out of range for multiplication; each number must be strictly between -200 and 200`
}

// msgNotWhole is feature 002's whole-number message, from its spec's "Error messages" table.
func msgNotWhole(text string) string {
	return `"` + text + `" is not a whole number; % accepts only whole numbers like 7 or 10`
}

// evalCase is one spec example. fr names the requirement it checks (constitution Principle II).
// A case expects either a result (want) or the exact list of problems (wantErr), never both.
type evalCase struct {
	fr      string
	input   string
	want    string
	wantErr []string
}

// evalCases holds every acceptance criterion and edge case of the spec that Evaluate decides.
var evalCases = []evalCase{
	// US1: compute a basic operation (T004).
	{fr: "FR-002", input: "2 + 3", want: "5"},
	{fr: "FR-002", input: "3 - 10", want: "-7"},
	{fr: "FR-002", input: "6 * 7", want: "42"},
	{fr: "FR-002", input: "9 / 2", want: "4.5"},
	{fr: "FR-002", input: "1.5 + 2.25", want: "3.75"},
	{fr: "FR-002", input: "0 / 5", want: "0"},
	{fr: "FR-001", input: "  2   +  3  ", want: "5"},
	{fr: "FR-001", input: "2+3", want: "5"},
	{fr: "FR-001", input: "2\t+\t3", want: "5"},
	{fr: "FR-001", input: "-3-2", want: "-5"},
	{fr: "FR-001", input: "5*-2", want: "-10"},
	{fr: "FR-001", input: "5--3", want: "8"},
	{fr: "FR-003", input: "0.5 * 2", want: "1"},
	{fr: "FR-003", input: "0 + 1", want: "1"},
	{fr: "FR-003", input: "-0 + 1", want: "1"},
	{fr: "FR-004", input: "0.1 + 0.2", want: "0.3"},
	{fr: "FR-004", input: "1 / 3", want: "0.33333"},
	{fr: "FR-004", input: "2 / 3", want: "0.66667"},
	{fr: "FR-004", input: "0.01 / 999", want: "0.00001"},
	{fr: "FR-004", input: "0.01 / 16", want: "0.00063"},
	{fr: "FR-004", input: "-0.01 / 16", want: "-0.00063"},
	{fr: "FR-004", input: "0 * -5", want: "0"},
	{fr: "FR-014", input: "999 + 999", want: "1998"},
	{fr: "FR-014", input: "-999.99 - 999.99", want: "-1999.98"},
	{fr: "FR-014", input: "999.75 + 0", want: "999.75"},

	// US2: informative message for bad input (T009).
	{fr: "FR-006", input: "abc + 3", wantErr: []string{msgNotNumber("abc")}},
	{fr: "FR-005", input: "4 ^ 2", wantErr: []string{msgUnknownOp("^")}},
	{fr: "FR-005", input: "4 +", wantErr: []string{msgMissingNumber}},
	{fr: "FR-007", input: "8 / 0", wantErr: []string{msgDivisionByZero}},
	{fr: "FR-014", input: "1000 + 1", wantErr: []string{msgTooManyDigits("1000")}},
	{fr: "FR-005", input: "2 * 10 + 15 - 50", wantErr: []string{msgTooMany}},
	{fr: "FR-015", input: "1000 * 5", wantErr: []string{msgTooManyDigits("1000"), msgMultLimit("1000")}},
	// The spec's "Which shape message is shown" table, row by row.
	{fr: "FR-005", input: "", wantErr: []string{msgNothingEntered}},
	{fr: "FR-005", input: "   ", wantErr: []string{msgNothingEntered}},
	{fr: "FR-005", input: "4 x 2", wantErr: []string{msgUnknownOp("x")}},
	{fr: "FR-005", input: "2 * 10 + 15", wantErr: []string{msgTooMany}},
	{fr: "FR-005", input: "2 3 + 4", wantErr: []string{msgTooMany}},
	{fr: "FR-005", input: "4 + abc 5", wantErr: []string{msgTooMany}},
	{fr: "FR-005", input: "+ 4", wantErr: []string{msgMissingNumber}},
	{fr: "FR-005", input: "4", wantErr: []string{msgMissingNumber}},
	{fr: "FR-005", input: "quit", wantErr: []string{msgMissingNumber}},
	{fr: "FR-005", input: "exit", wantErr: []string{msgMissingNumber}},
	{fr: "FR-005", input: "4 5", wantErr: []string{msgMissingOp}},
	{fr: "FR-005", input: "4 5 +", wantErr: []string{msgMisplacedOp}},
	{fr: "FR-005", input: "+5 + 3", wantErr: []string{msgMisplacedOp}},
	{fr: "FR-005", input: "2 + + 3", wantErr: []string{msgMisplacedOp}},
	// Not numbers.
	{fr: "FR-006", input: "NaN + 1", wantErr: []string{msgNotNumber("NaN")}},
	{fr: "FR-006", input: "Inf + 1", wantErr: []string{msgNotNumber("Inf")}},
	{fr: "FR-006", input: "1e3 + 1", wantErr: []string{msgNotNumber("1e3")}},
	{fr: "FR-006", input: "0x10 + 1", wantErr: []string{msgNotNumber("0x10")}},
	{fr: "FR-006", input: "1,000 + 1", wantErr: []string{msgNotNumber("1,000")}},
	{fr: "FR-006", input: ".5 + 1", wantErr: []string{msgNotNumber(".5")}},
	{fr: "FR-006", input: "5. + 1", wantErr: []string{msgNotNumber("5.")}},
	{fr: "FR-006", input: "-abc + 3", wantErr: []string{msgNotNumber("-abc")}},
	{fr: "FR-006", input: `"5" + 3`, wantErr: []string{msgNotNumber(`"5"`)}},
	// Leading zeros: reported instead of "not a number", and the number is not checked further.
	{fr: "FR-006", input: "007 + 1", wantErr: []string{msgLeadingZeros("007")}},
	{fr: "FR-006", input: "00.5 * 2", wantErr: []string{msgLeadingZeros("00.5")}},
	{fr: "FR-015", input: "0999 - 1", wantErr: []string{msgLeadingZeros("0999")}},
	{fr: "FR-006", input: "-007 + 1", wantErr: []string{msgLeadingZeros("-007")}},
	// Digit limits.
	{fr: "FR-014", input: "0.001 + 1", wantErr: []string{msgTooManyDigits("0.001")}},
	{fr: "FR-014", input: "1.125 * 2", wantErr: []string{msgTooManyDigits("1.125")}},
	// Division by zero.
	{fr: "FR-007", input: "0 / 0", wantErr: []string{msgDivisionByZero}},
	{fr: "FR-007", input: "5 / -0", wantErr: []string{msgDivisionByZero}},
	{fr: "FR-007", input: "5 / 0.00", wantErr: []string{msgDivisionByZero}},
	{fr: "FR-015", input: "5 / 0.000", wantErr: []string{msgTooManyDigits("0.000"), msgDivisionByZero}},
	// Several errors at once: grouped by kind, left number before right within a kind.
	{fr: "FR-015", input: "abc / 0", wantErr: []string{msgNotNumber("abc"), msgDivisionByZero}},
	{fr: "FR-015", input: "1000 * abc", wantErr: []string{msgNotNumber("abc"), msgTooManyDigits("1000"), msgMultLimit("1000")}},

	// US3: multiplication limit (T014).
	{fr: "FR-008", input: "199 * 199", want: "39601"},
	{fr: "FR-008", input: "200 * 200", wantErr: []string{msgMultLimit("200"), msgMultLimit("200")}},
	{fr: "FR-008", input: "5 * 250", wantErr: []string{msgMultLimit("250")}},
	{fr: "FR-008", input: "500 + 500", want: "1000"},
	{fr: "FR-008", input: "199.5 * 2", want: "399"},
	{fr: "FR-008", input: "-500 * 2", wantErr: []string{msgMultLimit("-500")}},
	{fr: "FR-008", input: "-200 * 2", wantErr: []string{msgMultLimit("-200")}},
	{fr: "FR-008", input: "-199 * 199", want: "-39601"},

	// Feature 002, US1: remainder of two whole numbers (feature 002's T003).
	// Acceptance criteria 1 to 9.
	{fr: "002/FR-001", input: "7 % 3", want: "1"},
	{fr: "002/FR-001", input: "4 % 2", want: "0"},
	{fr: "002/FR-001", input: "-7 % 3", want: "1"},
	{fr: "002/FR-001", input: "7 % -3", want: "1"},
	{fr: "002/FR-001", input: "-7 % -3", want: "1"},
	{fr: "002/FR-001", input: "999 % 250", want: "249"},
	{fr: "002/FR-001", input: "7.5 % 10", wantErr: []string{msgNotWhole("7.5")}},
	{fr: "002/FR-001", input: "10 % 7.5", wantErr: []string{msgNotWhole("7.5")}},
	{fr: "002/FR-001", input: "7.0 % 2", wantErr: []string{msgNotWhole("7.0")}},
	{fr: "002/FR-001", input: "5 % 0", wantErr: []string{msgDivisionByZero}},
	{fr: "002/FR-001", input: "5 % -0", wantErr: []string{msgDivisionByZero}},
	{fr: "002/FR-001", input: "5 % 0.000", wantErr: []string{msgTooManyDigits("0.000"), msgNotWhole("0.000"), msgDivisionByZero}},
	{fr: "002/FR-001", input: "4 ^ 2", wantErr: []string{msgUnknownOp("^")}},
	// Edge cases. 1.5 + 2.25 is not repeated here: feature 001's FR-002 case above checks it.
	{fr: "002/FR-001", input: "7%3", want: "1"},
	{fr: "002/FR-001", input: "7%-3", want: "1"},
	{fr: "002/FR-001", input: "3 % 7", want: "3"},
	{fr: "002/FR-001", input: "0 % 5", want: "0"},
	{fr: "002/FR-001", input: "-0 % 5", want: "0"},
	{fr: "002/FR-001", input: "999 % 999", want: "0"},
	{fr: "002/FR-001", input: "999 % 7", want: "5"},
	{fr: "002/FR-001", input: "1000 % 7", wantErr: []string{msgTooManyDigits("1000")}},
	{fr: "002/FR-001", input: "007 % 2", wantErr: []string{msgLeadingZeros("007")}},
	{fr: "002/FR-001", input: "7.5 % 2.5", wantErr: []string{msgNotWhole("7.5"), msgNotWhole("2.5")}},
	{fr: "002/FR-001", input: "2.5 % abc", wantErr: []string{msgNotNumber("abc"), msgNotWhole("2.5")}},
	{fr: "002/FR-001", input: "2.5 % 1000.5", wantErr: []string{msgTooManyDigits("1000.5"), msgNotWhole("2.5"), msgNotWhole("1000.5")}},
	{fr: "002/FR-001", input: "abc % 0", wantErr: []string{msgNotNumber("abc"), msgDivisionByZero}},
	{fr: "002/FR-001", input: "7.5 % 0", wantErr: []string{msgNotWhole("7.5"), msgDivisionByZero}},
	{fr: "002/FR-001", input: "7 % 0.0", wantErr: []string{msgNotWhole("0.0"), msgDivisionByZero}},
	{fr: "002/FR-001", input: "7 % 3 % 2", wantErr: []string{msgTooMany}},
	{fr: "002/FR-001", input: "50%", wantErr: []string{msgMissingNumber}},
}

// problems splits an error into its list of problems, one per line.
// The error returned for bad input lists its problems separated by newlines.
func problems(err error) []string {
	if err == nil {
		return nil
	}
	return strings.Split(err.Error(), "\n")
}

// check reports whether Evaluate gives the expected result or the expected problems for c.
func check(c evalCase) (got string, gotErr []string, ok bool) {
	got, err := Evaluate(c.input)
	gotErr = problems(err)
	if c.wantErr != nil {
		return got, gotErr, slices.Equal(gotErr, c.wantErr)
	}
	return got, gotErr, err == nil && got == c.want
}

// T004, T009, T014: every spec example gives the expected result or the exact messages.
func TestEvaluate(t *testing.T) {
	for _, c := range evalCases {
		t.Run(c.fr+"/"+c.input, func(t *testing.T) {
			got, gotErr, ok := check(c)
			if !ok {
				t.Errorf("Evaluate(%q) = %q, problems %q; want %q, problems %q",
					c.input, got, gotErr, c.want, c.wantErr)
			}
		})
	}
}

// T010, FR-010: every rejected input returns no result, and an error listing its problems.
func TestFR010NoResultOnError(t *testing.T) {
	for _, c := range evalCases {
		if c.wantErr == nil {
			continue
		}
		t.Run(c.fr+"/"+c.input, func(t *testing.T) {
			got, err := Evaluate(c.input)
			if got != "" {
				t.Errorf("Evaluate(%q) returned result %q; want none", c.input, got)
			}
			if !slices.Equal(problems(err), c.wantErr) {
				t.Errorf("Evaluate(%q) problems = %q; want %q", c.input, problems(err), c.wantErr)
			}
		})
	}
}

// T005, NFR-001: all examples together are answered correctly in under 1 second.
func TestNFR001Speed(t *testing.T) {
	start := time.Now()
	wrong := 0
	for _, c := range evalCases {
		if _, _, ok := check(c); !ok {
			wrong++
		}
	}
	if elapsed := time.Since(start); elapsed >= time.Second {
		t.Errorf("%d examples took %v; want under 1s", len(evalCases), elapsed)
	}
	if wrong > 0 {
		t.Errorf("%d of %d examples gave a wrong answer", wrong, len(evalCases))
	}
}
