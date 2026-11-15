# GitHub setup plan

Recommendation: begin with a private repository named `sky` under the authenticated account `Mickyj70`. This is a proposed destination; it has not been created. Confirm the destination and visibility before publishing the local source.

GitHub hosts source, issues, checks, and releases. It does not run the tracking process continuously or back up local SQLite data.

## Setup sequence

1. Create the repository with the selected visibility; do not initialize a second README/history remotely.
2. Add `origin` and push main plus the checkpoint tags.
3. Confirm the CI workflow runs successfully on GitHub. Local checks do not prove hosted CI success.
4. Create milestones M0–M5 using the roadmap names. Avoid due dates until the relevant design/provider dependencies are resolved.
5. Create labels and the project board described in M0-01. Use Backlog → Ready → In progress → In review → Done; track blockers with a label and dependency links.
6. Create M0–M1 issues from BACKLOG, assign milestones, and link dependent issues.
7. Enable a main-branch ruleset where supported: PR-based changes, required CI checks, blocked force pushes and deletion. Choose the actual check names from the first successful run.
8. For a solo project, use CI and a recorded self-review; do not configure mandatory approvals in a way that prevents the owner from merging. Revisit independent review when another maintainer joins.
9. Record repo, board, milestone and successful CI links in STATUS. Keep milestone completion based on acceptance evidence.

If the intended private destination is confirmed, the initial CLI operation would be:

```sh
# Run from the local project; this publishes committed source and documentation.
gh repo create Mickyj70/sky --private --source=. --remote=origin --push
# Push the already-created recovery tags explicitly.
git push origin checkpoint/prototype-2026-11-15 checkpoint/planning-2026-11-15
```

If the repository name is already taken or the owner differs, inspect that destination and choose the correct repo rather than overwriting unrelated history.

## Daily progress

- Keep one implementation issue In progress initially.
- Link the issue in its PR, including acceptance evidence and rollback instructions.
- Merge only with passing checks; update STATUS and CHANGELOG in the same change when applicable.
- At a milestone boundary, verify recovery, create an annotated checkpoint, and push the tag.
- Use Releases for actual downloadable versions once distribution is implemented. A milestone target is not automatically a release.

## Before public distribution

Choose the source-code license with the owner, verify provider attribution and redistribution terms, and review the installation/release instructions. No source license has been selected by this plan, and the provider's ODbL data license is not the code license.
