# Current project status

Updated: 2026-11-15. Current focus: **M0 — Project foundation**.

## Baseline

- Working Go prototype: registration/hex lookup, callsign lookup, nearby traffic, regional airline filtering, JSON/table output.
- Prototype commit: `b0ff308`.
- Recovery tag: `checkpoint/prototype-2026-11-15`.
- Planning checkpoint: `checkpoint/planning-2026-11-15` (created with the planning commit).
- Repository is local at present; no GitHub remote or public release exists.

## Progress

| Area | State | Evidence / remaining work |
| --- | --- | --- |
| Prototype | Done | Source and tests committed in `b0ff308`; live smoke checks completed 2026-11-15 |
| Local version control | Done | Main branch and prototype recovery tag |
| Roadmap and recovery process | Prepared | `ROADMAP.md`, `RECOVERY.md`, `CONTRIBUTING.md` |
| Automated checks | Local checks pass; hosted run pending | Workflow in `.github/workflows/ci.yml` |
| GitHub repository and board | Pending destination/visibility | Recommended: private `Mickyj70/sky` initially |
| Output contract and additional fixtures | Planned | M0-02 in `BACKLOG.md` |
| Watchlists and scheduler | Not started | M1-01 through M1-04 |
| History, commercial flight data, releases | Not started | M2–M5 |

## Verification of this planning checkpoint

- Local `go test -race ./...`, `go vet ./...`, and binary build passed with Go 1.27.4.
- Markdown links and workflow YAML were checked locally. Hosted CI and the minimum-version job remain unverified until GitHub setup.
- Recovery was demonstrated: created a separate detached worktree from the prototype tag, built it, ran `sky --version`, and removed the temporary worktree without altering main.

## Next work, in order

1. Confirm the GitHub destination and visibility; create remote, push commits/tags, and open the M0–M1 issues from `BACKLOG.md`.
2. Confirm the hosted CI run is green and configure required checks where supported.
3. Finish M0-02: define compatibility/error behavior and cover missing provider contract cases.
4. Design and implement M1-01: config paths, SQLite selection, schema, and backup/restore behavior.

## Open decisions

- GitHub owner/repository and visibility: private `Mickyj70/sky` is the recommendation, not an already-created repository.
- Code license: choose before public distribution; provider data license is separate.
- SQLite driver: decide with platform/release requirements in M1-01.
- Provider for scheduled flights: evaluate in M3; no paid provider is configured or required today.

## Update protocol

Update this file in each progress PR: what completed, validation performed, next issue, and any blockers. Keep it consistent with GitHub milestones. Avoid maintaining duplicate issue checklists here; `BACKLOG.md` is the initial seed until GitHub issues exist.
