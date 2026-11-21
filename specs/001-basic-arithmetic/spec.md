# Feature Specification: Basic Arithmetic

**Feature Branch**: `001-basic-arithmetic`

**Created**: 2026-11-18

**Status**: Draft

**Input**: User description: "Add feature to do basic math operations like addition, subtraction, multiplication, divison in the CLI and prompt the user whenever there is bad input or errors while computing the math operations, keep things default for now and for multiplication accept the numbers only upto 200 like 199 * 199 is valid 200 * 200 is bad input that should respond with proper user informative message"

> [!NOTE]
> The description above does not limit how many operands a calculation takes. This spec does:
> **exactly two operands** (`2 * 10` is supported, `2 * 10 + 15 - 50` is not). See Scope.

## Summary

A command-line calculator that adds, subtracts, multiplies, and divides two numbers, either once
per run or in an interactive session. Bad input and computation errors get a message that says what
was wrong and what is accepted. Multiplication accepts only operands strictly between -200 and 200.

## Scope

> [!IMPORTANT]
> **In scope**: exactly **two operands and one operator** per calculation, for example `2 * 10`,
> each with **at most 3 digits before and 2 digits after the decimal point** (up to `999.99`).
>
> **Out of scope**: expressions with more than two operands, such as `2 * 10 + 15 - 50`. They are
> rejected with the expected form `<number> <operator> <number>` (FR-001, FR-005). Also out of
> scope: parentheses, operator precedence, memory, and history.

## Clarifications

### Session 2026-11-21

- Q: When a result has a long or inexact decimal expansion, how should it be shown, for example
  `0.1 + 0.2` or `1 / 3`? → A: Round to 5 decimal places.
- Q: Must the user put spaces between the numbers and the operator, or should `2+3` and `-3-2` also
  be accepted? → A: Spaces optional; `2+3`, `-3-2`, and `5*-2` are valid.
- Q: What should happen when the user types a number that is itself too large for the calculator to
  hold? → A: Limit every number to 3 digits before the decimal point (up to 999) and, as the
  author added, 2 digits after it; reject longer ones. Results can then never be too large, so
  FR-009 is removed.
- Q: When one input breaks more than one rule, for example `1000 * 5` or `abc / 0`, which message
  should the user see? → A: All problems at once, under the headline `Errors:`, each problem on
  its own line as a list item.
- Q: In an interactive session, what text should the calculator show when it is waiting for input?
  → A: A welcome message once at start (exact text in FR-012); the user types each operation on its
  own line; the answer (`Result: <value>` or the `Errors:` block) follows on a new line, then a
  separator line and `Please enter your input.` (exact layout in FR-012).

## Functional Requirements

- **FR-001**: The calculator MUST accept one operation at a time in the form
  `<number> <operator> <number>`, where the operator is one of `+`, `-`, `*`, `/`. Spaces around
  numbers and operators are optional and MUST be ignored (`2+3` and `2 + 3` are the same); spaces
  inside a number are not allowed. A `-` written directly before the digits, at the start of the
  input or right after the operator, is the number's sign; any other `-` is the subtraction
  operator.
- **FR-002**: The calculator MUST compute addition, subtraction, multiplication, and division, and
  show the result.
- **FR-003**: Numbers MUST be written in plain decimal notation: an optional leading `-`, digits,
  and an optional decimal point followed by digits (for example `7`, `-3`, `2.5`). Any other form
  MUST be rejected as not a number.
- **FR-004**: Results MUST be rounded to at most 5 decimal places, then shown in plain decimal
  notation without trailing zeros or a trailing decimal point (`6`, not `6.0`; `0.3`, not
  `0.30000`), and never as `-0`.
- **FR-005**: When the input does not match FR-001 (missing number, extra input, unknown operator),
  the calculator MUST report what was wrong and show the expected form
  `<number> <operator> <number>` with the supported operators.
- **FR-006**: When a number is not valid under FR-003, the calculator MUST report the offending
  text and say that it is not a number.
- **FR-007**: Division by zero MUST be rejected with a message stating that division by zero is not
  allowed.
- **FR-008**: For multiplication, each operand MUST be strictly between -200 and 200; an operand
  of 200 or more, or -200 or less, MUST be rejected with a message that states the limit and names
  the offending value.
- **FR-009**: Removed (clarification 2026-11-21): with the 3-digit limit in FR-014, no result can be
  too large to represent.
