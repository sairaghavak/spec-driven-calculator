# spec-driven-calculator

A public demo of spec-driven development using GitHub Spec Kit and Claude Code, with a Go calculator as the example domain. The repo doubles as tech-talk material: the git history, tags, and `docs/journey.md` must tell the story on their own.

## Workflow
- Spec Kit phases, in order: constitution (once per project), then per feature: specify, clarify, plan, checklist, tasks, analyze, implement, converge. No phase is skipped.
- Each phase is reviewed by the author, then committed and tagged before the next starts. That commit is the approval.
- Write Go code (tests included) only to complete a task in a committed `specs/NNN-*/tasks.md`, and reference the task ID (e.g. `T012`) in the commit.
- When the output is wrong, fix the feature's `spec.md` or `plan.md` first, rerun the phases after it, then fix code. Record the correction in the journey log.
- A requirement change starts in `specs/NNN-*/spec.md` (edit the existing feature, or `/speckit-specify` a new one), then flows through the later phases into code. Never change code first.
- Spec Kit is pinned: `specify` 1.4.0, initialized with `--integration claude --extension git`. Upgrading is its own step with its own entry.

## Journey log (`docs/journey.md`)
- Append-only. Every step adds one entry, committed in the same commit as the work it describes.
- One job per place, no fact written in full twice: the commit message is the subject plus one line of why; the tag message is the step's one-line summary; the journey entry holds the detail.
- Entry format:
  ```
  ## Step NN · <Phase> · tag: step-NN-<slug>
  - Goal: what this step sets out to produce, in one line
  - Command/prompt: the input that started the step, verbatim: the Spec Kit command with its arguments, or the author's initial prompt. Follow-ups are summarized under Human review.
  - AI produced: files and versions
  - Human review: what was changed and why (write "accepted as-is" if nothing)
  - Lesson: the reusable takeaway for a human reader (a developer or talk attendee) to apply in their own project, in one line (not a summary of this step). A lesson that should change Claude's behavior also becomes a rule in this file.
  ```

## Commits, branches, tags
- Conventional Commits: `docs:`, `docs(spec):`, `docs(plan):`, `docs(tasks):`, `test:`, `feat:`, `chore:`, `refactor:`.
- One commit per phase, except implement, which is test-first and produces two: a red commit (failing tests) then a green commit (passing). So each feature has 9 commits and 9 tags.
- Every step gets an annotated tag `step-NN-<slug>` (two-digit, monotonic across features) whose message is the step's one-line summary.
- Features live on Spec Kit branches (`NNN-<name>`) and merge to `main` with `--no-ff`. The merge commit is not a phase, so it gets no tag.
- Never rewrite pushed history. Never commit, tag, or push without approval.

## Go
- Go 1.27, module `github.com/sairaghavak/spec-driven-calculator`, stdlib only unless a spec justifies a dependency.
- `gofmt`, `go vet`, and `go test ./...` must pass before any green commit.
- Every test case cites the requirement it verifies (e.g. `// FR-001`).
