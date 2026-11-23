# Feature Specification: [FEATURE NAME]

**Feature Branch**: `[###-feature-name]`

**Created**: [DATE]

**Status**: Draft

**Input**: User description: "$ARGUMENTS"

<!--
  Project override of Spec Kit's spec-template (resolved before the core template).
  Layout: requirements first, then user stories that point to them, in agile terms.
  Section mapping to Spec Kit's core template, for the speckit-* skills:
  - "Non-Functional Requirements" (NFR-###) plays the role of "Success Criteria" (SC-###):
    measurable, technology-agnostic quality targets. Put non-functional answers from
    /speckit-clarify here, and inventory NFR-### where a skill says SC-###.
  - "Definition of Done" replaces outcome criteria that only restate the acceptance criteria.
  - "User Stories" with "Acceptance Criteria" are Spec Kit's user stories with acceptance scenarios.
  - /speckit-clarify adds "## Clarifications" just after the Scope section.
  Write each fact once: a rule lives in one FR or NFR; stories and edge cases cite its ID and add
  examples, they do not restate it.
-->

## Summary

[Two or three sentences: what the feature does and why the user needs it.]

## Scope

<!--
  The feature's boundaries, at the top so no reader misses them. State the biggest limit first,
  with a concrete example of input that is NOT supported.
-->

> [!IMPORTANT]
> **In scope**: [what the feature does, with its main limit, e.g. "exactly two operands"]
>
> **Out of scope**: [what it deliberately does not do, with an example input]

## Functional Requirements

<!--
  One testable rule per FR, using MUST / MUST NOT. Number FR-001, FR-002, ... in order.
  Mark at most 3 open decisions as [NEEDS CLARIFICATION: question].
-->

- **FR-001**: [The system MUST ...]
- **FR-002**: [The system MUST ...]

## Non-Functional Requirements

<!--
  Measurable quality targets (performance, usability, reliability, security), stated as outcomes
  a user or tester can check, with no technology named. Remove the section if none apply.
-->

- **NFR-001**: [Measurable target, e.g. "Each result appears in under 1 second"]

## User Stories

<!--
  Ordered by priority. Each story is a slice that can be built, tested, and demonstrated on its own;
  its acceptance criteria are its independent test. "Covers" lists the FRs and NFRs it exercises.
  Every FR MUST be covered by at least one story. State the shared Given once, then one line per
  example: input -> expected output.
-->

### US1 · [Title] (P1)

As a [user], I want [goal], so that [benefit].

**Covers**: FR-001, FR-002

**Acceptance Criteria** (Given [shared context]):

1. When [action], then [expected outcome].
2. When [action], then [expected outcome].

### US2 · [Title] (P2)

As a [user], I want [goal], so that [benefit].

**Covers**: [FR-...]

**Acceptance Criteria** (Given [shared context]):

1. When [action], then [expected outcome].

## Edge Cases

<!--
  Only boundary inputs not already shown in a story, each with its expected outcome and the FR it
  follows from.
-->

- [Input]: [expected outcome] ([FR-...])

## Key Entities *(include if the feature involves data)*

- **[Entity]**: [what it represents, key attributes, no implementation]

## Assumptions

- [Default chosen where the input did not say, and why]

## Definition of Done

- Every acceptance criterion and edge case passes as an automated test that cites its FR or NFR.
- [Feature-specific completion checks, if any]
