# Kotlin Multiplatform as Mobile Framework

## Status

Accepted.

## Context

Need a mobile framework to support cross platform mobile application
development. Choices included:

* React Native
    * JavaScript/TypeScript
    * Renders native OS components
    * Massive ecosystem
* Kotlin Multiplatform (KMP)
    * Kotlin (Shared) + Swift (iOS UI)
    * Shared logic layer with native UIs
    * Growing ecosystem
* Flutter
    * Dart
    * Draws pixel-perfect custom UI on canvas
    * Large ecosystem

## Decision

Kotlin Multiplatform.

Although React Native is appealing with the write-it-once and native UI look,
maintenance and tooling remains a problem even with help of Expo. KMP solved it
by having a shared logic layer to avoid duplicated logic, while allowing native
UI layer with Swift/Compose to avoid the intermediary layers that caused RN the
issues.

## Consequences

* More work to implement and maintain native UI on both OSes.
* Overhead to learn Kotlin/Swift languages.

## Compliance

None needed.

## Notes

Author: bingbeann
Approved: bingbeann, 2026-09-21
Last Updated: 2026-09-21
