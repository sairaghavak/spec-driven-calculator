# Feature Specification: Remainder Operator

**Feature Branch**: `002-modulus`

**Created**: 2026-11-25

**Status**: Draft

**Input**: User description: "Add % to the calculator. It gives the remainder so 7 % 3 is 1 and 4 % 2 is 0. Only whole numbers are allowed: if either number has a decimal point, like 7.5 % 10, 10 % 7.5, or even 7.0 % 2, reject it with "7.5" is not a whole number; % accepts only whole numbers like 7 or 10. Ignore minus signs, so -7 % 3, 7 % -3 and -7 % -3 all give 1. If the right number is 0 or -0, show the same error as division: division by zero is not allowed. All the existing rules still apply, like the number format, leading zeros and the digit limit, but not the 200 limit for multiplication. When there are several errors, the whole-number error comes after the existing errors for that number (not a number, leading zeros, too many digits) and before division by zero, so 5 % 0.000 shows the digit-limit, whole-number and division-by-zero errors in that order. Add % to the list of operators in the welcome message and in the error messages, and use 4 ^ 2 as the example of an unknown operator. Keep it to 1 requirement and 1 user story."

> [!NOTE]
> The author asked for exactly 1 requirement and 1 user story. So [FR-001](#functional-requirements) holds several
> numbered rules, and [US1](#us1--remainder-of-two-whole-numbers-p1) checks all of them.

## Summary

The calculator gets a fifth operator, `%`, which gives the remainder after dividing two whole
numbers: `7 % 3` is `1`. Minus signs are ignored, a number with a decimal point is rejected with
its own message, and a right number of 0 is division by zero. Every rule of
[feature 001](../001-basic-arithmetic/spec.md) still applies, except the multiplication limit.

## Scope

> [!IMPORTANT]
> **In scope**: `%` between two **whole numbers written without a decimal point**, for example
> `7 % 3` or `-7 % 3`, under the same input rules as feature 001: two numbers, one operator, at
> most 3 digits.
>
> **Out of scope**: decimal numbers with `%`, such as `7.5 % 2` or even `7.0 % 2` (rejected,
> [FR-001](#functional-requirements)); negative remainders (`-7 % 3` gives `1`, not `-1`); `%` as a percent sign, such as
> `50%` (rejected as a missing number).

## Functional Requirements

- **FR-001**: The calculator MUST accept `%` as a fifth operator, under these rules:
  1. **Result**: the remainder after dividing the left number by the right one, with minus signs
     ignored. `7 % 3` gives `1` and `4 % 2` gives `0`; `-7 % 3`, `7 % -3`, and `-7 % -3` all give
     `1`. The result is never negative.
  2. **Whole numbers only**: a number written with a decimal point, on either side of `%`, MUST be
     rejected with the whole-number message, even when its value is whole: `7.5 % 10`,
     `10 % 7.5`, and `7.0 % 2` are all rejected.
  3. **Zero on the right**: a right number whose value is 0 (`0` or `-0`) MUST be rejected with
     feature 001's division-by-zero message ([FR-007](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001). A zero written with a decimal
     point, such as `0.0`, gets the whole-number message too.
  4. **Feature 001 rules**: every rule of feature 001 applies to `%` as to the other operators
     (input shape, spaces, signs, number format, leading zeros, digit limit, how results and errors
     are shown, both modes), except the multiplication limit ([FR-008](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001):
     `999 % 250` gives `249`.
  5. **Order of errors**: "not a whole number" is a new kind of error in feature 001's order
     ([FR-015](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001). The order for `%` is: not a number (including leading zeros),
     digit limit, not a whole number, division by zero; within one kind, the left number comes
     before the right. A number that is not a number is not checked further. A number that breaks
     the digit limit is still checked for the other two kinds, and a number that is not whole is
     still checked for division by zero. So `5 % 0.000` shows the digit-limit, whole-number, and
     division-by-zero messages, in that order.
  6. **Operator lists**: the welcome message and every error message that lists the operators
     MUST include `%`, and the example of an unknown operator becomes `4 ^ 2`. The exact texts are
     in [Texts that change in feature 001](#texts-that-change-in-feature-001).

### Error messages

The exact text of each message this feature adds to the `Errors:` block
([FR-015](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001). `<text>` is the number exactly as typed. Tests compare these letter
for letter.

| Problem | Requirement | Message |
|---|---|---|
| Not a whole number, e.g. `7.5` | [FR-001](#functional-requirements) | `"<text>" is not a whole number; % accepts only whole numbers like 7 or 10` |
| Right number is 0, e.g. `5 % 0` | [FR-001](#functional-requirements) | `division by zero is not allowed` (the same text as feature 001) |

### Texts that change in feature 001

These texts in feature 001's [`spec.md`](../001-basic-arithmetic/spec.md) are replaced by the new text below. Every other
text stays the same.

| Where in feature 001 | Old text | New text |
|---|---|---|
| List of operators in the input rule ([FR-001](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001) and in [US2](../001-basic-arithmetic/spec.md#us2--get-an-informative-message-for-bad-input-p2) criterion 2 | `+`, `-`, `*`, `/` | `+`, `-`, `*`, `/`, `%` |
| Welcome message, third line ([FR-012](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001) | `Operations: +, -, *, /` | `Operations: +, -, *, /, %` |
| End of all six input-shape messages ([Error messages](../001-basic-arithmetic/spec.md#error-messages) of feature 001) | `where <operator> is one of +, -, *, /` | `where <operator> is one of +, -, *, /, %` |
| Example of an unknown operator ([Error messages](../001-basic-arithmetic/spec.md#error-messages), [Which shape message is shown](../001-basic-arithmetic/spec.md#which-shape-message-is-shown), and [US2](../001-basic-arithmetic/spec.md#us2--get-an-informative-message-for-bad-input-p2) criterion 2 of feature 001) | `4 % 2`, naming `%` | `4 ^ 2`, naming `^` |

The new welcome line has 25 characters, still fewer than the 59 of the second line, so the `=` and
`-` lines stay 59 characters long.

## Non-Functional Requirements

None new. Feature 001's [NFR-001](../001-basic-arithmetic/spec.md#non-functional-requirements) (each answer in under 1 second) and [NFR-002](../001-basic-arithmetic/spec.md#non-functional-requirements) (the user can
fix bad input from the message alone) apply to `%` as to the other operators.

## User Stories

### US1 · Remainder of two whole numbers (P1)

As a user, I want to enter two whole numbers with `%` and see the remainder, so that I can check
whether one number divides another, or see what is left over, without working it out by hand.

**Covers**: [FR-001](#functional-requirements)

**Acceptance Criteria** (Given an interactive session, unless the criterion says one-shot mode):

1. When the user enters `7 % 3`, then it shows `Result: 1`.
2. When the user enters `4 % 2`, then it shows `Result: 0`.
3. When the user enters `-7 % 3`, `7 % -3`, or `-7 % -3`, then each shows `Result: 1`.
4. When the user enters `999 % 250`, then it shows `Result: 249` (no multiplication limit).
5. When the user enters `7.5 % 10`, then it shows the `Errors:` block, and no result:
   ```
   Errors:
   - "7.5" is not a whole number; % accepts only whole numbers like 7 or 10
   ```
6. When the user enters `10 % 7.5` or `7.0 % 2`, then it shows the same message, naming `7.5` or
   `7.0`.
7. When the user enters `5 % 0` or `5 % -0`, then it shows the `Errors:` block with the single
   item `- division by zero is not allowed`.
8. When the user enters `5 % 0.000`, then it shows:
   ```
   Errors:
   - "0.000" has too many digits; numbers may have at most 3 digits before and 2 digits after the decimal point (largest 999.99)
   - "0.000" is not a whole number; % accepts only whole numbers like 7 or 10
   - division by zero is not allowed
   ```
9. When the user enters `4 ^ 2`, then it shows:
   ```
   Errors:
   - "^" is not a supported operator; expected <number> <operator> <number>, where <operator> is one of +, -, *, /, %
   ```
10. When the session starts, then the welcome message's third line is `Operations: +, -, *, /, %`.
11. Given one-shot mode, when the calculator is started with `7 % 3`, then it shows `1` and exits,
    signaling success; when started with `7.5 % 10`, then it shows the `Errors:` block from
    criterion 5 on the error output and exits, signaling failure.

## Edge Cases

- `7%3` and `7%-3`: valid, both show `1`; spaces are optional and a `-` right after the operator is
  a sign ([FR-001](#functional-requirements) rule 4).
- `3 % 7`: shows `3`; `0 % 5` and `-0 % 5`: show `0`, not `-0` ([FR-001](#functional-requirements) rule 1).
- `999 % 999`: shows `0`; `999 % 7`: shows `5` ([FR-001](#functional-requirements) rule 1).
- `1000 % 7`: only the digit-limit message for `1000` ([FR-001](#functional-requirements) rule 4).
- `007 % 2`: only the leading-zeros message; a number that is not a number is not checked further
  ([FR-001](#functional-requirements) rule 5).
- `7.5 % 2.5`: two whole-number messages, `7.5` first ([FR-001](#functional-requirements) rule 5).
- `2.5 % abc`: the not-a-number message for `abc`, then the whole-number message for `2.5`
  ([FR-001](#functional-requirements) rule 5).
- `2.5 % 1000.5`: the digit-limit message for `1000.5`, then the whole-number messages for `2.5`
  and `1000.5` ([FR-001](#functional-requirements) rule 5).
- `abc % 0`: the not-a-number message for `abc`, then division by zero, the same as `abc / 0`
  ([FR-001](#functional-requirements) rule 5).
- `7.5 % 0`: the whole-number message for `7.5`, then division by zero ([FR-001](#functional-requirements) rule 5).
- `7 % 0.0`: the whole-number message for `0.0`, then division by zero ([FR-001](#functional-requirements) rules 3 and 5).
- `7.5 + 1`: shows `8.5`; the whole-number rule is for `%` only ([FR-001](#functional-requirements) rule 2).
- `7 % 3 % 2`: only the too-many-numbers message ([FR-001](#functional-requirements) rule 4).
- `50%`: only the missing-number message ([FR-001](#functional-requirements) rule 4).

## Key Entities

N/A: `%` adds no new data. The entities are feature 001's: operation, result, and error message.

## Assumptions

- Feature 001's [`spec.md`](../001-basic-arithmetic/spec.md) is not edited: it records feature 001 as it was built. This spec lists
  every feature 001 text it replaces ([Texts that change in feature 001](#texts-that-change-in-feature-001)); where the two differ,
  this spec wins.
- "The whole-number error comes after the existing errors for that number" is read with feature
  001's grouping by kind ([FR-015](../001-basic-arithmetic/spec.md#functional-requirements) of feature 001): the whole-number messages for both numbers come after
  the digit-limit messages for both numbers, as in the `2.5 % 1000.5` edge case. The other reading,
  all messages for the left number before any for the right, would break that existing rule.
- "Whole number" depends on how the number is written, not on its value: any decimal point makes it
  not whole, so `7.0` and `0.000` are rejected (from the author's `7.0 % 2` example).
- In one-shot mode, `%` needs no quotes in bash, the shell the demo uses, unlike `*` (see the note
  in feature 001's [US3](../001-basic-arithmetic/spec.md#us3--multiplication-operand-limit-p3)).
  Checked by running it; other shells were not checked.

## Definition of Done

- Every acceptance criterion and edge case passes as an automated test that cites [FR-001](#functional-requirements).
- Feature 001's tests still pass, with the texts listed in
  [Texts that change in feature 001](#texts-that-change-in-feature-001) updated to their new text.
