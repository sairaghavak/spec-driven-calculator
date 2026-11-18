<!--
Project rules for Claude Code.
What this file means for Claude: Claude Code loads it into context at the start of every session
in this repo, so every line is an instruction Claude follows. It covers how Claude works here:
the Spec Kit workflow, the journey log, and commits, tags, branches, and PRs.
The product's engineering principles are in .specify/memory/constitution.md, imported at the end.
-->
# spec-driven-calculator

A public demo of spec-driven development using GitHub Spec Kit and Claude Code, with a Go calculator as the example domain. The repo doubles as tech-talk material: the git history, tags, and `docs/journey.md` must tell the story on their own.

## Workflow
- Spec Kit phases, in order: constitution (once per project), then per feature: specify, clarify, plan, checklist, tasks, analyze, implement, converge. No phase is skipped.
- Each phase is reviewed by the author, then committed and tagged before the next starts. That commit is the approval.
- Write Go code (tests included) only to complete a task in a committed `specs/NNN-*/tasks.md`, and reference the task ID (e.g. `T012`) in the commit.
- When the output is wrong, fix the feature's `spec.md` or `plan.md` first, rerun the phases after it, then fix code. Record the correction in the journey log.
- A requirement change starts in `specs/NNN-*/spec.md` (edit the existing feature, or `/speckit-specify` a new one), then flows through the later phases into code. Never change code first.
- Spec Kit is pinned: `specify` 1.4.0, initialized with `--integration claude` and no extensions. Upgrading is its own step with its own entry.

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
  - Takeaway for readers: advice a developer or talk attendee can apply in their own project (not a summary of this step). If it should also change Claude's behavior, add it as a rule in this file.
  ```

## Commits, branches, tags
- Conventional Commits: `docs:`, `docs(spec):`, `docs(plan):`, `docs(tasks):`, `test:`, `feat:`, `chore:`, `refactor:`.
- One commit per phase, except implement, which is test-first and produces two: a red commit (failing tests) then a green commit (passing). So each feature has 9 commits and 9 tags.
- Every step gets an annotated tag `step-NN-<slug>` (two-digit, monotonic across features) whose message is the step's one-line summary.
- `main` changes only through pull requests: one branch and one PR per unit of work. Setup steps (00 to 02) live on `000-project-setup`; each feature gets a branch named after its spec folder (`git checkout -b 001-basic-arithmetic` for `specs/001-basic-arithmetic/`).
- Merge PRs with GitHub's "Create a merge commit" only (the same as `git merge --no-ff`). Never squash or rebase: both leave the step tags pointing at commits that are not on `main`, and rebase drops the author's signatures. The merge commit is not a phase, so it gets no tag.
- Never rewrite pushed history. Never commit, tag, or push without approval.

## Engineering principles
- The product's principles (spec-first traceability, test-first, Go stdlib only, explicit errors) and the Go quality gate are in the constitution, imported below so they apply in all work, not only inside Spec Kit phases.

@.specify/memory/constitution.md
