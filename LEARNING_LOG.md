# LEARNING_LOG.md

## Meta
- Session count: 1 (complete)
- Last updated: 2026-07-12
- Environment: macOS, Python 3.14.6, Go 1.26.4, 60 min/session
- Repo: https://github.com/Drazerr/python-go-coach (confirm pushed + go.mod path fixed at S2 start)

## Baseline Assessment (2026-07-09, updated S1)
- Python: 2 | Go: 1 | Linux/shell: 3 | Git: 2 | Testing: 2 | HTTP/APIs: 0 (priority gap, S4)
- Docker: 3 | AWS: 3 | Azure: 2 | GCP: 1 | K8s: 2
- Target: DevOps/SRE/Infra/Platform. Ops-strong, code-light. Ops contexts everywhere.

## Roadmap Summary (adaptive)
- S1 ✅ — S2: Config/env validator — S3: JSON↔YAML — S4: HTTP health checker — S5: Review #1
- S6–9: CSV reports, cert expiry, structured logging, CLI subcommands — S10: Review #2
- S11–14: Stage 2 (pagination, retries, stale resources, config diff; branch-per-exercise Git)
- S15: Review #3 — S16–19: SQLite, DI/interfaces, resilient client, first concurrency — S20: Review #4
- S21–24: Stage 3 — S25: Review #5 — S26–30: Capstone + final mock interview
- Adaptation note: S1 coding pace strong in both languages; if S2–3 land easily, compress
  Stage 1. Do NOT accelerate past communication practice — that's the lagging skill.

## Session 1 (2026-07-12) — Log error summarizer ✅
Python: v1 + dataclass refactor + pytest (verified). Go: struct, io.Reader seam,
(LogSummary, error) + scanner.Err(), sorted output helper, go test verified incl.
watched-it-fail drill. Comparison + interview reflection done. First commits made
(messages need imperative mood from S2 on).

## Progress
- Python concepts: file iteration, dict.get counting, argparse, dataclasses, OSError,
  exit codes, pytest basics, iterable seam for testability
- Go concepts: modules, os.Args, if err != nil, defer, bufio.Scanner + Err(),
  strings.Fields, strconv.Atoi, maps, structs (named-field literals), error returns,
  sort.Slice, strings.NewReader, go test + vet format checks
- Common mistakes: Fprintln/Println confusion; int-vs-string map key; scanner.Err()
  inside loop (dead code); capitalized locals; positional struct literal; copy-paste
  test assertions; pasting unrun code (coached: compiler first, always paste output)
- Topics needing revision: sort tie-breaking (exercise assigned); tests only catch
  what they assert (Q3 failed); streaming vs loading (Q2 property misnamed)
- Testing skills: writes tests in both languages, verifies runs, knows watched-it-fail;
  does NOT yet reason about coverage gaps (Q3). Rating holds at 2.
- Debugging: reads compiler errors with guidance
- Interview strengths: designs for testability instinctively (both languages)
- Interview weaknesses: PRIMARY GAP — verbal/written explanation of own code.
  Q1 two fragments vs 5–8 sentences; Q2 right conclusion wrong reason; Q3 wrong.
  Every session must include an explain-aloud drill until this moves.

## Current State
- Difficulty: Easy → Easy-medium for S2 (coding earned it; keep concepts contained)
- Rubric (S1 final): Correctness: Competent | Robustness: Developing | Readability:
  Competent | Structure: Competent | Testing: Developing | Error handling: Competent |
  Maintainability: Competent | Language usage: Py Competent / Go Developing |
  Operational thinking: Developing | Interview communication: Needs work
- Revision exercise before S2: (1) alphabetical tie-breaker in printSortedCounts +
  re-run go test; (2) write 5–8 sentence explanation of summarize unaided — reviewed
  at S2 warm-up

## Recommended Next Session
**S2 — Config/env validator (Easy-medium).** A deploy script needs required env vars /
a config file validated: presence, types, simple rules (e.g. port range), reporting ALL
problems not just the first, non-zero exit on failure. Python: dataclass config +
pytest.mark.parametrize. Go: struct validation, error wrapping with fmt.Errorf %w,
table-driven tests (the canonical Go pattern). Warm-up: review both revision-exercise
items; confirm repo pushed and go.mod path fixed.
