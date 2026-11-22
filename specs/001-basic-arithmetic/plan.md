# Implementation Plan: Basic Arithmetic

**Branch**: `001-basic-arithmetic` | **Date**: 2026-11-22 | **Spec**: [spec.md](spec.md)

**Input**: Feature specification from [`specs/001-basic-arithmetic/spec.md`](spec.md)

## Summary

A command-line calculator, `calc`, written in Go. It runs one operation and exits (one-shot mode),
or keeps asking for operations (interactive session). Numbers are checked against the spec's
format, then computed exactly, so `0.1 + 0.2` shows `0.3`. The choices behind this are in
[research.md](research.md); the exact output and error messages are in the spec
([FR-012](spec.md#functional-requirements), [Error messages](spec.md#error-messages)).

## Technical Context

- **Language/Version**: Go 1.27 (set by the [constitution](../../.specify/memory/constitution.md)).
- **Primary Dependencies**: none. Only Go's standard library, as the constitution requires.
- **Storage**: N/A. Nothing is saved; each operation stands alone.
- **Testing**: Go's built-in test tool (`go test ./...`). Each test names the requirement it checks,
  for example [FR-008](spec.md#functional-requirements) ([R6](research.md#r6--testing)).
- **Target Platform**: any system Go runs on. The demo runs in bash on Linux.
- **Project Type**: a single command-line program.
- **Performance Goals**: each answer in under 1 second
  ([NFR-001](spec.md#non-functional-requirements)). One operation takes far less than that.
- **Constraints**: output text exactly as the spec states
  ([FR-012](spec.md#functional-requirements), [Error messages](spec.md#error-messages));
  errors returned as values, never a crash (constitution, Principle IV).
- **Scale/Scope**: one user, one operation at a time, numbers up to `999.99`
  ([FR-014](spec.md#functional-requirements)).

## Constitution Check

*Checked at the start and again at the end of `/speckit-plan`, to catch a design choice that breaks
a principle.*

| Principle | Result |
|---|---|
| I. Spec-first traceability | Pass. Every behavior, including each exact error message, is in the spec. The two gaps found while planning were fixed there ([FR-004](spec.md#functional-requirements), [FR-015](spec.md#functional-requirements)). |
| II. Test-first | Pass. Failing tests are committed first (red), then the code that passes them (green). |
| III. Simplicity | Pass. Standard library only; one package; no interfaces, flags, or settings. |
| IV. Explicit errors | Pass. Mistakes come back as error values and are printed as the `Errors:` block; nothing crashes. |
| Quality gate | Pass. `gofmt`, `go vet`, and `go test ./...` must pass before the green commit. |

## Project Structure

### Documentation (this feature)

```text
specs/001-basic-arithmetic/
├── spec.md              # What the calculator does, including exact output and error messages
├── plan.md              # This file
├── research.md          # How it is built: decisions R1 to R6
├── checklists/
│   └── requirements.md  # Spec quality checklist
└── tasks.md             # Created next, by /speckit-tasks
```

Spec Kit's plan step also asks for `data-model.md`, `contracts/`, and `quickstart.md`. This
project skips them (see the project's [`CLAUDE.md`](../../CLAUDE.md)), because each would repeat another file: the entities and exact messages are in
the spec, and the input-check rules are in [R5](research.md#r5--finding-mistakes-in-the-input).
Spec Kit's later steps read those files only if they exist.

### Source Code (repository root)

```text
go.mod                   # Go module definition: github.com/sairaghavak/spec-driven-calculator
cmd/calc/
├── main.go              # Starts the program: passes the typed arguments and the terminal to run()
├── cli.go               # run(): chooses one-shot or interactive mode, prints output and errors
├── calc.go              # Evaluate(): checks the input, computes, and formats the result
├── cli_test.go          # Tests for both modes, output, and exit codes
└── calc_test.go         # One test per spec example, each naming its requirement
```

**Structure Decision**: one Go package in `cmd/calc/`, so building it produces a program named
`calc`. The code is split by file, not by package: `calc.go` only computes (no printing or reading
input), and `cli.go` does all the input and output. A separate package for the calculation was
considered and not chosen: no requirement needs it (constitution, Principle III). Built programs go
to `bin/`, which is kept out of git.

## Complexity Tracking

None. The design follows every principle, so there is nothing to justify.
