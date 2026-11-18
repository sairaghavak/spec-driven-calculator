<!--
Constitution: the engineering principles every version of this calculator must satisfy.
What this file means for Claude: Claude Code loads it in every session through the
import line at the end of CLAUDE.md, so these principles apply to all work. Every Spec Kit
skill (/speckit-specify through /speckit-converge) also reads it during its phase and checks
its output against it. Amend it only with /speckit-constitution, through a pull request,
with a semantic version bump.
-->
# Spec-Driven Calculator Constitution

## Core Principles

### I. Spec-First Traceability

Every behavior of the calculator **MUST** trace to a functional requirement ID (`FR-001`,
`FR-002`, ...) in a committed `specs/NNN-*/spec.md`. Behavior with no requirement behind it is a
defect, even if it works, and is removed or specified before release.

Rationale: the spec is the source of truth; untraced behavior cannot be reviewed or tested against
anything.

### II. Test-First

Tests **MUST** be written from the spec and seen failing before the implementation exists: a red
commit (failing tests) precedes the green commit (passing implementation). Every test case
**MUST** cite the requirement it verifies (e.g. `// FR-001`).

Rationale: a test that never failed proves nothing, and the citation keeps tests traceable to the
spec.

### III. Simplicity

Code **MUST** use only the Go standard library unless a spec explicitly justifies a dependency. No
abstraction (interface, package, configuration option) is added without a current requirement that
needs it.

Rationale: the calculator is a teaching example; every extra layer is something a reader must
understand before seeing the point.

### IV. Explicit Errors

Invalid input and edge cases (for example division by zero) **MUST** return Go `error` values with
behavior defined in the spec. Code **MUST NOT** panic for expected failures.

Rationale: errors as values make failure paths visible, testable, and part of the contract.

## Additional Constraints

- Language and module: Go 1.27, module `github.com/sairaghavak/spec-driven-calculator`.
- Quality gate: `gofmt`, `go vet`, and `go test ./...` **MUST** pass before any green commit.

## Development Workflow

- Spec Kit phases run in order: specify, clarify, plan, checklist, tasks, analyze, implement,
  converge.
- The author reviews each phase's output before the next phase starts; the commit of that output
  is the approval.

## Governance

- This constitution takes precedence over other practices in this repo. Runtime guidance for
  Claude Code lives in `CLAUDE.md` and **MUST NOT** contradict it.
- Amendments in this file **MUST** be made with `/speckit-constitution` and merged through a pull
  request.
- Versioning follows semantic versioning: MAJOR for removed or redefined principles, MINOR for
  added principles or materially expanded guidance, PATCH for wording fixes.
- Compliance is checked in `/speckit-analyze` for every feature; any violation is fixed or
  justified in the plan before implementation.

**Version**: 1.0.0 | **Ratified**: 2026-11-18 | **Last Amended**: 2026-11-18
