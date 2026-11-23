# Research: Basic Arithmetic

The [spec](spec.md) says *what* the calculator does. This file records *how* it is built: the
choices made while planning, R1 to R6 (R = research decision).

## Why these decisions were needed

None of them is a gap in the spec. A spec describes behavior, not how to build it, so every
feature needs decisions like these. They came from two sources: items the spec left to the plan,
and what the code needs in order to meet the spec and the
[constitution](../../.specify/memory/constitution.md).

| Decision | What caused it |
|---|---|
| R2, R3, R4 | The spec's [Assumptions](spec.md#assumptions) left them to the plan on purpose: how an operation is typed in one-shot mode, where output goes, and the exit codes. |
| R1 | To meet [FR-003](spec.md#functional-requirements) and [FR-004](spec.md#functional-requirements): `0.1 + 0.2` must show exactly `0.3`, and `1e3` must be rejected. |
| R5 | To meet [FR-005](spec.md#functional-requirements): each input mistake needs its own message. |
| R6 | The [constitution](../../.specify/memory/constitution.md)'s test-first principle and the spec's [Definition of Done](spec.md#definition-of-done): every spec example becomes an automated test. |

Planning did find two real gaps (how a tie like `0.000625` rounds, and the order of several
errors). Those are behavior, so they were fixed in the spec
([FR-004](spec.md#functional-requirements), [FR-015](spec.md#functional-requirements)), not here.

## Format

Each decision uses Spec Kit's three fields:

- **Decision**: what was chosen.
- **Rationale**: why it was chosen, shown with what a user or a script sees.
- **Alternatives considered**: the main other option, and why it lost.

## R1 · Exact results

- **Decision**: store and compute numbers as exact fractions, using Go's built-in `math/big.Rat`
  type (`0.1` is kept as exactly 1/10). Round to 5 decimal places only when showing the result
  ([FR-004](spec.md#functional-requirements)). Check each number's format against
  [FR-003](spec.md#functional-requirements) before computing.
- **Rationale**: results are exact: `0.1 + 0.2` shows `0.3`, as
  [FR-004](spec.md#functional-requirements) requires. The format check is needed because
  `math/big.Rat` also accepts `1e3` (meaning 1000), and
  [FR-003](spec.md#functional-requirements) says `1e3` is not a number.
- **Alternatives considered**: `float64`, Go's normal decimal type (Java's `double`). We did not
  choose it. It stores numbers in binary, where 0.1 never ends, just like 1/3 = 0.333... never
  ends in decimal. If we had used it, `0.1 + 0.2` would show `0.30000000000000004`, not `0.3`
  (checked by running it).

## R2 · Typing an operation in one-shot mode

- **Decision**: the calculator joins everything typed after `calc` into one operation. No quotes
  are needed, with or without spaces. The one exception is multiplication with spaces, which
  needs quotes (see [US3](spec.md#us3--multiplication-operand-limit-p3)).
- **Rationale**: the shell splits `calc 2 + 3` at the spaces and passes three pieces (`2`, `+`,
  `3`); joining them back gives `2 + 3`. So users can type naturally:

  | You type | Result |
  |---|---|
  | `calc 2 + 3` or `calc 2+3` | `5` |
  | `calc 9 / 2` or `calc 9/2` | `4.5` |
  | `calc 6*7` or `calc '6 * 7'` | `42` |
  | `calc 6 * 7` | error: the shell turns the lone `*` into file names before the calculator starts |

- **Alternatives considered**: make the calculator use only the first piece and ignore the rest.
  We did not choose this. If we had, `calc 2 + 3` would fail: the shell passes `2`, `+`, and
  `3`, but the calculator would use only `2`. Users would then need quotes around every
  operation with spaces, such as `calc '2 + 3'`.

## R3 · Where output goes in one-shot mode

- **Decision**: the result goes to normal output, errors to error output.
- **Rationale**: a script can capture the result alone. `x=$(calc 6*7)` stores `42` in `x`.
  `x=$(calc 6/0)` stores nothing, and the error still shows on screen.
- **Alternatives considered**: send errors to normal output too. We did not choose this. If we
  had, `x=$(calc 6/0)` would store the error text in `x`, and a script could mistake it for a
  result.

## R4 · How a script knows whether the operation worked

- **Decision**: exit code 0 when a result is shown, 1 when errors are shown. An exit code is the
  number every program hands back when it finishes; scripts read it, not the printed text
  ([FR-011](spec.md#functional-requirements)).
- **Rationale**: it is the standard every shell understands. `calc 6*7 && echo ok` prints `42`
  then `ok` (`&&` runs the next command only on exit code 0). `calc 8/0 || echo failed` prints the
  error then `failed` (`||` runs it only on a non-zero exit code). The code is not printed output:
  read it from `$?` right after the command, so `calc 6/0; echo $?` prints the error, then `1`.
  The interactive session ends with 0, or with 1 if reading the input itself fails (for example,
  the input stream breaks); that failure is printed to error output
  ([FR-016](spec.md#functional-requirements)).
- **Alternatives considered**: a different failure code for each error, for example 2 for a
  missing number and 3 for division by zero. We did not choose this. There is no need: it is just
  success or failure, 0 for success and 1 for failure. The error message says what went wrong.

## R5 · Finding mistakes in the input

- **Decision**: read the input piece by piece: number, operator, number. A `-` written directly
  before a number or word, at the start or right after the operator, is its sign
  ([FR-001](spec.md#functional-requirements)). When the pieces are not number, operator, number,
  the spec's [Which shape message is shown](spec.md#which-shape-message-is-shown) table decides
  the message.
- **Rationale**: the calculator can say exactly what is wrong
  ([FR-005](spec.md#functional-requirements)). For example, `-abc + 3` is reported as
  `"-abc" is not a number`, not as a vague input-shape problem.
- **Alternatives considered**: check the whole line against one pattern (a "regular expression")
  that describes a valid line: number, operator, number. We did not choose this. If we had,
  `4 +`, `4 % 2`, and `2 * 10 + 15` would all just fail to match. The calculator would know the
  line is invalid, but not which part is wrong (a missing number, an unknown operator, or too many
  numbers), so it could not give the specific message
  [FR-005](spec.md#functional-requirements) requires.

## R6 · Testing

- **Decision**: tests run the calculator's code directly with typed input, and check its output
  and exit code.
- **Rationale**: it is fast and needs no setup. `go test ./...` checks every spec example in both
  modes.
- **Alternatives considered**: build and start the `calc` program in each test. We did not choose
  this. If we had, the tests would be slower and would check nothing more.
