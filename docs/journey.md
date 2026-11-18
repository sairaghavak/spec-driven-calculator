# Journey log

How this repo was built, one step at a time, with GitHub Spec Kit and Claude Code. Each step has a git tag: run `git checkout step-NN-<slug>` to see the repo exactly as it was after that step.

## Pre-work (no repo changes)
- Refined the author's private, user-level `~/.claude/CLAUDE.md`: removed duplicate rules, added Go defaults, a public-repo guard, and a "review before commit" rule (show the changes and stop; ask to commit only after the author is satisfied). It is not part of this repo, which is the point: any rule this repo's spec-driven workflow depends on must live in the repo, or a fresh clone gets a different assistant.
- Chose the spec format. Options: plain Markdown requirements, Given/When/Then scenarios, or a tool-driven workflow. Picked GitHub Spec Kit because it is open source, agent-agnostic, and widely recognized. Trade-off accepted: it is labeled experimental, so the version is pinned.

## Step 00 · Project rules · tag: step-00-project-rules
- Goal: write the project rules in `CLAUDE.md` (how Claude Code works in this repo) and start this journey log, before any Spec Kit or calculator work.
- Command/prompt: "I wanted you to have a log of all the sequence of steps and finally draft an excalidraw diagram to understand the state of the project from init to featureX. And also the commit history should tell the story and must be demoable when I am trying to explain the spec-driven development as a tech talk"
- AI produced: `CLAUDE.md` (workflow, journey-log format, commit/tag conventions, Go rules), `docs/journey.md`
- Human review, in rounds:
  1. Removed a "public-repo hygiene" section: it protected the author's privacy but did not change how this repo gets built. Test: would a fresh clone behave wrongly without this line?
  2. Questioned "(non-negotiable)" and "no code without an approved spec". Emphasis words do not enforce anything (CLAUDE.md is a prompt, not a guardrail), and "approved" was undefined. Now: approval = reviewed, then committed and tagged; Go code only for a task in a committed `tasks.md`.
  3. Asked whether the phase order matched Spec Kit. Verified by running `specify init` 1.4.0 in a scratch folder outside this repo: the draft rules were missing `checklist`, `analyze`, and `converge`, used a nonexistent `--ai` flag, and assumed feature branches are created by default (they need `--extension git`).
  4. Kept `checklist` so no phase is skipped; corrected the step count to 8 phases and 9 commits per feature (implement = red + green).
  5. Added missing commit types (`docs:`, `docs(tasks):`) and made `--no-ff` merge commits untagged on purpose.
  6. Asked whether this journey log is needed at all, since git can hold the same text. Kept it: prompts, human review, and pre-work have no other home, and most readers never clone. Removed the duplication instead: commit message = subject plus one line of why, tag = one-line summary, journey = detail.
  7. Renamed this step from "Demo rules" (unclear: demo of what?) to "Project rules". Kept zero-padded numbers (`step-00`) because git sorts tags as text: unpadded, `step-10` lists before `step-2`.
  8. Asked what the "Spec Kit section markers" rule meant. It came from older Spec Kit versions, where `plan` ran `update-agent-context.sh` to rewrite a marked block in `CLAUDE.md`. Checked 1.4.0: that script is no longer shipped and nothing touches `CLAUDE.md`. Deleted the rule; if Spec Kit starts editing the file, the diff review will show it and the rule comes back based on observed behavior.
  9. Removed ambiguous wording, prompted by "does this mean X or Y?" questions: "demo" could mean the repo, the calculator, or the talk; "the spec" was being read as `CLAUDE.md`. Fixed by naming the thing: "spec" now only ever means `specs/NNN-*/spec.md`, rules point to exact files, ID examples (`FR-001`, `T012`) were checked against Spec Kit's templates, and every entry gained a Goal line and a precise definition of Command/prompt and Lesson.
- Lesson: Don't trust AI-written project rules (here, `CLAUDE.md`) until you've checked them against the tool they describe (here, Spec Kit), and treat every "what does this mean?" question from a reviewer as a wording bug to fix.

## Step 01 · Spec Kit init · tag: step-01-speckit-init
- Goal: install Spec Kit 1.4.0 into this repo with the Claude Code integration, so the later phases can run as `/speckit-*` commands.
- Command/prompt: `specify init --here --force --non-interactive --integration claude`
- AI produced: 29 new files; `specify init` changed no existing file (`CLAUDE.md` untouched by init, confirming Step 00 round 8).
  - `.claude/skills/speckit-*/SKILL.md` (9): one skill per phase: constitution, specify, clarify, plan, checklist, tasks, analyze, implement, converge.
  - `.specify/` (20): templates, helper scripts, the workflow definition, and a placeholder `memory/constitution.md`, filled in at Step 02.
  - `.gitignore`: added `.claude/settings.local.json` after init warned that agent folders can hold credentials. `.claude/skills/` stays committed because a fresh clone needs it.
- Human review, in rounds:
  1. First attempt also passed `--extension git` (as the Step 00 rules required) and produced 57 files, including auto-commit hooks around every phase (all disabled by default). The author rolled it back: too much for a newcomer to wade through, and the only benefit was branch creation, which `git checkout -b 001-<name>` already does. Updated `CLAUDE.md` to drop the extension and create feature branches by hand, then re-ran init.
  2. Considered skipping the tool and following Spec Kit's method by hand (zero files up front). Kept the tool: the talk should show a real, recognized toolkit the audience can install and replay.
  3. Asked whether each change should reach `main` through a PR. Yes, one PR per branch, not per step: the per-step review already happens before each commit, and the PR reviews the whole unit. Moved the unpushed Step 00 commit from `main` to a new `000-project-setup` branch, which will carry Steps 00 to 02. PRs must use "Create a merge commit", because squash or rebase would leave the step tags off `main`.
  4. Asked twice who a "Lesson" is for. Renamed the field to "Takeaway for readers" from this step on, since the label itself was the ambiguity. Step 00 above still keeps "Lesson" because the log is append-only.
- Takeaway for readers: Install only the parts of a tool you need, and keep the repo readable for newcomers. Here, Spec Kit's optional git extension (`--extension git`) added 28 files and auto-commit hooks, yet its only benefit was creating each feature branch automatically, which one command (`git checkout -b 001-basic-arithmetic`) does by hand.
