# Implementation Plan: Remainder Operator

**Branch**: `002-modulus` | **Date**: 2026-11-25 | **Spec**: [spec.md](spec.md)

**Input**: Feature specification from [`specs/002-modulus/spec.md`](spec.md)

## Summary

`%` is added to the existing calculator from feature 001; no new file or package. The input is
split at `%` like at the other operators, numbers written with a decimal point are rejected for
`%`, and the remainder is taken after dropping both minus signs, so `-7 % 3` shows `1`. The welcome
message and the input-shape messages list `%`. The choices behind this are in
[research.md](research.md); the exact texts are in the spec
([Error messages](spec.md#error-messages),
[Texts that change in feature 001](spec.md#texts-that-change-in-feature-001)).

## Technical Context

- **Language/Version**: Go 1.27 (set by the [constitution](../../.specify/memory/constitution.md)).
- **Primary Dependencies**: none. Only Go's standard library, as in feature 001.
- **Storage**: N/A. Nothing is saved; each operation stands alone.
- **Testing**: Go's built-in test tool (`go test ./...`). Feature 002's tests name their requirement
  as `002/FR-001` ([R5](research.md#r5--naming-the-requirement-in-tests-now-that-two-features-share-ids)).
- **Target Platform**: any system Go runs on. The demo runs in bash on Linux, where `%` needs no
  quotes in one-shot mode.
- **Project Type**: a single command-line program, unchanged.
- **Performance Goals**: each answer in under 1 second (feature 001's
  [NFR-001](../001-basic-arithmetic/spec.md#non-functional-requirements)). Numbers stay at most 3
  digits, so a remainder takes far less than that.
- **Constraints**: output text exactly as the spec states; errors returned as values, never a crash
  (constitution, Principle IV).
- **Scale/Scope**: one user, one operation at a time, whole numbers up to `999` for `%`.

## Constitution Check

*Checked at the start and again at the end of `/speckit-plan`, to catch a design choice that breaks
a principle.*

| Principle | Result |
|---|---|
| I. Spec-first traceability | Pass. Every behavior of `%`, including each message and its order, is in [FR-001](spec.md#functional-requirements). Planning changed no behavior, so the spec did not change. |
| II. Test-first | Pass. Failing tests are committed first (red), then the code that passes them (green). Each test names its requirement; feature 002's use `002/FR-001` so they cannot be confused with feature 001's FR-001 ([R5](research.md#r5--naming-the-requirement-in-tests-now-that-two-features-share-ids)). |
| III. Simplicity | Pass. Standard library only; the same package and files; no new type, interface, or setting. |
| IV. Explicit errors | Pass. The whole-number problem is one more message in the `Errors:` block; nothing crashes. |
| Quality gate | Pass. `gofmt`, `go vet`, and `go test ./...` must pass before the green commit. |

## Project Structure

### Documentation (this feature)

```text
specs/002-modulus/
├── spec.md              # What % does, including exact messages and the feature 001 texts it replaces
├── plan.md              # This file
├── research.md          # How it is built: decisions R1 to R6
├── checklists/
│   └── requirements.md  # Spec quality checklist
└── tasks.md             # Created next, by /speckit-tasks
```

Spec Kit's plan step also asks for `data-model.md`, `contracts/`, and `quickstart.md`. This
project skips them (see the project's [`CLAUDE.md`](../../CLAUDE.md)): `%` adds no data, and its
exact messages are in the spec.

### Source Code (repository root)

The same files as feature 001; the comments say what changes in each.

```text
cmd/calc/
├── calc.go              # Split at %; whole-number check; division by zero also for %; remainder; % in the operator list
├── cli.go               # Welcome message: % in the operator line
├── calc_test.go         # New cases for each 002 example; 001 cases get the new texts, and 4 % 2 becomes 4 ^ 2
└── cli_test.go          # New one-shot and welcome-line cases for 002; 001's expected welcome gets the new line
```

**Structure Decision**: no new file or package. `%` is one more operator in code that already
handles four, and each change sits next to the matching rule for the other operators: the split in
[`calc.go`](../../cmd/calc/calc.go) ([R2](research.md#r2--finding--in-the-input)), the checks in their existing order
([R4](research.md#r4--where-the-new-check-goes)), and the computation
([R1](research.md#r1--computing-the-remainder)). A separate file for `%` was considered and not
chosen: it would split one set of rules across two places, and no requirement needs it
(constitution, Principle III).

## Complexity Tracking

None. The design follows every principle, so there is nothing to justify.