- **FR-010**: Whenever input is rejected (FR-005 to FR-008, FR-014), no result MUST be shown for
  that operation.
- **FR-011**: When the calculator is started with an operation, it MUST run in one-shot mode:
  compute that one operation, show the result or the error, then exit, signaling success or failure
  to the caller so a script can tell them apart. A result MUST be shown as the bare value with no
  label (`5`, not `Result: 5`), so a script can capture it.
- **FR-012**: When the calculator is started without an operation, it MUST run an interactive
  session. It MUST first show this welcome message, exactly, as six lines (the fifth is 59 `=`
  characters, the length of the second line):
  ```
  Welcome to the interactive calculator session.
  A spec-driven development demo, built with GitHub Spec Kit.
  Operations: +, -, *, /
  Use exactly 2 numbers and 1 operator, for example 2+3
  ===========================================================
  Please enter your input. Example: 10/5
  ```
  Then, for each line the user enters (for example `2+3`), it MUST show on new lines: the answer,
  either `Result: <value>` (value as in FR-004) or the `Errors:` block (FR-015); then a separator
  line of 59 `-` characters (the same width as the `=` line); then `Please enter your input.`; and wait for the next line. A rejected
  operation MUST NOT end the session.
- **FR-013**: The interactive session MUST end when the user enters `quit` or `exit`, or when input
  ends (for example Ctrl+D). Blank lines MUST be ignored (wait for the next line, no error).
- **FR-014**: Each number MUST have at most 3 digits before the decimal point and at most 2 digits
  after it (largest `999.99`). A number with more MUST be rejected with a message that states the
  limit and names the offending value.
- **FR-015**: Rejected input MUST be reported as a headline line `Errors:` followed by one line per
  problem, each written as a list item starting with `- `, even when there is only one. If the input shape is wrong (FR-005), only that problem is reported.
  Otherwise every problem found is reported, left number before right, in this order: not a number
  (FR-006), digit limits (FR-014), multiplication limit (FR-008), division by zero (FR-007). A
  number that is not a number is not checked further.

## Non-Functional Requirements

- **NFR-001**: Each result or error appears in under 1 second.
- **NFR-002**: A user who enters bad input can correct it on the next attempt using only the error
  message, without reading other documentation (checked by reviewing every error message against
  the content rules in FR-005 to FR-008 and FR-014).

## User Stories

### US1 · Compute a basic operation (P1)

As a user, I want to enter two numbers and an operator and see the result, so that I can do quick
arithmetic at the command line.

**Covers**: FR-001, FR-002, FR-003, FR-004, NFR-001

**Acceptance Criteria** (Given the calculator is available):

1. When the user enters `2 + 3`, then it shows `5`.
2. When the user enters `3 - 10`, then it shows `-7`.
3. When the user enters `6 * 7`, then it shows `42`.
4. When the user enters `9 / 2`, then it shows `4.5`.
5. When the user enters `1.5 + 2.25`, then it shows `3.75`.

### US2 · Get an informative message for bad input (P2)

As a user, I want a clear message when my input is wrong or cannot be computed, so that I can fix
it instead of getting a wrong result or no response.

**Covers**: FR-005, FR-006, FR-007, FR-010, FR-014, FR-015, NFR-002

**Acceptance Criteria** (Given the calculator is available):

1. When the user enters `abc + 3`, then it reports that `abc` is not a number, and shows no result.
2. When the user enters `4 % 2`, then it reports that `%` is not a supported operator and lists
   `+`, `-`, `*`, `/`.
3. When the user enters `4 +`, then it reports that a number is missing and shows the expected form
   `<number> <operator> <number>`.
4. When the user enters `8 / 0`, then it reports that division by zero is not allowed, and shows no
   result.
5. When the user enters `1000 + 1`, then it reports that numbers may have at most 3 digits before
   and 2 digits after the decimal point, names `1000` as the rejected value, and shows no result.
6. When the user enters `2 * 10 + 15 - 50`, then it reports that only one operation with two
   numbers is accepted, shows the expected form `<number> <operator> <number>`, and shows no
   result.
7. When the user enters `1000 * 5`, then it shows `Errors:` followed by two list items: the
   digit-limit message for `1000`, then the multiplication-limit message for `1000`.

### US3 · Multiplication operand limit (P3)

As a user, I want multiplication to accept only operands strictly between -200 and 200, and to be
told the limit when I exceed it, so that I know why my input was rejected and what to enter.

**Covers**: FR-008, FR-010

