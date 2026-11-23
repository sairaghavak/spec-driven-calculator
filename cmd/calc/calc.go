package main

import (
	"math/big"
	"regexp"
	"strings"
)

// Message texts from the spec's "Error messages" table (FR-005 to FR-008, FR-014).
const formHint = "expected <number> <operator> <number>, where <operator> is one of +, -, *, /"

const (
	errNothingEntered = "no operation was entered; " + formHint
	errTooManyNumbers = "only one operation with two numbers is accepted; " + formHint
	errMissingNumber  = "a number is missing; " + formHint
	errMissingOp      = "the operator is missing; " + formHint
	errMisplacedOp    = "the operator must be between the two numbers; " + formHint
	errDivideByZero   = "division by zero is not allowed"
	numberExamples    = "numbers are written like 7, -3, or 2.5"
)

// InputError is returned for every rejected input. It holds each problem found, in the
// order the spec requires (FR-015).
type InputError struct {
	Problems []string
}

// Error lists the problems one per line.
func (e *InputError) Error() string {
	return strings.Join(e.Problems, "\n")
}

// numberFormat is the FR-003 format: an optional -, digits, and an optional . with digits.
var numberFormat = regexp.MustCompile(`^-?[0-9]+(\.[0-9]+)?$`)

// piece is one part of the input: a word (a candidate number) or an operator.
type piece struct {
	text       string
	isOperator bool
	start, end int // byte positions in the input, to tell whether two pieces touch
}

// Evaluate computes one operation of the form <number> <operator> <number>.
// It returns the formatted result, or an *InputError that lists every problem found.
func Evaluate(input string) (string, error) {
	pieces := split(input)
	if p := shapeProblem(pieces); p != "" {
		return "", &InputError{Problems: []string{p}}
	}
	left, op, right := pieces[0].text, pieces[1].text, pieces[2].text

	values, problems := checkNumbers(left, op, right)
	if len(problems) > 0 {
		return "", &InputError{Problems: problems}
	}
	return format(compute(values[0], op, values[1])), nil
}

// split breaks the input into words and operators (T006, FR-001). Spaces and tabs separate
// pieces; + - * / are operators. A - written directly before a word, at the start or right
// after an operator, is that word's sign.
func split(input string) []piece {
	var raw []piece
	wordStart := -1
	endWord := func(i int) {
		if wordStart >= 0 {
			raw = append(raw, piece{text: input[wordStart:i], start: wordStart, end: i})
			wordStart = -1
		}
	}
	for i := 0; i < len(input); i++ {
		switch c := input[i]; c {
		case ' ', '\t':
			endWord(i)
		case '+', '-', '*', '/':
			endWord(i)
			raw = append(raw, piece{text: string(c), isOperator: true, start: i, end: i + 1})
		default:
			if wordStart < 0 {
				wordStart = i
			}
		}
	}
	endWord(len(input))

	var pieces []piece
	for i := 0; i < len(raw); i++ {
		p := raw[i]
		isSign := p.text == "-" && i+1 < len(raw) && !raw[i+1].isOperator &&
			raw[i+1].start == p.end && (len(pieces) == 0 || pieces[len(pieces)-1].isOperator)
		if isSign {
			next := raw[i+1]
			pieces = append(pieces, piece{text: "-" + next.text, start: p.start, end: next.end})
			i++
			continue
		}
		pieces = append(pieces, p)
	}
	return pieces
}

// shapeProblem returns the one message for a badly shaped input, or "" if the shape is
// <number> <operator> <number> (T012, FR-005). The order of checks is the spec's
// "Which shape message is shown" table.
func shapeProblem(pieces []piece) string {
	if len(pieces) == 3 && !pieces[0].isOperator && pieces[1].isOperator && !pieces[2].isOperator {
		return ""
	}
	words, operators := 0, 0
	for _, p := range pieces {
		if p.isOperator {
			operators++
		} else {
			words++
		}
	}
	switch {
	case len(pieces) == 0:
		return errNothingEntered
	case words == 3 && operators == 0 && !strings.ContainsAny(pieces[1].text, "0123456789"):
		return `"` + pieces[1].text + `" is not a supported operator; ` + formHint
	case words >= 3:
		return errTooManyNumbers
	case words < 2:
		return errMissingNumber
	case operators == 0:
		return errMissingOp
	default:
		return errMisplacedOp
	}
}

// checkNumbers checks both numbers by kind, in the FR-015 order, left before right within a
// kind (T007, T013, T015). It returns their values if every check passed.
func checkNumbers(left, op, right string) ([2]*big.Rat, []string) {
	texts := [2]string{left, right}
	var values [2]*big.Rat
	var problems []string

	// 1. Not a number, or leading zeros (FR-003, FR-006). Such a number is not checked further.
	for i, t := range texts {
		switch {
		case !numberFormat.MatchString(t):
			problems = append(problems, `"`+t+`" is not a number; `+numberExamples)
		case hasLeadingZeros(t):
			problems = append(problems, `"`+t+`" has leading zeros, which are not allowed; `+numberExamples)
		default:
			values[i], _ = new(big.Rat).SetString(t) // safe: t matches numberFormat
		}
	}
	// 2. Digit limits (FR-014).
	for i, t := range texts {
		if values[i] != nil && !withinDigitLimits(t) {
			problems = append(problems, `"`+t+`" has too many digits; numbers may have at most 3 digits before and 2 digits after the decimal point (largest 999.99)`)
		}
	}
	// 3. Multiplication limit (FR-008): also for numbers that broke the digit limit.
	if op == "*" {
		limit := big.NewRat(200, 1)
		for i, t := range texts {
			if values[i] != nil && new(big.Rat).Abs(values[i]).Cmp(limit) >= 0 {
				problems = append(problems, `"`+t+`" is out of range for multiplication; each number must be strictly between -200 and 200`)
			}
		}
	}
	// 4. Division by zero (FR-007).
	if op == "/" && values[1] != nil && values[1].Sign() == 0 {
		problems = append(problems, errDivideByZero)
	}
	return values, problems
}

// hasLeadingZeros reports whether the part before the decimal point has more than one digit
// and starts with 0, as in 007 or 00.5 (FR-003).
func hasLeadingZeros(number string) bool {
	whole, _, _ := strings.Cut(strings.TrimPrefix(number, "-"), ".")
	return len(whole) > 1 && whole[0] == '0'
}

// withinDigitLimits reports whether a number has at most 3 digits before the decimal point
// and at most 2 after (FR-014). The sign is not a digit.
func withinDigitLimits(number string) bool {
	whole, fraction, _ := strings.Cut(strings.TrimPrefix(number, "-"), ".")
	return len(whole) <= 3 && len(fraction) <= 2
}

// compute applies the operator exactly (T008, FR-002). Division by zero never reaches here.
func compute(a *big.Rat, op string, b *big.Rat) *big.Rat {
	r := new(big.Rat)
	switch op {
	case "+":
		r.Add(a, b)
	case "-":
		r.Sub(a, b)
	case "*":
		r.Mul(a, b)
	default:
		r.Quo(a, b)
	}
	return r
}

// format shows a result as FR-004 requires: rounded to 5 decimal places with halves away from
// zero, without trailing zeros or a trailing point, and never as -0.
func format(r *big.Rat) string {
	s := strings.TrimRight(r.FloatString(5), "0")
	s = strings.TrimSuffix(s, ".")
	if s == "-0" {
		return "0"
	}
	return s
}
