# Mobile Application as Frontend Platform

## Status

Accepted.

## Context

Need a frontend. Time-sensitive push notification is extremely important for service/departure alerts, as public transit 
relies on real-time, rapidly changing conditions.

Platform choices included:

1. Web
2. Mobile
    1. Native
    2. Cross platform
    3. Web Wrapper

## Decision

Mobile application using a cross-platform framework (e.g. React Native or Flutter). Provides access to 

1. System notification APIs
2. Background execution
3. Local scheduling
4. Retains codebase portability across iOS and Android
5. More native UI/UX look compared to web wrapper

Most important is mobile provided superior push notification compared to web.

1. Users need to install the web app as a PWA before notification permission can be granted, and there's no reliable 
way to detect PWA eligibility and installed status.
2. On iOS for PWA, notification actions are not available (e.g. press-and-hold menu showing "Delete" or "Mark as read" 
actions).
3. For PWA, server need to keep track of active notification subscriptions, and dedup by unique device. No webhook to 
inform when a notification failed due to no more permission, need self check and delete.
4. PWA has limited background processing capability when compared to native mobile app.
5. Mobile support scheduling notifications, so users can receive them even if offline. PWA does not support as relies
on server sending Google/Apple a request when both parties are online.

## Consequences

1. Web platform cannot easily be supported in future without a new frontend codebase.
2. User need to manually install the app to start using, instead of visiting a URL and begin.
3. Harder app updates as need to pass Google/Apple review process which can be tedious and time consuming.
4. Backend APIs must maintain strict backward compatibility and if a breaking change is necessary, need to implement
a force upgrade mechanism to handle outdated client versions.
5. Requires a more sophisticated CI/CD pipeline to support code signing and building runners.
6. Ongoing App Store developer subscription cost.

## Compliance

None needed.

## Notes

Author: bingbeann
Approved: bingbeann, 2026-09-16
Last Updated: 2026-09-16