**Acceptance Criteria** (Given the calculator is available):

1. When the user enters `199 * 199`, then it shows `39601`.
2. When the user enters `200 * 200`, then it reports that multiplication accepts numbers strictly
   between -200 and 200, names `200` as the rejected value, and shows no result.
3. When the user enters `5 * 250`, then it rejects `250` with the same message, even though the
   other operand is valid.
4. When the user enters `500 + 500`, then it shows `1000` (the limit applies to multiplication
   only).

### US4 · One-shot and interactive modes (P4)

As a user, I want to run a single operation from the command line, or start a session and enter
several operations in a row, so that I can script the calculator or correct a mistake and try again
without restarting.

**Covers**: FR-011, FR-012, FR-013

**Acceptance Criteria**:

1. Given the calculator was started with no operation, then it shows the six-line welcome message
   from FR-012; when the user then enters `2+3`, it shows, on new lines:
   ```
   Result: 5
   -----------------------------------------------------------
   Please enter your input.
   ```
   and waits for more input.
2. Given the calculator was started with `6 * 7`, then it shows only `42` (no label) and exits,
   signaling success.
3. Given the calculator was started with `8 / 0`, then it shows the division-by-zero message and
   exits, signaling failure.
4. Given an interactive session, when the user enters `abc + 3`, then it shows, on new lines:
   ```
   Errors:
   - <not-a-number message for abc>
   -----------------------------------------------------------
   Please enter your input.
   ```
   and waits for the next line.
5. Given an interactive session, when the user enters `quit`, `exit`, or ends input, then the
   session ends.
6. Given an interactive session, when the user enters a blank line, then it waits for the next line
   with no error.

## Edge Cases

- `199.5 * 2`: valid, shows `399` (FR-008).
- `-500 * 2` and `-200 * 2`: rejected; `-199 * 199` is valid and shows `-39601` (FR-008).
- `abc / 0`: `Errors:` followed by the not-a-number message for `abc`, then the division-by-zero
  message (FR-015).
- `4 + abc 5`: only the input-shape message, because the shape is checked first (FR-015).
- `0 / 0`: reported as division by zero, the same as `8 / 0` (FR-007).
- `0 / 5`: valid, shows `0` (FR-002).
- `0 * -5`: shows `0`, not `-0` (FR-004).
- `0.1 + 0.2`: shows `0.3`; `1 / 3`: shows `0.33333`; `2 / 3`: shows `0.66667` (FR-004).
- `0.01 / 999`: shows `0.00001` (FR-004).
- `NaN`, `Inf`, `1e3`, `0x10`, `1,000`: rejected as not a number (FR-003, FR-006).
- `  2   +  3  ` and `2+3`: valid, both show `5` (FR-001).
- `-3-2`: shows `-5`; `5*-2`: shows `-10`; `5--3`: shows `8` (FR-001).
- `999 + 999`: shows `1998`; `-999.99 - 999.99`: shows `-1999.98`; `999.75 + 0`: valid, shows
  `999.75` (FR-014).
- `0.001 + 1` and `1.125 * 2`: rejected, because they have 3 digits after the decimal point
  (FR-014).
- `2 3 + 4`: rejected, with the expected form shown (FR-001, FR-005).

## Key Entities

- **Operation**: two operands and one operator, entered by the user.
- **Result**: the number produced by a valid operation.
- **Error message**: the text shown for rejected input, one line per problem under `Errors:`;
  states the problem, the offending value where there is one, and what is accepted.

## Assumptions

- "Keep things default" is read as: the four operators `+ - * /`, decimal numbers allowed, results
  shown in plain decimal form, and no other limits beyond the multiplication rule.
- The multiplication rule is taken from the author's example, not the phrase "up to 200": 199 is
  accepted and 200 is rejected. It applies to each operand separately, not to the result, and to
  negative operands too (author's answer: limit on size, so -199 is accepted and -200 rejected).
- Both modes come from the author's answer ("both"): one-shot when an operation is given at
  start, interactive otherwise. How a one-shot operation is passed on the command line (and how
  shells treat the `*` character) is a plan-phase decision.
- The ±200 limit applies to multiplication only; addition, subtraction, and division accept any
  number allowed by the 3-digit limit (FR-014).
- Error messages are in English.

## Definition of Done

- Every acceptance criterion and edge case passes as an automated test that cites its FR or NFR.
- Every rejected input produces a message that names the problem; no input produces a crash, a
  blank response, or a wrong result.
