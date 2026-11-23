<!--
Project rules for Claude Code.
What this file means for Claude: Claude Code loads it into context at the start of every session
in this repo, so every line is an instruction Claude follows. It covers how Claude works here:
the Spec Kit workflow, the journey log, and commits, tags, branches, and PRs.
The product's engineering principles are in .specify/memory/constitution.md, imported at the end.
-->
# spec-driven-calculator

A public demo of spec-driven development using GitHub Spec Kit and Claude Code, with a Go calculator as the example domain. The repo doubles as tech-talk material: the git history, tags, and [`docs/journey.md`](docs/journey.md) must tell the story on their own.

## Workflow
- Spec Kit phases, in order: constitution (once per project), then per feature: specify, clarify, plan, checklist, tasks, analyze, implement, converge. No phase is skipped.
- Each phase is reviewed by the author, then committed and tagged before the next starts. That commit is the approval.
- Write Go code (tests included) only to complete a task in a committed `specs/NNN-*/tasks.md`, and reference the task ID (e.g. `T012`) in the commit.
- Project state lives only in the repo: step tags, `specs/`, and the journey log. At session start, find the current step from the latest `step-NN` tag and `git status`; open decisions belong in the feature's `spec.md` or `plan.md`. Claude's memory holds only machine-local notes, never anything needed to continue the work.
- When the output is wrong, fix the feature's `spec.md` or `plan.md` first, rerun the phases after it, then fix code. Record the correction in the journey log.
- When a plan or research decision changes what a user or a script sees (output, exit codes, error handling), add or extend a functional requirement in `spec.md` in the same step. `/speckit-analyze` then checks both directions: every requirement has a task, and every task and test traces back to a requirement (not only to research).
- A requirement change starts in `specs/NNN-*/spec.md` (edit the existing feature, or `/speckit-specify` a new one), then flows through the later phases into code. Never change code first.
- Spec Kit is pinned: `specify` 1.4.0, initialized with `--integration claude` and no extensions. Upgrading is its own step with its own entry.
- Specs use the project override [`.specify/templates/overrides/spec-template.md`](.specify/templates/overrides/spec-template.md): requirements first, then user stories, in agile terms. Its "Non-Functional Requirements" (`NFR-###`) section plays the role of Spec Kit's "Success Criteria" (`SC-###`); when a `speckit-*` skill reads or writes Success Criteria, use the NFR section. Changing the template or other Spec Kit customization is a process step: its own commit, tag, and journey entry.

## Spec documents (`specs/NNN-*/`)
- Spec Kit's templates and command instructions are a starting point, not a checklist. Don't fill a section or create a file only because the template has it: for each one, decide whether it has a purpose in this feature. If a field does not apply, write "N/A" or "Not required" with a one-line reason (e.g. "Storage: N/A, nothing is saved"); if a whole file or section adds nothing another file does not already cover, skip it and record why in `plan.md` and the journey entry. Never write filler to satisfy a template.
- A feature's documents are `spec.md`, `plan.md`, `research.md`, `tasks.md`, and `checklists/`. When `/speckit-plan` says to generate `data-model.md`, `contracts/`, or `quickstart.md`, skip them (the project override [`.specify/templates/overrides/plan-template.md`](.specify/templates/overrides/plan-template.md) lists only the kept files): exact user-facing text (messages, output layout) goes in `spec.md`, and design rules go in `research.md`. Every file must have a purpose no other file already covers; later Spec Kit steps read the skipped files only if they exist.
- `research.md` uses Spec Kit's fields for each decision: Decision, Rationale, Alternatives considered. Write each alternative as clearly not chosen ("We did not choose this. If we had, ...").
- Every mention of an ID is a link to where it is defined, in every file under `specs/`, `spec.md` included, and every time it appears (a second mention in the same sentence is linked too). The only bare ID is the definition itself: the bold `**FR-001**:` label or the `### US1` heading. Targets, with examples from feature 001: `FR-###` and `NFR-###` link to the spec's requirement headings ([FR-004](specs/001-basic-arithmetic/spec.md#functional-requirements), [NFR-001](specs/001-basic-arithmetic/spec.md#non-functional-requirements)), `US#` to the story heading ([US3](specs/001-basic-arithmetic/spec.md#us3--multiplication-operand-limit-p3)), `R#` to the research decision ([R2](specs/001-basic-arithmetic/research.md#r2--typing-an-operation-in-one-shot-mode)), and `T###` to the task. Write the path relative to the file that holds the link: in `plan.md` the target is `spec.md#functional-requirements`; inside `spec.md` it is just `#functional-requirements`; in `checklists/requirements.md` it is `../spec.md#functional-requirements`. Anchors are GitHub's heading slugs. Before handing a document over, search it for bare IDs.
- Write for readers who do not know Go: each rule or decision states what the user or a script sees, with a concrete input and output. Go names appear only in parentheses. No filler: no preamble, no restating the spec.
- Write every sentence in this project for the author to read, not only for the AI. The AI can work with dense or jargon-heavy text; the author must be able to read each line once and understand it. Plain words, short sentences, no unexplained jargon.
- `tasks.md` layout: each task is one line `- [ ] T006 [US1] - Short title with its file`, followed by short sub-bullets, one fact each. Every task belongs to a user story: shared setup goes in the first story, final checks in the last. Each story lists its tasks under "Tests (red commit)" and "Code (green commit)". Don't repeat FR or research references in tasks: the spec already maps each story to its FRs. Keep Spec Kit's `- [ ] T0NN` and `[US#]` parts exactly; skip the optional `[P]` marker.
- `tasks.md` opens its task list with a summary: the number of user stories and tasks, and a table of tasks per story split into red and green commits. Count them from the task lines with a script, and recount whenever tasks change.
- Acceptance criteria name their mode (one-shot or interactive session) whenever the output differs between modes.

