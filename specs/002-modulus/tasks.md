# Tasks: Remainder Operator

Each user story lists its tasks, as in JIRA. This feature has one story.

**Where these tasks come from**

| Source | Gives the tasks |
|---|---|
| [spec.md](spec.md) | **What**: the story, the test examples, the exact messages, the feature 001 texts that change |
| [plan.md](plan.md) | **Where**: the four existing files in `cmd/calc/` |
| [research.md](research.md) | **How**: the six design decisions, [R1](research.md#r1--computing-the-remainder) to [R6](research.md#r6--feature-001-tests-whose-text-changes) |
| [constitution](../../.specify/memory/constitution.md) | **Order**: tests first, then code |
| [`CLAUDE.md`](../../CLAUDE.md) | **Layout**: two commits, the task format |

**How to read a task**

- `[ ]` becomes `[x]` when done.
- `T005` is the task ID. Commits mention it, with the feature number in the commit subject.
- `[US1]` is the user story the task belongs to.

**Two commits**

- **Red commit**: every "Tests" task. The tests build against today's code and fail, because `%` is not an operator yet and the texts are the old ones. No empty functions are needed.
- Every new test case names `002/FR-001`, so `go test -v` shows it. Feature 001's test cases keep their short names, such as `FR-005`.
- **Green commit**: every "Code" task and the final checks. All tests pass, feature 001's included.

All code goes in `cmd/calc/`.

**Summary**: 1 user story, 12 tasks (4 red, 8 green)

| Story | Red commit | Green commit | Tasks |
|---|---|---|---|
| [US1](spec.md#us1--remainder-of-two-whole-numbers-p1) · Remainder of two whole numbers | 4 | 8 | 12 |
| **Total** | **4** | **8** | **12** |

## [US1](spec.md#us1--remainder-of-two-whole-numbers-p1) · Remainder of two whole numbers

Done when the tests show `-7 % 3` gives `1`, `5 % 0.000` gives its three messages in order, and every feature 001 test passes with the new texts.

**Tests (red commit)**

- [ ] T001 [US1] - Update feature 001's expected texts in `cmd/calc/calc_test.go`
  - `expectedForm` ends `where <operator> is one of +, -, *, /, %`
  - The unknown-operator example `4 % 2` becomes `4 ^ 2`, expecting the message that names `^`
  - These cases keep their feature 001 names, such as `FR-005`
- [ ] T002 [US1] - Update feature 001's expected welcome in `cmd/calc/cli_test.go`
  - Third line of `welcome`: `Operations: +, -, *, /, %`
  - Every session test uses it and keeps its feature 001 name
- [ ] T003 [US1] - Write the remainder tests in `cmd/calc/calc_test.go`
  - Add a helper for the whole-number message, with the exact text from the spec's "Error messages" table
  - Add examples to `TestEvaluate`, each named `002/FR-001`
  - Results: `7 % 3` gives `1`; `4 % 2` gives `0`; `-7 % 3`, `7 % -3`, `-7 % -3` give `1`; `999 % 250` gives `249`
  - Whole-number message: `7.5 % 10` and `10 % 7.5` (naming `7.5`), `7.0 % 2` (naming `7.0`)
  - Division by zero: `5 % 0`, `5 % -0`
  - Three messages in order: `5 % 0.000`
  - Unknown operator: `4 ^ 2`
  - Every edge case in the spec's "Edge Cases" section, except `1.5 + 2.25`, which feature 001's tests already check (13 of 14)
- [ ] T004 [US1] - Write the mode tests in `cmd/calc/cli_test.go`
  - Session started with no input: the output is the welcome, whose third line is `Operations: +, -, *, /, %`
  - One-shot `calc 7 % 3`: prints `1`, exit code 0
  - One-shot `calc 7.5 % 10`: the whole-number `Errors:` block on the error output only, exit code 1
  - Each case named `002/FR-001`

**Code (green commit)**

- [ ] T005 [US1] - Split the input at `%` in `cmd/calc/calc.go`
  - `%` is an operator, like `+ - * /`, with or without spaces around it
  - Update the comment that lists the operators
- [ ] T006 [US1] - Add the whole-number check in `cmd/calc/calc.go`
  - Only for `%`
  - A number whose text has a decimal point is not whole; check the text, not the value
  - Runs after the digit limit and before division by zero; left number before right
  - Checks every number with the right format, including one over the digit limit
  - Use the exact message from the spec
- [ ] T007 [US1] - Check division by zero for `%` in `cmd/calc/calc.go`
  - Same check and message as for `/`
  - Also when the right number failed the whole-number check, as in `5 % 0.000`
- [ ] T008 [US1] - Compute the remainder in `cmd/calc/calc.go`
  - Drop both minus signs, then take the remainder (`big.Int.Rem`)
  - Only whole numbers, with a right number other than 0, reach this step
- [ ] T009 [US1] - Add `%` to the expected-form hint in `cmd/calc/calc.go`
  - The hint ends `where <operator> is one of +, -, *, /, %`
  - All six input-shape messages use it
- [ ] T010 [US1] - Add `%` to the welcome in `cmd/calc/cli.go`
  - Third line: `Operations: +, -, *, /, %`
  - The other five lines stay the same
- [ ] T011 [US1] - Run the quality checks
  - `gofmt -l .` prints nothing
  - `go vet ./...` passes
  - `go test ./...` passes, feature 001's tests included
- [ ] T012 [US1] - Try it by hand
  - Build: `go build -o bin/calc ./cmd/calc`
  - Run: `bin/calc 7 % 3` prints `1`; `bin/calc -7 % 3` prints `1`; `bin/calc 7.5 % 10` prints the whole-number message, exit code 1
  - Run one session: the welcome lists `%`; enter `4 ^ 2`, then `quit`
