# LEARNING_LOG.md

## Meta
- Session count: 0 (setup complete)
- Setup date: 2026-07-09
- Environment: macOS, Python 3.14.6, Go 1.26.4, 60 min/session
- Repo: https://github.com/Drazerr/python-go-coach (create in Session 1 if not yet created)

## Baseline Assessment (2026-07-09)
- Python: 1.5 (copied/modified examples; prior project not specified — ask in Session 1)
- Go: 0 (complete beginner)
- Linux/shell: 3 | Git: 2 (branches + merges) | Testing: 2 (has written unit tests)
- HTTP/APIs: 0 ← priority gap, introduced Session 4, reinforced constantly
- Docker: 3 | AWS: 3 | Azure: 2 | GCP: 1 | Kubernetes: 2
- Target role: DevOps / SRE / Infra / Platform
- Profile: ops-strong, code-light. Use ops contexts in all exercises. Skip Git/shell basics.

## Roadmap Summary (adaptive)
- S1: Log error summarizer (Easy) — files, dicts/maps, argparse, first Go program, first tests
- S2: Config/env validator (Easy) — structs, exit codes, error wrapping, table-driven tests
- S3: JSON↔YAML converter (Easy) — encoding/json, struct tags, malformed input
- S4: Sequential endpoint health checker (Easy-medium) — first HTTP; requests/net/http; mocking
- S5: Review #1
- S6–9: CSV reports, cert-expiry/dates, structured logging, CLI subcommands
- S10: Review #2
- S11–14: Pagination, retries/backoff, stale-resource detector, config comparison (Stage 2; branch-per-exercise Git starts)
- S15: Review #3
- S16–19: SQLite persistence, DI/interfaces, resilient client, first concurrency
- S20: Review #4
- S21–24: Worker pools, concurrent checker at scale, log pipeline, K8s auditor (Stage 3)
- S25: Review #5 + capstone selection
- S26–30: Capstone (own repo: README/tests/Dockerfile/CI) + final mock interview
- Adaptation triggers: compress Stage 1 if S1–3 easy; Go consolidation session if structs/errors cause repeated friction

## Progress
- Exercises completed: none
- Python concepts learned: —
- Go concepts learned: —
- Common mistakes: —
- Topics needing revision: —
- Testing skills demonstrated: — (baseline claim: 2/3)
- Debugging skills demonstrated: —
- Interview strengths: ops domain knowledge (expected)
- Interview weaknesses: — (expected: explaining code, API interaction)
- Projects completed: none

## Current State
- Difficulty level: Easy
- Rubric ratings: not yet evaluated
- Revision exercise before next session: none (first session)

## Recommended Next Session
**Session 1 — Log file error summarizer (Python, then Go).** Parse a sample
web-server log, count errors by type/status, print a summary report with a
CLI argument for the file path. Python first (file I/O, dicts, functions,
argparse, one pytest test), then the same task rebuilt in Go (go mod, file
reading, maps, explicit error handling, one Go test). Wrap-up: create the
monorepo, write .gitignore, first commits. Ask what the user previously
built in Python during warm-up.