## File links in Markdown
- In every Markdown file this project owns, each mention of a file that exists in the repo is a link to it, with the path relative to the file that holds the link. For example, this file links [`docs/journey.md`](docs/journey.md), while `plan.md` links the same kind of root file as `../../CLAUDE.md`. When the name is in backticks, keep the backticks inside the link text, as in [`docs/journey.md`](docs/journey.md). Never put a link inside backticks or a code block: it shows as plain text, not a link. Before handing a document over, search it for unlinked file names.
- Not linked: names inside code blocks; patterns and placeholders (`specs/NNN-*/spec.md`); generic names meaning "this file in any feature"; files not created yet or deleted (link them once they exist); the `@` import line below; journey entries already committed (append-only); Spec Kit's own files, including the templates in `.specify/templates/` (their text is copied into feature folders, where relative links would break); and the constitution, which changes only through `/speckit-constitution`.

## Journey log ([`docs/journey.md`](docs/journey.md))
- Append-only. Every step adds one entry, committed in the same commit as the work it describes.
- Write the entry when the step starts, not at commit time: Goal, Command/prompt, and AI produced first, then add each Human review round as soon as it closes. The takeaway is written last, when the step closes.
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
- One commit per phase, except implement, which is test-first and produces two: a red commit (failing tests) then a green commit (passing). The red commit holds the tests plus only what they need to run and fail on a wrong answer (the module file and empty versions of the functions they call); everything else waits for the green commit. So each feature has 9 phase commits and 9 tags, plus one per process step (see Workflow).
- Implement commits (red and green) name the user stories they cover in the subject, and list their task IDs grouped by story on a `Tasks:` line after the one line of why. Example:
  ```
  test: add failing tests for US1 to US4 of feature 001

  Tests come first: they turn the spec into checks that must fail before any code exists.

  Tasks: US1 T001-T005, US2 T009-T010, US3 T014, US4 T016-T018
  ```
- Every step gets an annotated tag `step-NN-<slug>` (two-digit, monotonic across features) whose message is the step's one-line summary.
- `main` changes only through pull requests: one branch and one PR per unit of work. Setup steps (00 to 02) live on `000-project-setup`; each feature gets a branch named after its spec folder (`git checkout -b 001-basic-arithmetic` for `specs/001-basic-arithmetic/`).
- Merge PRs with GitHub's "Create a merge commit" only (the same as `git merge --no-ff`). Never squash or rebase: both leave the step tags pointing at commits that are not on `main`, and rebase drops the author's signatures. The merge commit is not a phase, so it gets no tag.
- Never rewrite pushed history. Never commit, tag, or push without approval.

## Engineering principles
- The product's principles (spec-first traceability, test-first, Go stdlib only, explicit errors) and the Go quality gate are in the [constitution](.specify/memory/constitution.md), imported below so they apply in all work, not only inside Spec Kit phases.

@.specify/memory/constitution.md
