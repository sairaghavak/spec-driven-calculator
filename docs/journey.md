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
