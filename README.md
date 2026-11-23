# spec-driven-calculator

A small Go calculator, built step by step by [sairaghavak](https://github.com/sairaghavak) with
[GitHub Spec Kit](https://github.com/github/spec-kit) and Claude Code, to show spec-driven development from the first requirement to tested code.

The calculator is simple on purpose. The point is the process: every behavior starts as a written
requirement, becomes a test, and only then becomes code.

## Try it

**Prerequisites**

- [Go](https://go.dev/dl/) 1.27 or later. Check with `go version`.
- [Git](https://git-scm.com/downloads), to get the code.

```bash
git clone https://github.com/sairaghavak/spec-driven-calculator.git
cd spec-driven-calculator
go build -o bin/calc ./cmd/calc

bin/calc 2 + 3          # 5
bin/calc 0.1 + 0.2      # 0.3, exactly
bin/calc '6 * 7'        # 42 (quote * when it has spaces around it)
bin/calc 8/0            # Errors: division by zero is not allowed (exit code 1)
bin/calc                # interactive session; type quit to leave
```

Run the tests: `go test ./...` (136 cases, each naming the requirement it checks).

## How to read this repo

- [`docs/journey.md`](docs/journey.md): every step, with the command, what the AI produced, how [sairaghavak](https://github.com/sairaghavak) validated, reviewed, and constantly challenged the AI-generated specs, what changed as a result, and a takeaway.
- [`specs/001-basic-arithmetic/`](specs/001-basic-arithmetic/): the feature's documents, one or more per Spec Kit phase. [The steps](#the-steps) shows which step changed which file.
- [`.specify/memory/constitution.md`](.specify/memory/constitution.md): the project's rules: spec first, tests first, standard library only, errors as values.
- [`CLAUDE.md`](CLAUDE.md): how [sairaghavak](https://github.com/sairaghavak) instructed the AI to work in this repo.

Every step has a git tag. To see the repo exactly as it was after a step:

```bash
git tag -l 'step-*'
git checkout step-16-red    # the failing tests, before any calculator code
```

## The steps

In order. Each step links to its journey entry, and each commit link shows exactly what changed.
Rule steps turned [sairaghavak](https://github.com/sairaghavak)'s review feedback into written rules for the AI.
Short file names are in [`specs/001-basic-arithmetic/`](specs/001-basic-arithmetic/) (documents) or [`cmd/calc/`](cmd/calc/) (code).
Every step also adds its own journey entry, so the journey is not listed.

- **[Step 00](docs/journey.md#step-00--project-rules--tag-step-00-project-rules) · Project rules**
  - Changed in commit [766386e](https://github.com/sairaghavak/spec-driven-calculator/commit/766386ebdc800b53c55c01e48c4e2659e6e88b2e):
    - `CLAUDE.md`
- **[Step 01](docs/journey.md#step-01--spec-kit-init--tag-step-01-speckit-init) · Spec Kit setup** (`specify init`)
  - Changed in commit [2147cbc](https://github.com/sairaghavak/spec-driven-calculator/commit/2147cbcbb09614fbc8d249a173dc7f7021a992f9):
    - Spec Kit's own files: skills, scripts, templates (29 files)
    - `.gitignore`
    - `CLAUDE.md`
- **[Step 02](docs/journey.md#step-02--constitution--tag-step-02-constitution) · Constitution** (`/speckit-constitution`)
  - Changed in commit [a4cc03a](https://github.com/sairaghavak/spec-driven-calculator/commit/a4cc03ad13bb4d0dcf1701471691bdf4fd802860):
    - `.specify/memory/constitution.md`
    - `CLAUDE.md`
- **[Step 03](docs/journey.md#step-03--specify--tag-step-03-specify) · Specify** (`/speckit-specify`)
  - Changed in commit [9e41877](https://github.com/sairaghavak/spec-driven-calculator/commit/9e41877fcbc48348ec6912cb7f40d2f3ecd46a20):
    - `checklists/requirements.md`
    - `spec.md`
- **[Step 04](docs/journey.md#step-04--spec-template--tag-step-04-spec-template) · Rule: requirements-first spec template**
  - Changed in commit [ab37ac6](https://github.com/sairaghavak/spec-driven-calculator/commit/ab37ac65f597dc754678db5634e4a6f6ecd8ddf5):
    - `.specify/templates/overrides/spec-template.md`
    - `CLAUDE.md`
    - `checklists/requirements.md`
    - `spec.md`
- **[Step 05](docs/journey.md#step-05--clarify--tag-step-05-clarify) · Clarify** (`/speckit-clarify`)
  - Changed in commit [65f7f49](https://github.com/sairaghavak/spec-driven-calculator/commit/65f7f49f6926897a79e255f32e884acbe0ffe49f):
    - `spec.md`
- **[Step 06](docs/journey.md#step-06--spec-amendment--tag-step-06-leading-zeros) · Spec amendment: leading zeros**
  - Changed in commit [9f2679a](https://github.com/sairaghavak/spec-driven-calculator/commit/9f2679ad77c0ae860160533c00bcd460278d8704):
    - `spec.md`
- **[Step 07](docs/journey.md#step-07--process-rule--tag-step-07-repo-state) · Rule: project state lives in the repo**
  - Changed in commit [385a9d3](https://github.com/sairaghavak/spec-driven-calculator/commit/385a9d343a0943bef02b701d03db44d91f0ba59b):
    - `CLAUDE.md`
- **[Step 08](docs/journey.md#step-08--plan--tag-step-08-plan) · Plan** (`/speckit-plan`)
  - Changed in commit [392aaed](https://github.com/sairaghavak/spec-driven-calculator/commit/392aaed84d088c2093834b9df6ecc71e9c29e35d):
    - `checklists/requirements.md`
    - `plan.md`
    - `research.md`
    - `spec.md`
- **[Step 09](docs/journey.md#step-09--process-rules--tag-step-09-doc-rules) · Rules: spec documents, links, journey**
  - Changed in commit [6ed77cd](https://github.com/sairaghavak/spec-driven-calculator/commit/6ed77cdbb3442e561eaa6d001410828c8a62a94a):
    - `.specify/templates/overrides/plan-template.md`
    - `CLAUDE.md`
- **[Step 10](docs/journey.md#step-10--checklist--tag-step-10-checklist) · Checklist** (`/speckit-checklist`)
  - Changed in commit [95759c5](https://github.com/sairaghavak/spec-driven-calculator/commit/95759c5620b78542188db26e24aa14b1722364d1):
    - `checklists/input-output.md`
    - `research.md`
    - `spec.md`
- **[Step 11](docs/journey.md#step-11--tasks--tag-step-11-tasks) · Tasks** (`/speckit-tasks`)
  - Changed in commit [21b7c19](https://github.com/sairaghavak/spec-driven-calculator/commit/21b7c196a7641be650da9c4c537171938f0ade94):
    - `spec.md`
    - `tasks.md`
- **[Step 12](docs/journey.md#step-12--process-rules--tag-step-12-task-rules) · Rules: readable tasks, minimal red commit**
  - Changed in commit [081d4d4](https://github.com/sairaghavak/spec-driven-calculator/commit/081d4d46c42d0285c577f53afbbe35735a0af0f8):
    - `CLAUDE.md`
- **[Step 13](docs/journey.md#step-13--analyze--tag-step-13-analyze) · Analyze** (`/speckit-analyze`)
  - Changed in commit [63dcd3e](https://github.com/sairaghavak/spec-driven-calculator/commit/63dcd3e85f660e3549e404566b1cca28a0a4faca):
    - `tasks.md`
- **[Step 14](docs/journey.md#step-14--tasks-amendment--tag-step-14-tasks-fix) · Tasks amendment: one test task**
  - Changed in commit [884df83](https://github.com/sairaghavak/spec-driven-calculator/commit/884df8397e7b438680e186dcb4b7982d0d042866):
    - `tasks.md`
- **[Step 15](docs/journey.md#step-15--process-rule--tag-step-15-commit-stories) · Rule: user stories in implement commits**
  - Changed in commit [e538db8](https://github.com/sairaghavak/spec-driven-calculator/commit/e538db83a54a04c9dd69d2868def938bcf055dec):
    - `CLAUDE.md`
- **[Step 16](docs/journey.md#step-16--implement-red--tag-step-16-red) · Implement, red: failing tests** (`/speckit-implement`)
  - Changed in commit [91c5565](https://github.com/sairaghavak/spec-driven-calculator/commit/91c55651f6978939a3e6df25697be598c7e29c85):
    - `calc.go`
    - `calc_test.go`
    - `cli.go`
    - `cli_test.go`
    - `go.mod`
    - `tasks.md`
- **[Step 17](docs/journey.md#step-17--implement-green--tag-step-17-green) · Implement, green: code passes** (`/speckit-implement`)
  - Changed in commit [7309143](https://github.com/sairaghavak/spec-driven-calculator/commit/7309143e7569392d5af81c2270e3b51d636ea9fc):
    - `.gitignore`
    - `calc.go`
    - `cli.go`
    - `main.go`
    - `tasks.md`
- **[Step 18](docs/journey.md#step-18--converge--tag-step-18-converge) · Converge: one gap found** (`/speckit-converge`)
  - Changed in commit [aa93697](https://github.com/sairaghavak/spec-driven-calculator/commit/aa936971b87fdfa8d777742730b41c1851331a38):
    - `tasks.md`
- **[Step 19](docs/journey.md#step-19--process-rule--tag-step-19-trace-both-ways) · Rule: trace both ways**
  - Changed in commit [ce79caf](https://github.com/sairaghavak/spec-driven-calculator/commit/ce79cafdc297101b792f2773e68cbd880ae455b9):
    - `CLAUDE.md`
- **[Step 20](docs/journey.md#step-20--implement-fr-016--tag-step-20-fr016) · Implement: requirement for a failed read** (`/speckit-implement`)
  - Changed in commit [1751296](https://github.com/sairaghavak/spec-driven-calculator/commit/17512966a72e717cd4f25c0e06f828022b026c94):
    - `cli.go`
    - `cli_test.go`
    - `research.md`
    - `spec.md`
    - `tasks.md`
- **[Step 21](docs/journey.md#step-21--converge-second-run--tag-step-21-converge) · Converge: one more gap found** (`/speckit-converge`)
  - Changed in commit [15971cd](https://github.com/sairaghavak/spec-driven-calculator/commit/15971cde25bd98d1842f4767f6d3f59453094ddc):
    - `tasks.md`
- **[Step 22](docs/journey.md#step-22--implement-output-and-exit-codes--tag-step-22-streams) · Implement: output and exit codes** (`/speckit-implement`)
  - Changed in commit [69960c7](https://github.com/sairaghavak/spec-driven-calculator/commit/69960c70a6d4cc9b3e2f098a6967bbcff941662f):
    - `spec.md`
    - `tasks.md`
- **[Step 23](docs/journey.md#step-23--converge-third-run--tag-step-23-converged) · Converge: clean** (`/speckit-converge`)
  - Changed in commit [95d215f](https://github.com/sairaghavak/spec-driven-calculator/commit/95d215fddc92fe83dc93a0c8d78a74d11f50dcc8):
    - Only its journey entry

## License

[MIT](LICENSE)
