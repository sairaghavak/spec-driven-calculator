# Tasks: Basic Arithmetic

Each user story lists its tasks, as in JIRA.

**Where these tasks come from**

| Source | Gives the tasks |
|---|---|
| [spec.md](spec.md) | **What**: the stories, the test examples, the exact output |
| [plan.md](plan.md) | **Where**: the files and the Go version |
| [research.md](research.md) | **How**: the six design decisions, [R1](research.md#r1--exact-results) to [R6](research.md#r6--testing) |
| [constitution](../../.specify/memory/constitution.md) | **Order**: tests first, then code |
| [`CLAUDE.md`](../../CLAUDE.md) | **Layout**: two commits, the task format |

**How to read a task**

- `[ ]` becomes `[x]` when done.
- `T006` is the task ID. Commits mention it.
- `[US1]` is the user story the task belongs to.

**Two commits**

- **Red commit**: every "Tests" task, plus the minimum setup the tests need to run. The tests fail, because the code is not written yet.
- Every test case carries the ID of the requirement it checks (for example `FR-004`), in the test code. The ID is part of the test's name, so `go test -v` shows it. The constitution requires this.
- **Green commit**: every "Code" task and the final checks. The tests pass.

All code goes in `cmd/calc/`.

**Summary**: 4 user stories, 25 tasks (11 red, 13 green, 1 convergence)

| Story | Red commit | Green commit | Convergence | Tasks |
|---|---|---|---|---|
| [US1](spec.md#us1--compute-a-basic-operation-p1) · Compute a basic operation | 5 | 3 | 0 | 8 |
| [US2](spec.md#us2--get-an-informative-message-for-bad-input-p2) · Bad-input messages | 2 | 3 | 0 | 5 |
| [US3](spec.md#us3--multiplication-operand-limit-p3) · Multiplication limit | 1 | 1 | 0 | 2 |
| [US4](spec.md#us4--one-shot-and-interactive-modes-p4) · One-shot and session modes | 3 | 6 | 1 | 10 |
| **Total** | **11** | **13** | **1** | **25** |

## [US1](spec.md#us1--compute-a-basic-operation-p1) · Compute a basic operation

Done when the tests show `2 + 3` gives `5` and `0.1 + 0.2` gives `0.3`.

**Setup and tests (red commit)**

- [x] T001 [US1] - Create `go.mod`
  - Module `github.com/sairaghavak/spec-driven-calculator`
  - Go version `1.27`
  - Needed: without it, no test runs
- [x] T002 [US1] - Create an empty calculator in `cmd/calc/calc.go`
  - `func Evaluate(input string) (string, error)`
  - Returns no answer and a "not implemented" error
  - Needed: without it, the tests do not build
- [x] T003 [US1] - Create an empty command-line runner in `cmd/calc/cli.go`
  - `func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int`
  - Prints nothing, returns `1`
  - Needed: the one-shot and session tests call it
- [x] T004 [US1] - Write the compute tests in `cmd/calc/calc_test.go`
  - One test, `TestEvaluate`, runs a list of examples
  - Each example has: the input, the expected answer or errors, and the requirement it checks
  - Examples: every US1 acceptance criterion
  - Examples: spaces, tabs, and signs (`2+3`, `-3-2`, `5*-2`, `5--3`)
  - Examples: rounding (`0.1 + 0.2`, `1 / 3`, `0.01 / 16`, `0 * -5`)
  - Examples: the other result edge cases in the spec
- [x] T005 [US1] - Write the speed test in `cmd/calc/calc_test.go`
  - Runs every `TestEvaluate` example
  - Fails if any answer is wrong
  - Fails if all of them together take 1 second or more

**Code (green commit)**

- [x] T006 [US1] - Split the input into numbers and operators in `cmd/calc/calc.go`
  - Spaces and tabs separate the pieces
  - `+ - * /` are operators
  - A `-` right before a number or word, at the start or after the operator, is its sign
- [x] T007 [US1] - Read each number in `cmd/calc/calc.go`
  - Accept only the number format in the spec
  - Store it as an exact fraction (`math/big.Rat`)
- [x] T008 [US1] - Compute and show the result in `cmd/calc/calc.go`
  - Add, subtract, multiply, or divide exactly
  - Round to 5 decimal places; a 5 rounds away from zero
  - Drop trailing zeros and a trailing `.`
  - Show `-0` as `0`

## [US2](spec.md#us2--get-an-informative-message-for-bad-input-p2) · Informative message for bad input

Done when the tests show every bad input gets the spec's exact messages, in order.

**Tests (red commit)**

- [x] T009 [US2] - Write the error tests in `cmd/calc/calc_test.go`
  - Add examples to `TestEvaluate`, each with its exact expected messages
  - Every US2 acceptance criterion
  - Every row of the spec's "Which shape message is shown" table
  - Not numbers: `abc`, `1e3`, `.5`, `-abc`, `"5"`
  - Leading zeros: `007`, `00.5`, `0999`
  - Too many digits: `0.001`, `1000`
  - Division by zero: `5 / 0`, `5 / -0`, `5 / 0.000`
  - Several errors at once: `abc / 0`, `1000 * abc`
  - Every other error edge case in the spec
- [x] T010 [US2] - Write the no-answer test in `cmd/calc/calc_test.go`
  - Every bad input returns no answer
  - And an error that lists its problems
  - Checking both makes it fail before the code exists

**Code (green commit)**

- [x] T011 [US2] - Add an error that holds a list of problems in `cmd/calc/calc.go`
  - `type InputError struct{ Problems []string }`
  - Returned for every bad input; the program never crashes
- [x] T012 [US2] - Choose the message for a badly shaped input in `cmd/calc/calc.go`
  - Follow the spec's "Which shape message is shown" table
  - Exactly one message
- [x] T013 [US2] - Check each number in `cmd/calc/calc.go`
  - Check by kind, in this order; within each kind, left number before right
  - First: is it a number, without leading zeros? If not, stop checking it
  - Then: at most 3 digits before and 2 after the point
  - Then: division by zero
  - Use the exact messages from the spec

## [US3](spec.md#us3--multiplication-operand-limit-p3) · Multiplication limit

Done when the tests show `199 * 199` gives `39601` and `200 * 200` gives two limit messages.

**Tests (red commit)**

- [x] T014 [US3] - Write the multiplication tests in `cmd/calc/calc_test.go`
  - Add examples to `TestEvaluate`
  - Every US3 acceptance criterion
  - Edge cases: `199.5 * 2`, `-200 * 2`, `-199 * 199`, `1000 * 5`

**Code (green commit)**

- [x] T015 [US3] - Add the multiplication limit in `cmd/calc/calc.go`
  - Only for `*`
  - Each number must be strictly between -200 and 200
  - Checked after the digit limit, even for numbers that failed it

## [US4](spec.md#us4--one-shot-and-interactive-modes-p4) · One-shot and interactive modes

Done when `calc 6*7` prints `42`, and a session with `2+3` shows `Result: 5`.

**Tests (red commit)**

- [x] T016 [US4] - Write the one-shot tests in `cmd/calc/cli_test.go`
  - `calc 2 + 3` and `calc 2+3`: print `5`, exit code 0
  - `calc '6 * 7'`: prints `42`
  - `calc 8/0`: `Errors:` block on the error output only, exit code 1
  - `calc ""`: the nothing-entered message
  - `calc quit`: the missing-number message
- [x] T017 [US4] - Write the session tests in `cmd/calc/cli_test.go`
  - Compare the whole screen output, exactly
  - The six-line welcome
  - `2+3`: `Result: 5`, the `-` line, `Please enter your input.`
  - `abc + 3`: the `Errors:` block, the `-` line, the prompt
  - `6 * 7`: `Result: 42`
  - `quit`, `exit`, `  quit  `, or end of input: ends quietly, exit code 0
  - Blank lines and lines of only spaces: nothing printed
- [x] T018 [US4] - Write the read-failure test in `cmd/calc/cli_test.go`
  - Reading the keyboard input fails
  - The error goes to the error output, exit code 1

**Code and final checks (green commit)**

- [x] T019 [US4] - Add one-shot mode in `cmd/calc/cli.go`
  - Join the typed pieces with spaces; compute once
  - Answer: print it alone, exit code 0
  - Errors: print the `Errors:` block to the error output, exit code 1
- [x] T020 [US4] - Add the session in `cmd/calc/cli.go`
  - Used when nothing is typed after `calc`
  - Print the welcome exactly as in the spec
  - Read line by line; skip blank lines
  - Stop on `quit`, `exit`, or end of input
  - Otherwise print `Result: <answer>` or the `Errors:` block, the `-` line, the prompt
- [x] T021 [US4] - Create the program entry point in `cmd/calc/main.go`
  - Calls `run` with the real arguments, keyboard, and screen
  - Needed to build the `calc` program; the tests do not use it
- [x] T022 [US4] - Keep built programs out of git, in `.gitignore`
  - Add `/bin/`, before the first build
- [x] T023 [US4] - Run the quality checks
  - `gofmt -l .` prints nothing
  - `go vet ./...` passes
  - `go test ./...` passes
- [x] T024 [US4] - Try it by hand
  - Build: `go build -o bin/calc ./cmd/calc`
  - Run: `bin/calc 2 + 3`, `bin/calc '6 * 7'`, `bin/calc 8/0`
  - Run one session, end it with `quit`

## Convergence

Added by `/speckit-converge`, which checks the finished code against the spec. Each task below is something the code does not fully meet yet.

**Spec and test fix (its own commit)**

- [x] T025 [US4] - Add functional requirement [FR-016](spec.md#functional-requirements) to [`spec.md`](spec.md): what happens when reading the input fails
  - The behavior: the error goes to the error output, exit code 1
  - Today only research [R4](research.md#r4--how-a-script-knows-whether-the-operation-worked) says it
  - Why: the constitution says every behavior and every test must trace to a requirement
  - Then put `FR-016` on `TestInteractiveReadFailure` in `cmd/calc/cli_test.go`

## Convergence, second run

Added by `/speckit-converge` on its second run, which also checks that every behavior traces back to a requirement. Each task below is something the code does, that the spec does not state yet.

**Spec fix (its own commit)**

- [ ] T026 [US4] - Extend functional requirements [FR-011](spec.md#functional-requirements) and [FR-012](spec.md#functional-requirements) in [`spec.md`](spec.md): where output goes and the exit codes
  - One-shot: the typed pieces are joined with spaces ([R2](research.md#r2--typing-an-operation-in-one-shot-mode))
  - One-shot: the result goes to normal output, the `Errors:` block to error output ([R3](research.md#r3--where-output-goes-in-one-shot-mode))
  - One-shot: exit code 0 for a result, 1 for errors ([R4](research.md#r4--how-a-script-knows-whether-the-operation-worked))
  - Session: everything goes to normal output; it ends with exit code 0
  - Today only research says these; the code and tests already do them
  - Why: the rule from Step 19, behavior a script can see needs a requirement
