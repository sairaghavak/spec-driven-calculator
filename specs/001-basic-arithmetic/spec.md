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
> **In scope**: exactly **two operands and one operator** per calculation, for example `2 * 10`.
>
> **Out of scope**: expressions with more than two operands, such as `2 * 10 + 15 - 50`. They are
> rejected with the expected form `<number> <operator> <number>` (FR-001, FR-005). Also out of
> scope: parentheses, operator precedence, memory, and history.

## Functional Requirements

- **FR-001**: The calculator MUST accept one operation at a time in the form
  `<number> <operator> <number>`, where the operator is one of `+`, `-`, `*`, `/`. Extra spaces
  around numbers and operators MUST be ignored.
- **FR-002**: The calculator MUST compute addition, subtraction, multiplication, and division, and
  show the result.
- **FR-003**: Numbers MUST be written in plain decimal notation: an optional leading `-`, digits,
  and an optional decimal point followed by digits (for example `7`, `-3`, `2.5`). Any other form
  MUST be rejected as not a number.
- **FR-004**: Results MUST be shown in plain decimal notation, without trailing zeros or a trailing
  decimal point (`6`, not `6.0`), and never as `-0`.
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
- **FR-009**: A result that cannot be represented as a finite number MUST be reported as an error,
  not shown.
- **FR-010**: Whenever input is rejected (FR-005 to FR-009), no result MUST be shown for that
  operation.
- **FR-011**: When the calculator is started with an operation, it MUST run in one-shot mode:
  compute that one operation, show the result or the error, then exit, signaling success or failure
  to the caller so a script can tell them apart.
- **FR-012**: When the calculator is started without an operation, it MUST run an interactive
  session: show a prompt, read one operation per line, show the result or the error, and prompt
  again. A rejected operation MUST NOT end the session.
- **FR-013**: The interactive session MUST end when the user enters `quit` or `exit`, or when input
  ends (for example Ctrl+D). Blank lines MUST be ignored (prompt again with no error).

## Non-Functional Requirements

- **NFR-001**: Each result or error appears in under 1 second.
- **NFR-002**: A user who enters bad input can correct it on the next attempt using only the error
  message, without reading other documentation (checked by reviewing every error message against
  the content rules in FR-005 to FR-008).

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

**Covers**: FR-005, FR-006, FR-007, FR-009, FR-010, NFR-002

**Acceptance Criteria** (Given the calculator is available):

1. When the user enters `abc + 3`, then it reports that `abc` is not a number, and shows no result.
2. When the user enters `4 % 2`, then it reports that `%` is not a supported operator and lists
   `+`, `-`, `*`, `/`.
3. When the user enters `4 +`, then it reports that a number is missing and shows the expected form
   `<number> <operator> <number>`.
4. When the user enters `8 / 0`, then it reports that division by zero is not allowed, and shows no
   result.
5. When the user adds a number written as `1` followed by 308 zeros to itself, then it reports that
   the result is too large to represent, and shows no result.
6. When the user enters `2 * 10 + 15 - 50`, then it reports that only one operation with two
   numbers is accepted, shows the expected form `<number> <operator> <number>`, and shows no
   result.

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

1. Given the calculator was started with `6 * 7`, then it shows `42` and exits, signaling success.
2. Given the calculator was started with `8 / 0`, then it shows the division-by-zero message and
   exits, signaling failure.
3. Given an interactive session, when the user enters `abc + 3`, then it shows the not-a-number
   message and prompts again.
4. Given an interactive session, when the user enters `quit`, `exit`, or ends input, then the
   session ends.
5. Given an interactive session, when the user enters a blank line, then it prompts again with no
   error.

## Edge Cases

- `199.5 * 2`: valid, shows `399` (FR-008).
- `-500 * 2` and `-200 * 2`: rejected; `-199 * 199` is valid and shows `-39601` (FR-008).
- `0 / 0`: reported as division by zero, the same as `8 / 0` (FR-007).
- `0 / 5`: valid, shows `0` (FR-002).
- `0 * -5`: shows `0`, not `-0` (FR-004).
- `NaN`, `Inf`, `1e3`, `0x10`, `1,000`: rejected as not a number (FR-003, FR-006).
- `  2   +  3  `: valid, shows `5` (FR-001).

## Key Entities

- **Operation**: two operands and one operator, entered by the user.
- **Result**: the number produced by a valid operation.
- **Error message**: the text shown for rejected input; states the problem, the offending value
  where there is one, and what is accepted.

## Assumptions

- "Keep things default" is read as: the four operators `+ - * /`, decimal numbers allowed, results
  shown in plain decimal form, and no other limits beyond the multiplication rule.
- The multiplication rule is taken from the author's example, not the phrase "up to 200": 199 is
  accepted and 200 is rejected. It applies to each operand separately, not to the result, and to
  negative operands too (author's answer: limit on size, so -199 is accepted and -200 rejected).
- Both modes come from the author's answer ("both"): one-shot when an operation is given at
  start, interactive otherwise. How a one-shot operation is passed on the command line (and how
  shells treat the `*` character) is a plan-phase decision.
- The limit applies to multiplication only; addition, subtraction, and division have no operand
  limit beyond what can be represented.
- Error messages are in English.

## Definition of Done

- Every acceptance criterion and edge case passes as an automated test that cites its FR or NFR.
- Every rejected input produces a message that names the problem; no input produces a crash, a
  blank response, or a wrong result.
