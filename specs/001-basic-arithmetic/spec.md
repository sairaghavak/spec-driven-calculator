# Feature Specification: Basic Arithmetic

**Feature Branch**: `001-basic-arithmetic`

**Created**: 2026-11-18

**Status**: Draft

**Input**: User description: "Add feature to do basic math operations like addition, subtraction, multiplication, divison in the CLI and prompt the user whenever there is bad input or errors while computing the math operations, keep things default for now and for multiplication accept the numbers only upto 200 like 199 * 199 is valid 200 * 200 is bad input that should respond with proper user informative message"

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Compute a basic operation (Priority: P1)

A user enters two numbers and one of the four operators (addition, subtraction, multiplication,
division) at the command line and sees the result.

**Why this priority**: computing a correct result is the reason the calculator exists. Without it,
nothing else in this feature has value.

**Independent Test**: enter `2 + 3`, `10 - 4`, `6 * 7`, and `9 / 2` and confirm the results are
`5`, `6`, `42`, and `4.5`.

**Acceptance Scenarios**:

1. **Given** the calculator is available, **When** the user enters `2 + 3`, **Then** it shows `5`.
2. **Given** the calculator is available, **When** the user enters `3 - 10`, **Then** it shows `-7`.
3. **Given** the calculator is available, **When** the user enters `6 * 7`, **Then** it shows `42`.
4. **Given** the calculator is available, **When** the user enters `9 / 2`, **Then** it shows `4.5`.
5. **Given** the calculator is available, **When** the user enters `1.5 + 2.25`, **Then** it shows
   `3.75`.

---

### User Story 2 - Get an informative message for bad input (Priority: P2)

A user who makes a mistake (types a word instead of a number, uses an unknown operator, leaves out
a number, or divides by zero) gets a message that says what was wrong and what is accepted, instead
of a wrong result or no response.

**Why this priority**: the author asked explicitly that bad input and computation errors be
reported to the user. It builds on Story 1, which must exist for there to be input to reject.

**Independent Test**: enter each invalid input listed in the scenarios below and confirm each gets
a message naming the problem, and that no result is shown.

**Acceptance Scenarios**:

1. **Given** the calculator is available, **When** the user enters `abc + 3`, **Then** it reports
   that `abc` is not a number, and shows no result.
2. **Given** the calculator is available, **When** the user enters `4 % 2`, **Then** it reports
   that `%` is not a supported operator and lists `+`, `-`, `*`, `/`.
3. **Given** the calculator is available, **When** the user enters `4 +`, **Then** it reports that
   a number is missing and shows the expected form `<number> <operator> <number>`.
4. **Given** the calculator is available, **When** the user enters `8 / 0`, **Then** it reports
   that division by zero is not allowed, and shows no result.
5. **Given** the calculator is available, **When** the user adds a number written as `1` followed by
   308 zeros to itself, **Then** it reports that the result is too large to represent, and shows no
   result.

---

### User Story 3 - Multiplication operand limit (Priority: P3)

Multiplication accepts only operands below 200. A user who multiplies with an operand of 200 or
more gets a message stating the limit and the offending value.

**Why this priority**: it is a product rule the author set for this feature, narrower than the
general operations in Stories 1 and 2.

**Independent Test**: enter `199 * 199` and confirm the result `39601`; enter `200 * 200` and
confirm a message stating that multiplication accepts numbers below 200.

**Acceptance Scenarios**:

1. **Given** the calculator is available, **When** the user enters `199 * 199`, **Then** it shows
   `39601`.
2. **Given** the calculator is available, **When** the user enters `200 * 200`, **Then** it reports
   that multiplication accepts numbers below 200, names `200` as the rejected value, and shows no
   result.
3. **Given** the calculator is available, **When** the user enters `5 * 250`, **Then** it rejects
   `250` with the same message, even though the other operand is valid.
4. **Given** the calculator is available, **When** the user enters `500 + 500`, **Then** it shows
   `1000` (the limit applies to multiplication only).

---

### User Story 4 - Interactive session (Priority: P4)

A user starts the calculator with no operation, enters several operations one after another at a
prompt, sees each result or error, and quits when done.

**Why this priority**: lets the user correct a mistake and try again without restarting. It reuses
everything in Stories 1 to 3; only the way input arrives differs.

**Independent Test**: start the calculator with no operation, enter `2 + 3`, then `abc + 3`, then
`6 * 7`, then `quit`; confirm `5`, a not-a-number message, `42`, and a clean exit.

**Acceptance Scenarios**:

1. **Given** the calculator was started with an operation (`6 * 7`), **When** it runs, **Then** it
   shows `42` and exits, signaling success.
2. **Given** the calculator was started with an invalid operation (`8 / 0`), **When** it runs,
   **Then** it shows the division-by-zero message and exits, signaling failure.
3. **Given** an interactive session, **When** the user enters `abc + 3`, **Then** it shows the
   not-a-number message and prompts again.
4. **Given** an interactive session, **When** the user enters `quit`, `exit`, or ends input,
   **Then** the session ends.
5. **Given** an interactive session, **When** the user enters a blank line, **Then** it prompts
   again with no error.

---

### Edge Cases

- `199.5 * 2`: valid, because 199.5 is below 200; shows `399`.
- `-500 * 2`: rejected, because the limit applies in both directions (FR-008); `-199 * 199` is
  valid and shows `-39601`; `-200 * 2` is rejected.
- `0 / 0`: reported as division by zero, the same as `8 / 0`.
- `0 / 5`: valid, shows `0`.
- `-0` results: shown as `0`, never `-0`.
- Inputs such as `NaN`, `Inf`, `1e3`, `0x10`, or `1,000`: rejected as not a number (only plain
  decimal notation is accepted, see FR-003).
- Results too large to represent (for example adding two numbers near the largest representable
  value): reported as an error, never shown as "infinity" or a wrong number.
- Extra input such as `2 + 3 + 4`: rejected, with the expected form shown (one operation at a time).
- Extra spaces around numbers and operators: ignored.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The calculator MUST accept one operation at a time in the form
  `<number> <operator> <number>`, where the operator is one of `+`, `-`, `*`, `/`.
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

### Key Entities

- **Operation**: two operands and one operator, entered by the user.
- **Result**: the number produced by a valid operation.
- **Error message**: the text shown for rejected input; states the problem, the offending value
  where there is one, and what is accepted.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: 100% of the acceptance scenarios in Stories 1 to 4 produce exactly the stated result
  or message.
- **SC-002**: 100% of rejected inputs produce a message that names the problem; 0 inputs produce a
  crash, a blank response, or a wrong result.
- **SC-003**: A user who enters bad input can correct it on the next attempt using only the error
  message, without reading other documentation (checked by reviewing every error message against
  the content rules in FR-005 to FR-008).
- **SC-004**: Each result or error appears in under 1 second.

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
- One operation at a time; chained expressions, parentheses, operator precedence, memory, and
  history are out of scope.
- Error messages are in English.
