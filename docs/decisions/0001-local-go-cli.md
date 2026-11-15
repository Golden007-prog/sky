# ADR 0001: Go CLI with local execution

Date: 2026-11-15
Status: Accepted (Go was selected by the user; future storage/provider details remain open).

## Context

Sky tracks aircraft from external HTTP APIs. Current needs are lookups, terminal output, automation-friendly JSON, and future saved tracking. No shared service or continuous background hosting is required yet.

## Decision

- Use Go with Cobra for the CLI.
- Run provider requests locally with user-supplied credentials where needed.
- Keep provider HTTP handling outside the command package; introduce a provider-neutral contract before adding scheduler or additional providers.
- Use adsb.lol for the existing observed-aircraft commands. It does not provide a complete scheduled-flight or fleet model.
- Plan SQLite for saved tracking/history, but select the driver in a separate ADR after considering platform builds and backup behavior.
- Use Git commits, annotated checkpoints, and eventually GitHub PRs/issues/CI for change tracking.

## Consequences

The initial app needs no server deployment. Tracking runs only while the process runs. Each user's provider limits and credentials apply locally. A future hosted alerting service is a new architectural decision with operational and cost implications.

Provider data terms are distinct from the source-code license. The project source uses MIT under the owner's instruction to publish as open source.
