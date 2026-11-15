# Development and progress workflow

Start with [current status](docs/STATUS.md), then the [roadmap](docs/ROADMAP.md). New features should have an issue and acceptance criteria before implementation. The issue seed is [BACKLOG.md](docs/BACKLOG.md).

## One change at a time

1. Choose a Ready issue in the current milestone; mark it In progress.
2. Branch from updated main: `feat/watchlist-storage`, `fix/provider-timeout`, or `docs/recovery`.
3. Make small, coherent commits. Keep credentials, local databases, and binaries out of Git.
4. Add tests for important behavior or regressions; use local fixtures instead of live provider calls in CI.
5. Open a PR with the issue link, behavior change, verification evidence, and rollback instructions. Update STATUS and CHANGELOG as appropriate.
6. Mark In review, pass required CI, then review and squash-merge. Close the issue only after acceptance criteria are met.
7. Update milestone progress; create an annotated checkpoint at a tested milestone boundary.

Initial planning and the prototype were committed directly to establish the baseline. Subsequent implementation should use the PR workflow after GitHub setup. Main should stay buildable. Never rewrite shared main history to undo a feature; use a revert PR.

## Local checks

Requires Go 1.24+. CI runs the minimum version plus current stable Go on Linux. Cross-platform runtime support will be verified before the distribution milestone promises it.

```sh
gofmt -l cmd internal
go test -race ./...
go vet ./...
go build -o bin/sky ./cmd/sky
```

`gofmt -l` must print no paths. Live smoke checks are optional, consume provider quota, and should record the date and command without recording secrets. An API outage must not fail the offline test suite.

## Decisions and compatibility

Record significant choices under `docs/decisions/` with context, decision, consequences, and status. Examples: storage driver, provider contracts, JSON compatibility, paid provider selection, or a hosted backend.

During alpha development, document breaking changes before merging. Before 1.0, define the stable command/JSON contract and deprecation policy. Data migrations require the recovery plan in [RECOVERY.md](docs/RECOVERY.md).
