# Research: Remainder Operator

The [spec](spec.md) says *what* `%` does. This file records *how* it is built: the choices made
while planning, R1 to R6 (R = research decision).

## Why these decisions were needed

None of them is a gap in the spec. They are what the code and the tests need in order to meet the
spec and the [constitution](../../.specify/memory/constitution.md).

| Decision | What caused it |
|---|---|
| R1 | [FR-001](spec.md#functional-requirements) rule 1: `-7 % 3` must give `1`, but common languages give `-1` or `2`. |
| R2 | [FR-001](spec.md#functional-requirements) rule 4 and the edge cases `7%3` and `50%`: the calculator must find `%` in the input with or without spaces. |
| R3 | [FR-001](spec.md#functional-requirements) rule 2: `7.0` must count as not whole, although its value is whole. |
| R4 | [FR-001](spec.md#functional-requirements) rule 5: the new check must fit into feature 001's order of errors. |
| R5 | The constitution's Principle II: every test names its requirement, and now both features have a requirement with the same ID ([FR-001](spec.md#functional-requirements) here, [FR-001](../001-basic-arithmetic/spec.md#functional-requirements) in feature 001). |
| R6 | The spec's [Definition of Done](spec.md#definition-of-done): feature 001's tests must keep passing with the texts this feature replaces. |

Planning found no gap in the spec: every decision below only implements a rule the spec already
states, so no requirement changed.

## Format

Each decision uses Spec Kit's three fields:

- **Decision**: what was chosen.
- **Rationale**: why it was chosen, shown with what a user or a script sees.
- **Alternatives considered**: the main other option, and why it lost.

## R1 · Computing the remainder

- **Decision**: drop the minus signs of both numbers first, then take the remainder (Go's
  `big.Int.Rem`, like Java's `BigInteger.remainder`). The numbers are already stored exactly, as
  feature 001's [R1](../001-basic-arithmetic/research.md#r1--exact-results) chose; a whole number is
  read from that stored value.
- **Rationale**: `-7 % 3`, `7 % -3`, and `-7 % -3` all show `1`, as
  [FR-001](spec.md#functional-requirements) rule 1 requires. Only whole numbers with a right number
  other than 0 reach this step ([R4](#r4--where-the-new-check-goes)).
- **Alternatives considered**: use the language's own `%` on the signed numbers. We did not choose
  this. Go's `%` works like Java's: the result takes the sign of the left number. If we had used
  it, `-7 % 3` would show `-1`. Python's `%` takes the sign of the right number: `-7 % 3` would show
  `2` and `7 % -3` would show `-2`. Both break rule 1 (checked by running them).

## R2 · Finding `%` in the input

- **Decision**: `%` becomes an operator character, like `+`, `-`, `*`, and `/`. The input is split
  at it whether or not there are spaces around it. Feature 001's other input checks do not change.
- **Rationale**: users type `%` the same way as the other operators:

  | You type | Result |
  |---|---|
  | `7 % 3` or `7%3` | `1` |
  | `7%-3` | `1`: a `-` right after the operator is a sign, as in feature 001 |
  | `50%` | the missing-number message |
  | `7 % 3 % 2` | the too-many-numbers message |
  | `4 ^ 2` | the unknown-operator message, naming `^` |

- **Alternatives considered**: accept `%` only as a separate word, between spaces. We did not
  choose this. If we had, `7%3` would be one word, and the calculator would report a missing
  number, although feature 001 makes spaces optional for every operator.

## R3 · Whole numbers are decided by how they are written

- **Decision**: a number is whole when its text has no decimal point. The check reads the text,
  not the value.
- **Rationale**: `7.0 % 2` and `5 % 0.000` are rejected, as
  [FR-001](spec.md#functional-requirements) rule 2 requires, even though `7.0` equals 7.
- **Alternatives considered**: check the value (Go's `big.Rat.IsInt`). We did not choose this. It
  reports `7.0` as whole (checked by running it). If we had used it, `7.0 % 2` would show `1`
  instead of the whole-number message.

## R4 · Where the new check goes

- **Decision**: for `%`, the whole-number check runs after the digit-limit check and before the
  division-by-zero check. It checks every number that has the right format, including one that
  broke the digit limit. A number that fails it still has a value, so a right number of `0.000`
  still reaches the division-by-zero check, which now runs for both `/` and `%`.
- **Rationale**: the messages come out in the order of
  [FR-001](spec.md#functional-requirements) rule 5:
  - `5 % 0.000` shows the digit-limit, whole-number, and division-by-zero messages.
  - `2.5 % 1000.5` shows the digit-limit message for `1000.5`, then the whole-number messages for
    `2.5` and `1000.5`.
- **Alternatives considered**: treat a number that is not whole like one that is not a number,
  which stops all further checks on it. We did not choose this. If we had, `5 % 0.000` would show
  no division-by-zero message, which rule 5 requires.

## R5 · Naming the requirement in tests, now that two features share IDs

- **Decision**: tests for feature 002 name their requirement as `002/FR-001`. Feature 001's tests
  keep their short form, such as `FR-005`, which means feature 001.
- **Rationale**: each name still points to exactly one requirement. Searching the tests for
  `002/FR-` lists every test of feature 002; the short form lists feature 001's.
- **Alternatives considered**:
  - Use `FR-001` for both features. We did not choose this. If we had, a test naming `FR-001` could
    mean either feature, and a check that every requirement has a test could be fooled.
  - Rename feature 001's names too (`001/FR-005`). We did not choose this. If we had, all 90
    existing test cases would change, with no change in what they check.

## R6 · Feature 001 tests whose text changes

- **Decision**: the tests that compare a text this feature replaces get the new text and keep
  their feature 001 requirement. These are the welcome message, the six input-shape messages, and
  the unknown-operator example, which changes from `4 % 2` to `4 ^ 2`. Feature 002's own acceptance
  criteria get their own test cases. So `4 ^ 2` and the welcome line are each checked by one test
  of each feature.
- **Rationale**: feature 001's rules still decide *which* message appears
  ([FR-005](../001-basic-arithmetic/spec.md#functional-requirements) and
  [FR-012](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001); this feature
  only changes the wording ([Texts that change in feature 001](spec.md#texts-that-change-in-feature-001)).
  Keeping their names keeps those rules tested.
- **Alternatives considered**: move those tests to `002/FR-001`. We did not choose this. If we had,
  feature 001's [FR-012](../001-basic-arithmetic/spec.md#functional-requirements) would lose the
  tests that check the full session output, and a converge run would report it untested.
